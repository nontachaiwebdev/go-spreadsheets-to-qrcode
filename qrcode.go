package main

import qrcode "github.com/skip2/go-qrcode"

import (
  //"fmt"
  "bytes"
  "image"
  "image/png"
  "os"
)

func main() {
  //var qrImageByte []byte
  qrImageByte, _ := qrcode.Encode("https://google.co.th", qrcode.Medium, 256)
  outputFile, _ := os.Create("./test.png")
  img, _, _ := image.Decode(bytes.NewReader(qrImageByte))
  png.Encode(outputFile, img)
  outputFile.Close()
}
