package main

import (
<<<<<<< HEAD
=======
	. "fmt"

>>>>>>> 130e420... currrent httprouter work
	"context"
	"net/http"

	httprouter "github.com/julienschmidt/httprouter"
)

type Router struct {
	httprouter.Router
}

func NewRouter() *Router {

	baseRouter := httprouter.New()
<<<<<<< HEAD
=======
	baseRouter.HandleMethodNotAllowed = true
	baseRouter.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Println("Fuck")
		Println("Method:", r.Method)
	})
	// baseRouter.HandleOPTIONS = true
>>>>>>> 130e420... currrent httprouter work

	router := Router{
		*baseRouter,
	}
<<<<<<< HEAD
=======

>>>>>>> 130e420... currrent httprouter work
	router.initializeRoutes()

	return &router
}

func (router *Router) Get(path string, handler http.Handler) {
	router.GET(path, wrapHandler(handler))
}

func (router *Router) Post(path string, handler http.Handler) {
	router.POST(path, wrapHandler(handler))
}

func (router *Router) Put(path string, handler http.Handler) {
	router.PUT(path, wrapHandler(handler))
}

func (router *Router) Delete(path string, handler http.Handler) {
	router.DELETE(path, wrapHandler(handler))
}

func (router *Router) Options(path string, handler http.Handler) {
	router.OPTIONS(path, wrapHandler(handler))
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
<<<<<<< HEAD

		if runEnv == "dev" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
=======
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
>>>>>>> 130e420... currrent httprouter work

		ctx := context.WithValue(r.Context(), "params", ps)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}