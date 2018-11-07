package main

import (
	"fmt"
	"github.com/ryanuber/go-glob"
	"io/ioutil"
	"strings"
)

type Book struct {
	Authors         []string
	Title           string
	Publisher       string
	PublicationYear string
}

func ReadBooks(input string) []Book {
	var books []Book
	subs := strings.Split(input, "Book:\n")
	subs = append(subs[:0], subs[1:]...)
	for _, s := range subs {
		var book Book
		s = strings.TrimRight(s, "\n")
		cuts := strings.Split(s, "\n")
		for _, c := range cuts {
			els := strings.Split(c, ": ")
			switch els[0] {
			case "Author":
				book.Authors = append(book.Authors, els[1])
			case "Title":
				book.Title = els[1]
			case "Publisher":
				book.Publisher = els[1]
			case "Published":
				book.PublicationYear = els[1]
			default:
				panic("Bad input")
			}
		}
		books = append(books, book)
	}
	return books
}

func FindBooks(searchString string, books []Book) []Book {
	var selection []Book
	for _, b := range books {
		hit := true
		cursors := strings.Split(searchString, " & ")
		for _, c := range cursors {
			matched := setContains(c,
				strings.Join(b.Authors, ""),
				b.Title,
				b.Publisher,
				b.PublicationYear)
			hit = hit && matched
		}
		if hit {
			selection = append(selection, b)
		}
	}
	return selection
}

func setContains(searchString string, set ...string) bool {
	for _, s := range set {
		var matched bool
		if strings.ContainsRune(searchString, '*') {
			matched = glob.Glob(searchString, s)
		} else {
			matched = searchString == s
		}
		if matched {
			return true
		}
	}
	return false
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	s := string(bytes)
	books := ReadBooks(s)
	searchResult := FindBooks("*", books)
	fmt.Println(searchResult)
}
