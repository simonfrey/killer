# killer

**Works only on linux as it relies on `/proc`**

 A small go util that watches your linux processes and kills them if they are on the list. Helps to prevent procrastination.
 
*Kills only processes which name are at least 50% covered by your forbidden process name. e.g. You forbid `steam`, than `steam` will be killed
 but `steam-powered` not.*
 
 ## Installation
 
 ### Via go get
 `go get -u github.com/simonfrey/killer`
 
 ### Download binary
 
 `wget https://github.com/simonfrey/killer/raw/main/killer`
 
 ## Usage
 
`killer process1 process2`

Example: 
`killer thunderbird chrome steam`
 
 ## License
 
 MIT