package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func webpageText(url string) (string, error) {
	runOption := &playwright.RunOptions{
		SkipInstallBrowsers: true,
	}

	err := playwright.Install(runOption)
	if err != nil {
		return "", err
	}

	pw, err := playwright.Run()
	if err != nil {
		return "", err
	}

	chrome := "chrome"
	option := playwright.BrowserTypeLaunchOptions{
		Channel: &chrome,
	}

	browser, err := pw.Chromium.Launch(option)
	if err != nil {
		return "", err
	}

	page, err := browser.NewPage()
	if err != nil {
		return "", err
	}

	if _, err = page.Goto(url); err != nil {
		return "", err
	}

	body, err := page.QuerySelector("body")
	if err != nil {
		return "", err
	}

	text, err := body.TextContent()
	if err != nil {
		return "", err
	}

	text = strings.Replace(text, "\n", "", -1)
	text = strings.TrimSpace(text)

	if err = browser.Close(); err != nil {
		return "", err
	}

	if err = pw.Stop(); err != nil {
		return "", err
	}

	return text, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Error: URL is not specified")
	}

	url := os.Args[1]
	text, err := webpageText(url)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Success: fetch web page text")
	fmt.Println("============================")
	fmt.Println(text)
}
