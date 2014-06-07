package main

import (
	"log"
	"strings"
	"strconv"
	"net/http"
	"image"
	"image/jpeg"
	"image/draw"
	"image/color"
)

var (
        imageQuality = jpeg.Options{95}
	blue  color.Color = color.RGBA{0, 0, 255, 255}
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path[1:])
	size := strings.Split(r.URL.Path[1:], "x")
	width, err := strconv.Atoi(size[0])
	if err != nil {
		width = 0
	}
	height, err := strconv.Atoi(size[1])
	if err != nil {
		height = 0
	}
	log.Println(size)
	
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Cache-control", "public, max-age=259200")
	jpeg.Encode(w, m, &imageQuality)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
