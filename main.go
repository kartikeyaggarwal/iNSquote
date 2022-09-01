package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kartikeyaggarwal/printer"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
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

	req.Header.Add("X-RapidAPI-Key", "c0e89c3db7mshab12dcd7a726facp1ef543jsn57bdaa0b6784")
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
	fmt.Println(quotesace)
	texttosppech := quotesace.Content
	speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}
	speech.Speak(texttosppech)
	quotesaceprint := quotesace.Content + "\n \n \n -Vinayak Aggarwal "
	var (
		fontSize          = flag.Float64("fontSize", 60, "font fontSize in points")
		fontPath          = flag.String("fontPath", "assets/FiraSans-Light.ttf", "filename of the ttf font")
		backgroundImgPath = flag.String("bgImg", "assets/back.jpg", "image to use as background")
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

func instagrampost() {

}
