package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // 初始化 Gin 引擎
    router := gin.Default()

    // 路由定義
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    // 啟動伺服器
    router.Run(":8080")
}

