package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	_, methodExist := r.rules[method]
	handler, exist := r.rules[method][path]
	return handler, methodExist, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := r.FindHandler(request.Method, request.URL.Path)

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, request)
}
