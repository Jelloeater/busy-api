# busy-api

[![Test](https://github.com/Jelloeater/busy-api/actions/workflows/test.yml/badge.svg)](https://github.com/Jelloeater/busy-api/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jelloeater/busy-api)](https://goreportcard.com/report/github.com/jelloeater/busy-api)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/jelloeater/busy-api)

![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/jelloeater/busy-api/total)
![GitHub Release](https://img.shields.io/github/v/release/jelloeater/busy-api)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jelloeater/busy-api)
![GitHub Release Date](https://img.shields.io/github/release-date/jelloeater/busy-api)

**<https://libraries.io/go/github.com%2FJelloeater%2Fstampy>**

![Stampy](stampy.gif)

## Install

### Binary (eget)

Use <https://github.com/zyedidia/eget>

```shell
curl https://zyedidia.github.io/eget.sh | sh
sudo mv eget /usr/local/bin
sudo eget jelloeater/busy-api --to /usr/local/bin --asset=tar.gz
```

### Via Go

```shell
go install github.com/Jelloeater/busy-api@latest
```

## Build

Run

```shell
go build -o ./build .
./build/busy-api
```

or use <https://taskfile.dev/>
