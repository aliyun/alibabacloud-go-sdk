// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *BindIdentityProviderRequest
	GetClientToken() *string
	SetIdentityProviderType(v string) *BindIdentityProviderRequest
	GetIdentityProviderType() *string
	SetIdpMetadata(v string) *BindIdentityProviderRequest
	GetIdpMetadata() *string
	SetInstanceId(v string) *BindIdentityProviderRequest
	GetInstanceId() *string
	SetLoginEnabled(v bool) *BindIdentityProviderRequest
	GetLoginEnabled() *bool
	SetSyncEnabled(v bool) *BindIdentityProviderRequest
	GetSyncEnabled() *bool
}

type BindIdentityProviderRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	IdpMetadata          *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LoginEnabled *bool   `json:"LoginEnabled,omitempty" xml:"LoginEnabled,omitempty"`
	SyncEnabled  *bool   `json:"SyncEnabled,omitempty" xml:"SyncEnabled,omitempty"`
}

func (s BindIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s BindIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *BindIdentityProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *BindIdentityProviderRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *BindIdentityProviderRequest) GetIdpMetadata() *string {
	return s.IdpMetadata
}

func (s *BindIdentityProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BindIdentityProviderRequest) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *BindIdentityProviderRequest) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *BindIdentityProviderRequest) SetClientToken(v string) *BindIdentityProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *BindIdentityProviderRequest) SetIdentityProviderType(v string) *BindIdentityProviderRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *BindIdentityProviderRequest) SetIdpMetadata(v string) *BindIdentityProviderRequest {
	s.IdpMetadata = &v
	return s
}

func (s *BindIdentityProviderRequest) SetInstanceId(v string) *BindIdentityProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *BindIdentityProviderRequest) SetLoginEnabled(v bool) *BindIdentityProviderRequest {
	s.LoginEnabled = &v
	return s
}

func (s *BindIdentityProviderRequest) SetSyncEnabled(v bool) *BindIdentityProviderRequest {
	s.SyncEnabled = &v
	return s
}

func (s *BindIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
