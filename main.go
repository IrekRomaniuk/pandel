package main

import (
	"io"
	"net/http"
	"strings"
	//"fmt"
)

var feed []string

func main() {
	http.HandleFunc("/", input)
	http.HandleFunc("/feed", output)
	//http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	http.ListenAndServe(":8080", nil)
}

/*func form(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("foo\n"))
}*/

func input(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	feed = append(feed, v)
	//fmt.Println(feed)
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>` + v)
}

func output(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	//w.Write([]byte("bar\n"))
	io.WriteString(w, strings.Join(feed,"\n"))
}

// Go to https://localhost:10443/ or https://127.0.0.1:10443/
// list of TCP ports:
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

// Generate unsigned certificate
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
// for example
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost

// WINDOWS
// windows may have issues with go env GOROOT
// go run %(go env GOROOT)%/src/crypto/tls/generate_cert.go --host=localhost

// instead of go env GOROOT
// you can just use the path to the GO SDK
// wherever it is on your computer