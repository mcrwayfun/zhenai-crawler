package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	if resp == nil {
		panic("resp is nil")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",
			resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	printCityList(all)
}

func printCityList(contents []byte) {
	// 一定能编译通过（认为正则一定没问题）
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	// -1匹配所有
	match := re.FindAllSubmatch(contents, -1)
	for _, m := range match {
		fmt.Printf("City: %s, URL: %s\n ", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(match))
}
