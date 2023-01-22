// @program:     huangli
// @file:        main.go
// @author:      chengkenlee
// @create:      2022-09-25 20:43
// @description:
package main

import (
	"flag"
	"fmt"
	"huangli/run"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func usage() {

	fmt.Printf("\nUsage: %s [-d 日期(%s)] [-h]\n\n黄道吉日查询:\n",
		filepath.Base(os.Args[0]),
		time.Now().Format("2006-01-02"),
	)
	flag.PrintDefaults()
	fmt.Println()
}

func main() {
	var (
		date string
		help bool
	)

	flag.StringVar(&date, "d", "", fmt.Sprintf("日期(%s)", time.Now().Format("2006-01-02")))
	flag.BoolVar(&help, "h", false, "帮助")

	flag.Parse()
	flag.Usage = usage

	if help {
		flag.Usage()
		os.Exit(-1)
	}
	if len(date) != 10 || !strings.Contains(date, "-") {
		flag.Usage()
		os.Exit(-1)
	}

	run.Xpath(date)
}
