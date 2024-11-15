package server

import (
	"Debouncer/src"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func SMSServer(ctx context.Context, userID string) string {
	deBouncer := src.NewDeBouncer(1*time.Minute, "localhost:6379")
	locked, err := deBouncer.TryLock(ctx, userID, 5*time.Second)
	if err != nil || !locked {
		return "请求过于频繁，请稍后再试"
	}
	defer deBouncer.UnLock(ctx, userID)
	
	lastTimeStr, err := deBouncer.Redis.Get(ctx, userID).Result()
	if err == nil {
		lastTime, _ := time.Parse(time.RFC3339, lastTimeStr)
		if time.Since(lastTime) < deBouncer.Interval {
			return "请求过于频繁，请稍后再试"
		}
	} else if err != redis.Nil {
		return "Redis 错误，无法处理请求"
	}
	now := time.Now().Format(time.RFC3339)
	if err := deBouncer.Redis.Set(ctx, userID, now, deBouncer.Interval).Err(); err != nil {
		return "Redis 错误，无法更新请求时间"
	}
	// 发送服务的实现
	
	return "短信发送成功"
}
