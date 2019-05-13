# KubeCon + CloudNativeCon Europe 2019
> This repository contains the demos for the presentation `Reproducible development and deployment with Bazel and Telepresence`.

## Reproducible builds with Bazel
The first two commands from this demo are **Golang-oriented**. They have been included to show how Bazel creates a `dependency graph`.

- The first command will generate default `Golang BUILD files` for all the packages in the project.
- The second command will import the dependencies from the `Golang package manager`.
- The third command will execute the build of the application itself.

```sh
bazel run :gazelle
bazel run :gazelle -- update-repos -from_file=go.mod
bazel build //cmd/...
```

## Interactive development with Telepresence
- The first command will spawn a **new deployment** inside the remote cluster.
- The second command will **swap an existing deployment** ant take it over.
- The third command will spawn a new deployment via traffic forwarding from the `debian` Docker image.

```sh
telepresence
telepresence --swap-deployment demo-app
telepresence --docker-run -it debian
```

## Containerized development environment
- The first command will **prune** the local Docker instance from **all the resources**. This will ensure a clean workspace.
- The second command will build the `Bazel Docker image` and tag it with the version `0.25.2`.
- The third command will execute the image with no arguments (by default the version is printed).

```sh
docker system prune --all
docker build -t kubecon-bazel:0.25.2 -f build/Dockerfile .
docker run kubecon-bazel:0.25.2
```

## Run local containerized applications in the Cloud
- The first command will load the `Bazel distroless image` into the local Docker registry.
- The second command will open a **two-way communication proxy** between the local container and the remote cluster.

```sh
bazel run :register-image
telepresence --docker-run k8s.io/kubecon-bazel-telepresence:latest
```