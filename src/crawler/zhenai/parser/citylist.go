package parser

import (
	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) {
	// 一定能编译通过（认为正则一定没问题）
	re := regexp.MustCompile(cityListRe)
	// -1匹配所有
	match := re.FindAllSubmatch(contents, -1)
	for _, m := range match {
		fmt.Printf("City: %s, URL: %s\n ", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(match))
}
