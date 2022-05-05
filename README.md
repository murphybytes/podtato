#Podtato  
Test process for pods

This project allows you to build a container that is extended from busybox so it has loads of fun diagnostic tools on it and it also 
has a simple go program that will run a web server on port 8080 that will expose environment variables, and the working directory 
of the running pod.  

### Run the webserver pod
```
kubectl run podtato --port 8080 --image=optumpixeleastus93d9be61.azurecr.io/podtato:0.0.5
```
### Attach to the pod to use various diagnostic tools
```
kubectl exec --tty --stdin podtato -- /bin/sh 
```

