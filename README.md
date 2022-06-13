[![PkgGoDev](https://img.shields.io/badge/go.dev-docs-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/abergmeier/buildkit_ex)
[![Go Report Card](https://goreportcard.com/badge/github.com/abergmeier/buildkit_ex)](https://goreportcard.com/report/github.com/abergmeier/buildkit_ex)
[![codecov](https://codecov.io/gh/abergmeier/buildkit_ex/branch/main/graph/badge.svg)](https://codecov.io/gh/abergmeier/buildkit_ex)

# Buildkit (the missing parts)

This repository contains functionality that was missing from Buildkit.

- `digest.DigestOfFileAndAllInputs`

  Returns a combined digest of a Dockerfile and all its inputs. Useful e.g. for telling Terraform whether a
  image needs to be rebuilt.
