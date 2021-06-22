
# COMD

A command line tool for logging serial/com port traffic on linux. 


## Features

- Logs to a file
- configurable baud rate
- small footprint
  
## Usage

Below shows the available flags (You will mainly want to focus on `-target`, Unless you
have a very specific use-case and/or know what you're doing `-baud` default of 9600 should
work almost across universally).

```
$ comd [-target=/dev/ttyAMA0] [-logfile=./comd.log] [-baud=9600]
```

## Installation 

Build the binary with go build
```bash 
  git clone https://github.com/joshburnsxyz/comd
  go mod tidy
  go build
  ./comd
```
  
## License

[MIT](https://choosealicense.com/licenses/mit/)

  
