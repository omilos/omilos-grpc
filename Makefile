# Protobuf definitions
PROTO_FILES := $(shell find . \( -path "./node_modules" -o -path "./languages" -o -path "./specification" \) -prune -o -type f -name '*.proto' -print)
# Protobuf Go files
PROTO_GEN_GO_FILES = $(patsubst %.proto, %.pb.go, $(PROTO_FILES))
# Protobuf javascript files
PROTO_GEN_JS_FILES = $(patsubst %.proto, %_pb.js, $(PROTO_FILES))
# Protobuf python files
PROTO_GEN_PY_FILES = $(patsubst %.proto, %_pb2.py, $(PROTO_FILES))

# Protobuf generator
PROTO_GO_MAKER := protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

PROTO_JS_MAKER := npx protoc --ts_out=. --ts_opt long_type_number

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_UNIX=$(BINARY_NAME)_unix

PLUGIN_GEN_FILES = $(patsubst plugins/%.go, obj/%.so, $(wildcard plugins/*.go))

.PHONY: all build plugins test clean run trader manager

all: build

build: golang javascript

golang: $(PROTO_GEN_GO_FILES)

javascript: $(PROTO_GEN_JS_FILES)

%.pb.go: %.proto
	$(PROTO_GO_MAKER) --proto_path=. --proto_path=$(GOPATH)/src ./*.proto

%_pb.js: %.proto
	$(PROTO_JS_MAKER) --proto_path=. ./*.proto

%_pb2.py: %.proto
	$(PROTO_PY_MAKER) --proto_path=.


# }}} Protobuf end

# {{{ Cleanup
clean: protoclean

protoclean:
	rm -rf $(PROTO_GEN_GO_FILES) $(PROTO_GEN_PY_FILES) $(PROTO_GEN_JS_FILES)
# }}} Cleanup end
