package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve everything inside ./static
	fs := http.FileServer(http.Dir("./"))

	http.Handle("/", fs)

	log.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
