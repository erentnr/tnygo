package main

import (
	"net/http"

	"github.com/erentnr/tnygo/handlers"
)

func main() {
	http.Handle("/api/url", handlers.UrlHandler())
	http.Handle("/", handlers.RootHandler())

	s := &http.Server{
		Addr: ":8080",
	}
	s.ListenAndServe()
}
