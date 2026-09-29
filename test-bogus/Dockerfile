# Deliberately misconfigured Dockerfile. Exists only to verify reviewdog
# posts inline PR comments for trivy-config. Not part of the real esignet
# deployment.
FROM alpine:3.20
RUN apk add --no-cache curl
ADD https://example.com/some-script.sh /tmp/some-script.sh
RUN chmod 777 /tmp/some-script.sh
ENTRYPOINT ["/tmp/some-script.sh"]
