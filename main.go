package main

import (
	"Debouncer/handle"
	"net/http"
)

func main() {
	http.HandleFunc("/send_sms", handle.SMSHandle)
	http.ListenAndServe(":8080", nil)
}
