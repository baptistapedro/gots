FROM golang:1.19.1-buster as go-target
RUN apt-get update && apt-get install -y wget
ADD . /gots
WORKDIR /gots
RUN mkdir ./testdata
RUN wget https://tsduck.io/streams/test-patterns/test-1packet-01.ts
RUN wget https://tsduck.io/streams/test-patterns/test-2packets-02-03.ts
RUN wget https://tsduck.io/streams/test-patterns/test-3packets-04-05-06.ts
RUN mv *.ts ./testdata/
WORKDIR /gots/cli/
RUN go build

FROM golang:1.19.1-buster
COPY --from=go-target /gots/cli/cli /
COPY --from=go-target /gots/testdata/*.ts /testsuite/

ENTRYPOINT []
CMD ["/cli", "-f", "@@"]
