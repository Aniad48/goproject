package LetterDrawer

import (
	"github.com/fogleman/gg"
	"fmt"
	"image"
	"image/png"
	"os"
)

func drawLetters(letters string, font string)([]Word, ){ //добавить возврат изображения
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("/Library/Fonts/"+font+".ttf", 96); err != nil { //подтягивает шрифты
	 	panic(err)
	}
	dc.DrawStringAnchored(letters, S/2, S/2, 0.5, 0.5)
	dc.SavePNG("out.png")

}

func findLetter(lettersName string, startFindFrom int) Letter  {
	newLetter := Letter{0, 0, 0, 0, lettersName}



	return newLetter
}




// r, g, b, a := img.At(x, y).RGBA() координаты пикселя