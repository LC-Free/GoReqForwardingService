# GoReqForwardingService
Request forwarding service leveraging go and docker

Some public APIs constrain requests sent from frontend applications ( for example Quandl.com will block requests to their public APIs made from a frontend application ).

In order to work around this contraint this service forwards the requests sent to a certain localhost port to the desired public API as specified in the config.yaml's baseUrl parameter.

The script can be run by simply building and running the main.go file ( `go build` followed by `./fwdservice` in the fwdservice directory ), 
or using the Dockerfile found at the root. a simple usage example would be running `docker build -t fwdservice .` in the root directory of this repo,
then `docker run -d -p 8000:8000 fwdservice` to run that build container. 

The fwdservice tag applied to the build image and used in the run step is just an example, the image can be tagged with another name to the same effect, or not tagged at all.
Additionally the port mappings can be changed, but those would have to be changed in both the main.go file as well as the Dockerfile ( at least for now, port mappings may be added to the config.yaml at a later date along with a possible shell script to set it in the Dockerfile ).
