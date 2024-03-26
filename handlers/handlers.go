package handlers

import (
	logic "ascii-art-web/logic"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)


func GetHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	HandleHtml(w,"index")
}

type Data struct {
    Message string
}

var bigAssString string

func PostHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	// Access form data
	text := r.Form.Get("text")
	style := r.Form.Get("style")
	// Read the request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

    // Print the request body
    fmt.Println("Request Body:", string(body))

	bigAssString = logic.LogicAscii(text, style)


	fmt.Println("this is my bigAssString")
	fmt.Println(bigAssString)

	// this is my message
    data := Data{
        Message: bigAssString,
    }

    // Parse the HTML/EJS template
    tmpl, err := template.ParseFiles("./pages/destination.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Execute the template, passing the data to it
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// HandleHtml(w,"destination")
}

func InternalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// create the file
	filename := "output.txt"

	// Write the string to the file
	err := os.WriteFile(filename, []byte(bigAssString), 0644)
    if err != nil {
        fmt.Println("Error writing to file:", err)
		InternalServerError(w)
        return
    }
	
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		InternalServerError(w)
		return
	}
	defer file.Close()

	// Set the HTTP headers
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")

	// Send the file
	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Println("Error sending file:", err)
		InternalServerError(w)
		return
	}
}



func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

	HandleHtml(w,"404")
	
}

func HandleHtml(w http.ResponseWriter, page string) {
	// Read the HTML file
	htmlFile, err := os.ReadFile("./pages/"+ page + ".html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading HTML file: %s", err), http.StatusInternalServerError)
		return
	}

	// Write the HTML content to the response
	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlFile)
}