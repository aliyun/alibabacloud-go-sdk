// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderUdPushConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *GetIdentityProviderUdPushConfigurationRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *GetIdentityProviderUdPushConfigurationRequest
	GetInstanceId() *string
}

type GetIdentityProviderUdPushConfigurationRequest struct {
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

func (s GetIdentityProviderUdPushConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPushConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPushConfigurationRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderUdPushConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderUdPushConfigurationRequest) SetIdentityProviderId(v string) *GetIdentityProviderUdPushConfigurationRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationRequest) SetInstanceId(v string) *GetIdentityProviderUdPushConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
