package logger

import (
	"os"
	"strings"
	"time"
)

//判断文件是否已经存在(通用方法)
func isExist(name string) bool {
	if len(name) < 1 {
		return false
	}

	_, err := os.Stat(name)

	return err == nil || os.IsExist(err)
}

//定点闹钟
func (std *Logger) Clock(h, m, s, ns int) {
	// go func() {
	// 	for {
	// 		now := time.Now()

	// 		//计算下一个零点
	// 		next := now.Add(time.Hour * 24)
	// 		next = time.Date(next.Year(), next.Month(), next.Day(), h, m, s, ns, next.Location())

	// 		t := time.NewTimer(next.Sub(now))
	// 		<-t.C //must do

	// 		s := join(strings.Split(timeToStr(next), "-"))
	// 		err := std.ChangeFileTime(s)
	// 		if err != nil {
	// 			//务必切换成功
	// 			//切换失败，则尝试再次切换，并保证上次失败的切换没有产生垃圾文件
	// 		}
	// 	}
	// }()
}

//字符串拼接
func join(strs []string) string {
	return strings.Join(strs, "")
}

//Time转为string类型
func timeToStr(t time.Time) string {
	return t.Format("2006-01-02")
}

//判断是否偶数
func isEven(i int) bool {
	if i&0x1 == 0 {
		return true
	}

	return false
}
