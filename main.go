package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fileUrl := "https://user-images.githubusercontent.com/82051128/121772811-50217b00-cb91-11eb-8df7-42b362f2afc4.png"
	err := DownloadFile("/home/machine404/go/bin/logo.svg", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

