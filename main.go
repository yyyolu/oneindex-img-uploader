package main

import (
	"fmt"
	"log"
	"oneindex-img-uploader/file"
	"os"
	// "path/filepath"
)

// 此处存放所有oneindex的URL
var urls []string = []string{
	"https://share.yyyolu.com/images",
	"https://bb019.yyyolu.com/images",
}

func main() {
	// 此处存放文件夹路径
	var path string
	// var fileName string
	list := os.Args
	if len(list) == 2 {
		// 获取文件夹路径
		path = list[1]
		// 判断是否为文件夹
		if !file.IsDir(path) {
			fmt.Println("请输入文件夹路径，请不要输入文件路径")
			return
		}
		// 文件遍历操作
		err := file.File_deal(path, &urls)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		fmt.Println("请在可执行程序后加上你想要上传的文件夹路径，路径请用括号括起来，避免路径中的空格问题")
		return
	}
}
