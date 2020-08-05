# crypto

[![Go Report Card](https://goreportcard.com/badge/github.com/smallstep/crypto)](https://goreportcard.com/report/github.com/smallstep/crypto)
[![Build Status](https://travis-ci.com/smallstep/crypto.svg?branch=master)](https://travis-ci.com/smallstep/crypto)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Crypto is a collection of packages used in [smallstep](https://smallstep.com) products. See:

* [step](https://github.com/smallstep/cli): A zero trust swiss army knife for
  working with X509, OAuth, JWT, OATH OTP, etc.
* [step-ca](https://github.com/smallstep/certificates): A private certificate
  authority (X.509 & SSH) & ACME server for secure automated certificate
  management, so you can use TLS everywhere & SSO for SSH.

## x509util

Package  implements utilities to build X.509 certificates based on JSON
templates.