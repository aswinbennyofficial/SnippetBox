package main

import (
	"fmt"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprintf(w,"Display the home page ")
}

func snippetView(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprintf(w,"Display a specific snippet ")
}
func snippetCreate(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprintf(w,"Create a new snippet")
}


func main(){
	http.HandleFunc("/",home)
	http.HandleFunc("/snippet/view",snippetView)
	http.HandleFunc("/snippet/create",snippetCreate)

	http.ListenAndServe(":8080",nil)


}