package main

import (
	"giflive/ansimage"
	"fmt"
	"image/color"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	VT100_WIDTH  = 80
	VT100_HEIGHT = 24
)

// flags
const DITHERING_MODE = ansimage.NoDithering
const SCALE_MODE = ansimage.ScaleModeFit

var BACKGROUND_COLOUR = color.Black

func main() {
	e := echo.New()

	e.GET("/:GIFNAME", func(c echo.Context) error {
		gifName := c.Param("GIFNAME")

		// set image scale factor for ANSIPixel grid
		sfy, sfx := ansimage.BlockSizeY, ansimage.BlockSizeX // 8x4 --> with dithering
		if DITHERING_MODE == ansimage.NoDithering {
			sfy, sfx = 2, 1 // 2x1 --> without dithering
		}

		var image *ansimage.ANSImage
		var loadErr error
		var filename string = "reimu"

		switch gifName {
		case "reimu":
			filename = "./gifs/reimu.gif"
		case "chirno":
			filename = "./gifs/chirno.gif"
		case "cat":
			filename = "./gifs/cat.gif"
		default:
			return c.String(http.StatusNotFound,
				fmt.Sprintf("GIF image %s not found.\n", gifName))
		}

		image, loadErr = ansimage.NewScaledFromFile(
			filename,
			sfy*VT100_HEIGHT,
			sfx*VT100_WIDTH,
			BACKGROUND_COLOUR,
			SCALE_MODE,
			DITHERING_MODE)

		if loadErr != nil {
			return c.String(http.StatusInternalServerError,
				fmt.Sprintf("GIF image load error: %s.\n", loadErr.Error()))
		}

		// curl animation
		c.Response().Header().Set("Transfer-Encoding", "chunked")
		c.Response().WriteHeader(http.StatusOK)
		w := c.Response().Writer
		cn := w.(http.CloseNotifier)
		flusher := w.(http.Flusher)

		frame := 0
		for {
			select {
			// Handle client disconnect
			case <-cn.CloseNotify():
				log.Println("Client stopped listening")
				return nil
			default:
				// Clear screen
				clearScreen := "\033[2J\033[H"

				fmt.Fprint(w, clearScreen)

				// Print image
				fmt.Fprintln(w, image.RenderExt(frame, false))
				flusher.Flush()

				// GIF delay time
				time.Sleep(time.Millisecond * time.Duration(image.FrameDelay(frame)*10))
			}

			frame++
			if frame >= image.FrameCount() {
				frame = 0
			}
		}
	})

	e.Logger.Fatal(e.Start(":1323"))
}
