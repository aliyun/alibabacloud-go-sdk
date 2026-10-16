// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderName(v string) *GetIdentityProviderRequest
	GetIdentityProviderName() *string
}

type GetIdentityProviderRequest struct {
	// example:
	//
	// identity-provider-okta
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
}

func (s GetIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *GetIdentityProviderRequest) SetIdentityProviderName(v string) *GetIdentityProviderRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *GetIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
