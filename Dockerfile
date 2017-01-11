FROM golang:onbuild
RUN mkdir /overlay
ADD . /overlay/
WORKDIR /overlay
RUN go build -o o .
CMD ["/overlay/o"]