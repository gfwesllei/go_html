package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Titulo obtem o titulo de uma pagina http
func Titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			rg, _ := regexp.Compile("<title>(.*?)</title>")
			//fmt.Println(rg.FindStringSubmatch(string(html)))
			c <- rg.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}
