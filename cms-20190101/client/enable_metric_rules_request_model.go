// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricRulesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDryRun(v bool) *EnableMetricRulesRequest
  GetDryRun() *bool 
  SetRegionId(v string) *EnableMetricRulesRequest
  GetRegionId() *string 
  SetRuleId(v []*string) *EnableMetricRulesRequest
  GetRuleId() []*string 
}

type EnableMetricRulesRequest struct {
  // Specifies whether to perform a dry run only. Valid values:
  // 
  // - **true**: performs a dry run only and does not bind a security group to the instance. The system checks whether all required parameters are specified, the request format is valid, and the service limits are met. If the request fails the check, the corresponding error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
  // 
  // - **false*	- (default): sends a normal request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
  // 
  // example:
  // 
  // false
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the alert rule.
  // 
  // Valid values of N: 1 to 100.
  // 
  // For more information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
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

func (s *EnableMetricRulesRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableMetricRulesRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableMetricRulesRequest) GetRuleId() []*string  {
  return s.RuleId
}

func (s *EnableMetricRulesRequest) SetDryRun(v bool) *EnableMetricRulesRequest {
  s.DryRun = &v
  return s
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

