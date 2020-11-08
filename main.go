package main

import (
	"log"
	"net/http"

	servtatic "github.com/JakBaranowski/mini-blog/servers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	h := servtatic.NewHandler("templates/layout", "templates/bodies", "main")
	http.Handle("/", h)

	log.Print("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Failed to start listen server")
	}
}
