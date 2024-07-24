package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("pacman", "-Syu")
	cmd.Run()
	fmt.Println("欢迎使用archWSLsetHelper!")
	fmt.Println("请给root用户设置密码：")
	cmd = exec.Command("passwd", "root")
	cmd.Run()
	fmt.Println("请输入你想创建的用户名：（只允许用英文字母）")
	var username string
	fmt.Scanln(&username)
	cmd = exec.Command("useradd", username, "-m", "-G", "wheel", "-s", "/bin/bash")
	cmd.Run()
	fmt.Println("请给用户设置密码：")
	cmd = exec.Command("passwd", username)
	cmd.Run()
	// echo -e "[user]\ndefault = ws" >> /etc/wsl.conf
	cmd = exec.Command("echo", "-e", "[user]\ndefault =", username, ">>", "/etc/wsl.conf")
	cmd.Run()
	// 切换到用户
	cmd = exec.Command("su", username)
	cmd.Run()
	cmd = exec.Command("cd", "~")
	cmd.Run()
	// 安装yay
	cmd = exec.Command("git", "clone", "https://aur.archlinux.org/yay.git")
	cmd.Run()
	cmd = exec.Command("cd", "yay")
	cmd.Run()
	cmd = exec.Command("makepkg", "-si")
	cmd.Run()
}
