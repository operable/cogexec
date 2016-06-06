EXE                = cogexec
BUILD_DIR          = _build
SOURCES            = main.go messages/execution.go

.PHONY: all clean docker

all: $(BUILD_DIR)/$(EXE)

clean:
	rm -rf $(BUILD_DIR)

docker: clean $(SOURCES)
	GOOS=linux GOARCH=amd64 make all
	docker build -t operable/cogexec .

$(BUILD_DIR)/$(EXE): $(BUILD_DIR) $(SOURCES)
	go build -o $(BUILD_DIR)/$(EXE) github.com/operable/cogexec

$(BUILD_DIR):
	@mkdir -p $(BUILD_DIR)
