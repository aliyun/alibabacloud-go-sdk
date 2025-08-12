// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableEventRulesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetRegionId(v string) *EnableEventRulesRequest
  GetRegionId() *string 
  SetRuleNames(v []*string) *EnableEventRulesRequest
  GetRuleNames() []*string 
}

type EnableEventRulesRequest struct {
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // ruleName1
  RuleNames []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s EnableEventRulesRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableEventRulesRequest) GoString() string {
  return s.String()
}

func (s *EnableEventRulesRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableEventRulesRequest) GetRuleNames() []*string  {
  return s.RuleNames
}

func (s *EnableEventRulesRequest) SetRegionId(v string) *EnableEventRulesRequest {
  s.RegionId = &v
  return s
}

func (s *EnableEventRulesRequest) SetRuleNames(v []*string) *EnableEventRulesRequest {
  s.RuleNames = v
  return s
}

func (s *EnableEventRulesRequest) Validate() error {
  return dara.Validate(s)
}

