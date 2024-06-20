package main

import (
	"fmt"
	"net/url"
)

func main() {
	// url.Parse関数を使用してURL文字列を解析し、その結果をコンソールに出力
	base, err := url.Parse("http://example.com/")
	fmt.Println(base, err)

	// エラー発生
	base2, err2 := url.Parse("http://e xample.com/")
	fmt.Println(base2, err2)

}
