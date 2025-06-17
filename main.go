package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// 提示用户输入 Slug
	fmt.Println("【创建文章】")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入 Slug: ")
	slug, _ := reader.ReadString('\n')
	slug = strings.TrimSpace(slug)

	// 获取当前日期（格式化为 YYYY/MMdd）
	dateStr := time.Now().Format("2006/0102")

	// 创建文章
	cmd := exec.Command("hugo", "new", fmt.Sprintf("post/%s-%s/index.md", dateStr, slug))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("创建文章时出错： %v\n", err)
		return
	}

	fmt.Println("文章已创建！")
}
