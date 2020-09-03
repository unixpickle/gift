package gift

import (
	"image"
	"image/gif"
	"io"

	"github.com/unixpickle/smallpng/smallpng"
)

// EncodeAll writes an animated GIF to w.
//
// The images can have any color scheme, and will be
// quantized automatically.
//
// For each image, there should be a delay, in 100ths of a
// second.
//
// If loopCount is 0, the GIF will loop forever.
// If it is -1, it will be shown exactly once.
func EncodeAll(w io.Writer, images []image.Image, delays []int, loopCount int) error {
	if len(images) != len(delays) {
		panic("number of images must exactly match number of delays")
	}
	g := &gif.GIF{
		Delay:     delays,
		LoopCount: loopCount,
		Disposal:  make([]byte, len(delays)),
	}
	for i := range g.Disposal {
		g.Disposal[i] = gif.DisposalBackground
	}
	for _, img := range images {
		quantized := smallpng.PaletteImage(img, nil)
		g.Image = append(g.Image, quantized)
	}
	return gif.EncodeAll(w, g)
}
