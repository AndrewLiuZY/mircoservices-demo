package main

import (
	"fmt"
	"log"
	"net/http"

	"gin-blog/models"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Listening port: %d", setting.ServerSetting.HttpPort)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
