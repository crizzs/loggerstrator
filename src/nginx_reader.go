package loggerstratorpkg

import (
	"github.com/chrislusf/glow/flow"
	"fmt"
	//"strings"
	"flag"
	"github.com/dlclark/regexp2"
)

func Nginx_reader() {
	
	flag.Parse()
	//Regex for IP Address
	ipaddress_reg := regexp2.MustCompile(`((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))(?![\\d])`, 0)
	
	tokenizer := func(line string, ch chan string) {
		ipaddress_data,_ := ipaddress_reg.FindStringMatch(line)
		if ipaddress_data != nil {
			ch <- ipaddress_data.Groups()[1].Captures[0].String()
		}
	}

	//Read textfile for information
	f1 := flow.New()
	ipCount := f1.TextFile(
		"./sample/access.log", 3,
	).Map(tokenizer).Map(func(t string) (string, int) {
		return t, 1
	}).Sort(nil).LocalReduceByKey(func(x, y int) int {
		fmt.Println(x)
		return x + y
	});

	ipCount.Map(func(key string, left int) {
		println(key, ":", left)
	}).Run()
}