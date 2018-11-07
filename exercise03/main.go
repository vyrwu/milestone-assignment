package main

import (
	"github.com/gin-gonic/gin"
)

type Book struct {
	Room  int    `json:"room,Number"`
	Row   int    `json:"row,Number"`
	Shelf int    `json:"shelf,Number"`
	ISBN  string `json:"isbn"`
}

var books []Book

func findOne(isdn string) (Book, int, bool) {
	for i, b := range books {
		if b.ISBN == isdn {
			return b, i, true
		}
	}
	return Book{}, -1, false
}

func updateOne(isdn string, book Book) {
	_, i, exists := findOne(isdn)
	if !exists {
		books = append(books, book)
	} else {
		books[i] = book
	}
}

func main() {
	r := gin.Default()
	r.GET("/books", getBooks)
	r.GET("/book/:isbn", getBookByISBN)
	r.POST("/book", postBook)
	r.PUT("/book/:isbn", putBookByISBN)
	r.DELETE("/book/:isbn", deleteBookByISBN)
	r.Run()
}

func getBooks(c *gin.Context) {
	c.JSON(200, books)
}

func getBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	book, _, exists := findOne(isbn)
	if !exists {
		c.JSON(404, gin.H{"message": "not found"})
	} else {
		c.JSON(200, book)
	}
}

func postBook(c *gin.Context) {
	var book Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"message": "bad json body"})
	}
	if _, _, exists := findOne(book.ISBN); exists {
		c.JSON(400, gin.H{"message": "book already registered"})
	} else {
		books = append(books, book)
	}
}

func putBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	var book Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"message": "bad json body"})
	}
	updateOne(isbn, book)
}

func deleteBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	_, i, exists := findOne(isbn)
	if !exists {
		c.JSON(404, gin.H{"message": "not found"})
	} else {
		books = append(books[:i], books[i+1:]...)
	}
}
