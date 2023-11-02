package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const PORT = 8080

var books = []Book{
	{"The Lord of the Rings", "J. R. R. Tolkien"},
	{"The Little Prince", "Antoine de Saint-Exupéry"},
	{"Harry Potter and the Philosopher's Stone", "J. K. Rowling"},
}

type Book struct {
	Title  string
	Author string
}

func getBooksMap() map[string][]Book {
	m := map[string][]Book{
		"books": books,
	}
	return m
}

func addBook(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	books = append(books, Book{title, author})
	tmpl := fmt.Sprintf("<p>%s %s</p>", title, author)
	t, _ := template.New("t").Parse(tmpl)
	t.Execute(w, nil)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(("static")))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, getBooksMap())
	})
	http.HandleFunc("/add", addBook)

	log.Printf("Server running on port %d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
