package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

func main() {
	file, err := os.OpenFile("target.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	// 搜索条件
	// q := "qwer"
	var q string
	fmt.Print("想搜点啥：")
	fmt.Scanf("%s\n", &q)
	fmt.Println("你搜的是：", q)

	// 搜索条数
	// page := 10000
	var page int
	fmt.Print("想搜多少条：")
	fmt.Scanf("%v\n", &page)
	for i := -8; i < page; i++ {
		i = i + 9
		// fmt.Println(i)
		s1 := fmt.Sprintf("https://www.bing.com/search?q=%v&first=%v", q, i)
		resp, err := soup.Get(string(s1))
		if err != nil {
			log.Println(err.Error())
			return
		}
		doc := soup.HTMLParse(resp)
		tnum := doc.Find("div", "id", "b_content").FindAll("span", "class", "sb_count")
		for _, n := range tnum {
			fmt.Println(n.Text())
		}
		links := doc.Find("ol", "id", "b_results").FindAll("a")
		for _, link := range links {
			// fmt.Println(link.Attrs()["href"])
			target := link.Attrs()["href"]
			if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
				break
			} else {
				fmt.Println(link.Attrs()["href"])
				file.WriteString(link.Attrs()["href"] + "\n")
			}

		}
	}
}
