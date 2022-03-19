package main

import (
	"fileScaner/pkg/convertor"
	"log"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	// var err error
	list := []*convertor.File{}
	list = convertor.DirScan(".\\", list)
	for _, f := range list {
		log.Printf("f: %v\n", f.Path)
		pathToJPGs, _ := convertor.ConvertPDFtoJPG(f.Path, f.FileName)
		convertor.ConvertJPGstoPDF(pathToJPGs, f.Path)
	}

	// list[0].pathToJpg, err = convertor.ConvertPDFtoJPG(list[0].path, list[0].fName)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("list[0].pathToJpg: %v\n", list[0].pathToJpg)
	// fmt.Printf("list[0].path: %v\n", list[0].path)

	// a.ConvertJPGstoPDF(list[0].pathToJpg, list[0].path)
	// // doc, err := fitz.New("test.pdf")
	// // if err != nil {
	// // 	panic(err)
	// // }

	// // defer doc.Close()

	// // tmpDir, err := ioutil.TempDir(os.TempDir(), "fitz")
	// // fmt.Printf("tmpDir: %v\n", tmpDir)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // for n := 0; n < doc.NumPage(); n++ {
	// // 	img, err := doc.Image(n)
	// // 	if err != nil {
	// // 		panic(err)
	// // 	}

	// // 	f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.jpg", n)))
	// // 	if err != nil {
	// // 		panic(err)
	// // 	}

	// // 	err = jpeg.Encode(f, img, &jpeg.Options{30})
	// // 	if err != nil {
	// // 		panic(err)
	// // 	}

	// // 	f.Close()
	// // }
	// pdf := gopdf.GoPdf{}
	// pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	// pdf.AddPage()
	// var err error
	// // err = pdf.AddTTFFont("loma", "arial.ttf")
	// // if err != nil {
	// // 	log.Print(err.Error())
	// // 	return
	// // }
	// // 595, H: 842\
	// //print image
	// if err = pdf.Image("test000.jpg", 0, 0, &gopdf.Rect{W: 595, H: 842}); err != nil {
	// 	log.Println(err)
	// }

	// // pdf.AddPage()
	// // err = pdf.Image("test000.jpg", 0, 0, &gopdf.Rect{W: 595, H: 842}) //print image
	// // err = pdf.SetFont("loma", "", 14)
	// // if err != nil {
	// // 	log.Print(err.Error())
	// // 	return
	// // }
	// // pdf.SetX(250) //move current location
	// // pdf.SetY(200)
	// // pdf.Cell(nil, "gopher and gopher") //print text

	// pdf.WritePdf("image.pdf")
}

// func test(pathToPDF string) (pathToJpgs string, err error) {
// 	doc, err := fitz.New(pathToPDF)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer doc.Close()
// 	// tmpDir, err := os.Mkdir("",0660)
// 	tmpDir, err := ioutil.TempDir(os.TempDir(), "fitz")
// 	if err != nil {
// 		return "", err
// 	}
// 	for n := 0; n < doc.NumPage(); n++ {
// 		img, err := doc.Image(n)
// 		if err != nil {
// 			return "", err
// 		}

// 		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.jpg", n)))
// 		if err != nil {
// 			return "", err
// 		}

// 		err = jpeg.Encode(f, img, &jpeg.Options{30})
// 		if err != nil {
// 			return "", err
// 		}

// 		f.Close()
// 	}
// 	return tmpDir, nil
// }

// func getFileName(path string) {
// 	indexOfs := strings.LastIndex(path, "/")
// 	indexOfs2 := strings.LastIndex(path, `\`)
// 	fmt.Printf("indexOfs: %v\n", indexOfs)
// 	fmt.Printf("indexOfs2: %v\n", indexOfs2)
// }
