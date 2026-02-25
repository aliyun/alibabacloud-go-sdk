// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebCertConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *WebCertConfig
	GetCertName() *string
	SetCertificate(v string) *WebCertConfig
	GetCertificate() *string
	SetPrivateKey(v string) *WebCertConfig
	GetPrivateKey() *string
}

type WebCertConfig struct {
	// The certificate name.
	//
	// example:
	//
	// sae-certname
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The public key of the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	//  ……
	//
	// -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The private key of the certificate.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MIIEowIBAAKCAQEA5SIfpNCBoiDrZhX1H39CHwQMVD0kBNeBTWfP9xkeesvfzbOz ******	- POVNFfDf9h7pJtQ5fRZNTYTDs/d+cH62Z3+nS74mNnEfff0nkvne
	//
	// -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s WebCertConfig) String() string {
	return dara.Prettify(s)
}

func (s WebCertConfig) GoString() string {
	return s.String()
}

func (s *WebCertConfig) GetCertName() *string {
	return s.CertName
}

func (s *WebCertConfig) GetCertificate() *string {
	return s.Certificate
}

func (s *WebCertConfig) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *WebCertConfig) SetCertName(v string) *WebCertConfig {
	s.CertName = &v
	return s
}

func (s *WebCertConfig) SetCertificate(v string) *WebCertConfig {
	s.Certificate = &v
	return s
}

func (s *WebCertConfig) SetPrivateKey(v string) *WebCertConfig {
	s.PrivateKey = &v
	return s
}

func (s *WebCertConfig) Validate() error {
	return dara.Validate(s)
}
