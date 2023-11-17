//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/twpayne/go-jsonstruct/v2"
)

func main() {

	genFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	br := bytes.NewReader(genFile)

	gen := jsonstruct.NewGenerator()
	err = gen.ObserveJSONReader(br)
	if err != nil {
		panic(err)
	}
	out, err := gen.Generate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(out))
}
