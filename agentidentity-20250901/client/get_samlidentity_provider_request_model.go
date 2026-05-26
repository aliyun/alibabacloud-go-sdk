// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSAMLIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPoolName(v string) *GetSAMLIdentityProviderRequest
	GetUserPoolName() *string
}

type GetSAMLIdentityProviderRequest struct {
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetSAMLIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *GetSAMLIdentityProviderRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetSAMLIdentityProviderRequest) SetUserPoolName(v string) *GetSAMLIdentityProviderRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetSAMLIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
