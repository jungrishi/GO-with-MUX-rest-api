package models

import "fmt"

// type manipulate interface {
// 	getFullName() string
// }

//Book Struct(Model)
type Book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Books Init books var as a Slice Book struct
var Books []Book

//Author struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (a Author) getFullName() string {
	return a.Firstname + "_" + a.Lastname
}

func init() {
	fmt.Println("hello from MOdels")
	//Mock Data - @todo - imp DB
	Books = append(Books, Book{ID: "1", Name: "The Witcher", Title: "Ch-1", Author: &Author{Firstname: " Andrzej", Lastname: "Sapkowski"}})
	Books = append(Books, Book{ID: "2", Name: "LOR", Title: "Ch-1", Author: &Author{Firstname: " J.R.R", Lastname: "Tolkiens"}})
}
