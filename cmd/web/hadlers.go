package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request)  {

	if r.URL.Path != "/" {
		http.NotFound(w , r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
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