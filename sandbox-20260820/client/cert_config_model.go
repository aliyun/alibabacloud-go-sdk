// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *CertConfig
	GetCertName() *string
	SetCertificate(v string) *CertConfig
	GetCertificate() *string
	SetPrivateKey(v string) *CertConfig
	GetPrivateKey() *string
}

type CertConfig struct {
	// The certificate name, which is used to identify the certificate in the console.
	//
	// example:
	//
	// sandbox-example-com
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// The certificate public key content in PEM format, including the complete certificate chain.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDdzCCAl+gAwIBAgIEbGVzc29u
	//
	// -----END CERTIFICATE-----
	Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
	// The certificate private key content in PEM format. The private key is encrypted and stored on the server side, and is not returned in plaintext when queried.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----
	//
	// ****
	//
	// -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
}

func (s CertConfig) String() string {
	return dara.Prettify(s)
}

func (s CertConfig) GoString() string {
	return s.String()
}

func (s *CertConfig) GetCertName() *string {
	return s.CertName
}

func (s *CertConfig) GetCertificate() *string {
	return s.Certificate
}

func (s *CertConfig) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *CertConfig) SetCertName(v string) *CertConfig {
	s.CertName = &v
	return s
}

func (s *CertConfig) SetCertificate(v string) *CertConfig {
	s.Certificate = &v
	return s
}

func (s *CertConfig) SetPrivateKey(v string) *CertConfig {
	s.PrivateKey = &v
	return s
}

func (s *CertConfig) Validate() error {
	return dara.Validate(s)
}
