package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func GetSchedule(url string) {
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

// TODO: 配列や構造体に含めてグルーピングしたい
const MovieWalker = 1
const base_url = "http://movie.walkerplus.com"

type Theater struct {
	Url  string // 収集対象のurl
	Name string // 映画館の名前
}

// 映画館の情報をオブジェクトにして取得
// 1 = movie walker
func GetTheater(page_url string, page_type int) []Theater {
	var theaters []Theater
	switch {
	case page_type == MovieWalker:
		doc, _ := goquery.NewDocument(page_url)
		doc.Find("#rootAreaList > dl").Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Find("dt").Text())
			s.Find("dd > ul > li > a").Each(func(j int, l *goquery.Selection) {
				url, _ := l.Attr("href")
				fmt.Println(url)
				fmt.Println(l.Text())
				theaters = append(theaters,
					Theater{
						Url:  base_url + url,
						Name: "MovieWalker",
					})
			})
		})
	}
	return theaters
}

func main() {
	// url := "http://movie.walkerplus.com/th34/schedule.html"
	url := base_url + "/theater/"
	theaters := GetTheater(url, MovieWalker)
	fmt.Println(theaters)
	for _, t := range theaters {
		fmt.Println(t.Name)
		fmt.Println(t.Url)
	}
}

// 設計メモ
// クローラー対象のサイトごとにURLやルールを外部ファイルとして持っておきDSLで記述しプログラム実行時に読む
// mainのプログラムはサイトからの収集とRDSへの接続を担当する
