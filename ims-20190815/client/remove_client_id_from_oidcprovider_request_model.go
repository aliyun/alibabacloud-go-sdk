// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClientIdFromOIDCProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *RemoveClientIdFromOIDCProviderRequest
	GetClientId() *string
	SetOIDCProviderName(v string) *RemoveClientIdFromOIDCProviderRequest
	GetOIDCProviderName() *string
}

type RemoveClientIdFromOIDCProviderRequest struct {
	// The specified client ID.
	//
	// Format: The client ID can contain letters, digits, and the following special characters: `.-_:/`. It cannot start with a special character.
	//
	// Length: The client ID can be up to 128 characters in length.
	//
	// The `ClientId` must be bound to the OIDC IdP specified by `OIDCProviderName`. If it is not bound, you can first call `CreateOIDCProvider` by setting `ClientIds` or call `AddClientIdToOIDCProvider` to bind it.
	//
	// example:
	//
	// 498469743454717****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The name of the OIDC IdP that exists and is bound to the target `ClientId`. If the OIDC IdP has not been created or the client ID has not been bound, you can first call `CreateOIDCProvider` by setting `ClientIds` or call `AddClientIdToOIDCProvider`.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s RemoveClientIdFromOIDCProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveClientIdFromOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *RemoveClientIdFromOIDCProviderRequest) GetClientId() *string {
	return s.ClientId
}

func (s *RemoveClientIdFromOIDCProviderRequest) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *RemoveClientIdFromOIDCProviderRequest) SetClientId(v string) *RemoveClientIdFromOIDCProviderRequest {
	s.ClientId = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderRequest) SetOIDCProviderName(v string) *RemoveClientIdFromOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderRequest) Validate() error {
	return dara.Validate(s)
}
