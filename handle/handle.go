package handle

import (
	"Debouncer/server"
	"fmt"
	"net/http"
)

func SMSHandle(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id required", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	result := server.SMSServer(ctx, userID)
	_, err := fmt.Fprintln(w, result)
	if err != nil {
		return
	}
}
