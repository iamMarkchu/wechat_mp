package main

import (
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type WeixinRequest struct {
	XMLName      xml.Name `xml:"xml"`
	FromUserName string   `xml:"FromUserName"`
	ToUserName   string   `xml:"ToUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        int      `xml:"MsgId"`
}

type WeixinResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int
	MsgType      CDATA
	Content      CDATA
}

type CDATA struct {
	Text string `xml:",cdata"`
}

func main() {
	log.Println("开始")
	server := http.Server{
		Addr:           ":8000",
		Handler:        &httpHandler{},
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 0,
	}

	server.ListenAndServe()
}

func homeGet(w http.ResponseWriter, r *http.Request) {
	log.Println("HOME GET")
	str := r.URL.Query().Get("echostr")
	log.Println(str)
	w.Write([]byte(str))
}

func homePost(w http.ResponseWriter, r *http.Request) {
	log.Println("HOME POST")

	res, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var data WeixinRequest
	xml.Unmarshal(res, &data)

	var ret = WeixinResponse{
		ToUserName:   CDATA{data.FromUserName},
		FromUserName: CDATA{data.ToUserName},
		CreateTime: int(time.Now().Unix()),
		MsgType: CDATA{data.MsgType},
		Content: CDATA{data.Content},
	}
	test, _ := xml.Marshal(ret)
	fmt.Println(string(test))
	w.Write(test)
}

func verify(r *http.Request) (bool, error) {

}
