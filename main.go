package main

import (
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	mux := http.NewServeMux()
	serv := &http.Server{Handler: jsonContentTypeMiddleware(mux), Addr: ":8080"}
	serv.ListenAndServe()
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("listening to server: 8080")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
