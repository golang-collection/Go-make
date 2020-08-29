SHELL := /bin/bash
BASEDIR = $(shell pwd)

ldflags="-w -X versionDir.gitTag={gitTag} -X versionDir.buildDate={buildDate} -X versionDir.gitCommit={gitCommit} -X versionDir.gitTreeState={gitTreeState}"

all: gotool
	@go build -v -ldflags ${ldflags} .
gotool:
	gofmt -w .

