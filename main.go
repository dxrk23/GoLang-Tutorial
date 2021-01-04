package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request)  {

	if r.URL.Path != "/" {
		http.NotFound(w , r)
		return
	}

	w.Write([]byte("Hello from snippet box!"))
}

func showSnippet(w http.ResponseWriter, r *http.Request){
	id , err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specififc snippet with ID %d..." , id)
}

func createSnippet(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w , "Method not Allowed" , 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("snipppet/create" , createSnippet)

	log.Println("Starting server on :4000")
	err:= http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}