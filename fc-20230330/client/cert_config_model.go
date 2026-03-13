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
	// The name of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-cert
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// The certificate in the .pem format.
	//
	// This parameter is required.
	//
	// example:
	//
	// PEM format
	Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
	// The private key in the .pem format.
	//
	// This parameter is required.
	//
	// example:
	//
	// PEM format
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
