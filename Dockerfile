FROM gliderlabs/alpine:latest

MAINTAINER Kevin Smith <kevin@operable.io>

RUN mkdir -p /operable/cogexec/bin

COPY _build/cogexec /operable/cogexec/bin

RUN chmod +x /operable/cogexec/bin/cogexec

VOLUME /operable/cogexec