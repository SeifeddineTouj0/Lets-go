package main

import (
	// "fmt"
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello From SnippetBox"))
}

func main() {
    // fmt.Println("Hello world!")
	mux:=http.NewServeMux()
	mux.HandleFunc("/",home)


	log.Print("starting server on :4000")
	err:=http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}
