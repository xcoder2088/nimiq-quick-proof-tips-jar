package main

import (
	"log"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Without this, Nimiq Pay's embedded mini-app WebView (or any other
	// intermediate cache) can keep serving a stale copy of this page
	// indefinitely after a content update, since http.ServeFile alone
	// doesn't forbid caching — exactly what happened with the old
	// "Built for the Nimiq Mini Apps Competition" footer text still
	// showing inside Nimiq Pay after it was already fixed here.
	w.Header().Set("Cache-Control", "no-cache")

	http.ServeFile(w, r, "./templates/index.html")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	log.Printf("NIM Tip Jar mini app listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
