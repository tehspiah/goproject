package main

import (
	"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/gif"
	"image/png"
	"log"
	"net/http"
	"os"
)

var root = flag.String("root", ".", "file system path")

func main() {

	fmt.Print("Gif creation tool (PNG EDITION): \n")
	var input [50]string
	var num int

	fmt.Print("How many images in the gif? \n")
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Print("Enter image # ", i+1, " name: ")
		fmt.Scanln(&input[i])
	}

	dst := gif.GIF{
		Image: []*image.Paletted{},
	}

	for i := 0; i < num; i++ {
		file, _ := os.Open(input[i])
		j, _ := png.Decode(file)
		defer file.Close()

		r := j.Bounds()

		original := image.NewPaletted(r, palette.WebSafe)
		for x := r.Min.X; x < r.Max.X; x++ {
			for y := r.Min.Y; y < r.Max.Y; y++ {
				original.Set(x, y, j.At(x, y))
			}
		}
		dst.Image = append(dst.Image, original)
	}

	dst.Delay = make([]int, len(dst.Image))
	dst.LoopCount = 100

	var input7 string

	for input7 != "z" {
		f, _ := os.OpenFile("meme.gif", os.O_WRONLY|os.O_CREATE, 0600)
		defer f.Close()
		gif.EncodeAll(f, &dst)
		http.Handle("/", http.FileServer(http.Dir(*root)))

		log.Println("Gif creation successful. \nListening on 1040")
		//Type in "localhost:1040/meme.gif" in browser URL bar
		//Can also type in 1.jpg (etc) to view jpgs used
		// or just "localhost:1040
		err := http.ListenAndServe(":1040", nil)
		if err != nil {
			log.Fatal("ListenAndServe:", err)
		}
		fmt.Println("Press z to exit")
		fmt.Scanln(&input7)
	}

}
