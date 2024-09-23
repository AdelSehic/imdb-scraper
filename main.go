package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	if len(os.Args) < 2 {
		panic("Give me a name!")
	}
	id := os.Args[1]

	c := colly.NewCollector(
		colly.AllowedDomains("www.imdb.com"),
	)
	RegisterLogs(c)

	linkfmt := "https://www.imdb.com/title/%s/episodes/?season=%d"
	var episodes []*Episode
	season := 1

	var exp Exporter
	exp = &CsvExport{}

	if err := exp.Open(); err != nil {
		panic(err.Error())
	}
	defer exp.Close()

	c.OnHTML(".episode-item-wrapper", func(h *colly.HTMLElement) {
		episode := &Episode{}
		nameInfo := h.ChildText(".ipc-title__text")
		episode.Rating = h.ChildText(".ipc-rating-star--rating")
		episode.CountRatings = h.ChildText(".ipc-rating-star--voteCount")
		episode.parseEpisodeName(nameInfo)

		episodes = append(episodes, episode)
		if err := exp.Write(episode); err != nil {
			panic(err.Error())
		}
	})

	c.OnHTML("button#next-season-btn", func(h *colly.HTMLElement) {
		fmt.Println("Next season avaliable")
		season++
	})

	c.Visit(fmt.Sprintf(linkfmt, id, season))
	for i := 1; i <= season; i++ {
		c.Visit(fmt.Sprintf("https://www.imdb.com/title/%s/episodes/?season=%d", id, i))
	}
}
