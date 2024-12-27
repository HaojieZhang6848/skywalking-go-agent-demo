build:
	@echo "Makefile path: $(CURDIR)/Makefile"
	@cd serviceA && go build -toolexec="$(CURDIR)/sw-go-agent" -a -o "$(CURDIR)/bin/serviceA" .
	@cd serviceB && go build -toolexec="$(CURDIR)/sw-go-agent" -a -o "$(CURDIR)/bin/serviceB" .

run-serviceA:
	@echo "Running serviceA"
	@SW_AGENT_NAME=serviceA ./bin/serviceA

run-serviceB:
	@echo "Running serviceB"
	@SW_AGENT_NAME=serviceB ./bin/serviceB