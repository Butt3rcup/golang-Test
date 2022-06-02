package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("starting")

	r := strings.NewReader("hello world")

	buf := make([]byte, 20)

	r.Read(buf)

	fmt.Printf("buf: %v\n", string(buf))

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}

}
