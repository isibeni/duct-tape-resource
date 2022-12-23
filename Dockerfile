# Copyright © 2022 The Homeport Team
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

FROM golang:1.19 as bootstrap
WORKDIR /go/src/github.com/homeport/duct-tape-resource
COPY . .

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
RUN --mount=type=cache,target=/root/.cache/go-build \
    mkdir -p /tmp/dist/opt/resource && \
    go build \
      -trimpath \
      -ldflags "-s -w -extldflags '-static'" \
      -o /tmp/dist/opt/resource \
      ./cmd/...


FROM registry.access.redhat.com/ubi9-minimal:latest
RUN \
    microdnf --refresh --assumeyes --best --nodocs --noplugins --setopt=install_weak_deps=0 upgrade && \
    microdnf --assumeyes --nodocs install git jq && \
    microdnf clean all && \
    rm -rf /var/cache/yum

COPY --from=bootstrap /tmp/dist /
