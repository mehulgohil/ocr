package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"io"
	"net/http"
	"os"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	fatherName := r.FormValue("fatherName")
	defer file.Close()

	tmpfile, err := os.Create("./" + handler.Filename)
	defer tmpfile.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("./" + handler.Filename)
	text, _ := client.Text()

	if strings.Contains(strings.ToLower(text), fatherName) {
		fmt.Fprintf(w, "true")
		return
	}
	fmt.Fprintf(w, "false")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	setupRoutes()
}