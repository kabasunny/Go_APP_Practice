package main

import (
	"fmt"
	"net/url"
)

func main() {

	base, _ := url.Parse("http://example.com/aaa")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
}
