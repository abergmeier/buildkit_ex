# Buildkit (the missing parts)

This repository contains functionality that was missing from Buildkit.

- `digest.DigestOfFileAndAllInputs`

  Returns a combined digest of a Dockerfile and all its inputs. Useful e.g. for telling Terraform whether a
  image needs to be rebuilt.
