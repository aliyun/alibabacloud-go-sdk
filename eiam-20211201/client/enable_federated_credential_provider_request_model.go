// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFederatedCredentialProviderRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFederatedCredentialProviderId(v string) *EnableFederatedCredentialProviderRequest
  GetFederatedCredentialProviderId() *string 
  SetInstanceId(v string) *EnableFederatedCredentialProviderRequest
  GetInstanceId() *string 
}

type EnableFederatedCredentialProviderRequest struct {
  // 联邦凭证提供方ID
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // fcp_mkv7rgt4d7i4u7zqtzev2mxxxx
  FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableFederatedCredentialProviderRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableFederatedCredentialProviderRequest) GoString() string {
  return s.String()
}

func (s *EnableFederatedCredentialProviderRequest) GetFederatedCredentialProviderId() *string  {
  return s.FederatedCredentialProviderId
}

func (s *EnableFederatedCredentialProviderRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableFederatedCredentialProviderRequest) SetFederatedCredentialProviderId(v string) *EnableFederatedCredentialProviderRequest {
  s.FederatedCredentialProviderId = &v
  return s
}

func (s *EnableFederatedCredentialProviderRequest) SetInstanceId(v string) *EnableFederatedCredentialProviderRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableFederatedCredentialProviderRequest) Validate() error {
  return dara.Validate(s)
}

