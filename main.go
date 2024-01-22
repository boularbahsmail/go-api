package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var BooksList = []Book{
	{ID: "1", Title: "Tha DaVinci Code", Author: "Dan Brown", Quantity: 10},
	{ID: "2", Title: "Valley Of The Forgotten Wolves", Author: "Amrou Abdelhamid", Quantity: 20},
	{ID: "3", Title: "Land Of Zicola", Author: "Amrou Abdelhamid", Quantity: 15},
}

func checkQuery(context *gin.Context) (*Book, error) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter!!"})
		// return
	}

	book, err := getBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found!!"})
		// return
	}

	return book, err
}

// GET => /books
func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, BooksList)
}

// GET => /books/:id
func bookById(context *gin.Context) {
	// Targeting books by id
	id := context.Param("id")
	book, err := getBookById(id)

	// Return "Book Not Found!!" message if book not found
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found!!"})
		return
	}

	// Returning the indented JSON of the book
	context.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*Book, error) {
	// Looping through the books list and returning the book item from the books list
	for index, book := range BooksList {
		if book.ID == id {
			return &BooksList[index], nil
		}
	}
	return nil, fmt.Errorf("Book Not Found")
}

// POST => books/{}
func createBook(context *gin.Context) {
	var newBook Book

	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	BooksList = append(BooksList, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

// CheckOut a book (Query parameter style)
func checkOutBook(context *gin.Context) {
	book, err := checkQuery(context)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found!!"})
		return
	}

	if book.Quantity <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book Is Not Available!!"})
		return
	}

	book.Quantity -= 1
	context.IndentedJSON(http.StatusOK, book)
}

// Returning a book
func returnBook(context *gin.Context) {
	book, err := checkQuery(context)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found!!"})
		return
	}

	if book.Quantity <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book Is Not Available!!"})
		return
	}

	book.Quantity += 1
	context.IndentedJSON(http.StatusOK, book)
}

func main() {
	// Setup GIN Router
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)

	router.Run("localhost:8080")
}
