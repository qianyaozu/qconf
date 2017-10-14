package qconf

import (
	"testing"
	"fmt"
)

func TestQconf(t *testing.T) {
	conf, err := LoadConfiguration("conf.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)
	fmt.Println(conf.kv)
	fmt.Println(conf.getInteger("WorkerHeartTime"))
	fmt.Println(conf.getInteger("vpshearttime"))
	fmt.Println(conf.getString("redispassword"))
}
