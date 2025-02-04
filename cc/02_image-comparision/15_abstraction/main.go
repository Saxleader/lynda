package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"time"
)

// at terminal:
// go run -race main.go

type pixel struct {
	r, g, b, a uint32
}

func main() {
	start := time.Now()

	images, err := getImages()
	if err != nil {
		log.Println("Error getting images", err)
	}

	// range over the [] holding the []pixel - eg, give me each img
	//     range over the []pixel hold the pixels - eg, give me each pixel
	for i, img := range images {
		for j, pixel := range img {
			fmt.Println("Image", i, "\t pixel", j, "\t r g b a:", pixel)
			if j == 10 {
				break
			}
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func getImages() ([][]pixel, error) {

	paths, err := getPaths()
	if err != nil {
		log.Println("Error getting paths", err)
	}

	var images [][]pixel
	for _, path := range paths {
		pixels := getPixels(path)
		images = append(images, pixels)
	}

	return images, nil
}

func getPaths() ([]string, error) {
	const dir = "../00_images/"
	var paths []string

	wf := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		paths = append(paths, path)
		return nil
	}

	if err := filepath.Walk(dir, wf); err != nil {
		return nil, err
	}

	return paths, nil
}

func getPixels(path string) []pixel {
	img := loadImage(path)
	bounds := img.Bounds()
	fmt.Println(bounds.Dx(), " x ", bounds.Dy()) // debugging
	pixels := make([]pixel, bounds.Dx()*bounds.Dy())

	for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()
		r, g, b, a := img.At(x, y).RGBA()
		pixels[i].r = r
		pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a
	}

	return pixels
}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img
}