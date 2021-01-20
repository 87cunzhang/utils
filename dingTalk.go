package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var webHook, secretKey string

func InitDingTalkConf(webHookConf string, secretKeyConf string) {
	webHook = webHookConf
	secretKey = secretKeyConf
}

//发送钉钉消息
func SendDingMsg(content string) string {
	httpUrl := makeHttpUrl()
	postStr := "{\"msgtype\":\"text\",\"text\":{\"content\":\"" + content + "\"},\"at\":{\"atMobiles\":[]}}\n"
	resp, _ := http.Post(httpUrl, "application/json;charset=utf-8", strings.NewReader(postStr))
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

//拼装请求地址
func makeHttpUrl() string {
	currTime := time.Now().UnixNano() / 1e6
	httpUrl := webHook + "&timestamp=" + strconv.FormatInt(currTime, 10) + "&sign=" + computeSignature(currTime, secretKey)
	return httpUrl
}

//计算签名
func computeSignature(timestamp int64, secret string) string {
	b := &[]byte{}
	*b = append(*b, strconv.FormatInt(timestamp, 10)...)
	*b = append(*b, '\n')
	*b = append(*b, secret...)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(*b)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
