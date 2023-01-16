package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator) //存放路径分隔符

type Entry interface {
	/*
		寻找和加载class文件
		参数是class文件的相对路径,返回读到的字节数据,最终定位到class文件的Entry以及错误信息
	*/
	readClass(className string) ([]byte, Entry, error)
	//相当于Java中的toString()方法
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ",JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
