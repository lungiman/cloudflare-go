//go:build tools
// +build tools

package tools

//go:generate go install golang.org/x/tools/gopls@latest
//go:generate go install github.com/orijtech/structslop/cmd/structslop
//go:generate go install github.com/ramya-rao-a/go-outline
//go:generate go install github.com/cweill/gotests/gotests
//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go install github.com/uudashr/gopkgs/v2/cmd/gopkgs
//go:generate go install github.com/go-delve/delve/cmd/dlv
//go:generate go install golang.org/x/lint/golint
//go:generate go install github.com/hashicorp/go-changelog/cmd/changelog-build
//go:generate go install github.com/google/go-github/github
//go:generate go install golang.org/x/oauth2

import (
	// local development tooling for linting and debugging.
	_ "github.com/cweill/gotests/gotests"
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/hashicorp/go-changelog/cmd/changelog-build"
	_ "github.com/orijtech/structslop/cmd/structslop"
	_ "github.com/ramya-rao-a/go-outline"
	_ "github.com/uudashr/gopkgs/v2/cmd/gopkgs"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/gopls"

	// used for changelog-check tooling
	_ "github.com/google/go-github/github"
	_ "golang.org/x/oauth2"
)
