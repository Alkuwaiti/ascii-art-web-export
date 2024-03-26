package main

import (
	handler "ascii-art-web/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.GetHandler)
	http.HandleFunc("/ascii-art", handler.PostHandler)
	http.HandleFunc("/download", handler.DownloadHandler)

	http.ListenAndServe(":3000", nil)
}