package src

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type DeBouncer struct {
	Interval time.Duration
	Redis    *redis.Client
}

func NewDeBouncer(interval time.Duration, redisAddr string) *DeBouncer {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0,
	})
	return &DeBouncer{
		Interval: interval,
		Redis:    client,
	}
}
func (b *DeBouncer) TryLock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	result, err := b.Redis.SetNX(ctx, key, "locked", expiration).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}
func (b *DeBouncer) UnLock(ctx context.Context, key string) error {
	return b.Redis.Del(ctx, key).Err()
}
