FROM ARG_FROM

MAINTAINER vishbin@github.com

ADD bin/ARG_ARCH/ARG_BIN /ARG_BIN

USER nobody:nobody
ENTRYPOINT ["/ARG_BIN"]
ENV SERVICE_NAME service-queue-go
