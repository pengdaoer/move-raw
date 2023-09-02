package file

import (
	"os"
	"path/filepath"
)

// Type 文件类型
var Type = ".CR3"

//变小写 string.ToLower("ABC")

// CreateRawDir 若没有RAW文件夹则新建文件夹
func CreateRawDir(path string) (string, error) {
	if err := os.MkdirAll(filepath.Dir(path+"RAW/"), 0777); err != nil {
		return path, err
	}
	return path + "RAW/", nil
}

// FindCR3Files 遍历文件夹中的所有文件，并找出后缀名为.CR3的文件
func FindCR3Files(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// 判断文件名是否以 ".CR3" 为结尾
		if filepath.Ext(path) == Type {
			//获取文件名，使其不带绝对路径
			name := filepath.Base(path)
			files = append(files, name)
		}
		// 判断文件名是否以 ".CR3" 为结尾
		//if strings.HasSuffix(path, Type) {
		//	fmt.Println(path)
		//}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// MoveCR3Files 将查到的.CR3文件移动到新建的RAW文件夹中
func MoveCR3Files(path string, files []string) error {
	// 遍历文件名的集合
	for _, file := range files {
		// 源文件的路径
		src := filepath.Join(path, file)

		// 目标文件的路径
		dest := filepath.Join(path+"RAW", file)

		// 移动文件
		err := os.Rename(src, dest)
		if err != nil {
			return err
		}
	}
	return nil
}
