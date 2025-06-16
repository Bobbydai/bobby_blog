#!/bin/bash

echo "【创建文章】"
read -p "请输入 Slug: " slug

# 获取当前日期
date_str=$(date +"%Y/%m%d")

# 创建文章
hugo new "post/${date_str}-${slug}/index.md"

echo "文章已创建！" 