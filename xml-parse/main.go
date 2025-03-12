package main

import (
	"embed"
	"fmt"
	"log"
	"xml-parse/xml"

	"github.com/beevik/etree"
)

//go:embed *.xml
var xmlFs embed.FS

func main() {
	//Parse("35112d6d43c823000142dm1.xml")
	//Parse("6. Literature with attachements (EDQM_revision).xml")
	xml.Generate()
}

// Parse
//
//	@author dongdong.zhang
//	@param filename
//	@return {}
func Parse(filename string) {
	content, err := xmlFs.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}

	doc := etree.NewDocument()
	err = doc.ReadFromBytes(content)
	if err != nil {
		log.Panic(err)
	}

	root := doc.SelectElement("MCCI_IN200100UV01")

	n11Node := root.SelectElement("name")
	n11 := n11Node.SelectAttrValue("code", "")
	fmt.Println(n11)
	batchNumberNode := root.SelectElement("id")
	batchNumber := batchNumberNode.SelectAttrValue("extension", "")
	fmt.Println(batchNumber)

	for _, v := range root.SelectElements("PORR_IN049016UV") {
		identifierNode := v.SelectElement("id")
		identifier := identifierNode.SelectAttrValue("extension", "") //N.2.r.1
		fmt.Println(identifier)

		creationTimeNode := v.SelectElement("creationTime")
		creationTime := creationTimeNode.SelectAttrValue("value", "")
		fmt.Println(creationTime)

		// controlActProcess := v.SelectElement("controlActProcess")
		// subjectNode := controlActProcess.SelectElement("subject")
		// investigationEventNode := subjectNode.SelectElement("investigationEvent")
		// referenceNodes := investigationEventNode.SelectElements("reference")

		// fmt.Println(len(referenceNodes))
		// referenceNode1 := referenceNodes[0]
		// documentNode := referenceNode1.SelectElement("document")
		// textNode := documentNode.SelectElement("text")
		// text := textNode.Text()
		// mediaType := textNode.SelectAttrValue("mediaType", "")
		// compression := textNode.SelectAttrValue("compression", "")

		// //fmt.Println(text)
		// fmt.Println(mediaType)
		// fmt.Println(compression)

		// //base64解码
		// // 使用 base64.StdEncoding.DecodeString 进行解码
		// decodedBytes, err := base64.StdEncoding.DecodeString(text)
		// if err != nil {
		// 	fmt.Println("解码失败:", err)
		// 	return
		// }
		// var fileBuffer bytes.Buffer
		// if compression == "DF" {
		// 	reader := flate.NewReader(bytes.NewReader(decodedBytes))
		// 	_, err := io.Copy(&fileBuffer, reader)
		// 	if err != nil {
		// 		fmt.Println("解压失败:", err)
		// 		return
		// 	}
		// 	reader.Close()
		// } else {
		// 	bytes.NewBuffer(decodedBytes)
		// }

		// //DF

		// // 创建一个读取器，用于解压 DEFLATE 数据

		// // // 读取解压后的数据
		// // var decompressedData bytes.Buffer
		// // _, err := io.Copy(&decompressedData, reader)
		// // if err != nil {
		// // 	fmt.Println("解压失败:", err)
		// // 	return
		// // }

		// // 打开或创建目标文件
		// file, err := os.Create("D:\\output.pdf") // 创建或覆盖文件
		// if err != nil {
		// 	fmt.Println("无法创建文件:", err)
		// 	return
		// }
		// defer file.Close() // 确保函数结束时关闭文件

		// // 将 io.Reader 的内容写入文件
		// _, err = io.Copy(file, &fileBuffer)
		// if err != nil {
		// 	fmt.Println("写入文件失败:", err)
		// 	log.Panic(err)
		// }

	}

}
