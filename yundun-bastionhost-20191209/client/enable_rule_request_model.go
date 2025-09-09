// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableRuleRequest
  GetInstanceId() *string 
  SetRegionId(v string) *EnableRuleRequest
  GetRegionId() *string 
  SetRuleId(v string) *EnableRuleRequest
  GetRuleId() *string 
}

type EnableRuleRequest struct {
  // The bastion host ID.
  // 
  // >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // bastionhost-cn-78v1ghxxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The region ID of the bastion host.
  // 
  // > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the authorization rule that you want to enable.
  // 
  // >  You can call the [ListRules](https://help.aliyun.com/document_detail/2758868.html) operation to query the authorization rule ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 3
  RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s EnableRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableRuleRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableRuleRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableRuleRequest) GetRuleId() *string  {
  return s.RuleId
}

func (s *EnableRuleRequest) SetInstanceId(v string) *EnableRuleRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableRuleRequest) SetRegionId(v string) *EnableRuleRequest {
  s.RegionId = &v
  return s
}

func (s *EnableRuleRequest) SetRuleId(v string) *EnableRuleRequest {
  s.RuleId = &v
  return s
}

func (s *EnableRuleRequest) Validate() error {
  return dara.Validate(s)
}

