// Mandelbrot emits a PNG image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img) // Note: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		c := cmplx.Abs(v)
		if c > 2 {
			// return color.Gray{255 - contrast - n}
			return color.RGBA{
				(n * 5) % 255,
				(n * 10) % 255,
				(n * 15) % 255,
				255,
			}
		}
	}
	return color.Black
}
