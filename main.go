package main

import (
	v "Go-make/version"
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

/**
* @Author: super
* @Date: 2020-08-29 16:42
* @Description: 支持显示版本信息
**/
var (
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	fmt.Println("hello world")
	pflag.Parse()
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}
}
