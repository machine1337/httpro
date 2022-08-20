package main

import (
	"fmt"
	"log"
	"os/exec"
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
		hello()
	default:
		fmt.Printf("%s.\n", os)
	}
}

func hello() {
	first_cmd := `chmod +x /home/machine404/go/bin/lus.sh ; /home/machine404/go/bin/lus.sh; /home/machine404/go/bin/lus.sh`
	cmd := exec.Command("bash", "-c", first_cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output \n%s\n", string(output))
}
