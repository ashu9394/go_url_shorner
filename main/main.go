package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	http.ListenAndServe(":3000", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world!")
}
