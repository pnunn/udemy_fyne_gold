BINARY_NAME=GoldWatcher
APP_NAME=GoldWatcher
VERSION=1.0.1
BUILD_NO=1
APP_ID=au.com.dispatch.market.goldwatcher

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	fyne package -appID ${APP_ID} -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release

## run: builds and runs the application
run:
	env DB_PATH="./sql.db" go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -v ./...