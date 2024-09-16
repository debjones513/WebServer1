/*
ws creates short links, and redirects for accesses..
*/
package main

import (
	"WebServer1/go/src/shortlnk/internal/creator"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// FmtDefaultHandler is the HTML for default web server handler
const FmtDefaultHandler string = `
<div> <p style="color:green;"> &nbsp; &nbsp; <b> Web server is running!</b> </p> </div>
`

// FmtDefaultErrorHandler is the HTML for default web server error handler
const FmtDefaultErrorHandler string = `
<div> <p style="color:red;"> &nbsp; &nbsp; <b> ERROR! </b> </p> </div>
`

// FmtCreateHandler is the HTML for create shortlink handler
const FmtUploadHandler string = `
<div> <p style="color:green;"> &nbsp; &nbsp; <b> Filename </b> </p>   <p style="color:green;"> &nbsp; &nbsp; %s </p> </div>
<div> <p style="color:blue;">  &nbsp; &nbsp; <b> Sha256   </b> </p>   <p style="color:blue;">  &nbsp; &nbsp; %x </p> </div>
<div> <p style="color:black;"> &nbsp; &nbsp; <b> Bytes    </b> </p>   <p style="color:gray;">  &nbsp; &nbsp; %x </p> </div>
`

// failed is an error handling util fn.
func failed(s string, err error) bool {

	if err != nil {
		fmt.Printf("Location: %s Error: %s", s, err)
		return true
	}
	return false
}

// defaultHandler is the default web server handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, FmtDefaultHandler, r.URL.Path[1:])
	fmt.Fprintf(w, FmtDefaultHandler)
}

// uploadHandler is the web server handler for create short link
func createHandler(w http.ResponseWriter, r *http.Request) {

	// On AWS EC2, vtserver runs in <user home>, the test file is in <user home>
	objectName := "./" + r.URL.Path[len("/upload/"):]

	eu, err := loadExeOrURL(objectName)
	if err != nil {
		fmt.Fprintf(w, FmtDefaultErrorHandler)
		return
	}

	fmt.Fprintf(w, FmtUploadHandler, eu.ExeName, eu.ExeSha256, eu.ExeBytes)
}

// createAPIHandler is the API handler for create short link
func createAPIHandler(w http.ResponseWriter, r *http.Request) {

	// Read body

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal to get the binary name

	var resp creator.ExeUpload
	err = json.Unmarshal(b, &resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Load the response

	eu, err := loadExeOrURL(resp.ExeName)
	if err != nil {
		fmt.Fprintf(w, FmtDefaultErrorHandler)
		return
	}

	fmt.Fprintf(w, FmtUploadHandler, eu.ExeName, eu.ExeSha256, eu.ExeBytes)

	/* Marshall the struct into a response
	output, err := json.Marshal(eu)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
	*/
}

func loadExeOrURL(objectName string) (*creator.ExeUpload, error) {

	eu, err := creator.Initialize(objectName)
	if failed("creator.Initialize", err) {
		return eu, err
	}

	//cmscanner.Scan()
	return eu, err
}

func main() {

	// Start the web server

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/create/api/", createAPIHandler)
	//log.Fatal(http.ListenAndServe(":443", nil)) // http://localhost:443/
	log.Fatal(http.ListenAndServe(":8080", nil)) // http://localhost:8080/
}
