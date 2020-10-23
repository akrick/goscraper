package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func ScrapeZwd(url string) map[string]string {
	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})

	params := make(map[string]string)

	c.OnHTML("#J_goodsForm > div.goods-item-title > span", func(e *colly.HTMLElement) {
		params["goods_name"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("#J_goodsForm > div.goods-parameter-outside-container > div.goods-page-server-container > div.goods-page-top > div > span.goods-value", func(e *colly.HTMLElement) {
		params["original_price"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("#goods-pifa-price", func(e *colly.HTMLElement) {
		params["wholesale_price"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("#J_goodsForm > div.goods-parameter-outside-container > div.goods-all-parameter-container > div:nth-child(1) > div.parameter-right > a", func(e *colly.HTMLElement) {
		params["store_goods_no"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("#J_goodsForm > div.goods-parameter-outside-container > div.goods-all-parameter-container > div:nth-child(2) > div.parameter-right > a", func(e *colly.HTMLElement) {
		params["store_onsale_time"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("body > div.web-container > div.item-clear-container.main-function > div.promote-goods-page-container > div.goods-shop-info-wrap > div > div.upper-part > div.new-shop-panel-header > div.left", func(e *colly.HTMLElement) {
		params["store_name"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("body > div.web-container > div.item-clear-container.main-function > div.promote-goods-page-container > div.goods-shop-info-wrap > div > div.second-part > ul > li:nth-child(6) > div.new-shop-panel-item-value.new-shop-panel-item-value-showall", func(e *colly.HTMLElement) {
		params["store_address"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("body > div.web-container > div.item-clear-container.main-function > div.promote-goods-page-container > div.goods-shop-info-wrap > div > div.second-part > ul > li:nth-child(2) > div.new-shop-panel-item-value > a", func(e *colly.HTMLElement) {
		params["store_tel"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("body > div.web-container > div:nth-child(13) > div.detail-left > div.promote-goods-details-container > div.promote-goods-details-right.pull-right > div.details-right-content", func(e *colly.HTMLElement) {
		params["details"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("#sku-color-selector > span", func(e *colly.HTMLElement) {
		colors := ""
		e.ForEach("span", func(i int, el *colly.HTMLElement){
			fmt.Println(el.Attr("title")+"here we go\n")
			if colors == "" {
				colors += el.Attr("title")
			}else{
				colors += "|"+el.Attr("title")
			}
		})
		params["colors"] = colors
	})

	c.OnHTML("#sku-size-selector > ul > li > div:nth-child(1)", func(e *colly.HTMLElement) {
		sizes := ""
		e.ForEach("div", func(i int, el *colly.HTMLElement){
			if sizes == "" {
				sizes += el.Text
			}else{
				sizes += "|"+el.Text
			}
		})
		params["sizes"] = sizes
	})

	c.OnHTML("#sku-size-selector > ul > li > div.sku-item-info.price", func(e *colly.HTMLElement) {
		sizePrices := ""
		e.ForEach("div", func(i int, el *colly.HTMLElement){
			if sizePrices == "" {
				sizePrices += el.Text
			}else{
				sizePrices += "|"+el.Text
			}
		})
		params["sizePrices"] = sizePrices
	})

	c.OnHTML("body > div.web-container > div.item-clear-container.main-function > div.promote-goods-page-container > div.view-BigPicture > div.BigPicture-content > div.left-area > div.imgWrap > a", func(e *colly.HTMLElement) {
		pics := ""
		e.ForEach("div", func(i int, el *colly.HTMLElement){
			if pics == "" {
				pics += el.Attr("href")
			}else{
				pics += "|"+el.Attr("href")
			}
		})
		params["pics"] = pics
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	c.Visit(url)
	c.Wait()
	return params
}

func ScrapeVvic(url string) map[string]string {
	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})
	params := make(map[string]string)

	c.OnHTML("", func(e *colly.HTMLElement) {
		params["goods_name"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		params["original_price"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		params["wholesale_price"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		params["store_goods_no"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		params["store_onsale_time"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		params["store_name"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		params["store_address"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		params["store_tel"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		params["details"] = strings.ReplaceAll(strings.ReplaceAll(e.Text, " ", ""), "\n", "")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		colors := ""
		e.ForEach("span", func(i int, el *colly.HTMLElement){
			fmt.Println(el.Attr("title")+"here we go\n")
			if colors == "" {
				colors += el.Attr("title")
			}else{
				colors += "|"+el.Attr("title")
			}
		})
		params["colors"] = colors
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		sizes := ""
		e.ForEach("div", func(i int, el *colly.HTMLElement){
			if sizes == "" {
				sizes += el.Text
			}else{
				sizes += "|"+el.Text
			}
		})
		params["sizes"] = sizes
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		sizePrices := ""
		e.ForEach("div", func(i int, el *colly.HTMLElement){
			if sizePrices == "" {
				sizePrices += el.Text
			}else{
				sizePrices += "|"+el.Text
			}
		})
		params["sizePrices"] = sizePrices
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
		pics := ""
		e.ForEach("div", func(i int, el *colly.HTMLElement){
			if pics == "" {
				pics += el.Attr("href")
			}else{
				pics += "|"+el.Attr("href")
			}
		})
		params["pics"] = pics
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	c.Visit(url)
	c.Wait()
	return params
}
func main()  {

	url := "https://gz.17zwd.com/item.htm?GID=116697047&spm=46158f5bdb2f1bf4.42.101.0.0.140470.0"
	params := ScrapeZwd(url)
	fmt.Println(params)
	vurl := ""
	data := ScrapeVvic(vurl)
	fmt.Println(data)
}
