# 介绍
简单的 tcp 端口转发程序

# 编译
go build -ldflags '-w -s' tcpforward.go  
go build -ldflags '-w -s -linkmode external -extldflags "-static"' tcpforward.go   #静态链接


# 使用
```
./tcpforward -l 0.0.0.0:443 -r 8.8.8.8:443
```
