package main

import (
	"log"
	"net/http"
)

//1.create a handler
//2.create a router for the coresponding path and handler
//3.create a web server to listen the port and router

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the snippet"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the view"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the snippet creation"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("The server is starting ... ")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
