//撰寫程式

package main  //可執行程式必須使用main封包
import "fmt"  //建立內建的fmt封包，用來做基本輸出輸入
func main() {  //建立main函式。程式的進入點
	fmt.Print("Hello Golang")  //執行程式從這邊開始
}
//建置(Build):go build 程式的名稱
//執行程式:輸入.\執行檔的名稱

//資料型態
//整數int:5
//浮點數float64:5.2
//字串string:"apple"
//布林值bool:true、false
//字元rune:'a'
package main
import "fmt"
func main(){  //輸出訊息到終端:fmt.Println(資料,資料)
	fmt.Println(3)  //整數
	fmt.Println(3.14)  //浮點數
	fmt.Println("Hello Golang")  //字串
	fmt.Println(true)  //布林值
	fmt.Println('a')  //字符(以整數儲存)
}

//變數的使用流程
//宣告變數:var 變數名稱 資料型態名稱
var x int
//指定資料:變數名稱=對應型態的資料
x=4
//使用變數:使用變數名稱代替資料來做運算
fmt.println(x)
//指定新的資料，覆蓋新的資料
x=10
fmt.println(x)
//宣告變數順便放進資料
var f float64=3.14
fmt.println(f)
var s string="你好"
fmt.println(s)
var b bool=true
fmt.println(b)
var r rune='c'
fmt.println(r)

//基本輸出語法
package main
import "fmt"
func main(){ 
	fmt.Println(3) 
}
//基本輸入語法
package main  
import "fmt"  
func main(){
var x int  //宣告變數整數x
var y int  //宣告變數整數y
fmt.Print("請輸入一個整數:")
fmt.Scanln(&x)  //輸入變數x
fmt.Print("請輸入一個整數:")
fmt.Scanln(&y)  //輸入變數y
var z int=x+y  //宣告變數整數z=x+y
fmt.Println(z)  //輸出變數z
}

package main  
import "fmt"  
func main(){
var x int  //宣告變數整數x
var y int  //宣告變數整數y
fmt.Print("請輸入兩個整數，用空白隔開:")
fmt.Scanln(&x,&y)  //輸入變數x及變數y
var z int=x+y  //宣告變數整數z=x+y
fmt.Println(z)  //輸出變數z
}

//基本運算符號
//算術運算:"+"、"-"、"*"、"/"
package main  
import "fmt"  
func main(){
	var x int
	x=6+3
	fmt.Println(x)
	x=6-3
	fmt.Println(x)
	x=6/3
	fmt.Println(x)
	x=6*3
	fmt.Println(x)
}
//指定運算:"="、"+="、"-="、"*="、"/="
package main  
import "fmt"  
func main(){
	var x int
	x=5
	fmt.Println(x)
	x+=1
	fmt.Println(x)
	x-=1
	fmt.Println(x)
	x*=1
	fmt.Println(x)
	x/=1
	fmt.Println(x)

}
//單元運算:"++"、"--"
package main  
import "fmt"  
func main(){
	x++  //資料x+1
	fmt.Println(x)
	x--  //資料x-1
	fmt.Println(x)
}
//比較運算:">"、"<"。">="、"<="、"=="
package main  
import "fmt"  
func main(){
	var x bool=4>3
	fmt.Println(x)  //印出true
	var x bool=4<3
	fmt.Println(x)  //印出flase
	var x bool=4==4  //4和4是否相等
	fmt.Println(x)  //印出true
}
//邏輯運算:"!"、"&&"、"||"
package main  
import "fmt"  
func main(){
	var x bool
	x=!true  //!是布林值的相反
	fmt.Println(x)  //印出false
	x=true&&false  //&&是且的意思(若有負責得負)
	fmt.Println(x)  //正正得正；正負得負；負正得負；負負得負
	x=true||false  //||是或的意思(若有正得正))
	fmt.Println(x)  //正正得正；正負得正；負正得正；負負得負
}

//if 判斷式
package main  
import "fmt"  
func main(){
	var x int
	fmt.Print("請輸入數字:")
	fmt.Scanln(&x)
	if x>100{
		fmt.Println("x>100")
	}else if x==100{
		fmt.Println("x=100")
	}else{
		fmt.Println("x<100")
	}
}

//for 迴圈
package main  
import "fmt"  
func main(){
	var x int=0
	for x<11{
		fmt.Println(x)
		x+=1
	}
}
//初始化命令迴圈(在一開始就設定好條件)
package main  
import "fmt"  
func main(){
	var x int
	for x=0;x<11;x++{
		fmt.Println(x)
	}
}
//實作練習:1+2+3+...+50
package main  
import "fmt"  
func main(){
	var sum int=0
	var x int=0
	for x<51 {
		sum=sum+x
		fmt.Println("x=",x,"總和是",sum)
		x++
	}
}
//break應用
package main  
import "fmt"  
func main(){
	var x  int=0
	for x<5 {
		if x==3{
			break
		}
	fmt.Println(x)
	x++
}
}
//continue應用
package main  
import "fmt"  
func main(){
	var x  int=0
	for x<5 {
		if x==3{
			continue
		}
	fmt.Println(x)
	x++
}
}

//函式的基礎與應用
package main  
import "fmt"
func msg(x string){  //函式1:印出字串
	fmt.Println(x)
}  
func x(n1 int,n2 int){  //函式2:相乘
	var y int=n1*n2
	fmt.Println(y)
}
func sum(z int){  //函式3:連加
	var result int=0
	var n int
	for n=1;n<=z;n++{
		result=result+n
	}
	fmt.Println("result:",result)
}
func main(){  //主程式
	msg("你好")  //呼叫函式1(msg)
	x(10,5)  //呼叫函式2(x)
	x(5,-3)  //呼叫函式2(x)
	sum(20)  //呼叫函式3(sum)
	sum(5)  //呼叫函式3(sum)
}

//回傳值的應用
package main  
import "fmt"
func msg(x string){
	return  //return在Print之前，所以不會印出結果
	fmt.Println(x)
} 
func add(n1 int,n2 int)int{
	var x int=n1+n2
	return x
}
func test()(int,string){  //同時使用多個回傳值
	return 5,"test"
}
func main(){
	msg("你好")  //基本的return操作
	var y =add(3,4)  //透過return，讓加法函式可以做額外運算
	fmt.Println(y*20)
	var i int  
	var s string
	i,s=test()
	fmt.Println(i,s)
}

//指標基礎
package main  
import "fmt"
func main(){
	//建立存放資料的變數
	var x int=5
	fmt.Println("原來的資料:",x)
	//取得記憶體位址:&變數名稱
	fmt.Println("資料的記憶體位址:",&x)
	//存放到指標變數 注意指標資料型態:*資料型態
	var xPtr *int=&x
	fmt.Println("資料的記憶體位址:",xPtr)
	//反解指標變數:*指標變數名稱
	fmt.Println("反解指標回原來的資料:",*xPtr)
}

// Pass by Value 和 Pass by Pointer 函式參數傳遞
//Pass by Value:正常情況，x不相連
package main  
import "fmt"
func add(x int){
	x=x+10
	fmt.Println("add:",x)
}
func main(){
	var x int=10
	add(x)
	fmt.Println("main:",x)  //不同func的x可視為不同的變數，add(x)和main(x)是不同的
}
//pass by pointer:利用指標建立連動性
package main  
import "fmt"
func add(xPtr *int){
	*xPtr=*xPtr+10
	fmt.Println("add:",*xPtr)
}
func main(){
	var x int=10
	add(&x)
	fmt.Println("main:",x) 
	var msg string
	fmt.Println(&msg)  //傳遞字串變數的指標(記憶體位址)
	fmt.Println(msg)
}

//結構Struc基礎
package main  
import "fmt"
type point struct{
	x int
	y int
} 
func main(){
	var p1 point=point{3,4}  //沒特別指定就照順序
	var p2 point=point{y:5,x:2}  //指定可以不照順序
	fmt.Println("(",p1.x,",",p1.y,")")
	fmt.Println("(",p2.x,",",p2.y,")")
}

package main  
import "fmt"
type person struct{
	name string
	age int
} 
func main(){
	var p1 person=person{"Andy",14}
	var p2 person=person{"Amy",12}
	p2.name="小美"  //可以覆蓋資料
	fmt.Println(p1.name,p1.age)
	fmt.Println(p2.name,p2.age)
}

//陣列資料型態
package main  
import "fmt"
func main(){
	var arr1 [4]int  //預設值為0
	var arr2 [2]string=[2]string{"apple","book"}
	var arr3 [3]bool=[3]bool{true,false,true}
	var arr4 [3]int
	arr4[0]=-4
	arr4[1]=2
	arr4[2]=6
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(len(arr1))  //取得陣列長度
	var arr [4]int=[4]int{3,-10,5,8}
	var x int 
	for x=0;x<len(arr);x++{  //利用迴圈逐一列出陣列資料
		fmt.Println(arr[x])
	}
}