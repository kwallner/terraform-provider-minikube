.PHONY: init
init:
	go mod tidy

.PHONY: configure
configure: build set-rc
	go generate ./...

.PHONY: clean
clean:
	rm bin/* || true
	rm tests/terraform.tfstate || true
	rm tests/terraform.tfstate.backup || true
	minikube delete -p terraform-provider-minikube
	minikube delete -p terraform-provider-minikube-acc

.PHONY: test
test:
	go clean -testcache
	go test ./... 

.PHONY: acceptance
acceptance:
	TF_ACC=true go test ./minikube -run "TestClusterCreation" -v -p 1 --timeout 10m

.PHONY: test-stack
test-stack: set-rc
	terraform -chdir=examples/resource/minikube_cluster init || true 
	terraform -chdir=examples/resource/minikube_cluster apply --auto-approve
	terraform -chdir=tests destroy --auto-approve

.PHONY: build
build:
	go build -o bin/terraform-provider-minikube

.PHONY: set-rc
set-rc: build
	touch ~/.terraformrc
	echo "provider_installation { dev_overrides { \"hashicorp/minikube\" = \"$$(pwd)/bin\" } direct {}}" > ~/.terraformrc 

SED_FLAGS := -i
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
		SED_FLAGS += -e 
endif
ifeq ($(UNAME_S),Darwin)
		SED_FLAGS += ''
endif
.PHONY: set-version
set-version:
	$(eval VERSION := $(shell cat minikube/version/version.go | grep Version | tr -d "[:space:]" | sed 's/Version\="//g' | sed 's/"\/\/.*//g'))
	sed $(SED_FLAGS) 's/VERSION=".*"/VERSION="$(VERSION)"/g' bootstrap/install-driver.sh
