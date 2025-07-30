// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationClientSecretRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableApplicationClientSecretRequest
  GetApplicationId() *string 
  SetInstanceId(v string) *EnableApplicationClientSecretRequest
  GetInstanceId() *string 
  SetSecretId(v string) *EnableApplicationClientSecretRequest
  GetSecretId() *string 
}

type EnableApplicationClientSecretRequest struct {
  // The ID of the application.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // app_mkv7rgt4d7i4u7zqtzev2mxxxx
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // The ID of the instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The client key ID of the application.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // sci_k52x2ru63rlkflina5utgkxxxx
  SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s EnableApplicationClientSecretRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationClientSecretRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationClientSecretRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableApplicationClientSecretRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableApplicationClientSecretRequest) GetSecretId() *string  {
  return s.SecretId
}

func (s *EnableApplicationClientSecretRequest) SetApplicationId(v string) *EnableApplicationClientSecretRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableApplicationClientSecretRequest) SetInstanceId(v string) *EnableApplicationClientSecretRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableApplicationClientSecretRequest) SetSecretId(v string) *EnableApplicationClientSecretRequest {
  s.SecretId = &v
  return s
}

func (s *EnableApplicationClientSecretRequest) Validate() error {
  return dara.Validate(s)
}

