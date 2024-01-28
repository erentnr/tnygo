package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"regexp"
)

func validateUrl(shortUrl string) bool {
	regex, err := regexp.Compile("^[a-zA-Z0-9]{5}$")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}

	return regex.MatchString(shortUrl)
}

func getOriginalUrl(shortUrl string) string {
	return "https://eren.sh"
}

func redirectToHomepage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://127.0.0.1:8080", http.StatusFound)
}

func renderPage(w http.ResponseWriter, templateFile string) {
	filePath := path.Join("templates", templateFile)
	homeTemplate, err := template.ParseFiles(filePath)
	if err != nil {
		fmt.Fprint(w, "Error", http.StatusInternalServerError)
		fmt.Printf("%s", err.Error())
		return
	}

	if err := homeTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RootHandler() http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {

		shortUrl := r.URL.Path[1:]
		if len(shortUrl) == 0 {
			renderPage(w, "index.html")
			return
		}

		isValidUrl := validateUrl(shortUrl)
		if !isValidUrl {
			redirectToHomepage(w, r)
			return
		}

		originalUrl := getOriginalUrl(shortUrl)
		http.Redirect(w, r, originalUrl, http.StatusFound)
	}

	return http.HandlerFunc(handleFunc)
}
