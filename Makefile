PLATFORM = linux

build:
	@for dir in `ls handler`; do \
		GOOS=$(PLATFORM) go build -o dist/handler/$$dir ./handler/$$dir; \
	done

local:
	@sam local start-api