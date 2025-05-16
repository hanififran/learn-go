package main

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"
)

func getBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	
	fmt.Fprintf(w, "You've requested the book: %s on the page %s\n", title, page)
}
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", getBook).Methods("GET")

	http.ListenAndServe(":8080",r)
}
