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
  // The authorization rule ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // arrule_01kf143ug06fg7m9f43u7vahxxxx
  AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
  // The client token that is used to ensure the idempotence of the request. You can use the client to generate a parameter value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see References: [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // client-token-example
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The instance ID.
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

