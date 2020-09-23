package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// PathExists 文件存在检测
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// BuildDir 创建目录
func BuildDir(abs_dir string) error {
	return os.MkdirAll(path.Dir(abs_dir), os.ModePerm) //生成多级目录
}

// RemoveFile 删除文件或文件夹
func RemoveFile(abs_dir string) error {
	return os.RemoveAll(abs_dir)
}

// GetPathDirs 获取目录所有文件夹
func GetPathDirs(abs_dir string) (re []string) {
	if PathExists(abs_dir) {
		files, _ := ioutil.ReadDir(abs_dir)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// GetPathFiles 获取目录所有文件
func GetPathFiles(abs_dir string) (re []string) {
	if PathExists(abs_dir) {
		files, _ := ioutil.ReadDir(abs_dir)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// GetCurrentDir 获取程序运行路径
func GetCurrentDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

// GetExeDir 获取可执行文件的路径文件夹
func GetExeDir() (dir string, err error) {
	path, err := os.Executable()
	if err != nil {
		return
	}
	dir = filepath.Dir(path)
	return
}
