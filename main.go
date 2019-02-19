package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	var mandelbrotImage *image.RGBA = buildMandelbrotImage(500, 500, 256, 180, -0.74, -0.139)
	file, err := os.OpenFile("mandelbrot.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	png.Encode(file, mandelbrotImage)
}

func buildMandelbrotImage(width, height, maxIterations int, zoom, moveX, moveY float64) *image.RGBA {
	var mandelbrotImage *image.RGBA = image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var pixelReal float64 = 1.5*(float64(x-width)/2)/(0.5*zoom*float64(width)) + moveX
			var pixelImaginary float64 = (float64(y-height)/2)/(0.5*zoom*float64(height)) + moveY
			var newReal float64 = 0
			var newImaginary float64 = 0
			var i int = 0

			for ; i < maxIterations; i++ {
				var oldReal float64 = newReal
				var oldImaginary float64 = newImaginary

				newReal = (math.Pow(oldReal, 2)) - (math.Pow(oldImaginary, 2)) + pixelReal
				newImaginary = 2*oldReal*oldImaginary + pixelImaginary

				if math.Pow(newReal, 2)+math.Pow(newImaginary, 2) > 4 {
					break
				}
			}

			var hue float64 = 255.0 * float64(i) / float64(maxIterations)
			var saturation float64 = 255.0
			var value float64 = 0.0
			if i < maxIterations {
				value = 255.0
			}

			mandelbrotImage.Set(x, y, hsvToRGB(hue, saturation, value))
		}
	}

	return mandelbrotImage
}

func hsvToRGB(h, s, v float64) color.RGBA {
	var r, g, b float64 = 0, 0, 0

	var i = math.Floor(h * 6)
	var f = h*6 - i
	var p = v * (1 - s)
	var q = v * (1 - f*s)
	var t = v * (1 - (1-f)*s)
	var mod = math.Mod(i, 6)

	switch mod {
	case 0:
		r = v
		g = t
		b = p
		break
	case 1:
		r = q
		g = v
		b = p
		break
	case 2:
		r = p
		g = v
		b = t
		break
	case 3:
		r = p
		g = q
		b = v
		break
	case 4:
		r = t
		g = p
		b = v
		break
	case 5:
		r = v
		g = p
		b = q
		break
	}

	return color.RGBA{
		uint8(r * 255),
		uint8(g * 255),
		uint8(b * 255),
		255,
	}
}
