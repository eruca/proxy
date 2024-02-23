package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	App       string `toml:"app"`
	ServerUrl string `toml:"server_url"`
}

var (
	configFile = flag.String("f", "proxy.toml", "配置文件TOML位置")
	pOpenProxy = flag.Bool("s", false, "unset the proxy")
)

func findPath(app string) (string, error) {
	paths := os.Getenv("PATH")

	for _, p := range strings.Split(paths, path_sep) {
		entris, err := os.ReadDir(p)
		if err != nil {
			continue
		}
		for _, ent := range entris {
			if ent.IsDir() {
				continue
			}
			if ent.Name() == app {
				return p, nil
			}
		}
	}
	return "", fmt.Errorf("app not found: %s", app)
}

func main() {
	flag.Parse()

	appName := os.Args[0]
	p, err := findPath(appName)
	if err != nil {
		// 如果没有找到，就是在local, Getwd
		p, _ = os.Getwd()
	}

	var config Config
	if _, err := toml.DecodeFile(filepath.Join(p, *configFile), &config); err != nil {
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
