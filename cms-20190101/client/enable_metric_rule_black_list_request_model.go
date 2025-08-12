// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricRuleBlackListRequest interface {
  dara.Model
  String() string
  GoString() string
  SetId(v string) *EnableMetricRuleBlackListRequest
  GetId() *string 
  SetIsEnable(v bool) *EnableMetricRuleBlackListRequest
  GetIsEnable() *bool 
  SetRegionId(v string) *EnableMetricRuleBlackListRequest
  GetRegionId() *string 
}

type EnableMetricRuleBlackListRequest struct {
  // The IDs of the blacklist policies. Separate multiple IDs with commas (,). You can specify up to 50 IDs.
  // 
  // For information about how to obtain the ID of a blacklist policy, see [DescribeMetricRuleBlackList](https://help.aliyun.com/document_detail/457257.html).
  // 
  // > You can also set this parameter to a JSON array. Example: `["a9ad2ac2-3ed9-11ed-b878-0242ac12****","5cb8a9a4-198f-4651-a353-f8b28788****"]`.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // a9ad2ac2-3ed9-11ed-b878-0242ac12****
  Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
  // Specifies whether to enable the blacklist policy. Valid values:
  // 
  // 	- true: The blacklist policy is enabled.
  // 
  // 	- false (default): The blacklist policy is disabled.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // true
  IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableMetricRuleBlackListRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableMetricRuleBlackListRequest) GoString() string {
  return s.String()
}

func (s *EnableMetricRuleBlackListRequest) GetId() *string  {
  return s.Id
}

func (s *EnableMetricRuleBlackListRequest) GetIsEnable() *bool  {
  return s.IsEnable
}

func (s *EnableMetricRuleBlackListRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableMetricRuleBlackListRequest) SetId(v string) *EnableMetricRuleBlackListRequest {
  s.Id = &v
  return s
}

func (s *EnableMetricRuleBlackListRequest) SetIsEnable(v bool) *EnableMetricRuleBlackListRequest {
  s.IsEnable = &v
  return s
}

func (s *EnableMetricRuleBlackListRequest) SetRegionId(v string) *EnableMetricRuleBlackListRequest {
  s.RegionId = &v
  return s
}

func (s *EnableMetricRuleBlackListRequest) Validate() error {
  return dara.Validate(s)
}

