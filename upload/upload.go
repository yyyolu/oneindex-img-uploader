package upload

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

// 文件上传
func PostFile(filename string, targetUrl string, fileName string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		os.Exit(-1)
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		os.Exit(-1)
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		os.Exit(-1)
	}

	tr := &http.Transport{
		// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bodyWriter.Close()
	req, err := http.NewRequest("POST", targetUrl, bodyBuf)
	if err != nil {
		fmt.Println("发送报文失败", err)
		os.Exit(-1)
	}

	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/87.0")
	client := &http.Client{Transport: tr, Timeout: time.Second * time.Duration(60)}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(-1)
	}
	// 获取返回的HTML内容
	str1 := string(resp_body)
	// 获取出现http的最后一次出现
	num1 := strings.LastIndex(str1, "https")
	// 获取文件名的最后一次出现
	num2 := strings.LastIndex(str1, fileName)
	str := str1[num1 : num2+len(fileName)]
	// 写入文件
	local, err := os.Getwd()
	if err != nil {
		return err
	}
	localFile := local + "/img.txt"
	file, err := os.OpenFile(localFile, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.WriteString("\n")
	writer.Flush()
	file.Close()
	return nil
}
