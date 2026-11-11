MODULE := github.com/nathanverse/mygo
PROTO_FILES := $(shell find proto -name '*.proto')

apis: $(PROTO_FILES)
	rm -rf apis
	protoc \
		--go_out=. --go_opt=module=$(MODULE) \
		--go-grpc_out=. --go-grpc_opt=module=$(MODULE) \
		$(PROTO_FILES)
	touch apis

.PHONY: apis-clean
apis-clean:
	rm -rf apis
