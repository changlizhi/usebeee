# learbee
把饿了么的前端实现同beego整合起来

## 文件描述

* 在learnvue项目完成之后执行了npm run build之后生成了dist文件
* 把dist文件夹下的文件复制到bee new出来的此项目相同文件夹下
* 发现learnvue下读取的data.json没有获取，然后又在main.go提供了一个获取data.json的controller
* 更换获取数据的链接为main.go提供的restapi
* 重新npm run build
* done.
