package file

import (
	"fmt"
	"log"
	"math/rand"
	"oneindex-img-uploader/upload"
	"os"
	"path/filepath"
	"time"
)

var Global_path []string
var Global_url []string

// 此处对文件夹进行处理
// 接收的是一个文件夹(绝对路径)
// 1. 遍历文件夹(此处不会管其中是否存在多个文件夹，仅限单文件夹)
func File_deal(path string, urls *[]string) error {
	Global_url = *urls
	err := filepath.Walk(path, walkFunc)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 文件夹遍历操作
func walkFunc(path string, info os.FileInfo, err error) error {
	if info == nil {
		// 文件名称超过限定长度等其他问题也会导致info == nil
		// 如果此时return err 就会显示找不到路径，并停止查找。
		println("can't find:(" + path + ")")
		return nil
	}
	// 如果是文件夹
	if info.IsDir() {
		// 不打印根目录
		if !IsDoubleFolder(path, &Global_path) {
			// 里面存在，表明已经走过这了，直接返回，跳过此目录
			return filepath.SkipDir
		}
		fmt.Println(path)
		// 递归遍历
		filepath.Walk(path, walkFunc)
		return nil
	} else {
		// 如果是文件，在此处执行文件上传操作
		// 文件大小判断，大于4M直接跳过
		if info.Size() > 4*1024*1024 {
			return nil
		}
		fmt.Println(path)
		// 文件上传
		// 生成随机的链接
		// 我们一般使用系统时间的不确定性来进行初始化
		rand.Seed(time.Now().Unix())
		temp_url_index := rand.Intn(len(Global_url))
		temp_url := Global_url[temp_url_index]
		// 文件上传
		// 文件路径 URL 文件名称
		// 根据后缀判断
		err := upload.PostFile(path, temp_url, filepath.Ext(info.Name()))
		if err != nil {
			log.Println(path, "上传失败")
			return nil
		}
		// 休息一秒
		time.Sleep(time.Second)
		return nil
	}
}
