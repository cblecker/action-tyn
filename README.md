# action-tyn

**NOTE:** DO NOT USE THIS. This is awful and has no tests.

GitHub Action to comment on closed PRs with a sassy message.

Input:
- Must be in a GitHub Action for a pull_request event
- Requires `GITHUB_TOKEN` be enabled
- Pulls in event data from `GITHUB_EVENT_JSON`

Usage:
```
workflow "tyn-workflow" {
  on = "pull_request"
  resolves = "tyn"
}

action "tyn" {
  uses = "cblecker/action-tyn@master"
  secrets = ["GITHUB_TOKEN"]
}
```
