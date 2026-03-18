// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCredentialProviderRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCredentialProviderId(v string) *EnableCredentialProviderRequest
  GetCredentialProviderId() *string 
  SetInstanceId(v string) *EnableCredentialProviderRequest
  GetInstanceId() *string 
}

type EnableCredentialProviderRequest struct {
  // 认证令牌提供商ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
  CredentialProviderId *string `json:"CredentialProviderId,omitempty" xml:"CredentialProviderId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableCredentialProviderRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCredentialProviderRequest) GoString() string {
  return s.String()
}

func (s *EnableCredentialProviderRequest) GetCredentialProviderId() *string  {
  return s.CredentialProviderId
}

func (s *EnableCredentialProviderRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableCredentialProviderRequest) SetCredentialProviderId(v string) *EnableCredentialProviderRequest {
  s.CredentialProviderId = &v
  return s
}

func (s *EnableCredentialProviderRequest) SetInstanceId(v string) *EnableCredentialProviderRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableCredentialProviderRequest) Validate() error {
  return dara.Validate(s)
}

