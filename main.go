package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed ui/build/*
var content embed.FS

func clientHandler() http.Handler {
	fsys := fs.FS(content)
	contentStatic, _ := fs.Sub(fsys, "ui/build")
	return http.FileServer(http.FS(contentStatic))
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("PORT set to %v\n", port)
	}

	mux := http.NewServeMux()
	mux.Handle("/", clientHandler())
	log.Printf("http://localhost:%v\n", port)
	http.ListenAndServe(":"+port, mux)
}
