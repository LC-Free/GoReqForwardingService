FROM golang
VOLUME /var/log/service
COPY fwdservice /fwdservice
WORKDIR /fwdservice
RUN go get gopkg.in/yaml.v2
RUN go build
ENTRYPOINT ./fwdservice
EXPOSE 8000

