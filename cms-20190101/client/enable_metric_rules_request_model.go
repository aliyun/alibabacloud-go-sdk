// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricRulesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableMetricRulesRequest
  GetClientToken() *string 
  SetDryRun(v bool) *EnableMetricRulesRequest
  GetDryRun() *bool 
  SetRegionId(v string) *EnableMetricRulesRequest
  GetRegionId() *string 
  SetRuleId(v []*string) *EnableMetricRulesRequest
  GetRuleId() []*string 
}

type EnableMetricRulesRequest struct {
  // The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, and you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
  // 
  // example:
  // 
  // 6868E9A1-56E6-5FD7-9684-E2963BBAFB5F
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // Specifies whether to precheck only this request. Valid values:
  // 
  // - **true**: sends a check request and does not associate a security group with the instance. The check items include whether the required parameters are specified, the request format is valid, and the service limits are met. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
  // 
  // - **false*	- (default): sends a regular request. If the check succeeds, an HTTP 2xx status code is returned and the operation is directly performed.
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

func (s *EnableMetricRulesRequest) GetClientToken() *string  {
  return s.ClientToken
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

func (s *EnableMetricRulesRequest) SetClientToken(v string) *EnableMetricRulesRequest {
  s.ClientToken = &v
  return s
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

