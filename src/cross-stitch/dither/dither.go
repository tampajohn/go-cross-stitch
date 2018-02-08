package dither

import (
  "os"
  "image"
  "image/png"
  "image/jpeg"
  "csv"
)

type Dither struct {
  Filter [][] float32
}

func Open(filename string) (image.Image, error) {
  file, err := os.Open(filename)
  if err != nil {
     return nil, err
  }
  defer file.Close()

  image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
  image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)

  img, _, err := image.Decode(file)
  return img, err
}

func Greyscale(img image.Image, outputLoc string) (*image.Gray, error) {
  bounds := img.Bounds()
  greyImg := image.NewGray(bounds)

  for x := bounds.Min.X; x < bounds.Dx(); x++ {
    for y := bounds.Min.Y; y < bounds.Dy(); y++ {
      pix := img.At(x, y)
      greyImg.Set(x, y, pix)
    }
  }

  place, err := os.Create(outputLoc)
  if err != nil {
    return greyImg, err
  }
  defer place.Close()

  err = png.Encode(place, greyImg)

  return greyImg, err
}

func ConvertToDMC(img image.Image, output string) (image.Image, error {
  file, err := os.Open("palette/dmc-floss.csv")
  if err != nil {
     return nil, err
  }
  defer file.Close()

  // convert dmc data to csv hash
  reader := csv.NewReader(file)
  reader.Comma = ','

  bounds := img.Bounds()
  dmcImg = image.NewRGBA(bounds)

  for x := bounds.Min.X; x < bounds.Dx(); x++ {
    for y := bounds.Min.Y; y < bounds.Dy(); y++ {
      // Euclidean distance
    }
  }

  return dmcImg, nil
}
