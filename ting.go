package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"./color"
)

var version string = "1.0.0"

// 监听端口
var port int

var help bool

func PrintCopyRight() {
	fmt.Printf(color.Yellow(fmt.Sprintf(`
=======================================
Name   : Ting Static Server
Version: %s
Author : http://junqiang.wang/
=======================================
`, version)))
}

func Init() bool {

	flag.BoolVar(&help, "h", false, "show help info")
	flag.IntVar(&port, "p", 80, "server port, default 80")
	flag.Parse()

	if help {
		PrintCopyRight()
		fmt.Println("Usage: ting -p=12345")
		flag.PrintDefaults()
		fmt.Println("")
		return false
	} else {
		return true
	}
}

func main() {
	if !Init() {
		return
	}
	PrintCopyRight()
	file, _ := exec.LookPath(os.Args[0])
	current_path, _ := filepath.Abs(file)
	path := "./"

	fmt.Println("Server Starting...")
	fmt.Println(color.Yellow(fmt.Sprintf("http://127.0.0.1:%d", port)))
	fmt.Println(color.Cyan(fmt.Sprintf("Path\t%s\nPort\t%d\n", current_path, port)))

	http.Handle("/", http.FileServer(http.Dir(path)))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
