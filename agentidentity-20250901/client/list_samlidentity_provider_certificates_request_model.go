// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSAMLIdentityProviderCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPoolName(v string) *ListSAMLIdentityProviderCertificatesRequest
	GetUserPoolName() *string
}

type ListSAMLIdentityProviderCertificatesRequest struct {
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s ListSAMLIdentityProviderCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSAMLIdentityProviderCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListSAMLIdentityProviderCertificatesRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListSAMLIdentityProviderCertificatesRequest) SetUserPoolName(v string) *ListSAMLIdentityProviderCertificatesRequest {
	s.UserPoolName = &v
	return s
}

func (s *ListSAMLIdentityProviderCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
