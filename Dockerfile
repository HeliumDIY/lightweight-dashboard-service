# Dockerfile
FROM alpine
COPY lightweight-dashboard \
    /lightweight-dashboard
ENTRYPOINT ["/lightweight-dashboard"]