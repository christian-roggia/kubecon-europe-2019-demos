FROM debian:9.8

# Install required packages
RUN apt-get update && apt-get install -y curl build-essential

# Set Bazel version
ENV BAZEL_VERSION="0.25.2" \
    BAZEL_SHA256SUM="c6b9f3a3350419521461c6d9bdbfd09b4f745af564e1728e2fbf9d4d343f93b6  bazel-0.25.2-linux-x86_64"

# Download Bazel
RUN curl -sLO https://github.com/bazelbuild/bazel/releases/download/${BAZEL_VERSION}/bazel-${BAZEL_VERSION}-linux-x86_64 && \
    echo ${BAZEL_SHA256SUM} | sha256sum --check && \ 
    chmod +x bazel-${BAZEL_VERSION}-linux-x86_64 && \
    mv bazel-${BAZEL_VERSION}-linux-x86_64 /usr/local/bin/bazel

# Set the default entrypoint and cmd
CMD [ "version" ]
ENTRYPOINT [ "bazel" ]