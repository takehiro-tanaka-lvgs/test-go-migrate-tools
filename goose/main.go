package main

import "net/http"

func main() {
	s := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	s.ListenAndServe()
}
