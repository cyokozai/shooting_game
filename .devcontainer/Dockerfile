# TinyGoの公式イメージを使用
FROM tinygo/tinygo:latest

WORKDIR /src/app

RUN apt-get update && apt-get install -y git usbutils

ENV PATH=$PATH:/usr/local/go/bin