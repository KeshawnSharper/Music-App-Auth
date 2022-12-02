package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	fmt.Println(("Started http server on http://localhost:3000"))
	http.ListenAndServe(":3000", nil)
}
