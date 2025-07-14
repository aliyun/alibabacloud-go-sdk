// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationScalingRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppId(v string) *EnableApplicationScalingRuleRequest
  GetAppId() *string 
  SetScalingRuleName(v string) *EnableApplicationScalingRuleRequest
  GetScalingRuleName() *string 
}

type EnableApplicationScalingRuleRequest struct {
  // The application ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 7171a6ca-d1cd-4928-8642-7d5cfe69****
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  // The name of the auto scaling policy.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // timer-0800-2100
  ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s EnableApplicationScalingRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleRequest) GetAppId() *string  {
  return s.AppId
}

func (s *EnableApplicationScalingRuleRequest) GetScalingRuleName() *string  {
  return s.ScalingRuleName
}

func (s *EnableApplicationScalingRuleRequest) SetAppId(v string) *EnableApplicationScalingRuleRequest {
  s.AppId = &v
  return s
}

func (s *EnableApplicationScalingRuleRequest) SetScalingRuleName(v string) *EnableApplicationScalingRuleRequest {
  s.ScalingRuleName = &v
  return s
}

func (s *EnableApplicationScalingRuleRequest) Validate() error {
  return dara.Validate(s)
}

