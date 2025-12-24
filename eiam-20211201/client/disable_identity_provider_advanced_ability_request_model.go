// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableIdentityProviderAdvancedAbilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *DisableIdentityProviderAdvancedAbilityRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *DisableIdentityProviderAdvancedAbilityRequest
	GetInstanceId() *string
}

type DisableIdentityProviderAdvancedAbilityRequest struct {
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
	// idaas_12323131xzxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableIdentityProviderAdvancedAbilityRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableIdentityProviderAdvancedAbilityRequest) GoString() string {
	return s.String()
}

func (s *DisableIdentityProviderAdvancedAbilityRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *DisableIdentityProviderAdvancedAbilityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableIdentityProviderAdvancedAbilityRequest) SetIdentityProviderId(v string) *DisableIdentityProviderAdvancedAbilityRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *DisableIdentityProviderAdvancedAbilityRequest) SetInstanceId(v string) *DisableIdentityProviderAdvancedAbilityRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableIdentityProviderAdvancedAbilityRequest) Validate() error {
	return dara.Validate(s)
}
