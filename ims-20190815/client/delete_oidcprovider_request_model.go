// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOIDCProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOIDCProviderName(v string) *DeleteOIDCProviderRequest
	GetOIDCProviderName() *string
}

type DeleteOIDCProviderRequest struct {
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s DeleteOIDCProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteOIDCProviderRequest) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *DeleteOIDCProviderRequest) SetOIDCProviderName(v string) *DeleteOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

func (s *DeleteOIDCProviderRequest) Validate() error {
	return dara.Validate(s)
}
