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
  SetProjectId(v int64) *EnableRuleRequest
  GetProjectId() *int64 
  SetRegionId(v string) *EnableRuleRequest
  GetRegionId() *string 
  SetRuleId(v string) *EnableRuleRequest
  GetRuleId() *string 
}

type EnableRuleRequest struct {
  // The ID of the bastion host instance.
  // 
  // > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // bastionhost-cn-78v1ghxxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The project ID.
  // 
  // example:
  // 
  // 3
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // The region ID of the bastion host.
  // 
  // > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the authorization rule to enable.
  // 
  // > You can call the [ListRules](https://help.aliyun.com/document_detail/2758868.html) operation to query the rule ID.
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

func (s *EnableRuleRequest) GetProjectId() *int64  {
  return s.ProjectId
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

func (s *EnableRuleRequest) SetProjectId(v int64) *EnableRuleRequest {
  s.ProjectId = &v
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

