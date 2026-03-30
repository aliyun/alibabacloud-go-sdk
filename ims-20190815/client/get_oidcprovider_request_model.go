// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOIDCProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOIDCProviderName(v string) *GetOIDCProviderRequest
	GetOIDCProviderName() *string
}

type GetOIDCProviderRequest struct {
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s GetOIDCProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *GetOIDCProviderRequest) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *GetOIDCProviderRequest) SetOIDCProviderName(v string) *GetOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

func (s *GetOIDCProviderRequest) Validate() error {
	return dara.Validate(s)
}
