package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	//router.HandleFunc("/todos/{todoID}",TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Showing list of all Todos")
}

/*
func TodoShow(w http.Req)

*/
