FROM golang:latest
MAINTAINER Anusha
WORKDIR ~/mon1
ADD . .
EXPOSE 80
RUN go get gopkg.in/mgo.v2
CMD ["go", "run", "myc.go"]
#ADD . ./
#RUN go get gopkg.in/mgo.v2
#RUN go build -o myc
#ENV PORT 80
#EXPOSE 80
#CMD [“/myc”]
#RUN mkdir -p /myc

#WORKDIR /myc

#ADD . /myc
#RUN go get gopkg.in/mgo.v2
#RUN cd /myc; go build -o myc.go; 
#RUN go build ./myc.go
#ENTRYPOINT ["./myc"]
#CMD ["./myc"]
#EXPOSE 80
