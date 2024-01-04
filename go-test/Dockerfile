# 使用 golang 官方映像檔作為基礎
FROM golang:latest

# 複製程式碼到容器中的 /app 資料夾
COPY . /app

# 設定工作目錄為 /app
WORKDIR /app

# 編譯程式碼
RUN go build -o main .

# 預設的容器啟動指令
CMD ["./main"]

