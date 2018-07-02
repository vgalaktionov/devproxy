#!/bin/sh
cmd="go run devproxy.go"
osascript -e "do shell script \"$cmd\" with prompt \"Devproxy requires admin privileges to:\n \t - bind to port 80\n \t - append /etc/hosts\" with administrator privileges"
