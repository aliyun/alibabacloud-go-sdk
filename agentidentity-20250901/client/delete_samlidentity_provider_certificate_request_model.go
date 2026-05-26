// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSAMLIdentityProviderCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *DeleteSAMLIdentityProviderCertificateRequest
	GetCertificateId() *string
	SetUserPoolName(v string) *DeleteSAMLIdentityProviderCertificateRequest
	GetUserPoolName() *string
}

type DeleteSAMLIdentityProviderCertificateRequest struct {
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxx
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s DeleteSAMLIdentityProviderCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSAMLIdentityProviderCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSAMLIdentityProviderCertificateRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *DeleteSAMLIdentityProviderCertificateRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *DeleteSAMLIdentityProviderCertificateRequest) SetCertificateId(v string) *DeleteSAMLIdentityProviderCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *DeleteSAMLIdentityProviderCertificateRequest) SetUserPoolName(v string) *DeleteSAMLIdentityProviderCertificateRequest {
	s.UserPoolName = &v
	return s
}

func (s *DeleteSAMLIdentityProviderCertificateRequest) Validate() error {
	return dara.Validate(s)
}
