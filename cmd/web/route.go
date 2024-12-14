package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.getHome)
	mux.HandleFunc("GET /posts/create", app.createPostForm)
	mux.HandleFunc("POST /posts/create", app.createPost)
	return mux
}
