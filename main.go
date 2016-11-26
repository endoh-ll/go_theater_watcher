package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func GetPage(url string) {
	doc, _ := goquery.NewDocument(url)
	// fmt.Println(doc.Find("div#container > #content > #main > div.mwb > div.movie > div.clearfix > div.movieSchedule").Html())
	doc.Find("div#container > #content > #main > div.mwb > div.movie").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find(".movieTitle > h2 > a").Text())
		fmt.Println(s.Find(".titleIcon").Text())

		// 日付、曜日、時間をそれぞれ日付でまとめる
		for j := 0; j < 7; j++ {
			s.Find(".clearfix > div > table > tbody > tr").Each(func(k int, l *goquery.Selection) {
				var elm *goquery.Selection
				if l.Has("th").Size() > 0 {
					elm = l.Find("th")
				} else {
					elm = l.Find("td")
				}
				fmt.Println(elm.Eq(j).Text())
			})
		}
	})
}

func main() {
	url := "http://movie.walkerplus.com/th34/schedule.html"
	GetPage(url)
}
