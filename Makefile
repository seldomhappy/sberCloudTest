.PHONY: test
.SILENT:

help:   ## show this help
	@echo 'make run - сборка и выполнение приложения'
	@echo 'make build - сборка приложения в дирректорию ./.bin'
	@echo 'make test - запуск тестов'

run: build
	echo " - Выполнение приложения"
	./.bin/shortString

build:
	echo " - Бинарный файл собран в ./.bin"
	go build -o ./.bin/shortString ./cmd/main.go

test:
	echo " - Результат теста:"
	go test -v ./test

