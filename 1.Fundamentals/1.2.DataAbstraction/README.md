# 目录

## 练习解答
* [1.2.1](./1.2.1/main.go)
* [1.2.2](./1.2.2/main.go)
* [1.2.3](./1.2.3/main.go)
* [1.2.4](./1.2.4/main.go)
* [1.2.5](./1.2.5/main.go)
* [1.2.6](./1.2.6/main.go)
* [1.2.7](./1.2.7/main.go)
* 1.2.8 在Go语言中也一样，详情请查阅Go语言切片的定义。
* [1.2.9](./1.2.9/main.go)
* [1.2.10](./1.2.10/main.go)
* 1.2.11 略
* 1.2.12 略
* 1.2.13 略
* 1.2.14 略
* 1.2.15 略
* [1.2.16](./1.2.16/main.go)
* [1.2.17](./1.2.16/main.go) 我在1.2.16的答案中，尽可能地对各种错误情况进行了预处理与报错处理。以下是Golang官网对为何Golang没有断言功能的回复。

>[Why does Go not have assertions?](https://golang.org/doc/faq#assertions)
>
>Go doesn't provide assertions. They are undeniably convenient, but our experience has been that programmers use them as a crutch to avoid thinking about proper error handling and reporting. Proper error handling means that servers continue operation after non-fatal errors instead of crashing. Proper error reporting means that errors are direct and to the point, saving the programmer from interpreting a large crash trace. Precise errors are particularly important when the programmer seeing the errors is not familiar with the code.
>
>We understand that this is a point of contention. There are many things in the Go language and libraries that differ from modern practices, simply because we feel it's sometimes worth trying a different approach.
* [1.2.18](./1.2.18/main.go) 我利用[百度百科](http://baike.baidu.com/item/%E6%A0%87%E5%87%86%E5%81%8F%E5%B7%AE)页面介绍的求样本方差和标准差的方法，进行了对比求解，发现两者的对比不大，反而是利用了"math/big"包进行精确计算的结果，与两者相差较大。我不知道是由于Golang和Java的表示方式不一样，还是什么原因引起的。
* 1.2.19 略
