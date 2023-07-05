package main

// https://www.cnblogs.com/xxnn/p/10875693.html
import "fmt"

// ********************************************************************************** 对象 ********************************************************************************** \\
/*
Java中的对象 class -> new -> instance
Go中的对象 从语法上来说 没有class 没有Object
那么到底什么是对象？看Go给的答案
*/

// 类型声明 struct
type Rect struct {
	width  int
	height int
}

// 给这个struct声明一个方法
// 编译器会再给我生成一个不是指针的方法
func (receiver *Rect) Area() int {
	return receiver.width * receiver.height
}

// 可以声明任何类型
type Rects []*Rect

// 同样可以给这个声明一个方法
func (receiver Rects) Area() int {
	var a int
	for _, rect := range receiver {
		a += rect.Area()
	}
	return a
}

// 甚至可以声明一个函数类型
type Foo func() int

// 同样的，给这个（函数）类型声明一个方法
func (receiver Foo) Add(x int) int {
	return receiver() + x
}

// ********************************************************************************** 继承 ********************************************************************************** \\
type Person struct {
	Name string
	// Go中通过嵌入类型来实现组合
	Address
}

type Address struct {
	City string
}

func (receiver Address) String() string {
	return receiver.City
}

func main() {
	rect := Rect{width: 10, height: 5}
	fmt.Println(rect.Area())

	r1 := Rect{width: 10, height: 5}
	r2 := Rect{width: 10, height: 5}
	rects := Rects{&r1, &r2}
	fmt.Println(rects.Area())

	var x Foo
	x = func() int {
		return 1
	}
	fmt.Println(x.Add(3))

	//使用组合字面量声明一个Struct
	p := Person{
		Name: "zhangsan",
		Address: Address{
			City: "Beijing",
		},
	}
	//Go中通过嵌入类型来实现组合
	fmt.Println(p.String())
}
