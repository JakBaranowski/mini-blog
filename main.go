package main

import (
	"log"
	"net/http"

	servetastic "github.com/JakBaranowski/servetastic"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	h := servetastic.NewHandler("templates/layout", "templates/content", "bulma")
	http.Handle("/", h)

	log.Print("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Failed to start listen server")
	}
}
