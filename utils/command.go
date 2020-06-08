package utils

import (
	"fmt"
	"os/exec"
)

func RunCommand(name string, arg ...string) error {
	command := exec.Command(name, arg...)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := command.StdoutPipe()
	command.Stderr = command.Stdout

	if err != nil {
		return err
	}

	if err = command.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
	if err = command.Wait(); err != nil {
		return err
	}
	return nil
}
