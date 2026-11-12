// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationSSLRequest
	GetApplicationId() *string
	SetCaPem(v string) *ModifyApplicationSSLRequest
	GetCaPem() *string
	SetCertPem(v string) *ModifyApplicationSSLRequest
	GetCertPem() *string
	SetKeyPem(v string) *ModifyApplicationSSLRequest
	GetKeyPem() *string
	SetSslEnabled(v bool) *ModifyApplicationSSLRequest
	GetSslEnabled() *bool
}

type ModifyApplicationSSLRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The server certificate chain in PEM format. This parameter is optional.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICa
	//
	// -----END CERTIFICATE-----
	CaPem *string `json:"CaPem,omitempty" xml:"CaPem,omitempty"`
	// The server certificate in PEM format. This parameter is provided when you use a bring-your-own (BYO) certificate. It must match KeyPem and override the instance connection domain name. This parameter is allowed only when CA SSL is enabled for the instance.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICert
	//
	// -----END CERTIFICATE-----
	CertPem *string `json:"CertPem,omitempty" xml:"CertPem,omitempty"`
	// The server private key in PEM format. PKCS#1, PKCS#8, and EC SEC1 formats are supported.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	//
	// MIIKey
	//
	// -----END PRIVATE KEY-----
	KeyPem *string `json:"KeyPem,omitempty" xml:"KeyPem,omitempty"`
	// Specifies whether SSL is enabled.
	//
	// This parameter is required.
	SslEnabled *bool `json:"SslEnabled,omitempty" xml:"SslEnabled,omitempty"`
}

func (s ModifyApplicationSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSSLRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationSSLRequest) GetCaPem() *string {
	return s.CaPem
}

func (s *ModifyApplicationSSLRequest) GetCertPem() *string {
	return s.CertPem
}

func (s *ModifyApplicationSSLRequest) GetKeyPem() *string {
	return s.KeyPem
}

func (s *ModifyApplicationSSLRequest) GetSslEnabled() *bool {
	return s.SslEnabled
}

func (s *ModifyApplicationSSLRequest) SetApplicationId(v string) *ModifyApplicationSSLRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationSSLRequest) SetCaPem(v string) *ModifyApplicationSSLRequest {
	s.CaPem = &v
	return s
}

func (s *ModifyApplicationSSLRequest) SetCertPem(v string) *ModifyApplicationSSLRequest {
	s.CertPem = &v
	return s
}

func (s *ModifyApplicationSSLRequest) SetKeyPem(v string) *ModifyApplicationSSLRequest {
	s.KeyPem = &v
	return s
}

func (s *ModifyApplicationSSLRequest) SetSslEnabled(v bool) *ModifyApplicationSSLRequest {
	s.SslEnabled = &v
	return s
}

func (s *ModifyApplicationSSLRequest) Validate() error {
	return dara.Validate(s)
}
