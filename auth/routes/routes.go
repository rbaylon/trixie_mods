package authroutes

import (
	"fmt"
	"net/http"
)

func GetRouter() *http.ServeMux {
	Router := http.NewServeMux() // 3

	Router.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Customer Sign in")
	})

	Router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Get all users")
	})

	Router.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Get user with id: %d", r.PathValue("id"))
	})

	return Router
}
