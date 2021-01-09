package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct

type Book struct {
	ID string 'json:"id"'
	Title string 'json:"title"'
	Author string 'json:"author"'
}

//Get Books

func getBooks(w http.ResponseWriter, r *http.Request)
{

}
//Get All Books

func getBooks(w http.ResponseWriter, r *http.Request)
{
	
}
//Get Single Books

func getBook(w http.ResponseWriter, r *http.Request)
{
	
}
//Create Books

func createBook(w http.ResponseWriter, r *http.Request)
{
	
}
func updateBook(w http.ResponseWriter, r *http.Request)
{
	
}
func deleteBook(w http.ResponseWriter, r *http.Request)
{
	
}

func main()
{
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))

}