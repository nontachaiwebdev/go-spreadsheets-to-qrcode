package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
	"./spreadsheets"
)

func main() {
	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1Qf3XKIoD-25utbfaBuZ3M-YU35GAdCv3lWVwSZYlnGA"
	readRange := "userEvent"
	srv, err := spreadsheets.GetSheets()
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			var buf bytes.Buffer
			name := fmt.Sprintf("%s", row[0])
			buf.WriteString("https://choop.me?name=")
			buf.WriteString(name)
			qrImageByte, _ := qrcode.Encode(buf.String(), qrcode.Medium, 256)
			buf.Reset()
			fmt.Println(row[0])
			buf.WriteString("./qrcode/")
			buf.WriteString(name)
			buf.WriteString(".png")
			fmt.Println(buf.String())
			outputFile, _ := os.Create(buf.String())
			img, _, _ := image.Decode(bytes.NewReader(qrImageByte))
			png.Encode(outputFile, img)
			outputFile.Close()
		}
	} else {
		fmt.Print("No data found.")
	}

}
