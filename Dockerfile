#################
# Base image
#################
FROM alpine:3.12 as postgresweekly-base

USER root

RUN addgroup -g 10001 postgresweekly && \
    adduser --disabled-password --system --gecos "" --home "/home/postgresweekly" --shell "/sbin/nologin" --uid 10001 postgresweekly && \
    mkdir -p "/home/postgresweekly" && \
    chown postgresweekly:0 /home/postgresweekly && \
    chmod g=u /home/postgresweekly && \
    chmod g=u /etc/passwd

ENV USER=postgresweekly
USER 10001
WORKDIR /home/postgresweekly

#################
# Builder image
#################
FROM golang:1.17-alpine AS postgresweekly-builder
RUN apk add --update --no-cache alpine-sdk
WORKDIR /app
COPY . .
RUN make build

#################
# Final image
#################
FROM postgresweekly-base

COPY --from=postgresweekly-builder /app/bin/postgresweekly /usr/local/bin

# Command to run the executable
ENTRYPOINT ["postgresweekly"]
