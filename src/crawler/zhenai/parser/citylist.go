package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	// 一定能编译通过（认为正则一定没问题）
	re := regexp.MustCompile(cityListRe)
	// -1匹配所有
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		// m[2]是City Name
		result.Items = append(result.Items, string(m[2]))
		// m[1]是City Url
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
