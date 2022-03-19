package convertor

import (
	"fmt"
	"testing"
)

func Test_DirScan(t *testing.T) {

	// getFileName("..\\..\\test.pdf759475763\\", "test.pdf")
	err := ConvertJPGstoPDF(".\\test2.pdf876329263", "12.pdf")
	// _, err := a.ConvertPDFtoJPG("1.pdf", "test2.pdf")
	fmt.Printf("err: %v\n", err)
}
