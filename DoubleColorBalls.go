// goGetJpg
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func main() {
	index := 1
	redBall := [6]int{0}
	blueBall := 0
	ballIndex := 0
	x, _ := goquery.NewDocument("http://zx.caipiao.163.com/trend/ssq_basic.html?beginPeriod=2000001&endPeriod=2014002&historyPeriod=2014003&year=")
	db, err := sql.Open("mysql", "root:test@tcp(10.100.0.202:3306)/test")
	fmt.Println(err)
	x.Find("#cpdata tr").Each(func(i int, s *goquery.Selection) {

		openDate, err := s.Find("td").First().Attr("title")
		//fmt.Println(err)
		if err {
			fmt.Println(index)
			openIndexs, _ := s.Find("td").First().Html()
			openIndex, _ := strconv.Atoi(openIndexs)
			openDate = strings.Split(openDate, ":")[1]
			fmt.Println(openIndex, openDate)
			ballIndex = 0
			s.Find("td.chartBall01").Each(func(j int, c *goquery.Selection) {
				ball, _ := c.Html()
				redBall[ballIndex], _ = strconv.Atoi(ball)
				ballIndex++
			})
			s.Find("td.chartBall07").Each(func(j int, c *goquery.Selection) {
				ball, _ := c.Html()
				redBall[ballIndex], _ = strconv.Atoi(ball)
				ballIndex++
			})
			s.Find("td.chartBall02").Each(func(j int, c *goquery.Selection) {
				ball, _ := c.Html()
				blueBall, _ = strconv.Atoi(ball)
			})

			for i := 0; i < len(redBall); i++ {
				for j := i + 1; j < len(redBall); j++ {
					if redBall[i] > redBall[j] {
						redBall[i], redBall[j] = redBall[j], redBall[i]
					}
				}
			}
			fmt.Println(redBall, blueBall)

			stmt, err := db.Prepare("update ball set date=? where id=?")
			fmt.Println(err)
			res, err := stmt.Exec(openDate, openIndex)
			fmt.Println(res, err)
			index++
		}

	})
	fmt.Println("下载完成！")
}
