// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateIdentityProviderRequest
	GetClientToken() *string
	SetIdentityProviderType(v string) *UpdateIdentityProviderRequest
	GetIdentityProviderType() *string
	SetIdpMetadata(v string) *UpdateIdentityProviderRequest
	GetIdpMetadata() *string
	SetInstanceId(v string) *UpdateIdentityProviderRequest
	GetInstanceId() *string
	SetLoginEnabled(v bool) *UpdateIdentityProviderRequest
	GetLoginEnabled() *bool
	SetSyncEnabled(v bool) *UpdateIdentityProviderRequest
	GetSyncEnabled() *bool
}

type UpdateIdentityProviderRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	IdpMetadata          *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LoginEnabled *bool   `json:"LoginEnabled,omitempty" xml:"LoginEnabled,omitempty"`
	SyncEnabled  *bool   `json:"SyncEnabled,omitempty" xml:"SyncEnabled,omitempty"`
}

func (s UpdateIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIdentityProviderRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *UpdateIdentityProviderRequest) GetIdpMetadata() *string {
	return s.IdpMetadata
}

func (s *UpdateIdentityProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateIdentityProviderRequest) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *UpdateIdentityProviderRequest) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *UpdateIdentityProviderRequest) SetClientToken(v string) *UpdateIdentityProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetIdentityProviderType(v string) *UpdateIdentityProviderRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetIdpMetadata(v string) *UpdateIdentityProviderRequest {
	s.IdpMetadata = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetInstanceId(v string) *UpdateIdentityProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetLoginEnabled(v bool) *UpdateIdentityProviderRequest {
	s.LoginEnabled = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetSyncEnabled(v bool) *UpdateIdentityProviderRequest {
	s.SyncEnabled = &v
	return s
}

func (s *UpdateIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
