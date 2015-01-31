# ting
>Ting, tiny static file server for quick demo preview.

```
Name   : Ting Static Server
Version: 1.0.0
Author : http://junqiang.wang/
```  

一个小巧的demo预览工具，主要解决html demo快速预览，但是依赖服务器环境的情况下。

在要预览的目录下执行
```
$cd /path/to/demo
$ting -p=端口号
```
即可,默认监听80端口

```
$ting -h
Usage:	 ting -p=12345
-h=false: show help info
-p=80: server port, default 80
```

golang命令行彩色文字代码来源：
[[github.com/xcltapestry/xclpkg/clcolor]]

>绝对不要用在生产环境，真心没测试过，玩具代码
