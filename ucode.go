package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

var arg1 = flag.String("e", "", "ucode -e https://sadsecure.blogspot.com")
var arg2 = flag.String("d", "", "ucode -d  https%3A%2F%2Fsadsecure.blogspot.com")

func main() {

	flag.Parse()

	if os.Args[1] == "-e" {
		res := URLEncode(*arg1)
		fmt.Println("URL Encode:", res)
		return
	}

	if os.Args[1] == "-d" {
		res := URLDecode(*arg2)
		fmt.Println("URL Decode:", res)
		return
	}

}

func URLEncode(payload string) string {

	res := url.QueryEscape(payload)
	res = strings.Replace(res, "+", "%20", -1)
	return res
}

func URLDecode(payload string) string {
	res, _ := url.QueryUnescape(payload)

	return res
}
