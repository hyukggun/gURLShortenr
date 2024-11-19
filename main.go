package main

import (
	"gURLShortener/util"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var urlStore = make(map[string]string)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.naver.com",
	}

	for _, url := range urls {
		enc := util.ShortUrl(url)
		urlStore[enc] = url
	}

	r := chi.NewRouter()

	r.Post("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
			return
		}

		// Form Data를 읽어온다.
		targetUrl := r.FormValue("url")
		key := util.ShortUrl(targetUrl)
		w.Write([]byte("http://localhost:8080/rd/" + key))
	})

	r.Get("/rd/{shortUrl}", func(w http.ResponseWriter, r *http.Request) {
		shortUrl := chi.URLParam(r, "shortUrl")
		url, ok := urlStore[shortUrl]
		if !ok {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, url, http.StatusSeeOther)
	})

	http.ListenAndServe(":8080", r)
}
