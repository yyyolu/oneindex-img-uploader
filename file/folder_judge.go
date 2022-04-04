package file

import "os"

// 此处用于判断是否为文件夹
func IsDir(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// 用于判断文件夹遍历时是否重复打印文件夹内容
func IsDoubleFolder(path string, global_path *[]string) bool {
	for _, v := range *global_path {
		if v == path {
			// 如果找到了，直接离开此目录
			return false
		}
	}
	*global_path = append(*global_path, path)
	return true
}
