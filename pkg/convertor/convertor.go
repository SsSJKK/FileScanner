package convertor

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/go-fitz"
	"github.com/signintech/gopdf"
)

type File struct {
	FileName  string
	Path      string
	PathToJpg string
}

func DirScan(dir string, fileList []*File) []*File {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Printf("err: %v\n", err)
		return fileList
	}
	for _, fi := range files {
		if strings.HasSuffix(strings.ToLower(fi.Name()), ".pdf") {
			fileList = append(fileList, &File{
				FileName: fi.Name(),
				Path:     fmt.Sprintf("%s%s", dir, fi.Name()),
			})
		}
		if fi.IsDir() {
			fileList = DirScan(dir+fi.Name()+"\\", fileList)
		}
	}
	return fileList
}

func ConvertPDFtoJPG(pathToPDF string, fileName string) (pathToJpgs string, err error) {
	doc, err := fitz.New(pathToPDF)
	if err != nil {
		return "", err
	}
	defer doc.Close()
	// tmpDir, err := os.Mkdir("",0660)
	os.Mkdir(".\\temp",0660)
	tmpDir, err := ioutil.TempDir(".\\temp\\", fileName)
	if err != nil {
		return "", err
	}
	log.Printf("doc.NumPage(): %v\n", doc.NumPage())
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			return "", err
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("%03d.jpg", n)))
		if err != nil {
			return "", err
		}

		err = jpeg.Encode(f, img, &jpeg.Options{30})
		if err != nil {
			return "", err
		}

		f.Close()
	}
	return tmpDir, nil
}
func ConvertJPGstoPDF(pathToJPGs string, pathToPDF string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	var err error
	files, err := ioutil.ReadDir(pathToJPGs)
	if err != nil {
		log.Printf("err: %v\n", err)
		return err
	}
	for _, fi := range files {
		pdf.AddPage()
		if err = pdf.Image(pathToJPGs+"\\"+fi.Name(), 0, 0, &gopdf.Rect{W: 595, H: 842}); err != nil {
			log.Println(err)
			return err
		}
		log.Println(pathToJPGs + "\\" + fi.Name())
	}
	pdf.WritePdf(pathToPDF)
	pdf.Close()
	return err
}
