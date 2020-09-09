package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// init variables
	directory := "/home/sysk1ll3r/"
	var bindAddr string

	// Lets ask the user
	fmt.Printf("[*] Enter Bind Address: ")
	fmt.Scanln(&bindAddr)

	// Parsing bind address
	bino := fmt.Sprintf("%s:9000", bindAddr)

	// Start
	fmt.Printf("\n[+] HTTP file server started on %s:9000\n", bindAddr)
	fmt.Printf("[+] Serving: %s\n", string(directory))

	// Handling server
	fileSystem := http.FileServer(http.Dir(directory))
	log.Fatal(http.ListenAndServe(bino, fileSystem))
}
