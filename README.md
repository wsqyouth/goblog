### 个人博客系统

cd cmd
go build -o goblog main.go
./goblog & 
基于Go语言的一个简易markdown博客系统 

示例：[codwiki.cn](https://codwiki.cn) 

源码：[github.com/MilkyMoon/goblog](https://github.com/MilkyMoon/goblog)


#### 项目目的
通过实现一个简单的博客系统来学习Go语言的开发，在个人学习的同时也给其他正在学习Golang的同学一个参考。本项目将不依赖数据库存储，主要通过md文件的形式存储，也方便编辑。

#### 项目需求
1. ~~实现一个简单的博客网站~~
2. 实现一个简单的markdown文档阅读器
3. 实现一个简单的缓存管理
4. 通过文件夹实现阅读器的目录结构
5. ~~通过github的Webhooks实现自动更新~~
6. 图床工具
7. 日志记录

#### 更新日志
1. 完成网站基本内容  2020-05-18
2. 实现webhooks自动更新  2020-05-26
3. 实现分页功能  2020-06-02
4. 新增gomod依赖管理 2020-11-12

#### 现存问题
1. ~~暂无分页，后续实现~~
2. linux下无法获取文件创建时间

#### 项目依赖
- [github.com/kataras/iris](https://github.com/kataras/iris) Web框架
- [github.com/pelletier/go-toml](https://github.com/pelletier/go-toml) toml配置文件读取
- [github.com/google/uuid](https://github.com/google/uuid) 生成全局唯一id

#### 项目运行
下载源码
```shell script
git clone https://github.com/MilkyMoon/goblog.git
cd goblog/cmd
go run main.go
```
下载依赖

```shell
go mod tidy
```

编译项目

```shell script
go build -o goblog main.go

#mac下编译linux执行文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goblog main.go
```
运行项目
```shell script
./goblog
#注意执行文件要与静态资源文件路径一致，项目根目录为当前运行目录
```
浏览器访问
[http://localhost:8181](http://localhost:8181)

#### 文档编写
1. 文档存放在resources/docs，在该目录下创建分类文件夹即可
2. 文档中`<br description/>` 文章描述 `<br description/>`标签内的内容将作为列表的描述
3. 文档中非空的第一行文字如果是标题（# ## ###），则将作为文档标题，否则以文件名作为标题



#### coopers博客搭建
1. 运行该服务,待服务启动后检查该服务是否正常运行
```
/home/coopers/workspace/goblog/cmd
nohup ./goblog > wsq.log &
http://159.75.38.33:8181
```
2. nginx配置转发
首先检查腾讯云域名是否请求该ip,之后配置端口转发
```
cd /usr/local/nginx/sbin
./nginx
conf配置:
location / {
    #proxy_pass http://www.baidu.com;
    #proxy_pass http://www.wsqyouth.cn:8181;
    proxy_pass http://127.0.0.1:8181;
        root   html;
        index  index.html index.htm;
    }
防火墙配置:
https://console.cloud.tencent.com/lighthouse/instance/detail?searchParams=rid%3D1&rid=1&id=lhins-rc1v9izt&tab=firewall

ping www.wsqyouth.cn 检查防火墙是否配置正常
www.wsqyouth.cn
```


### 为啥使用Iris而不是Gin
Kataras/Iris 和 Gin 都是优秀的 Go Web 框架，它们都具有高性能和易用性。
如果你需要更丰富的功能和模板引擎支持，可以选择 Iris；如果你更关注稳定性和社区支持，可以选择 Gin。
* 性能：Iris 的一个主要关注点是性能。它的路由器和中间件处理速度非常快，可以处理大量的并发请求，这使得 Iris 非常适合构建高性能的 Web 服务。
* 模块化设计：Iris 使用模块化的设计，允许开发者根据需要添加或删除功能。这种设计使得 Iris 非常灵活，可以适应各种不同的开发需求。
* 中间件支持：Iris 支持中间件，这使得开发者可以轻松地添加各种功能，如日志记录、身份验证、跨域资源共享（CORS）等。

关键字: Kataras/Iris 博客
相关文档：https://www.cnblogs.com/huiyichanmian/p/15251697.html