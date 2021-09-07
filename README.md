# properties

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/obity/properties?color=peru)](https://github.com/obity/properties/releases/latest)
[![Released API docs](https://img.shields.io/badge/godoc-reference-blue)](github.com/obity/properties)
[![Build](https://img.shields.io/github/workflow/status/obity/properties/Go.svg)](#)
[![Go Report Card](https://goreportcard.com/badge/github.com/obity/properties?color=pink)](https://goreportcard.com/report/github.com/obity/properties)
[![Lines of code](https://img.shields.io/tokei/lines/github/obity/properties.svg?color=beige)](#)
[![Downloads of releases](https://img.shields.io/github/downloads/obity/properties/total.svg?color=lavender)](https://github.com/obity/properties/releases/latest)
[![Languages](https://img.shields.io/github/languages/top/obity/properties.svg?color=yellow)](#)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/obity/properties)](#)

Properties 是一个用于读写属性文件的Go库。Properties 可保存在流中或从流中加载。
属性列表中每个键及其对应值都是一个字符串。它是线程安全的：多个线程可以共享单个 Properties 对象而无需进行外部同步。

Properties is a Go library for reading and writing properties files.
The Properties model represents a persistent set of properties. The Properties can be saved to a stream or loaded from a stream. Each key and its corresponding value in the property list is a string.

It is thread-safe: multiple threads can share a single Properties object without the need for external synchronization.

# Doc

See this document at [GoDoc](https://pkg.go.dev/github.com/obity/properties)

# Install
    
    go get -u github.com/obity/properties@latest

