// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricRulesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetRegionId(v string) *EnableMetricRulesRequest
  GetRegionId() *string 
  SetRuleId(v []*string) *EnableMetricRulesRequest
  GetRuleId() []*string 
}

type EnableMetricRulesRequest struct {
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The IDs of the alert rules.
  // 
  // Valid values of N: 1 to 100.
  // 
  // For information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ab05733c97b7ce239fb1b53393dc1697c123****
  RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s EnableMetricRulesRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableMetricRulesRequest) GoString() string {
  return s.String()
}

func (s *EnableMetricRulesRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableMetricRulesRequest) GetRuleId() []*string  {
  return s.RuleId
}

func (s *EnableMetricRulesRequest) SetRegionId(v string) *EnableMetricRulesRequest {
  s.RegionId = &v
  return s
}

func (s *EnableMetricRulesRequest) SetRuleId(v []*string) *EnableMetricRulesRequest {
  s.RuleId = v
  return s
}

func (s *EnableMetricRulesRequest) Validate() error {
  return dara.Validate(s)
}

