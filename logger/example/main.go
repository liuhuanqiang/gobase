package main

import "github.com/liuhuanqiang/gobase/logger"

func main() {
	logger.New("test")

	logger.Info("this is a info log.")

	logger.Warn("this is a Warn log.")

	logger.Debug("this is a debug log.")

	logger.Error("this is a error log.")

}
