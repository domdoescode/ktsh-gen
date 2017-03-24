package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

func OpenTemplate(file string) image.Image {
	body, err := os.Open(file)
	if err != nil {
		log.Fatalf("Could not read %s because of %v", file, err)
	}
	image, _, err := image.Decode(body)
	if err != nil {
		log.Fatalf("Could not decode %s because of %v", file, err)
	}
	return image
}

func main() {
	const fontSize = 52

	text := flag.String("text", "", "the text to use")
	flag.Parse()

	memeText := *text

	if memeText == "" {
		fmt.Printf("No text! %s", string(memeText))
		os.Exit(1)
	}

	memeText = strings.ToUpper(memeText)

	path := fmt.Sprintf("./memes/%d-%d.png", time.Now().Unix(), rand.Intn(1000))
	args := flag.Args()
	if len(args) > 0 {
		path = args[0]
	}

	img := OpenTemplate("joezone_temp.png")
	r := img.Bounds()
	w := r.Dx()
	h := r.Dy()

	m := gg.NewContext(w, h)
	m.DrawImage(img, 0, 0)
	m.LoadFontFace("/Library/Fonts/Impact.ttf", fontSize)

	// Apply black stroke
	m.SetHexColor("#000")
	strokeSize := 6
	for dy := -strokeSize; dy <= strokeSize; dy++ {
		for dx := -strokeSize; dx <= strokeSize; dx++ {
			// give it rounded corners
			if dx*dx+dy*dy >= strokeSize*strokeSize {
				continue
			}
			x := float64(w/2 + dx)
			y := float64(h - fontSize + dy)
			m.DrawStringAnchored(memeText, x, y, 0.5, 0.5)
		}
	}

	// Apply white fill
	m.SetHexColor("#FFF")
	m.DrawStringAnchored(memeText, float64(w)/2, float64(h)-fontSize, 0.5, 0.5)
	m.SavePNG(path)

	fmt.Printf("Saved to %s\n", path)
}
