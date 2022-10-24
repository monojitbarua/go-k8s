# go-k8s

## Sample rest-api project build with go-lang and deploy in kubernetes

* After deployment it has <b>ip:4040</b>, <b>ip:4040?name=sample-name</b>, <b>ip:4040?customers</b> rest end point to test


## Installation

* To install and run, please execute below command with kubnestes cluster connection setup

 - This will create the deployment: <b>kubectl apply -f k8s-deployment.yaml</b>
 - This will create the service: <b>kubectl apply -f k8s-service.yaml</b>
 - If you have minikube cluster then excute this command to get the application url: <b>minikube service go-hello-world-service --url</b>
