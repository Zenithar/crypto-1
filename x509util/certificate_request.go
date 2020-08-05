package x509util

import (
	"crypto/rsa"
	"crypto/x509"
)

// CertificateRequest is the JSON representation of an X.509 certificate. It is
// used to build a certificate request from a template.
type CertificateRequest struct {
	Version            int                     `json:"version"`
	Subject            Subject                 `json:"subject"`
	DNSNames           MultiString             `json:"dnsNames"`
	EmailAddresses     MultiString             `json:"emailAddresses"`
	IPAddresses        MultiIP                 `json:"ipAddresses"`
	URIs               MultiURL                `json:"uris"`
	Extensions         []Extension             `json:"extensions"`
	PublicKey          interface{}             `json:"-"`
	PublicKeyAlgorithm x509.PublicKeyAlgorithm `json:"-"`
	Signature          []byte                  `json:"-"`
	SignatureAlgorithm x509.SignatureAlgorithm `json:"-"`
}

func newCertificateRequest(cr *x509.CertificateRequest) *CertificateRequest {
	return &CertificateRequest{
		Version:            cr.Version,
		Subject:            newSubject(cr.Subject),
		DNSNames:           cr.DNSNames,
		EmailAddresses:     cr.EmailAddresses,
		IPAddresses:        cr.IPAddresses,
		URIs:               cr.URIs,
		Extensions:         newExtensions(cr.Extensions),
		PublicKey:          cr.PublicKey,
		PublicKeyAlgorithm: cr.PublicKeyAlgorithm,
		Signature:          cr.Signature,
		SignatureAlgorithm: cr.SignatureAlgorithm,
	}
}

// GetCertificate returns the Certificate representation of the
// CertificateRequest.
func (c *CertificateRequest) GetCertificate() *Certificate {
	return &Certificate{
		Subject:            c.Subject,
		DNSNames:           c.DNSNames,
		EmailAddresses:     c.EmailAddresses,
		IPAddresses:        c.IPAddresses,
		URIs:               c.URIs,
		Extensions:         c.Extensions,
		PublicKey:          c.PublicKey,
		PublicKeyAlgorithm: c.PublicKeyAlgorithm,
	}
}

// GetLeafCertificate returns the Certificate representation of the
// CertificateRequest, including KeyUsage and ExtKeyUsage extensions.
func (c *CertificateRequest) GetLeafCertificate() *Certificate {
	keyUsage := x509.KeyUsageDigitalSignature
	if _, ok := c.PublicKey.(*rsa.PublicKey); ok {
		keyUsage |= x509.KeyUsageKeyEncipherment
	}

	cert := c.GetCertificate()
	cert.KeyUsage = KeyUsage(keyUsage)
	cert.ExtKeyUsage = ExtKeyUsage([]x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth,
		x509.ExtKeyUsageClientAuth,
	})
	return cert
}