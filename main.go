package main

import (
	"flag"
	"bufio"
	"os"

    "github.com/viert/go-lame"
)

func main() {
	var inph string
	var outph string

	flag.StringVar(&inph , "i", "", "Input file")
	flag.StringVar(&outph , "o", "", "Output file")
	flag.Parse()

	of, err := os.Create(outph)
	if err != nil {
		panic(err)
	}
	defer of.Close()
	enc := lame.NewEncoder(of)
	defer enc.Close()

	inf, err := os.Open(inph)
	if err != nil {
		panic(err)
	}
	defer inf.Close()

	r := bufio.NewReader(inf)
	r.WriteTo(enc)
}