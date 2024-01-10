package main

import (
	"log"
	"os"
	"path/filepath"
	"plugin"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	p, err := plugin.Open(exPath + "/plug.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		log.Fatalln("No args[1]")
	}
	*v.(*string) = os.Args[1]
	a, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	f.(func(A string))(a)
}
