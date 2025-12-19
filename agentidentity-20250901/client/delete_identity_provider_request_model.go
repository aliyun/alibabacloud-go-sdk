// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderName(v string) *DeleteIdentityProviderRequest
	GetIdentityProviderName() *string
}

type DeleteIdentityProviderRequest struct {
	// example:
	//
	// identity-provider-okta
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
}

func (s DeleteIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *DeleteIdentityProviderRequest) SetIdentityProviderName(v string) *DeleteIdentityProviderRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *DeleteIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
