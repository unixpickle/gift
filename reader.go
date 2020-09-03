package gift

import (
	"image"
	"image/color"
	"image/gif"
)

// Bounds computes the bounds for the GIF image.
//
// This is not necessarily the bounds of the first frame.
func Bounds(g *gif.GIF) image.Rectangle {
	if g.Config.Width == 0 && g.Config.Height == 0 {
		result := g.Image[0].Bounds()
		for _, frame := range g.Image {
			result = result.Union(frame.Bounds())
		}
		return result
	}
	return image.Rect(0, 0, g.Config.Width, g.Config.Height)
}

// Frames computes the disposal-aware frames from g.
func Frames(g *gif.GIF) []image.Image {
	var results []image.Image
	out := image.NewRGBA(Bounds(g))
	previous := image.NewRGBA(Bounds(g))
	for i, frame := range g.Image {
		disposal := g.Disposal[i]
		switch disposal {
		case 0:
			clearImage(out)
			clearImage(previous)
			drawImageWithBackground(out, frame, out)
			drawImageWithBackground(previous, frame, previous)
		case gif.DisposalNone:
			drawImageWithBackground(out, frame, out)
		case gif.DisposalPrevious:
			drawImageWithBackground(out, frame, previous)
		case gif.DisposalBackground:
			bgColor := g.Config.ColorModel.(color.Palette)[g.BackgroundIndex]
			fillImage(out, bgColor)
			drawImageWithBackground(out, frame, out)
		}
		frameCopy := image.NewRGBA(out.Bounds())
		copy(frameCopy.Pix, out.Pix)
		results = append(results, frameCopy)
	}
	return results
}

func drawImageWithBackground(dst *image.RGBA, src, bg image.Image) {
	b := src.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			px := src.At(x, y)
			_, _, _, a := px.RGBA()
			if a == 0 {
				px = bg.At(x, y)
			}
			dst.Set(x, y, px)
		}
	}
}

func clearImage(dst *image.RGBA) {
	for i := range dst.Pix {
		dst.Pix[i] = 0
	}
}

func fillImage(dst *image.RGBA, c color.Color) {
	b := dst.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			dst.Set(x, y, c)
		}
	}
}
