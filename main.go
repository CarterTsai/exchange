package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"gopkg.in/toast.v1"
)

func transformWord(message string) string {
	big5Word, _, _ := transform.String(traditionalchinese.Big5.NewEncoder(), message)
	return big5Word
}

func cathaybk() {
	doc, err := goquery.NewDocument("https://www.cathaybk.com.tw/cathaybk/personal/exchange/product/currency-billboard/")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	table := "#layout_0_rightcontent_1_firsttab01_0_tab_rate_realtime tbody > tr"
	doc.Find(table).Each(func(i int, s *goquery.Selection) {
		g := s.Find("font")

		if strings.Contains(g.Eq(0).Text(), "日圓(JPY)") {
			price := fmt.Sprintf("賣出價 %s \n", g.Eq(2).Text())
			Title := transformWord("日圓(JPY)")
			Message := transformWord(price)
			fmt.Println(price)

			notification := toast.Notification{
				AppID:   "Example App",
				Title:   Title,
				Message: Message,
				Icon:    "", // This file must exist (remove this line if it doesn't)
				Actions: []toast.Action{},
			}
			err := notification.Push()
			if err != nil {
				log.Fatalln(err)
			}
		}
	})
}

func main() {
	cathaybk()
}
