#!/usr/bin/env python
# coding=utf-8

import RPi.GPIO as GPIO
import datetime
import SimpleHTTPServer
import SocketServer
import json

class MyRequestHandler(SimpleHTTPServer.SimpleHTTPRequestHandler):
    def do_GET(self):
        data = {}
        if GPIO.input(6) == GPIO.LOW:
            data['occupied'] = 1
            data['time'] = 0
        else:
            data['occupied'] = 0
            data['time'] = 0

        self.send_response(200)

        # Custom headers, if need be
        self.send_header('Content-type', 'application/json')
        self.end_headers()

        # Custom body
        self.wfile.write(json.dumps(data))

def init_GPIO(pin = 6):
    GPIO.setmode(GPIO.BCM)
    GPIO.setup(pin, GPIO.IN, pull_up_down=GPIO.PUD_UP)

def init_webserver(port = 80):
    httpd = SocketServer.TCPServer(('0.0.0.0', port), MyRequestHandler)
    httpd.serve_forever()

if __name__ == "__main__":
    init_GPIO()
    init_webserver()
