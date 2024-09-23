package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Exporter interface {
	Open() error
	Write(*Episode) error
	Close()
}

type CsvExport struct {
	file *os.File
	writer *csv.Writer
}

func (ex *CsvExport) Open() error {
	file, err := os.Create("episodes.csv")
	if err != nil {
		return err
	}
	ex.file = file

	writer := csv.NewWriter(file)

	headers := []string{
		"Title",
		"Season",
		"Episode Number",
		"Rating",
		"Rating count",
	}
	writer.Write(headers)

	ex.writer = writer

	return nil
}

func (ex *CsvExport) Write(episode *Episode) error {
	defer ex.writer.Flush()
	return ex.writer.Write([]string{
		episode.Title,
		strconv.Itoa(episode.Season),
		strconv.Itoa(episode.Number),
		episode.Rating,
		episode.CountRatings,
	})
}

func (ex *CsvExport) Close() {
	ex.file.Close()
}
