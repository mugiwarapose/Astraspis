package xml

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
)

//go:embed *.xml
var xmTemplate embed.FS

func Generate() {

	patient := map[string]string{
		"name": "John Doe",
		"age":  "30",
		//"ageUnit": html.EscapeString("adf©'&"),
		"ageUnit": "adf©'&",
	}
	//tmpl, err := template.New("xml").ParseFS(xmTemplate, "template.xml")
	tmpl, err := template.ParseFS(xmTemplate, "template.xml")
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	buffer.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	err = tmpl.Execute(buffer, patient)
	if err != nil {
		panic(err)
	}

	// 打开或创建目标文件
	file, err := os.Create("D:\\zxd.xml") // 创建或覆盖文件
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close() // 确保函数结束时关闭文件

	// 将 io.Reader 的内容写入文件
	_, err = io.Copy(file, buffer)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		log.Panic(err)
	}
}
