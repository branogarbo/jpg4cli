package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"imgcli/util"
	"io"
	"log"
	"os"

	"github.com/anthonynsimon/bild/transform"
)

func main() {
	var (
		file            string
		imgWidth        int
		imgHeight       int
		printWidth      int
		isWebImg        bool
		err             error
		img             io.ReadCloser
		imgData         image.Image
		isPrintSaved    bool
		printSaveTo     string
		isPrintInverted bool
		printMode       string
	)

	// process flags/args

	flag.IntVar(&printWidth, "width", 100, "the number of characters in each row of the printed image")
	flag.BoolVar(&isWebImg, "web", false, "whether the image is in the filesystem or fetched from the web")
	flag.BoolVar(&isPrintSaved, "save", false, "whether or not the the print will be written to a text file")
	flag.BoolVar(&isPrintInverted, "invert", false, "whether or not the the print will be inverted")
	flag.StringVar(&printMode, "mode", "gray", "the mode the image will be printed in. (color, ascii, or gray)")

	flag.Parse()

	if printMode != "gray" && printMode != "ascii" && printMode != "color" {
		fmt.Println("please provide a valid print mode (color, ascii, or gray)")
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		fmt.Println("please provide an image file or address(url) to print")
		os.Exit(1)
	}

	file = flag.Args()[0]

	if isPrintSaved {
		if printMode == "color" {
			fmt.Println("cannot save print in color mode.")
			os.Exit(1)
		} else {
			if len(flag.Args()) == 1 {
				printSaveTo = "./print.txt"
			} else {
				printSaveTo = flag.Args()[1]
			}
		}
	}

	if len(file) < 3 {
		fmt.Println("please provide an image file or address(url) to print")
		os.Exit(1)
	}

	// process image

	if isWebImg {
		img = util.GetImgByUrl(file)
	} else {
		img = util.GetImgByFilePath(file)
	}
	defer img.Close()

	imgData, _, err = image.Decode(img)
	if err != nil {
		log.Fatal(err)
	}

	imgData = transform.Resize(imgData, printWidth, printWidth*imgData.Bounds().Max.Y/imgData.Bounds().Max.X*9/20, transform.Linear)

	imgWidth = imgData.Bounds().Max.X
	imgHeight = imgData.Bounds().Max.Y

	// draw image

	util.DrawPixels(imgData, imgWidth, imgHeight, isPrintSaved, printSaveTo, isPrintInverted, printMode)
}
