package main

import (
	"net/http"
)

func main() {
	app := http.NewServeMux()

	app.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":5000", app)
}