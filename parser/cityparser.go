package parser

import (
	"github.com/gocrawl/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, c := range all {
		name:=string(c[2])
		result.Items = append(result.Items, "User:"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseUserInfo(c, name)
			},
		})
	}
	return result
}
