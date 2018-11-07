package main

import (
	"encoding/json"
	"github.com/emicklei/forest"
	"net/http"
	"strconv"
	"testing"
)

var test = forest.NewClient(
	"http://localhost:8080",
	http.DefaultClient)

func TestOperations(t *testing.T) {
	for i := 0; i < 1; i++ {
		itoa := strconv.Itoa(i)
		book := &Book{
			Room:  3,
			Row:   3,
			Shelf: 3,
			ISBN:  itoa,
		}
		jbody, _ := json.Marshal(book)
		r := test.POST(t, forest.Path("/book").Body(string(jbody)))
		forest.ExpectStatus(t, r, 200)

		r = test.GET(t, forest.Path("/books"))
		forest.ExpectStatus(t, r, 200)
		forest.ExpectJSONArray(t, r, func(array []interface{}) {
			if len(array) != 1 {
				t.Error("Bad JSON response")
			}
		})

		book = &Book{
			Room:  4,
			Row:   4,
			Shelf: 4,
			ISBN:  itoa,
		}
		r = test.PUT(t, forest.Path("/book/" + itoa).Body(string(jbody)))
		forest.ExpectStatus(t, r, 200)

		r = test.GET(t, forest.Path("/book/"+itoa))
		forest.ExpectStatus(t, r, 200)
		forest.ExpectJSONDocument(t, r, &Book{4, 4, 4, itoa})

		r = test.DELETE(t, forest.Path("/book/"+itoa))
		forest.ExpectStatus(t, r, 200)

		r = test.GET(t, forest.Path("/books"))
		forest.ExpectStatus(t, r, 200)
		forest.ExpectJSONArray(t, r, func(array []interface{}) {
			if len(array) != 0 {
				t.Error("Bad JSON response")
			}
		})
	}
}
