package main

import (
	"net/http"
	"text/template"
)

func (app *app) getHome(w http.ResponseWriter, r *http.Request) {
	posts, err := app.post.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t, err := template.ParseFiles("./assets/templates/home.page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, map[string]any{"Posts": posts})
}

func (app *app) createPostForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/post.create.page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, nil)
}

func (app *app) createPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = app.post.Insert(
		r.PostForm.Get("title"),
		r.PostForm.Get("content"),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
