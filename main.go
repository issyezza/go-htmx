package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	// Hello world, the web server
	fmt.Println("hello  world")

	h1 := func(w http.ResponseWriter, req *http.Request) {
		tmp := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "foo", Director: "bar"},
				{Title: "foo", Director: "bar"},
				{Title: "foo", Director: "bar"},
			},
		}
		tmp.Execute(w, films)
	}
	addFilm := func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(1 * time.Second)
		title := req.PostFormValue("title")
		director := req.PostFormValue("director")
		htmlString := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s -%s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlString)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", addFilm)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
