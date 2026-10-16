// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSAMLIdentityProviderCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPoolName(v string) *AddSAMLIdentityProviderCertificateRequest
	GetUserPoolName() *string
	SetX509Certificate(v string) *AddSAMLIdentityProviderCertificateRequest
	GetX509Certificate() *string
}

type AddSAMLIdentityProviderCertificateRequest struct {
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDdz...
	//
	// -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s AddSAMLIdentityProviderCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSAMLIdentityProviderCertificateRequest) GoString() string {
	return s.String()
}

func (s *AddSAMLIdentityProviderCertificateRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *AddSAMLIdentityProviderCertificateRequest) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *AddSAMLIdentityProviderCertificateRequest) SetUserPoolName(v string) *AddSAMLIdentityProviderCertificateRequest {
	s.UserPoolName = &v
	return s
}

func (s *AddSAMLIdentityProviderCertificateRequest) SetX509Certificate(v string) *AddSAMLIdentityProviderCertificateRequest {
	s.X509Certificate = &v
	return s
}

func (s *AddSAMLIdentityProviderCertificateRequest) Validate() error {
	return dara.Validate(s)
}
