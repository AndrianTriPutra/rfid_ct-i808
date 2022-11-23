#!/bin/bash

sleep 2
socat TCP-LISTEN:8080,fork,reuseaddr FILE:/dev/ttyUSB0,b57600,raw