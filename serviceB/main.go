package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/apache/skywalking-go"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	http.HandleFunc("/helloB", func(writer http.ResponseWriter, request *http.Request) {
		logrus.Debug("This is a debug log in serviceB")
		logrus.Info("This is a info log in serviceB")
		logrus.Warn("This is a warn log in serviceB")
		logrus.Error("This is a error log in serviceB")

		// random sleep duration to simulate the delay, from 1s to 10s
		sleepDuration := time.Duration(rand.Intn(9000) + 1000) * time.Millisecond
		time.Sleep(sleepDuration)

		s := fmt.Sprintf("Hello %s!", gofakeit.Name())
		writer.Write([]byte(s))
	})
	logrus.Info("ServiceB is running on port 8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}
}
