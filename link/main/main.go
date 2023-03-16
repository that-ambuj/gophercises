package main

import (
	"bufio"
	"fmt"
	"link"
	"os"
)

var exampleHtml = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
	  <span>some span</span>
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>

`

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
