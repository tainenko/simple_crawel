package main

import (
	"github.com/gocrawl/engine"
	"github.com/gocrawl/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})

}
