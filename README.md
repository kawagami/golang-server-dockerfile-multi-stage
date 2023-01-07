# 概述
* 練習 dockerfile multi-stage build 的產物
* clone 此專案後移動到專案下，執行下列指令
```
docker build -t kawa/golang:alpine .
```
* 編譯完後執行下列指令
```
docker run --rm -p 80:9090 -d --name go-server kawa/golang:alpine
```
* 即可在 http://localhost/ 看到 golang server 的回應
