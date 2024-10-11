CNPG_VERSION ?= "1.24.0"
POSTGRES_VERSION ?= "17"

# common env setup
export POSTGRES_VERSION

build:
	go mod tidy && \
	go generate ./... && \
	go build && \
	go mod tidy

minikube-cleanup:
	@if minikube status > /dev/null 2>&1; then \
		minikube stop; \
		minikube delete; \
	fi

minikube-setup: minikube-cleanup
	minikube start --cpus 3 --memory 4096

cnpg-controller-setup:
	kubectl apply --server-side -f \
		https://raw.githubusercontent.com/cloudnative-pg/cloudnative-pg/main/releases/cnpg-$(CNPG_VERSION).yaml
	@echo -e "\n\e[0;32mInstalled CNPG controller on the cluster :)\n\e[0m"
	sleep 30
	kubectl get deployment -n cnpg-system cnpg-controller-manager

cnpg-cluster-setup: cnpg-controller-setup
	kubectl create ns postgres-cluster
	kubectl apply -f acceptance/cnpg-cluster.yaml
	@echo -e "\n\e[0;32mCreated CNPG cluster on the cluster :)\n\e[0m"
	sleep 30
	kubectl get cluster -n postgres-cluster

pg-setup: minikube-setup cnpg-cluster-setup
	
acceptance-test: pg-setup
	docker-compose -f acceptance/docker-compose.yaml up