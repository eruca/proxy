package main

import (
	"flag"
	"log"
	"os/exec"
)

const server_port = "sock://127.0.0.1:9981"

var pOpenProxy = flag.Bool("s", false, "unset the proxy")

func main() {
	flag.Parse()

	// Specify the application path
	appPath := "/Applications/Ghelper.app" // Replace this with the path to the application you want to open

	if !*pOpenProxy {

		// Open the application
		cmd := exec.Command("open", "-a", appPath)
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		if err := exec.Command("git", "config", "--global", "http.proxy", server_port).Run(); err != nil {
			panic(err)
		}
		if err := exec.Command("git", "config", "--global", "https.proxy", server_port).Run(); err != nil {
			panic(err)
		}
		log.Println("Set Successed")
		return
	}

	if err := exec.Command("git", "config", "--global", "--unset", "http.proxy").Run(); err != nil {
		panic(err)
	}
	if err := exec.Command("git", "config", "--global", "--unset", "https.proxy").Run(); err != nil {
		panic(err)
	}

	log.Println("Unset Successed")
}
