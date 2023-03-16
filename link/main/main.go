package main

import (
	"bufio"
	"fmt"
	"link"
	"os"
)

func main() {
	file, err := os.OpenFile("../ex3.html", os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(file)
	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Println("Link:", link.Href)
		fmt.Println("Text:", link.Text)
	}
}
