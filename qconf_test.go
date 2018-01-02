package qconf

import (
	"fmt"
	"testing"
)

func TestQconf(t *testing.T) {
	conf, err := LoadConfiguration("conf.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)
	fmt.Println(conf.kv)
	fmt.Println(conf.GetString("DialHours"))
	fmt.Println(conf.GetInteger("WorkerHeartTime"))
	fmt.Println(conf.GetInteger("vpshearttime"))
	fmt.Println(conf.GetInteger("redispassword"))
}
