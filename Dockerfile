FROM golang:1.11

LABEL "name"="thank u, next"
LABEL "maintainer"="Christoph Blecker <admin@toph.ca>"
LABEL "version"="0.0.1"

LABEL "com.github.actions.name"="thank u, next"
LABEL "com.github.actions.description"="Comment on closed PRs with a sassy message"
LABEL "com.github.actions.icon"="x-circle"
LABEL "com.github.actions.color"="red"

COPY main.go go.mod go.sum /

ENTRYPOINT ["go"]
CMD ["run", "/main.go"]
