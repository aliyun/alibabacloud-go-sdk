// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableClientPublicKeyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableClientPublicKeyRequest
  GetApplicationId() *string 
  SetClientPublicKeyId(v string) *EnableClientPublicKeyRequest
  GetClientPublicKeyId() *string 
  SetClientToken(v string) *EnableClientPublicKeyRequest
  GetClientToken() *string 
  SetInstanceId(v string) *EnableClientPublicKeyRequest
  GetInstanceId() *string 
}

type EnableClientPublicKeyRequest struct {
  // IDaaS的应用资源ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // app_mkv7rgt4d7i4u7zqtzev2mxxxx
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // 应用ClientPublicKey的ID
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // KEYEqDnDJhztiEAwSin7MZoxGcihzCAuxxxx
  ClientPublicKeyId *string `json:"ClientPublicKeyId,omitempty" xml:"ClientPublicKeyId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // client-token-example
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableClientPublicKeyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableClientPublicKeyRequest) GoString() string {
  return s.String()
}

func (s *EnableClientPublicKeyRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableClientPublicKeyRequest) GetClientPublicKeyId() *string  {
  return s.ClientPublicKeyId
}

func (s *EnableClientPublicKeyRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableClientPublicKeyRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableClientPublicKeyRequest) SetApplicationId(v string) *EnableClientPublicKeyRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableClientPublicKeyRequest) SetClientPublicKeyId(v string) *EnableClientPublicKeyRequest {
  s.ClientPublicKeyId = &v
  return s
}

func (s *EnableClientPublicKeyRequest) SetClientToken(v string) *EnableClientPublicKeyRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableClientPublicKeyRequest) SetInstanceId(v string) *EnableClientPublicKeyRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableClientPublicKeyRequest) Validate() error {
  return dara.Validate(s)
}

