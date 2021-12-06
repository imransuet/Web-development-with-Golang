package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// To pipeliene the output into a html file, type this command "go run main.go > index.html", here the ouput will in a file namedindex.html

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + text + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
}
