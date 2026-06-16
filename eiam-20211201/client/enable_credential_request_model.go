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
  // A client-generated token that ensures the idempotence of the request. Make sure that the token is unique for each request. The ClientToken parameter supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // client-token-example
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The credential ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cred_mkv7rgt4d7i4u7zqtzev2mxxxx
  CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
  // The instance ID.
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

