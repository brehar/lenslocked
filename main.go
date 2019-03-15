package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<p>To get in touch, please send an e-mail to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.</p>")
	} else if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	}
}
