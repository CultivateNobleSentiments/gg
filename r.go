package main
import (
	"fmt"
	"regexp"
	"strings"
	"os"
	"strconv"
)
type calculater interface{
	calc()
}
type add struct{
	complx1 complex128
	complx2 complex128
}
func (a add)calc(){

	 fmt.Println("结果是：",a.complx1+a.complx2)
}
type sub struct{
	complx1 complex128
	complx2 complex128
}
func (s sub)calc(){

	fmt.Println("结果是：",s.complx1-s.complx2)
}
type multi struct{
	complx1 complex128
	complx2 complex128
}
func (m multi)calc(){

	fmt.Println("结果是：",m.complx1*m.complx2)
}
type divide struct{
	complx1 complex128
	complx2 complex128
}
func (d divide)calc(){

	fmt.Println("结果是：",d.complx1/d.complx2)
}

////////////////////////////////////////////////////////////////////////////////////////////
var inputSave[2] string 
var Symbol int
var SymbolString string
var inputSaveCompare[2] string
var part0,part1 bool
var itAfterProcess[2] complex128
// var a *complex128
///////////////////////////////////////////////////////////////////////////////////////////
func main(){
	
	for{
		input()
		
		Calculate(JudgeSymbol())//1:加 2：减 3：乘 4：除 
		 //fmt.Println(itAfterProcess[0],itAfterProcess[1])
		
	}
	
}
func input(){
	fmt.Println("本程序可以进行运算两个虚数/*-+操作")
	
	judgeInput(0)
	judgeInput(1)
	
}
func judgeInput(i int){
	for{
		fmt.Println("第",i+1,"个数（包括实部虚部，没虚部可以不写，也可以只写虚部）：")
		fmt.Scanln(&inputSave[i])
		part0=false 
		part1=false    //MustCompile
		line1:=regexp.MustCompile(`(?:[+\-]?(?:0(?:\.\d+)?)|(?:[1-9]\d*(?:\.\d+)?))(?:[+\-](?:(?:0(?:\.\d+)?)|(?:[1-9]\d*(?:\.\d+)?))?i)?`)//实数或由实数构成的虚数***part0***
		line2:=regexp.MustCompile(`[-]?(?:(?:0(?:\.\d+)?)|(?:[1-9]\d*(?:\.\d+)?))?i`)//纯虚数***part1***
		inputSave[i]=strings.Replace(inputSave[i]," ","",-1)
		if inputSave[i]==""{
			fmt.Println("请输入正确格式")
			continue
		}
/////////////////////////////////////////////////////////////////////////////////////
		inputSaveCompare[0]=line1.FindString(inputSave[i])
		if inputSaveCompare[0]==inputSave[i]{
			part0=true
		}
		inputSaveCompare[1]=line2.FindString(inputSave[i])
		if inputSaveCompare[1]==inputSave[i]{
			part1=true
		}
		if part0||part1{
			fmt.Println("录入成功！")
			break
		}else {
			fmt.Println("输入格式错误，请重新输入")
			continue
		}
	}
	processData(i)
}
func processData(i int){
	if part0{
		findI:=strings.Index(inputSave[i],"i")
		// fmt.Println("1bufen")
		if findI==-1 {
			s1,_:=strconv.ParseFloat(inputSave[i],64)
			itAfterProcess[i]=complex(s1,0)
			fmt.Println("你的输入格式为**实数**")

		}else{
			
			strArrTwo := strings.FieldsFunc(inputSave[i], checkSpiltRune)

			s2,_:=strconv.ParseFloat(strArrTwo[0],64)
			
			var s3 float64
			strArrTwo[1]=strings.Replace(strArrTwo[1],"i","",-1)
			if strArrTwo[1]==""{
				s3=1
			}else{
			s3,_=strconv.ParseFloat(strArrTwo[1],64)
			}
			// fmt.Println(strArrTwo[1])
			itAfterProcess[i]=complex(s2,s3)
			fmt.Println("你的输入格式为**a+bi**")

		}
		
	}else if part1{
		inputSave[i]=strings.Replace(inputSave[i],"i","",-1)
		var s4 float64
		if inputSave[i]==""{
			s4=1
		}else{
		s4,_=strconv.ParseFloat(inputSave[i],64)
		}
		itAfterProcess[i]=complex(0,s4)
		fmt.Println("你的输入格式为**纯虚数**")

	}else{
		fmt.Println("恭喜你发现了一处bug，没法玩了，直接退吧！")
		os.Exit(1)
	}
}
func checkSpiltRune(r rune)bool{
	if r=='-'||r=='+'{
		return true
	}
	return false
}
func JudgeSymbol()int{
	for {
		fmt.Println("请输入对应数字来选择操作")
		fmt.Println("1:加 2：减 3：乘 4：除 ")
		fmt.Scanln(&SymbolString)
		if len(SymbolString)==1{
		fff:=[]byte(SymbolString)
		var gt byte =fff[0]
		if gt>=49&&gt<=53{
			switch gt {
			case 49:return 1
			case 50:return 2
			case 51:return 3
			case 52:return 4
			case 53:return 5
			default:fmt.Println("李在赣神魔？？？？？？")
			}	
		}else {
			fmt.Println("李在赣神魔？？？？？？")
		continue
		}	
	}else{
		fmt.Println("李在赣神魔？？？？？？")
		continue
		}
	}
}
func Calculate(t int){
	g:=add{itAfterProcess[0],itAfterProcess[1]}
	gg:=sub{itAfterProcess[0],itAfterProcess[1]}
	ggg:=multi{itAfterProcess[0],itAfterProcess[1]}
	gggg:=divide{itAfterProcess[0],itAfterProcess[1]}
	calculate:=[]calculater{&g,&gg,&ggg,&gggg}
	
	calculate[t-1].calc()

}