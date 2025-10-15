// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationFederatedCredentialRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationFederatedCredentialId(v string) *EnableApplicationFederatedCredentialRequest
  GetApplicationFederatedCredentialId() *string 
  SetApplicationId(v string) *EnableApplicationFederatedCredentialRequest
  GetApplicationId() *string 
  SetInstanceId(v string) *EnableApplicationFederatedCredentialRequest
  GetInstanceId() *string 
}

type EnableApplicationFederatedCredentialRequest struct {
  // 应用联邦凭证Id
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // afc_aaaaa1111
  ApplicationFederatedCredentialId *string `json:"ApplicationFederatedCredentialId,omitempty" xml:"ApplicationFederatedCredentialId,omitempty"`
  // IDaaS的应用资源ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // app_mkv7rgt4d7i4u7zqtzev2mxxxx
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableApplicationFederatedCredentialRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationFederatedCredentialRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialId() *string  {
  return s.ApplicationFederatedCredentialId
}

func (s *EnableApplicationFederatedCredentialRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableApplicationFederatedCredentialRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialId(v string) *EnableApplicationFederatedCredentialRequest {
  s.ApplicationFederatedCredentialId = &v
  return s
}

func (s *EnableApplicationFederatedCredentialRequest) SetApplicationId(v string) *EnableApplicationFederatedCredentialRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableApplicationFederatedCredentialRequest) SetInstanceId(v string) *EnableApplicationFederatedCredentialRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableApplicationFederatedCredentialRequest) Validate() error {
  return dara.Validate(s)
}

