APP_NAME=wifi-card

default: build


build-ui: clean-ui
	@cd ui && yarn && yarn build && cd -

clean-ui:
	@rm -rf ui/build

build: clean
	@go build -o bin/$(APP_NAME)

clean:
	@rm -rf bin

run:
	@./bin/$(APP_NAME)

