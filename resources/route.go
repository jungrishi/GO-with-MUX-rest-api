package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jungrishi/rest_api/models"
	"math/rand"
	"net/http"
	"strconv"
)

//GetAuthor author of a book
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range models.Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item.Author.GetFullName())
			return
		}
	}
}

//GetBooks All Book
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Books)
}

//GetBook Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get the params

	//loop through books and find the id
	for _, item := range models.Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})
}

//CreateBook A Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000)) //Mock ID
	models.Books = append(models.Books, book)
	json.NewEncoder(w).Encode(book)
}

//UpdateBooks a Book
func UpdateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.Books {
		if item.ID == params["id"] {
			models.Books = append(models.Books[:index], models.Books[index+1:]...)
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			models.Books = append(models.Books, book)
			return
		}
	}
	json.NewEncoder(w).Encode(models.Books)
}

//DeleteBooks book
func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.Books {
		if item.ID == params["id"] {
			models.Books = append(models.Books[:index], models.Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.Books)
}
