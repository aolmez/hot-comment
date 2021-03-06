package main

import (
	"fmt"
	"github.com/goalong/hot-comment/router"
)

// @title 云音乐歌曲、评论搜索API
// @version 1.0
// @description 使用Go、Gin、Elasticsearch开发的一个网易云音乐歌曲、评论搜索API，可以在web上点击发请求哦，能查到数据，不过当然不是全部的数据啦


// @host 127.0.0.1:8080
// @BasePath /
func main() {
	newRouter := router.InitRouter()
	fmt.Println(newRouter)
	newRouter.Run()
	
}
