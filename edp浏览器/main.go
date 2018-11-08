// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp-master"
) //"github.com/chromedp/chromedp"

// C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe" 可以安放 绿色版chrome

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	err = c.Run(ctxt, click())
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func click() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.baidu.com`),
		//chromedp.WaitVisible(`#footer`),
		//chromedp.Click(`#pkg-overview`, chromedp.NodeVisible),
		chromedp.Sleep(15 * time.Second),
	}
}
