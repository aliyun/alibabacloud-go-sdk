// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEventBusName(v string) *EnableRuleRequest
  GetEventBusName() *string 
  SetRuleName(v string) *EnableRuleRequest
  GetRuleName() *string 
}

type EnableRuleRequest struct {
  // The name of the event bus.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // MyEventBus
  EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
  // The name of the event rule.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ramrolechange
  RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s EnableRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableRuleRequest) GetEventBusName() *string  {
  return s.EventBusName
}

func (s *EnableRuleRequest) GetRuleName() *string  {
  return s.RuleName
}

func (s *EnableRuleRequest) SetEventBusName(v string) *EnableRuleRequest {
  s.EventBusName = &v
  return s
}

func (s *EnableRuleRequest) SetRuleName(v string) *EnableRuleRequest {
  s.RuleName = &v
  return s
}

func (s *EnableRuleRequest) Validate() error {
  return dara.Validate(s)
}

