FROM registry.access.redhat.com/ubi8/ubi:latest
ADD dummy-web-server /bin
ENTRYPOINT "/bin/dummy-web-server"