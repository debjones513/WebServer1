package main

/*
Vtserver runs a virus scan on a binary.
It stores the binary in S3, and uses third party scanners.
*/

import (
	"fmt"
	"log"
	"net/http"
)

// FmtDefaultHandler is the HTML for default web server handler
const FmtDefaultHandler string = `
<div> <p style="color:green;"> &nbsp; &nbsp; <b> Web server is running!</b> </p> </div>
<div> <p style="color:green;"> &nbsp; &nbsp; To test on AWS EC2: http://44.234.131.118:8080/upload/vtserver_test </p> </div>
`

// FmtDefaultErrorHandler is the HTML for default web server error handler
const FmtDefaultErrorHandler string = `
<div> <p style="color:red;"> &nbsp; &nbsp; <b> ERROR! </b> </p> </div>
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

// testHandler is the web server test handler
func testHandler(w http.ResponseWriter, r *http.Request) {

	// On AWS EC2, server runs in <user home>, the test file is in <user home>
	//objectName := "./" + r.URL.Path[len("/test/"):]

}

// uploadAPIHandler is the API upload handler
func testAPIHandler(w http.ResponseWriter, r *http.Request) {

	// Read body
	/*
				b, err := io.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			// Unmarshal to get the binary name

		var resp uploader.ExeUpload
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


			// Marshall the struct into a response
			output, err := json.Marshal(eu)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("content-type", "application/json")
			w.Write(output)
	*/
}

func main() {

	// Start the web server

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/test/", testHandler)
	http.HandleFunc("/test/api/", testAPIHandler)
	//log.Fatal(http.ListenAndServe(":443", nil)) // http://localhost:443/
	log.Fatal(http.ListenAndServe(":8080", nil)) // http://localhost:8080/
}
