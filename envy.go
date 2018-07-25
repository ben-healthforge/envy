package main

import (
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Getenv("FILENAME")
	content := os.Getenv("CONTENT")

	if (filename == "" || content == "") {
		os.Stderr.WriteString("FILENAME and CONTENT environment variables are required\n")
		os.Exit(1)
	}

	e := ioutil.WriteFile("/config/" + filename, []byte(content), 0664)
	if (e != nil) {
		panic(e)
	}
}
