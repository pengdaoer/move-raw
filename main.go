package main

import (
	"fileMove/file"
	. "fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	//如果接收不到参数，获取当前程序所在路径，且默认移动的文件为 .CR3
	if len(args) == 0 {
		path := filepath.Dir(os.Args[0]) + "/"
		moveAction(path)
	} else if len(args) == 1 {
		//如果接收一个参数，那这个参数必须是后缀名，表示移动特定类型的文件
		path := filepath.Dir(os.Args[0]) + "/"
		file.Type = args[0]
		moveAction(path)
	} else {
		//如果接收大于一个参数，第一个是文件类型，从第二个开始就是文件夹路径
		file.Type = args[0]
		//文件夹路径数组重新分割
		paths := args[1:]
		// 遍历参数的集合
		for _, path := range paths {
			Println("正在处理文件夹：", path)
			path += "/"
			moveAction(path)
			Println(path, "文件夹处理完成")
		}
	}
}

func moveAction(path string) {

	Println("开始进行文件移动")
	Println("正在创建RAW文件夹")
	_, err := file.CreateRawDir(path)
	if err != nil {
		Println("文件夹创建失败，程序运行结束")
		return
	}
	Println("文件夹创建成功")
	Println("正在查找文件")
	files, err := file.FindCR3Files(path)
	Println("获取到所有文件数量为", len(files))
	if err != nil {
		Println("文件查找过程中出现错误，程序运行结束")
		return
	}
	Println("文件查找完成")
	Println("正在移动文件")
	if err := file.MoveCR3Files(path, files); err != nil {
		Println("文件移动过程中出现错误，程序运行结束")
		return
	}
	Println("移动成功，程序运行结束")
}
