// @program:     huangli
// @file:        xpath.go
// @author:      chengkenlee
// @create:      2022-09-25 20:50
// @description:
package run

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"net/http"
)

func Xpath(date string) {

	y, d := newReplaceVar(date)

	url := fmt.Sprintf("https://rili.ximizi.com/wannianli/%s/%s.html", y, d)

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	doc, _ := htmlquery.Parse(resp.Body)
	list := htmlquery.Find(doc, "//div[1]/div[6]/div[2]")

	fmt.Println(fmt.Sprintf("\n黄道吉日"))
	for _, item := range list {
		fmt.Println(fmt.Sprintf("%s", htmlquery.InnerText(item)))
	}
}
