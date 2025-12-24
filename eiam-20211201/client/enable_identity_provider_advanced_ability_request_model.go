// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableIdentityProviderAdvancedAbilityRequest interface {
  dara.Model
  String() string
  GoString() string
  SetIdentityProviderId(v string) *EnableIdentityProviderAdvancedAbilityRequest
  GetIdentityProviderId() *string 
  SetInstanceId(v string) *EnableIdentityProviderAdvancedAbilityRequest
  GetInstanceId() *string 
}

type EnableIdentityProviderAdvancedAbilityRequest struct {
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
  // idaas_111ccc111xx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableIdentityProviderAdvancedAbilityRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableIdentityProviderAdvancedAbilityRequest) GoString() string {
  return s.String()
}

func (s *EnableIdentityProviderAdvancedAbilityRequest) GetIdentityProviderId() *string  {
  return s.IdentityProviderId
}

func (s *EnableIdentityProviderAdvancedAbilityRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableIdentityProviderAdvancedAbilityRequest) SetIdentityProviderId(v string) *EnableIdentityProviderAdvancedAbilityRequest {
  s.IdentityProviderId = &v
  return s
}

func (s *EnableIdentityProviderAdvancedAbilityRequest) SetInstanceId(v string) *EnableIdentityProviderAdvancedAbilityRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableIdentityProviderAdvancedAbilityRequest) Validate() error {
  return dara.Validate(s)
}

