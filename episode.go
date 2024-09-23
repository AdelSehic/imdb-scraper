package main

import (
	"regexp"
	"strconv"
	"fmt"
)

type Episode struct {
	Title        string
	Season       int
	Number       int
	Rating       string
	CountRatings string
}

func (ep *Episode) parseEpisodeName(input string) error {
	// Regular expression to match the pattern: S1.E1 ∙ Title
	re := regexp.MustCompile(`S(\d+)\.E(\d+)\s+∙\s+(.+)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		return fmt.Errorf("input string format is incorrect")
	}

	// Convert season and episode number from string to int
	season, err := strconv.Atoi(matches[1])
	if err != nil {
		return err
	}

	number, err := strconv.Atoi(matches[2])
	if err != nil {
		return err
	}

	// Return the parsed episode struct
	ep.Title = matches[3]
	ep.Number = number
	ep.Season = season
	return nil
}
