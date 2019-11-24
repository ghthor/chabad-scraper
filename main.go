package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const hebrew27 = "אבגדהוזחטיכלמנסעפצקרשתךםןףץ"

func ExampleScrape() {
	// Load HTML
	f, err := os.Open("./data_src/0_genesis/1.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	var chapter bytes.Buffer

	// Find the review items
	doc.Find("tr.Co_Verse td.hebrew").Each(func(i int, s *goquery.Selection) {
		verseText := s.Find("span.co_VerseText").Text()

		for _, r := range verseText {
			if strings.ContainsAny(string(r), hebrew27) {
				chapter.WriteRune(r)
			}
		}

		chapter.WriteString("\n")
	})

	// fmt.Print(chapter.String())

	scanner := bufio.NewScanner(bytes.NewReader(chapter.Bytes()))

	maxLen := 0
	for scanner.Scan() {
		length := len([]rune(scanner.Text()))
		if length > maxLen {
			maxLen = length
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(bytes.NewReader(chapter.Bytes()))
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Printf(fmt.Sprint("%3d%", maxLen, "s\n"), len([]rune(txt)), txt)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	ExampleScrape()
}
