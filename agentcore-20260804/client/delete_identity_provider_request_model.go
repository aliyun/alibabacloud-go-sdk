// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIdentityProviderRequest
	GetClientToken() *string
}

type DeleteIdentityProviderRequest struct {
	// Not supported.
	//
	// example:
	//
	// Not supported
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeleteIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIdentityProviderRequest) SetClientToken(v string) *DeleteIdentityProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
