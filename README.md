# Go-make

> [https://www.ruanyifeng.com/blog/2015/02/make.html](https://www.ruanyifeng.com/blog/2015/02/make.html)
> 
> 作者： 阮一峰
> 
> 日期： 2015年2月20日

代码变成可执行文件，叫做编译（compile）；

先编译这个，还是先编译那个（即编译的安排），叫做构建（build）。

Make是最常用的构建工具，诞生于1977年，主要用于C语言的项目。但是实际上 ，任何只要某个文件有变化，就要重新构建的项目，都可以用Make构建。
