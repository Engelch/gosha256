# gosha256

## About

Very simple quick hack to check of the go calculation of a SHA256 digest results the same value as the usual CLI commands such as `shasum -a 256` or `sha256sum`.

It is kept in a repository to easily do a retest if required.

Currently, March 2022, the compiled binary outputs the same string as the stanard CLI commands.

## Synopsis

``
gosha256 <<filename>>
``

## Limitations

1. The application reads the complete file into a buffer. No protection exists against buffer overflows.
2. No usage handling.
3. Limited error handling, no beautification.
4. Documentation is very, very poor.

## Licenses

| Software/Package        | License                |
|-------------------------|------------------------|
| urfave CLI package      | MIT                    |
| GO programming language | https://go.dev/LICENSE |