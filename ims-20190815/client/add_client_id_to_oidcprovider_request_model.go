// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClientIdToOIDCProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *AddClientIdToOIDCProviderRequest
	GetClientId() *string
	SetOIDCProviderName(v string) *AddClientIdToOIDCProviderRequest
	GetOIDCProviderName() *string
}

type AddClientIdToOIDCProviderRequest struct {
	// The client ID that you want to add.
	//
	// The client ID can contain letters, digits, and special characters and cannot start with the special characters. The special characters are `periods, (.), hyphens (-), underscores (_), colons (:), and forward slashes (/)`.``
	//
	// The client ID can be up to 128 characters in length.
	//
	// example:
	//
	// 598469743454717****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s AddClientIdToOIDCProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s AddClientIdToOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *AddClientIdToOIDCProviderRequest) GetClientId() *string {
	return s.ClientId
}

func (s *AddClientIdToOIDCProviderRequest) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *AddClientIdToOIDCProviderRequest) SetClientId(v string) *AddClientIdToOIDCProviderRequest {
	s.ClientId = &v
	return s
}

func (s *AddClientIdToOIDCProviderRequest) SetOIDCProviderName(v string) *AddClientIdToOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

func (s *AddClientIdToOIDCProviderRequest) Validate() error {
	return dara.Validate(s)
}
