package main

import (
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		delay := r.URL.Query().Get("delay")
		if delay != "" {
			delayTime, err := strconv.Atoi(delay)
			if err == nil && delayTime > 0 {
				time.Sleep(time.Duration(delayTime) * time.Millisecond)
			}
		}
		
		w.Header().Set("Content-Type", "application/json")

		shouldReturned := r.URL.Query().Get("shouldReturned")
		if shouldReturned != "" {
			w.Write([]byte(shouldReturned))
			return;
		}

		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8000", mux)
}