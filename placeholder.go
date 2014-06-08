package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"io/ioutil"
	"net/http"
	"image"
	"image/jpeg"
	"image/draw"
	"image/color"
	"code.google.com/p/freetype-go/freetype"
)

var (
        imageQuality = jpeg.Options{95}
	blue  color.Color = color.RGBA{170, 170, 170, 1}
	fontSize = 30
	fontSpacing = 1.5
)

func handler(w http.ResponseWriter, r *http.Request) {
	size := strings.Split(r.URL.Path[1:], "x")
	if len(size) < 2 {
		return
	}
	width, err := strconv.Atoi(size[0])
	if err != nil {
		width = 0
	}
	height, err := strconv.Atoi(size[1])
	if err != nil {
		height = 0
	}
	
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	fontBytes, err := ioutil.ReadFile("/usr/share/fonts/TTF/luxisr.ttf")
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(float64(fontSize))
	c.SetClip(m.Bounds())
	c.SetDst(m)
	c.SetSrc(image.White)
	c.SetHinting(freetype.FullHinting)

	text := fmt.Sprintf("%d X %d", width, height)
	x0 := width/2 - len(text) * fontSize/3
	y0 := height/2
	pt := freetype.Pt(x0, y0)
	c.DrawString(text, pt)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Cache-control", "public, max-age=259200")
	jpeg.Encode(w, m, &imageQuality)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
