package main

import (
	"io/ioutil"
	"net/http"

	_ "github.com/apache/skywalking-go"
	"github.com/sirupsen/logrus"
)

var n = 5

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		logrus.Debug("This is a debug log in serviceA before calling serviceB")
		logrus.Info("This is a info log in serviceA before calling serviceB")
		logrus.Warn("This is a warn log in serviceA before calling serviceB")
		logrus.Error("This is a error log in serviceA before calling serviceB")
		c := make(chan string, n)
		for i := 0; i < n; i++ {
			go func() {
				c <- callServiceB()
			}()
		}
		s := ""
		for i := 0; i < n; i++ {
			s += <-c
			if i < n-1 {
				s += "\n"
			}
		}
		defer close(c)
		logrus.Debug("This is a debug log in serviceA after calling serviceB")
		logrus.Info("This is a info log in serviceA after calling serviceB")
		logrus.Warn("This is a warn log in serviceA after calling serviceB")
		logrus.Error("This is a error log in serviceA after calling serviceB")
		writer.Write([]byte(s))
	})
	logrus.Info("ServiceA is running on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func callServiceB() string {
	resp, err := http.Get("http://localhost:8001/helloB")
	if err != nil {
		logrus.Error("Failed to call serviceB")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("Failed to read response body")
	}
	return string(body)
}