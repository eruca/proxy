package main

import (
	"flag"
	"log"
	"os/exec"

	"github.com/BurntSushi/toml"
)

type Config struct {
	App       string `toml:"app"`
	ServerUrl string `toml:"server_url"`
}

var (
	configFile = flag.String("f", "config.toml", "配置文件TOML位置")
	pOpenProxy = flag.Bool("s", false, "unset the proxy")
)

func main() {
	flag.Parse()

	var config Config
	if _, err := toml.DecodeFile(*configFile, &config); err != nil {
		log.Fatal("toml.DecodeFile failed: ", err)
	}

	if !*pOpenProxy {
		// Open the application
		cmd := exec.Command("open", "-a", config.App)
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		if err := exec.Command("git", "config", "--global", "http.proxy", config.ServerUrl).Run(); err != nil {
			panic(err)
		}
		if err := exec.Command("git", "config", "--global", "https.proxy", config.ServerUrl).Run(); err != nil {
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
