// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAuthorizationRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuthorizationRuleId(v string) *EnableAuthorizationRuleRequest
  GetAuthorizationRuleId() *string 
  SetClientToken(v string) *EnableAuthorizationRuleRequest
  GetClientToken() *string 
  SetInstanceId(v string) *EnableAuthorizationRuleRequest
  GetInstanceId() *string 
}

type EnableAuthorizationRuleRequest struct {
  // 授权规则标识。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // arrule_01kf143ug06fg7m9f43u7vahxxxx
  AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
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

func (s EnableAuthorizationRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAuthorizationRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableAuthorizationRuleRequest) GetAuthorizationRuleId() *string  {
  return s.AuthorizationRuleId
}

func (s *EnableAuthorizationRuleRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableAuthorizationRuleRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *EnableAuthorizationRuleRequest {
  s.AuthorizationRuleId = &v
  return s
}

func (s *EnableAuthorizationRuleRequest) SetClientToken(v string) *EnableAuthorizationRuleRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableAuthorizationRuleRequest) SetInstanceId(v string) *EnableAuthorizationRuleRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableAuthorizationRuleRequest) Validate() error {
  return dara.Validate(s)
}

