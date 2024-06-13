package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files from the root directory
	fs := http.FileServer(http.Dir("."))
	mux.Handle("/", fs)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
