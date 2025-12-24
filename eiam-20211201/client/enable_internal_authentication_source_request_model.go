// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInternalAuthenticationSourceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuthenticationSourceId(v string) *EnableInternalAuthenticationSourceRequest
  GetAuthenticationSourceId() *string 
  SetInstanceId(v string) *EnableInternalAuthenticationSourceRequest
  GetInstanceId() *string 
}

type EnableInternalAuthenticationSourceRequest struct {
  // 内部认证源ID，比如 ia_password, ia_otp_sms 等
  // 
  // example:
  // 
  // ia_password
  AuthenticationSourceId *string `json:"AuthenticationSourceId,omitempty" xml:"AuthenticationSourceId,omitempty"`
  // IDaaS EIAM的实例id
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_111ccc11xxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableInternalAuthenticationSourceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableInternalAuthenticationSourceRequest) GoString() string {
  return s.String()
}

func (s *EnableInternalAuthenticationSourceRequest) GetAuthenticationSourceId() *string  {
  return s.AuthenticationSourceId
}

func (s *EnableInternalAuthenticationSourceRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableInternalAuthenticationSourceRequest) SetAuthenticationSourceId(v string) *EnableInternalAuthenticationSourceRequest {
  s.AuthenticationSourceId = &v
  return s
}

func (s *EnableInternalAuthenticationSourceRequest) SetInstanceId(v string) *EnableInternalAuthenticationSourceRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableInternalAuthenticationSourceRequest) Validate() error {
  return dara.Validate(s)
}

