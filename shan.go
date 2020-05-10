package main
import (
	"fmt"
	"net/http"
	"net/url"
	"flag"
	"strings"
	"os"
	)

func main() {
	banner:= `
                    ____   _   _                 
                   / ___| | | | |    __ _  _ __  
                   \___ \ | |_| |   / _' || '_ \ 
                    ___) ||  _  | _| (_| || | | |
                   |____/ |_| |_|(_)\__'_||_| |_|
                   `
	bannerTxt:=`[Security Headers Analyzer]`

	fmt.Print("\033[34;1m",banner,"\033[0m")
	fmt.Println("\033[1;32m",bannerTxt,"\033[0m")
	fmt.Println("\t\t\t\t\t\033[38;5;31mCoded\033[0m\033[1;37;41m©Sahad.Mk\033[0m")
	fmt.Println("\n")
	missHead()

}

//Function for checking Secure Headers
func missHead() {

	hostarg := flag.String("host", "", "option -host is missing! \nuse -h/-help for Help message")

	flag.Usage= func(){
		fmt.Println("\n      Usage: go run shan.go -host <http://hostname or https://hostname>\n")
		fmt.Println("      Example:  go run shan.go -host http://example.com \n")
		fmt.Println("                go run shan.go -host http://example.com \n")
		}
	flag.Parse()

	 if *hostarg=="" {

		flag.PrintDefaults()
		os.Exit(1)
	 }


	urlRet :=  *hostarg
	urlHost,_ := url.Parse(urlRet)

	if urlHost.Scheme==""||urlHost.Host==""{
		    fmt.Println("\033[0;31mError:\033[0m \033[1;37m Enter a valid url!\033[0m\n")
            os.Exit(1)
			}

	res, err := http.Get(urlRet)

	if err != nil {
		fmt.Println("\033[0;31m Error:\033[0m \033[1;37m Check your hostname!",err,"\033[0m")
		return
	}
	  fmt.Println("\n \033[45;1m Server Info  \033[0m")

	  fmt.Println("\n \033[38;5;130m Host name :", urlHost.Host, "\n  Server    :", res.Header.Get("Server"), "\n  Date      :", res.Header.Get("Date"),"\033[0m\n")

	//Checking for missing secure headers

	   HeaderArray := []string{"Content-Security-Policy", "X-Frame-Options", "X-Xss-Protection", "X-Content-Type-Options", "Strict-Transport-Security", "Referrer-Policy","Feature-Policy"}

	   var missHeader []string

	   fmt.Println("\n \033[45;1m                          Analyzing Secure Headers                      \033[0m\n")

	for i := 0; i < len(HeaderArray); i++ {
		if res.Header.Get(HeaderArray[i]) == "" {
			missHeader = append(missHeader, HeaderArray[i])
			fmt.Println("\033[0;31m ! Missing Security Header:\033[0m","\033[1;37m",HeaderArray[i],"\033[0m")
		}
	}

	 //Description
	  policy_map:=map[string]string{
		"Content-Security-Policy"  :"  : It is an effective measure to protect your site from XSS attacks.",
		"X-Frame-Options"          :"          : It defends your site from attacks like Click jacking.",
		"X-Xss-Protection"         :"         : It can prevent some level of XSS  attacks.",
		"X-Content-Type-Options"   :"   : It stops a browser from trying to MIME-sniff the content type.",
		"Strict-Transport-Security":": It ensure all communication from a browser is sent over HTTPS.",
		"Referrer-Policy"          :"          : It controls how much referrer information should be included with requests.",
		"Feature-Policy"           :"           : It allows a site to control which features and APIs can be used in the browser."}

	  if missHeader!=nil{
		fmt.Println("\n \033[4;94m Additional Info:\033[0m")
		 for i:=range missHeader{
			header_desc, _ := policy_map[missHeader[i]]
			fmt.Println( "\n",missHeader[i],"\033[38;5;251m",header_desc,"\033[0m")
		   }
	  }

	 if missHeader == nil {
		fmt.Println("\033[38;5;49m Status:\033[0m \033[1;32m All security Headers are present! \033[0m\n")
	 } else
	   {
		if len(missHeader) == len(HeaderArray) {
			//fmt.Println("length:", len(missHeader), len(HeaderArray))
			fmt.Println("\033[38;5;49m \n Status:\033[0m \033[38;5;196m All security Headers are missing!\033[0m\n")
		}
	  }

	//checking for insecure csp
		csp := res.Header.Get(HeaderArray[0])
		if res.Header.Get("Content-Security-Policy") != "" {
			scpCheck(csp)
		}

}

//Function for Evaluating CSP Policy

func scpCheck(cspVal string){
              var dirArr[] string
              var j int
              var insecPol[] string

              spl:= strings.Split(cspVal,";")
			  fmt.Println("\n \033[45;1m                         Evaluating CSP Policy                         \033[0m\n")
			for i:=range spl {
				cp_value := spl[i]
				cspField := strings.Fields(cp_value)
				cspDir := cspField[0]
				dirArr = append(dirArr,cspDir)
				cspPolicy := []string{"unsafe-inline", "unsafe-eval"," *"}

				for j=0;j<len(cspPolicy);j++ {
					        if strings.Contains(cp_value, cspPolicy[j]) {
							fmt.Println("\033[38;5;196m ! Insecure policy found:\033[0m","\033[1;37m",cspDir,"contains", cspPolicy[j],"\033[0m")
							insecPol = append(insecPol,cspPolicy[j])
						     }
				}

			}

			//Description for CSP Policies
			     csp_map:=map[string]string{

				           "unsafe-inline" :": allows the execution of unsafe in-page scripts and event handlers",
		                    "unsafe-eval"  :"  : allows the execution of code injected into DOM APIs such as evals",
		                      " *"         : "           : allows everything without restrictions"}
			      if insecPol!=nil{
				     fmt.Println("\n \033[4;94m Additional Info:\033[0m\n")
				     //fmt.Println("insecpol array",insecPol)
				       for i:= range insecPol {
				     	 policy_desc,_:=csp_map[insecPol[i]]

						 fmt.Println( "",insecPol[i],"\033[38;5;251m",policy_desc,"\n\033[0m")
					    }
			      }

			      if insecPol==nil{
				            fmt.Println("\033[38;5;49m Status:\033[0m \033[1;32m CSP is secure!\033[0m \n")
			                 }

}








