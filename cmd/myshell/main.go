package main

import (
	"bufio"
	"fmt"
	"os/exec"
  "os"
  "strings"
  "github.com/mattn/go-shellwords"
  "bytes"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)

func main() {
	// Uncomment this block to pass the first stage
  for i:=0;;i++{
	 fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
  readed,err := bufio.NewReader(os.Stdin).ReadString('\n')
  clear_in :=strings.TrimRight(readed,"\n")
  cmds :=strings.Split(clear_in," ")
  if(err!=nil){
    fmt.Fprint(os.Stdout,"error occured")
  }
  readed=readed[:len(readed)-1]
  check :="exit"
  
  read:=cmds[0]

if (strings.Contains(readed,">") || strings.Contains(readed,"1>")) && !strings.Contains(readed,"2>") && !strings.Contains(readed,"2>>") { 
     var key int
     for i:=0;i<len(cmds);i++{
       if cmds[i]==">" || cmds[i]=="1>" || cmds[i]=="1>>" || cmds[i]==">>"{
         key=i
       }
     }
     var stdout bytes.Buffer
     my_command:=exec.Command(cmds[0],cmds[1:key]...)
      
     file,err := os.OpenFile(cmds[key+1],os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
       if err!=nil {
         fmt.Fprint(os.Stdout,err)
       }
      my_command.Stderr=os.Stderr 
      my_command.Stdout=&stdout


       err=my_command.Run()
    

       strout:=stdout.String()
       strout=strings.ReplaceAll(strout,"\"","")
       _,e:=file.WriteString(strout)

       if e!=nil{
         fmt.Fprint(os.Stdout,e)
       }
       if err!=nil{
         
       }

        continue 
   }else if strings.Contains(readed,"2>") || strings.Contains(readed,"2>>"){
    var key int
     for i:=0;i<len(cmds);i++{
       if cmds[i]=="2>" || cmds[i]=="2>>"{
         key=i
       }
     }

       my_command:=exec.Command(cmds[0],cmds[1:key]...)
       var stdout,stderr bytes.Buffer
     file,err := os.OpenFile(cmds[key+1],os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
       if err!=nil {
         fmt.Fprint(os.Stdout,err)
       }
      
      my_command.Stderr=&stderr 
      my_command.Stdout=&stdout
        err=my_command.Run()
        stdstr:=stdout.String()
        stdstr=strings.ReplaceAll(stdstr,"'","")
        stdstr=strings.ReplaceAll(stdstr,"\"","")

        strerr:=stderr.String()
        strerr=strings.ReplaceAll(strerr,"cant open ","")
        _,e:=file.WriteString(strerr)
        if e!=nil{
          fmt.Fprint(os.Stdout,e)
        }


    fmt.Fprint(os.Stdout,stdstr)
       if err!=nil{
         //do something
       }

        continue 


   }

if  !strings.Contains(readed,"\"") && strings.Contains(readed,"cat") && !strings.Contains(readed,"'") && !strings.Contains(readed,"type"){
    my_command:=exec.Command(cmds[0],cmds[1:]...)
      var stdout bytes.Buffer
       my_command.Stderr=os.Stdout
       my_command.Stdout=&stdout
       err=my_command.Run()
     
       stdstr:=stdout.String()
       stdstr=strings.ReplaceAll(stdstr,"'","")
       stdstr=strings.ReplaceAll(stdstr,"cant open ","")
        fmt.Fprint(os.Stdout,stdstr)

    if err!=nil{
      fmt.Fprint(os.Stdout,err)
    }
    continue
}



  if read==check{
    os.Exit(0)
  }else if read=="echo"{
  
     target:= string(strings.TrimPrefix(readed, "echo "))
    tarArr,err:=shellwords.Parse(readed)
  
    if(string(target[0])=="'"){
  target=strings.ReplaceAll(target,"'","")
    fmt.Fprint(os.Stdout,target+"\n")
    continue
  }

if string(target[0])=="\"" && !(strings.Contains(target,"\\\"") || strings.Contains(target,"\\"+"\\")){
  var store []int
      
    for i:=0; i<len(target); i++{
      if(string(target[i])=="\""){
        store=append(store,i)

      }
    }
    for i:=0;i<len(store); i=i+2{
      fmt.Fprint(os.Stdout,target[store[i]+1:store[i+1]]+" ")
    }
    fmt.Fprint(os.Stdout,"\n")
  

  continue
}


  if strings.Contains(target,"\\") && !strings.Contains(target,"\"") && !strings.Contains(target,"'"){
      target=strings.ReplaceAll(target,"\\","")
       fmt.Fprint(os.Stdout,target+"\n")
       continue
  }else if  strings.Contains(target,"\\") && !strings.Contains(target,"\"") && strings.Contains(target,"'"){
    target=strings.ReplaceAll(target,"'","")
    fmt.Fprint(os.Stdout,target+"\n")
    continue
  }else if strings.Contains(target,"\\") && strings.Contains(target,"\""){
    
    for i:=0; i<len(target) ; i++ {
      if string(target[i])=="\"" {
        continue
      }

      if string(target[i])=="\\" {
        if  string(target[i+1])=="\"" || string(target[i+1])=="\\"{
          fmt.Fprint(os.Stdout,string(target[i+1]))
          i++
        }
        continue
      }
    fmt.Fprint(os.Stdout,string(target[i]))
 }
   
  fmt.Fprint(os.Stdout,"\n")
  continue
   
  }else if string(target[0])=="\"" && !strings.Contains(target,"\\"){
    var store []int
      
    for i:=0; i<len(target); i++{
      if(string(target[i])=="\""){
        store=append(store,i)

      }
    }
    for i:=0;i<len(target); i=i+2{
      fmt.Fprint(os.Stdout,target[store[i]+1:store[i+1]]+" ")
    }
    fmt.Fprint(os.Stdout,"\n")
  }

    if strings.Contains(target,"\""){
    if err!=nil{
      fmt.Fprint(os.Stdout,err)
    } 
    for i,val:=range tarArr{
       if i!=0{
      fmt.Fprint(os.Stdout,val+" ")
    }
    }
    fmt.Fprint(os.Stdout,"\n")
     continue
    }else if strings.Contains(target,"'") {
     target= strings.ReplaceAll(target,"'","")
   }else{
 target =strings.Join(strings.Fields(target)," ")
   }
   fmt.Fprint(os.Stdout, target+"\n")
  continue;
  }else if read=="type"{
    shell_command:=readed[5:]
    fname,err:=exec.LookPath(shell_command)
      
         if shell_command=="echo" || shell_command=="type" || shell_command=="exit" || shell_command=="pwd"{
       fmt.Fprint(os.Stdout,shell_command+" is a shell builtin\n")
     }else{
         if(err!=nil){
         fmt.Fprint(os.Stdout,shell_command+ ": not found\n")
         continue
    }
    fmt.Fprint(os.Stdout,shell_command+ " is "+ fname+ "\n")   
    }
     continue
  }else if readed=="pwd"{
  curr_dir,_:=os.Getwd()
  fmt.Fprint(os.Stdout,curr_dir+"\n")
  continue
  }else if read=="cd" {
    if(cmds[1]=="~"){
      home,err:= os.UserHomeDir()
      if err!=nil {
        fmt.Fprint(os.Stdout,err)
      }
      err = os.Chdir(home)
    
    }else{
    err = os.Chdir(cmds[1])
    if err!=nil {
      fmt.Fprint(os.Stdout, cmds[0]+": "+cmds[1]+": No such file or directory\n")
    }
  }
    continue
  }else if read=="cat" {

    target:=strings.TrimPrefix(readed,"cat ")


 if string(target[0])=="\"" && strings.Contains(readed,"\\") && !(strings.Contains(readed,"\\\"") || strings.Contains(readed,"\\"+"\\")){

   var spacedCat []int
    
   for i:=0;i<len(readed);i++{
      if string(readed[i])=="\""{
      spacedCat=append(spacedCat,i)    
      }
   }
   for i:=0; i<len(spacedCat); i=i+2{
     file, err := os.Open(readed[spacedCat[i]+1:spacedCat[i+1]])
    if err != nil {
 fmt.Fprint(os.Stdout,err,spacedCat[i],spacedCat[i+1])
    }
    defer func() {
        if err = file.Close(); err != nil {
  fmt.Fprint(os.Stdout,err)
}
    }()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {             // internally, it advances token based on sperator
        fmt.Fprint(os.Stdout,scanner.Text())  // token in unicode-char
    }
  }
  fmt.Fprint(os.Stdout,"\n")
continue
 }else if strings.Contains(readed,"\"") && strings.Contains(readed,"'"){
     var store []int 
     for i:=4;i<len(readed);i++{
       var check bool
      if(i!=len(readed)-1){
       check= string(readed[i+1])==" " || string(readed[i-1])==" "
     }else{
       check=true
     }
       if (string(readed[i])=="\"" || string(readed[i])=="'") && check{
         store=append(store,i)
       }
     }
     
     for i:=0; i<len(store); i=i+2{
    if string(readed[store[i]])=="\""{
       matcher:=readed[store[i]+1:store[i+1]]
       var result string
       var contn  bool=false
      for i:=0 ;i<len(matcher);i++{
          if string(matcher[0])=="\"" {
        continue
      }

      if contn{
        result=result+string(matcher[i])
        continue
      }
        if string(matcher[i])=="'"{
          contn=!contn
        }

      if string(matcher[i])=="\\" {
        if  string(matcher[i+1])=="\"" || string(matcher[i+1])=="\\" || string(matcher[i+1])=="'" {
          result=result+string(matcher[i+1])
          
          i++
        }
        continue
      }
      result=result+string(matcher[i])
      
      }

     file, err := os.Open(result)
      if err != nil {
       fmt.Fprint(os.Stdout,err)
     }
      defer func() {
        if err = file.Close(); err != nil {
           fmt.Fprint(os.Stdout,err,result)
          }
         }()

          scanner := bufio.NewScanner(file)

    for scanner.Scan() {             // internally, it advances token based on sperator
        fmt.Fprint(os.Stdout,scanner.Text())  // token in unicode-char
    }
    result=""
   }else {
     file, err := os.Open(readed[store[i]+1:store[i+1]])
      if err != nil {
       fmt.Fprint(os.Stdout,err)
     }
      defer func() {
        if err = file.Close(); err != nil {
           fmt.Fprint(os.Stdout,err)
          }
         }()

          scanner := bufio.NewScanner(file)

    for scanner.Scan() {             // internally, it advances token based on sperator
        fmt.Fprint(os.Stdout,scanner.Text())  // token in unicode-char
    }

   }
 }
   fmt.Fprint(os.Stdout,"\n")
 continue 
 }
    
	p := shellwords.NewParser()
	p.ParseBacktick = true
	splits, _ := p.Parse(readed)

    for i:=1; i<len(splits); i++{
   //  cmds[i]= strings.ReplaceAll(cmds[i],"'","")
   file, err := os.Open(splits[i])
    if err != nil {
 fmt.Fprint(os.Stdout,err,splits[i])
    }
    defer func() {
        if err = file.Close(); err != nil {
  fmt.Fprint(os.Stdout,err)
}
    }()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {             // internally, it advances token based on sperator
        fmt.Fprint(os.Stdout,scanner.Text())  // token in unicode-char
    }

  }
  fmt.Fprint(os.Stdout,"\n")
    continue
   }else if strings.Contains(readed,"exe") && (strings.Contains(readed,"'") || strings.Contains(readed,"\"")){
     var store []int
     
     for i:=0;i<len(readed);i++{
        if string(readed[0])=="\"" && string(readed[i])=="\""{
           store=append(store,i)
        }else if string(readed[0])=="'" && string(readed[i])=="'"{
           store=append(store,i)
        }
      }
         result:=readed[store[len(store)-1]+2:]

 file, err := os.Open(result)
    if err != nil {
 fmt.Fprint(os.Stdout,err)
    }
    defer func() {
        if err = file.Close(); err != nil {
  fmt.Fprint(os.Stdout,err)
}
    }()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {             // internally, it advances token based on sperator
        fmt.Fprint(os.Stdout,scanner.Text())  // token in unicode-char
    }


      fmt.Fprint(os.Stdout,"\n")

      continue
   } 
 
   
command:=exec.Command(cmds[0],cmds[1:]...)
command.Stderr= os.Stderr
command.Stdout=os.Stdout
err =command.Run()
if err!=nil {
  fmt.Printf("%s: command not found\n",cmds[0])
}


  }
}



