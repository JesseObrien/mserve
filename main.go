package main

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/russross/blackfriday"
)

type Page struct {
	Title string
	Body string
}

// Load a markdown file's contents into memory
func loadFile(title string) (*Page, error) {
	filename := title + ".md"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// Convert the body into markdown
	mdbody := string(blackfriday.MarkdownBasic(body)[:])
	return &Page{Title: title, Body: mdbody}, nil
}

// Load and display any markdown file
func filehandler(writer http.ResponseWriter, request *http.Request) {
	
}

// Any markdown extension we can think of
var extensions = []string{".md", ".markdown", ".mdown", ".mkdn", ".mkd", ".mdtext", ".mdtxt", ".txt", ".text"}

// Call the index file, this is a required
func index(writer http.ResponseWriter, request *http.Request) {
	page, _ := loadFile("index")
	files, _ := ioutil.ReadDir("./")

	var directoryList string
	for _, f := range files {
		for _, ext := range extensions {
			// If the file name has a valid markdown extension, add it to the list
			if strings.Contains(f.Name(), ext) {
				directoryList += "<li><a href=\"" + f.Name() + "\">" + f.Name() + "</a></li>"
			}
		}
	}
	
	// Print the body output to the response
	fmt.Fprintf(writer, page.Body + directoryList)
}

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8008", nil)
}
