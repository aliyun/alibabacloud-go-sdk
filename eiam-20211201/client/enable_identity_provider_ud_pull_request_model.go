// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableIdentityProviderUdPullRequest interface {
  dara.Model
  String() string
  GoString() string
  SetIdentityProviderId(v string) *EnableIdentityProviderUdPullRequest
  GetIdentityProviderId() *string 
  SetInstanceId(v string) *EnableIdentityProviderUdPullRequest
  GetInstanceId() *string 
}

type EnableIdentityProviderUdPullRequest struct {
  // Identity provider ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idp_my664lwkhpicbyzirog3xxxxx
  IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
  // The ID of the instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableIdentityProviderUdPullRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableIdentityProviderUdPullRequest) GoString() string {
  return s.String()
}

func (s *EnableIdentityProviderUdPullRequest) GetIdentityProviderId() *string  {
  return s.IdentityProviderId
}

func (s *EnableIdentityProviderUdPullRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableIdentityProviderUdPullRequest) SetIdentityProviderId(v string) *EnableIdentityProviderUdPullRequest {
  s.IdentityProviderId = &v
  return s
}

func (s *EnableIdentityProviderUdPullRequest) SetInstanceId(v string) *EnableIdentityProviderUdPullRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableIdentityProviderUdPullRequest) Validate() error {
  return dara.Validate(s)
}

