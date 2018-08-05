# goblog

> 个人博客goblog后端

[![Go Report Card](https://goreportcard.com/badge/github.com/1377195627/goblog)](https://goreportcard.com/report/github.com/1377195627/goblog)
[![GoDoc](https://godoc.org/github.com/1377195627/goblog?status.svg)](https://godoc.org/github.com/1377195627/goblog)
[![Sourcegraph](https://sourcegraph.com/github.com/1377195627/goblog/-/badge.svg)](https://sourcegraph.com/github.com/1377195627/goblog@2.0)

## 前端

[goblog-web](https://github.com/xiayesuifeng/goblog-web.git)

## 重构计划

现已使用 [gin](https://github.com/gin-gonic/gin) , [gorm](https://github.com/jinzhu/gorm) , [gin-sessions](https://github.com/gin-contrib/sessions) 进行后端重构，前端使用react重构，使用axios和后端 API 进行交互。

- [X] 架构搭建；
- [X] 登录；
- [X] 获取article；
- [X] 新建article；
- [X] 编辑article；
- [X] 删除article;
- [X] 获取分类;
- [X] 新建分类;
- [X] 编辑分类;
- [X] 删除分类;
- [ ] 插件机制;

## 编译

```
go get github.com/xiayesuifeng/goblog
go build -ldflags "-s -w" github.com/xiayesuifeng/goblog
```

## 配置

> 前端

```
git clone https://github.com/xiayesuifeng/goblog-web.git
npm install
npm build
```

> 后端

```
cp goblog goblog-web
wget -O goblog-web/config.json https://github.com/xiayesuifeng/goblog/raw/master/config.default.json
```

> config.json

```
{
  "mode":"debug",           //调试模式，正式部署改为release
  "name": "goblog",         //个人博客名
  "password": "0925e15d0ae6af196e6295923d76af02b4a3420f",   //登录密码,当前为admin
  "useCategory": true,      //使用分类功能,不使用分类改为false
  "dataDir":"data",         //数据存放路径,当前为goblog运行路径下的data下，用于存储article和图片等
  "database":{              //数据库信息(暂只支持mysql,敬请期待别的数据库支持)
    "username":"root",      //数据库用户名
    "password":"",          //数据库密码
    "dbname":"goblog",      //数据库名(需要手动创建)
    "address":"127.0.0.1",  //数据库地址，当前为本地
    "port":"3306"           //数据库端口
  },
  "smtp":{                  //邮箱配置,用于当article有新评论时发送邮件通知(尚未支持,无需配置,敬请期待)
    "username":"",
    "password": "",
    "host": ""
  }
}
```

> caddy配置实例

```
your {
    root /your/path/goblog-web
    gzip
    
    rewrite {
        if {path} not_match ^/api
        to {path} {path} /index.html
    }
    proxy /api localhost:20181
}
```

## 加密密码生成

```
echo -n yourpassword | openssl dgst -md5 -binary | openssl dgst -sha1
```

## Contributors

- [Contributors](https://github.com/xiayesuifeng/goblog/graphs/contributors)

## License

goblog is licensed under [GPLv3](LICENSE).
