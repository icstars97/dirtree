package main

import (

	"fmt"
	"os"
	"bufio"
	"sort"
	"path"
	"strings"
)

type input struct{
	path string;
	dironly bool;
}
func read() *input {
	input:=new(input);
	var s, sdironly string;
	in := bufio.NewScanner(os.Stdin);
	in.Scan();s=in.Text();
	l:=len(s);
	sdironly=s[l-3:l];
	if sdironly==" -f"{
		input.dironly=true;
		input.path=strings.Trim(s, " -f");
		//fmt.Println(input.path);
		//fmt.Println(sdironly);

	}else{
		input.dironly=false;
		input.path=s;
		//fmt.Println(input.path);
		//fmt.Println(sdironly);
	}
	return input;


}

type count struct {
	nd int;
	nf int;

}

func (count *count) dircount(dirpath string) (*count){
	f,_:=os.Open(dirpath);
	prop,_:=f.Stat();
	if prop.IsDir(){
		count.nd++;
	}else{
		count.nf++;
	}
	f.Close();

	return count;
}
func dirnames(dirpath string) []string{
	f,_:=os.Open(dirpath);
	names,_:=f.Readdirnames(-1);
	sort.Strings(names);
	f.Close();
	return names;

}
func draw(count *count, dirpath string, prefix string,dironly bool){
	names:=dirnames(dirpath);
	for index,name := range names {
		isfile:=false;
		if name[0] == '.' {
			continue;
		}
		subpath:=path.Join(dirpath,name)
		f,_:=os.Open(subpath);
		prop,_:=f.Stat();
		if !prop.IsDir() {
			isfile=true;
		}
		count.dircount(subpath);


		if index == len(names)-1 {
			if isfile&&dironly {
				draw(count, subpath, prefix+"\t",dironly);
			}else{
				fmt.Println(prefix+"└──", name)
				draw(count, subpath, prefix+"\t",dironly)
			}
			
		} else {
			if isfile&&dironly {
				draw(count, subpath, prefix+"│\t",dironly);
			}else{
				fmt.Println(prefix+"├──", name)
				draw(count, subpath, prefix+"│\t",dironly);
			}

		}
	}
	}
func (count *count) out(dironly bool){
	if dironly {
		fmt.Println("количество папок: ",count.nd);
	}else{
		fmt.Println("количество папок: ",count.nd,"  количество файлов: ",count.nf,);
	}

}




func main(){
	input:=read();
	count:=new(count);
	draw(count,input.path,"",input.dironly);
	//count.out(input.dironly);



	}