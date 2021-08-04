#!bin/zsh
cluster=$(terraform output -raw kubernetes_cluster_name)
gcloud container clusters  get-credentials $cluster --region us-central1

#Dashboard

kubectl apply -f ./recommended.yaml

# Admin user

kubectl apply -f ./admin.yaml

# Hello

# Admin user

kubectl apply -f ./helloword.yml

# Cred

kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep service-controller-token | awk '{print $1}')

# Proxy
kubectl proxy


