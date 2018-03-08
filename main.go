package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"os/exec"
)

func main() {
	root := "."
	if len(os.Args) == 1 {
		root = getWd()
	} else if len(os.Args) > 2 || os.Args[1] == "-h" || os.Args[1] == "help" || os.Args[1] == "--help" {
		fmt.Println("Usage: ./serve directory-name\nExample: ./serve ~/files")
		return
	} else {
		root = os.Args[1]
	}

	root = trimLeadSlash(root)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file := root + "/" + r.URL.Path[1:]
		fmt.Println("downloading ", file)
		http.ServeFile(w, r, file)
	})
	fmt.Println("👂  at " + getIP() + ":8000 " + "& serving " + root)
	panic(http.ListenAndServe(":8000", nil))
}

func getWd() string{
    pwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return pwd
}

func trimLeadSlash(str string) string {
	if strings.HasSuffix(str, "/") {
	   return str[:len(str)-1]
	}
	return str
}

func getIP() string{
	out, err := exec.Command("bash", "-c", "ifconfig | pcregrep -M -o '^[^\\t:]+:([^\\n]|\\n\\t)*status: active' | egrep -o -m 1 'inet\\s*(.*)\\s*netmask' | cut -d' ' -f2 | tr -d '[:space:]'").Output()
	if err != nil {
		fmt.Printf("Error", err)
	}
	return string(out)
}
