package handlers

import (
	"fmt"
	"net/http"
)

func UrlHandler() http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "UrlHandler\n")
	}
	return http.HandlerFunc(handleFunc)
}
