package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", loginHandler)
	mux.HandleFunc("/student", student)
	mux.HandleFunc("/save", save)
	mux.HandleFunc("/admin", adminAccountSelector)
	mux.HandleFunc("/admin/schedule", adminHandler)
	mux.HandleFunc("/login", loginHandlerPost)

	fmt.Println("Server is running on http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
