package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	router := pat.New()
	server := NewPostServer()
	router.Post("/post/create", http.HandlerFunc(server.createPost))
	router.Get("/post/:id", http.HandlerFunc(server.getPost))
	router.Del("/post/:id", http.HandlerFunc(server.deletePost))
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", router)
	log.Fatal(err)
}
