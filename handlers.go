package main

import (
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"strconv"
)


func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (p *postServer) createPost(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling post create at %s\n", r.URL.Path)

	
	type RequestPost struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}


	// Enforce JSON 
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var ps RequestPost
	if err := dec.Decode(&ps); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post := p.repo.CreatePost(ps.Title, ps.Content)
	renderJSON(w, post)

}

func (p *postServer) deletePost(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling delete post at %s\n", r.URL.Path)

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	err = p.repo.DeletePost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (p *postServer) getPost(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get post at %s\n", r.URL.Path)

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	post, err := p.repo.GetPost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(w, post)
}
