VERSION = $(shell grep 'version =' utils/version.go | sed -E 's/.*"(.+)"$$/\1/')

all: 
	go build -buildmode=plugin -o ./plugins/eng/eng.so ./plugins/eng/greeter.go
	go build -buildmode=plugin -o ./plugins/chi/chi.so ./plugins/chi/greeter.go
	go build -o pluginexample

version:
	@echo $(VERSION)

.PTHONY: all build version