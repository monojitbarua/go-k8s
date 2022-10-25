# go-k8s

## Sample rest-api project build with go-lang and deploy in kubernetes

* After deployment it has <b>ip:4040</b>, <b>ip:4040?name=sample-name</b>, <b>ip:4040?customers</b> rest end point to test

## Local image build and push to own registry 
* docker build -t go-k8s .
* docker tag <docker id> monojit2012/go-k8s:<tag-version>
* docker push monojit2012/go-k8s:<tag-version>


## Installation

* To install and run, please execute below command with kubnestes cluster connection setup

 - This will create the deployment: <b>kubectl apply -f k8s-deployment.yaml</b>
 - This will create the service: <b>kubectl apply -f k8s-service.yaml</b>
 - If you have minikube cluster then excute this command to get the application url: <b>minikube service go-k8s --url</b>
