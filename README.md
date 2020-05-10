<p align="center"> <img src="https://github.com/sahad-mk/SH.an/blob/master/screenshots/banner.png" height="50%" width="75%"></p>
<p align="center"><b><i> SH.an - Security Headers Analyzer </i> </b> </p>

## SH.an v1.0

![SH.an](https://img.shields.io/badge/version-1.0-success)   ![go](https://img.shields.io/badge/go-v1.1-important)   ![Tested](https://img.shields.io/badge/Tested%20On-Ubuntu%2018.04-green)  [![Linkedin](https://img.shields.io/badge/Linkedin-/Sahadmk-blue)](https://www.linkedin.com/in/sahadmk)

SH.an is a go script for finding missing secure headers and insecure CSP policies. If you want to add or change the Secure Headers or CSP policies,check the correspondent array or map structure in the script and make the change.

## Prerequisites
          1. go
          
           
           
## Installation
          • clone the SH.an repo,   git clone https://github.com/sahad-mk/SH.an
          
          • To make executable of the script, go build shan.go (it will build the executable 'shan')
          
          
          
## Usage
          ➢  go run shan.go -host <http://hostname or https://hostname>
          
          ➢ 'go run shan.go -h' for help
 
   Examples:
                                                                                                                                             
          • go run shan.go -host https://example.com 
              
          • go run -host http://example.com
                                                         
           
                                                         
             
