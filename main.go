package main

import (
	"github.com/fogleman/gg"
	"fmt"
    "image"
    "image/png"
    "os"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	// if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 96); err != nil { подтягивает шрифты 
	// 	panic(err)
	// }
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	dc.SavePNG("out.png")
}




// r, g, b, a := img.At(x, y).RGBA() координаты пикселя