// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCredentialRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableCredentialRequest
  GetClientToken() *string 
  SetCredentialId(v string) *EnableCredentialRequest
  GetCredentialId() *string 
  SetInstanceId(v string) *EnableCredentialRequest
  GetInstanceId() *string 
}

type EnableCredentialRequest struct {
  // 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // client-token-example
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // 凭据ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cred_mkv7rgt4d7i4u7zqtzev2mxxxx
  CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableCredentialRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCredentialRequest) GoString() string {
  return s.String()
}

func (s *EnableCredentialRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableCredentialRequest) GetCredentialId() *string  {
  return s.CredentialId
}

func (s *EnableCredentialRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableCredentialRequest) SetClientToken(v string) *EnableCredentialRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableCredentialRequest) SetCredentialId(v string) *EnableCredentialRequest {
  s.CredentialId = &v
  return s
}

func (s *EnableCredentialRequest) SetInstanceId(v string) *EnableCredentialRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableCredentialRequest) Validate() error {
  return dara.Validate(s)
}

