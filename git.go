package main

import (
	"fmt"
	"github.com/go-playground/webhooks/v6/gitlab"
	"os"
	"os/exec"
	"strings"
)

func PullEvent(event gitlab.PushEventPayload) {

	//判断文件夹是否建立

	if _, err2 := os.Stat(event.Project.Name); os.IsNotExist(err2) {

		//如果没有创建，就执行git clone
		//git clone http://username:password@ip:port/xx/uem-vis-realtime.git
		fullcmd := strings.Replace(event.Project.GitHTTPURL, "http://", "git clone http://i4color%40qq.com:88888@", -1)
		println(fullcmd)
		gocmd(fullcmd)
	}

	fmt.Println("进入目录,并拉取最新代码")

	gocmd("cd " + event.Project.Name + "&&git pull")
	//
	//fmt.Println("拉取最新代码")
	//gocmd("git pull")
	gocmd("pwd")

}

func gocmd(cmdargs string) {
	//返回一个 cmd 对象
	cmd, err := exec.Command("/bin/bash", "-c", cmdargs).Output()
	if err != nil {
		fmt.Println("error %s", err)
	}
	output := string(cmd)
	fmt.Printf("output: %s", output)
}
