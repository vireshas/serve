package main

import (
    "fmt"
    "net/http"
	"os"
)

func main() {
    if len(os.Args) != 2 { 
    	fmt.Println("Usage: ./serve directory-name\nExample: ./serve ~/files")
    } else {
    	root := os.Args[1]
    	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    		file := root + "/" + r.URL.Path[1:]
    		fmt.Println("downloading ", file)
    		http.ServeFile(w, r, file)
    	})
    	panic(http.ListenAndServe(":8000", nil))
    }
}
