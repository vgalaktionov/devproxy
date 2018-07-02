package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"syscall"
)

func main() {

	proxyURL, err := url.Parse("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	if err := syscall.Setgid(80); err != nil {
		log.Fatal(err)
	}

	log.Println("starting devproxy on localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
