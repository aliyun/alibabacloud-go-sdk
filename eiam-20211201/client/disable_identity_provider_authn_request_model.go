// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableIdentityProviderAuthnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *DisableIdentityProviderAuthnRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *DisableIdentityProviderAuthnRequest
	GetInstanceId() *string
}

type DisableIdentityProviderAuthnRequest struct {
	// IDaaS的身份提供方主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_11111
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableIdentityProviderAuthnRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableIdentityProviderAuthnRequest) GoString() string {
	return s.String()
}

func (s *DisableIdentityProviderAuthnRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *DisableIdentityProviderAuthnRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableIdentityProviderAuthnRequest) SetIdentityProviderId(v string) *DisableIdentityProviderAuthnRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *DisableIdentityProviderAuthnRequest) SetInstanceId(v string) *DisableIdentityProviderAuthnRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableIdentityProviderAuthnRequest) Validate() error {
	return dara.Validate(s)
}
