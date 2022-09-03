package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kartikeyaggarwal/printer"
)

type quotes struct {
	Content string
}

//
func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	url := "https://quotes15.p.rapidapi.com/quotes/random/?language_code=en"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "YOU API KEY HERE") // REPLACE API KEY
	req.Header.Add("X-RapidAPI-Host", "quotes15.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var quotesace quotes
	fmt.Println(quotesace)
	err := json.Unmarshal([]byte(body), &quotesace)
	if err != nil {
		fmt.Println(err)
	}
	quotesaceprint := quotesace.Content 
	var (
		fontSize          = flag.Float64("fontSize", 60, "font fontSize in points")
		fontPath          = flag.String("fontPath", "assets/Antonio-Bold.ttf", "filename of the ttf font")
		backgroundImgPath = flag.String("bgImg", "assets/image.png", "image to use as background")
		text              = flag.String("text", quotesaceprint, "text to print on the image")
		outputPath        = flag.String("output", "cool_img.png", "output path for the resulting image")
	)
	flag.Parse()
	img, err := printer.TextOnImg(
		printer.Request{
			BgImgPath: *backgroundImgPath,
			FontPath:  *fontPath,
			FontSize:  *fontSize,
			Text:      *text,
		},
	)
	if err != nil {
		return err
	}

	if err := printer.Save(img, *outputPath); err != nil {
		return err
	}

	log.Println("image saved on [", *outputPath, "]")
	return nil
}

