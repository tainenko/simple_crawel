package parser

import (
	"fmt"
	"github.com/gocrawl/engine"
	"github.com/gocrawl/model"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
)

var re=regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseUserInfo(contents []byte,name string) engine.ParseResult{
	match := re.FindSubmatch(contents)
	if len(match)>=2{
		json:=match[1]
		profile:= parseJson(json)
		profile.Name=name
		fmt.Printf("json: %s\n",profile)
	}
	return engine.ParseResult{}

}

func parseJson(json []byte) model.User{
	res,err:=simplejson.NewJson(json)
	if err!=nil{
		log.Print("Failed to parse json.")
	}
	infos, err := res.Get("objectInfo").Get("basicInfo").Array()
	var profile model.User
	for k,v :=range(infos){
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}
	}
	return profile
}