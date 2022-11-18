#!/bin/bash

sleep 3
stty -F /dev/ttyS0 9600 cs7 -cstopb -parenb
sleep 2
socat TCP-LISTEN:8080,fork,reuseaddr FILE:/dev/ttyS0,b57600,raw