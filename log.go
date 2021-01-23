package utils
import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

//记录日志
func LogErr(content string, err error) {
	logData := fmt.Sprintf("%s err: %s, content: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error(), content)
	fileName := DefaultConf().String("errLogPath") + "_" + time.Now().Format("2006-01-02")
	if err := ioutil.WriteFile(fileName, []byte(logData), 0644); err != nil {
		log.Println("write file err:", err)
	}
}
