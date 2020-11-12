package common

import (
	"goblog/config"
	"path/filepath"
)

func GetRootPath() string {
	return filepath.Join(config.Root,config.Conf.Get("res.root_dir").(string))
}

//获取文章文件夹根目录
func GetDocsPath() string {
	return filepath.Join(config.Root,config.Conf.Get("res.docs_dir").(string))
}

//获取文档文件夹根目录
func GetBookPath() string  {
	return filepath.Join(config.Root,config.Conf.Get("res.books_dir").(string))
}

//获取图片文件夹根目录
func GetImagePath() string {
	return filepath.Join(config.Root,"web",config.Conf.Get("image_host.img_dir").(string))
}

//为了兼容windows对大小写不敏感，所以使用36进制
func BaseTo36(num int64) string {
	var char = [36]string{"0","1","2","3","4","5","6","7","8","9","a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

	res := ""
	for num != 0 {
		res = char[num % 36] + res
		num = num / 36
	}

	return res
}

type Error struct {
	code  int
	msg   string
}

func (e *Error) Error() string {
	return e.msg
}