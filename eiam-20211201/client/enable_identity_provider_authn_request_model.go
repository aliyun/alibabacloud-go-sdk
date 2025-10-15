// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableIdentityProviderAuthnRequest interface {
  dara.Model
  String() string
  GoString() string
  SetIdentityProviderId(v string) *EnableIdentityProviderAuthnRequest
  GetIdentityProviderId() *string 
  SetInstanceId(v string) *EnableIdentityProviderAuthnRequest
  GetInstanceId() *string 
}

type EnableIdentityProviderAuthnRequest struct {
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

func (s EnableIdentityProviderAuthnRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableIdentityProviderAuthnRequest) GoString() string {
  return s.String()
}

func (s *EnableIdentityProviderAuthnRequest) GetIdentityProviderId() *string  {
  return s.IdentityProviderId
}

func (s *EnableIdentityProviderAuthnRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableIdentityProviderAuthnRequest) SetIdentityProviderId(v string) *EnableIdentityProviderAuthnRequest {
  s.IdentityProviderId = &v
  return s
}

func (s *EnableIdentityProviderAuthnRequest) SetInstanceId(v string) *EnableIdentityProviderAuthnRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableIdentityProviderAuthnRequest) Validate() error {
  return dara.Validate(s)
}

