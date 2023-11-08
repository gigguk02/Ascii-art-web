package main

import (
	"acsii-art-fs/internal"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ArtWork struct {
	Result string
	Text   string
}

func Error(w http.ResponseWriter, err string, number int) {
	d := struct {
		ErrorCode int
		ErrorText string
	}{
		ErrorCode: number,
		ErrorText: err,
	}
	fileError, err1 := template.ParseFiles("templates/err.html")
	if err1 != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	er := fileError.ExecuteTemplate(w, "err.html", d)
	if er != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

}

func ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		Error(w, internal.BadRequest.Error(), http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	banner := r.FormValue("banner")
	text := r.FormValue("text")
	if len(text) > 1000 {
		Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	a, err := internal.Args1(banner, text)
	if err != nil {
		Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// if a == "" {
	// 	Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	b := ArtWork{a, text}
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	err1 := tmpl.Execute(w, b)
	if err1 != nil {
		Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func home_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		Error(w, "Bad Request", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func handleRequest() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	http.HandleFunc("/", home_page)
	http.HandleFunc("/ascii-art", ascii)
	fmt.Println("Server is running at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequest()
}
