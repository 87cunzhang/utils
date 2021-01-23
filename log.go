package utils
import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func LogErr(data []byte, err error) {
	logPath := DefaultConf().String("errLogPath")
	logData := fmt.Sprintf("%s dberr: %s,params: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error(), string(data))
	if err := ioutil.WriteFile(logPath, []byte(logData), 0644); err != nil {
		log.Println("write file err:", err)
	}
}
