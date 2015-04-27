# Robins

Robins is a simple TCP proxy that proxies many-ports on the local machine to a
host given from a list of available hosts. The picking of the host is random and
will try another host if a connection fails.

## Installation

`go get github.com/jwaldrip/robins`

## Usage

`robins host1.com,host2.com 3000,3001`

## Docker Usage

`docker run -it jwaldrip/robins host1.com,host2.com 3000,3001`