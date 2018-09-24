package logger

import (
	"bytes"
	"fmt"
	"time"
)

//待log格式确定:remore_addr request method status httpversion time_local level path line
func (l *Logger) formatHeader(buf *bytes.Buffer, t time.Time, file string, line int, lvl int, reqId string) {

}

//格式化msg：k1=v1 k2=v2 k3=v3
func formatMsg(v ...interface{}) string {
	var buf bytes.Buffer

	if len(v) == 1 {
		buf.WriteString(fmt.Sprint(v[0]))
		return buf.String()
	}

	for i, item := range v {
		if isEven(i) {
			s := join([]string{fmt.Sprint(item), "="})
			buf.WriteString(s)
		} else {
			s := join([]string{fmt.Sprint(item), " "})
			buf.WriteString(s)
		}
	}

	return buf.String()
}
