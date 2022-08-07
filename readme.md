<!--
                    ***

                ~/README.md
    https://github.com/michalspano/clocrt
                @michalspano

                    ***
-->

# `clocrt` - Count Lines of Code _Redefined Tables_

<!-- GitHub Shields -->
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Contributors][contributors-shield]][contributors-url]
[![Release][release-shield]][release-url]
[![MIT License][license-shield]][license-url]

## Table of Contents

  * [Introduction](#introduction)
  * [Requirements](#requirements)
    * [Additional](#additional)
  * [Installation](#installation)
    * [Build from the source](#build-from-the-source)
  * [Example](#example)
  * [Usage](#usage)
    * [`cloc` docs](#cloc-docs)

## Introduction

A simple, lightweight __CLI__ tool to transform a `cloc` output into a `Markdown` table written in `Go`. 

## Requirements

1. A `unix`-like operating system (with either `bash` or `sh`, etc.);
2. The [`cloc`](https://github.com/AlDanial/cloc) command-line utility;
3. The [`wget`](https://savannah.gnu.org/git/?group=wget) command-line utility.

\*The `wget` command-line utility is used to download the `clocrt` binary. Indeed, `clocrt` is not a fork of `cloc`, and therefore, `cloc` is a direct dependency of `clocrt`.

### Additional

A copy of the `Go` binary is needed to install the `clocrt` binary from the source. The `Go` binary is available [here](https://golang.org/dl/).

## Installation

The __automated installation__ is carried with a single command:

```sh
$ curl -s https://raw.githubusercontent.com/michalspano/clocrt/main/install | sh
```

#### Build from the source

```sh
$ git clone https://github.com/michalspano/clocrt.git && cd clocrt
$ go build -o clocrt src/main.go
$ ./clocrt --version  # validates the installation
```

## Example

Suppose the following __working tree__ with some _dummy_ files:

```text
$ tree ./test

test
├── docs
│   └── foo.md
├── foo.go
└── foo.py
```

Running `cloc` would yield the following output:

```text
$ cloc ./test

       3 text files.
       3 unique files.                              
       0 files ignored.

github.com/AlDanial/cloc v 1.94  T=0.00 s (626.3 files/s, 1878.9 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               1              2              0              5
Markdown                         1              0              0              1
Python                           1              0              0              1
-------------------------------------------------------------------------------
SUM:                             3              2              0              7
-------------------------------------------------------------------------------
```

Which we transform into a `Markdown` table with `clocrt`:

```sh
$ clocrt "`cloc ./test`"
```

### github.com/AlDanial/cloc v 1.94  T=0.00 s (626.3 files/s, 1878.9 lines/s)
| Language | files | blank | comment | code |
| :------: | :---: | :---: | :-----: | :--: |
| Go       | 1     | 2     | 0       | 5    |
| Markdown | 1     | 0     | 0       | 1    |
| Python   | 1     | 0     | 0       | 1    |
| SUM:     | 3     | 2     | 0       | 7    |

\*`Markdown` table syntax: [link](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#tables)

_Voila!_ :confetti_ball:

## Usage

```text
$ clocrt -h 
$ clocrt --help

NAME
        clocrt - count lines of code redefined tables

DESCRIPTION
        A simple, lightweight CLI tool to transform a cloc output into a Markdown table.

USAGE
        clocrt "`[CLOC-PATTERN]`" [OPTIONS]

OPTIONS
        --help, -h                       show the usage
        --version, -v                    fetch and display the current version

        --print, -pr                     print the result to stdout
        --cell-align|-ca=<align>         center|c (default), left|l, right|r
        --output-path|-op=<path>         output file path (default: out.md)
BUGS
         Don't forget to quote the pattern per the example: "`[CLOC-PATTERN]`"

```

#### `cloc` docs

Documentation of the `cloc` command-line utility can be found [here](https://github.com/AlDanial/cloc#options-).

<!-- GitHub Shields -->
[contributors-shield]: https://img.shields.io/github/contributors/michalspano/clocrt.svg?style=for-the-badge
[contributors-url]: https://github.com/michalspano/clocrt/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/michalspano/clocrt.svg?style=for-the-badge
[forks-url]: https://github.com/michalspano/clocrt/network/members
[stars-shield]: https://img.shields.io/github/stars/michalspano/clocrt.svg?style=for-the-badge
[stars-url]: https://github.com/michalspano/clocrt/stargazers
[issues-shield]: https://img.shields.io/github/issues/michalspano/clocrt.svg?style=for-the-badge
[issues-url]: https://github.com/michalspano/clocrt/issues
[license-shield]: https://img.shields.io/github/license/michalspano/clocrt.svg?style=for-the-badge
[license-url]: https://github.com/michalspano/clocrt/blob/main/LICENSE
[release-shield]: https://img.shields.io/github/tag/michalspano/clocrt.svg?style=for-the-badge
[release-url]: https://github.com/michalspano/clocrt/tags/latest