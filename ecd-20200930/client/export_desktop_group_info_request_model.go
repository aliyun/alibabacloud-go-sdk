// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDesktopGroupInfoRequest interface {
  dara.Model
  String() string
  GoString() string
  SetChargeType(v string) *ExportDesktopGroupInfoRequest
  GetChargeType() *string 
  SetDesktopGroupId(v []*string) *ExportDesktopGroupInfoRequest
  GetDesktopGroupId() []*string 
  SetDesktopGroupName(v string) *ExportDesktopGroupInfoRequest
  GetDesktopGroupName() *string 
  SetEndUserId(v []*string) *ExportDesktopGroupInfoRequest
  GetEndUserId() []*string 
  SetExpiredTime(v string) *ExportDesktopGroupInfoRequest
  GetExpiredTime() *string 
  SetLangType(v string) *ExportDesktopGroupInfoRequest
  GetLangType() *string 
  SetMaxResults(v int32) *ExportDesktopGroupInfoRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *ExportDesktopGroupInfoRequest
  GetNextToken() *string 
  SetOfficeSiteId(v string) *ExportDesktopGroupInfoRequest
  GetOfficeSiteId() *string 
  SetPolicyGroupId(v string) *ExportDesktopGroupInfoRequest
  GetPolicyGroupId() *string 
  SetRegionId(v string) *ExportDesktopGroupInfoRequest
  GetRegionId() *string 
  SetTag(v []*ExportDesktopGroupInfoRequestTag) *ExportDesktopGroupInfoRequest
  GetTag() []*ExportDesktopGroupInfoRequestTag 
}

type ExportDesktopGroupInfoRequest struct {
  // The billing method of the cloud computer share.
  // 
  // Valid values:
  // 
  // 	- PostPaid: pay-as-you-go.
  // 
  // 	- PrePaid: subscription.
  // 
  // example:
  // 
  // PrePaid
  ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
  // The IDs of the cloud computer shares.
  DesktopGroupId []*string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" type:"Repeated"`
  // The name of the cloud computer share.
  // 
  // example:
  // 
  // test
  DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
  // The IDs of the users to be authorized.
  EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
  // The expiration date of the subscription cloud computer share.
  // 
  // example:
  // 
  // 2022-12-31T15:59Z
  ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
  // The language of the response.
  // 
  // example:
  // 
  // zh-CN
  LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
  // The number of entries to return on each page.
  // 
  // Maximum value: 100.
  // 
  // Default value: 10.
  // 
  // example:
  // 
  // 10
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // The token that determines the start point of the next query. If this parameter is left empty, all results are returned.
  // 
  // example:
  // 
  // caeba0bbb2be03f84eb48b699f0a4883
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  // The ID of the office network.
  // 
  // example:
  // 
  // cn-hangzhou+dir-467671****
  OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
  // The ID of the security policy.
  // 
  // example:
  // 
  // pg-53iyi2aar0nd6****
  PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
  // The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by Elastic Desktop Service.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The tags. You can specify up to 20 tags.
  Tag []*ExportDesktopGroupInfoRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ExportDesktopGroupInfoRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportDesktopGroupInfoRequest) GoString() string {
  return s.String()
}

func (s *ExportDesktopGroupInfoRequest) GetChargeType() *string  {
  return s.ChargeType
}

func (s *ExportDesktopGroupInfoRequest) GetDesktopGroupId() []*string  {
  return s.DesktopGroupId
}

func (s *ExportDesktopGroupInfoRequest) GetDesktopGroupName() *string  {
  return s.DesktopGroupName
}

func (s *ExportDesktopGroupInfoRequest) GetEndUserId() []*string  {
  return s.EndUserId
}

func (s *ExportDesktopGroupInfoRequest) GetExpiredTime() *string  {
  return s.ExpiredTime
}

func (s *ExportDesktopGroupInfoRequest) GetLangType() *string  {
  return s.LangType
}

func (s *ExportDesktopGroupInfoRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *ExportDesktopGroupInfoRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportDesktopGroupInfoRequest) GetOfficeSiteId() *string  {
  return s.OfficeSiteId
}

func (s *ExportDesktopGroupInfoRequest) GetPolicyGroupId() *string  {
  return s.PolicyGroupId
}

func (s *ExportDesktopGroupInfoRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportDesktopGroupInfoRequest) GetTag() []*ExportDesktopGroupInfoRequestTag  {
  return s.Tag
}

func (s *ExportDesktopGroupInfoRequest) SetChargeType(v string) *ExportDesktopGroupInfoRequest {
  s.ChargeType = &v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetDesktopGroupId(v []*string) *ExportDesktopGroupInfoRequest {
  s.DesktopGroupId = v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetDesktopGroupName(v string) *ExportDesktopGroupInfoRequest {
  s.DesktopGroupName = &v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetEndUserId(v []*string) *ExportDesktopGroupInfoRequest {
  s.EndUserId = v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetExpiredTime(v string) *ExportDesktopGroupInfoRequest {
  s.ExpiredTime = &v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetLangType(v string) *ExportDesktopGroupInfoRequest {
  s.LangType = &v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetMaxResults(v int32) *ExportDesktopGroupInfoRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetNextToken(v string) *ExportDesktopGroupInfoRequest {
  s.NextToken = &v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetOfficeSiteId(v string) *ExportDesktopGroupInfoRequest {
  s.OfficeSiteId = &v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetPolicyGroupId(v string) *ExportDesktopGroupInfoRequest {
  s.PolicyGroupId = &v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetRegionId(v string) *ExportDesktopGroupInfoRequest {
  s.RegionId = &v
  return s
}

func (s *ExportDesktopGroupInfoRequest) SetTag(v []*ExportDesktopGroupInfoRequestTag) *ExportDesktopGroupInfoRequest {
  s.Tag = v
  return s
}

func (s *ExportDesktopGroupInfoRequest) Validate() error {
  return dara.Validate(s)
}

type ExportDesktopGroupInfoRequestTag struct {
  // The tag key. You cannot specify an empty string as a tag key. A tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
  // 
  // example:
  // 
  // TestKey
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // The tag value. You can specify an empty string as a tag key. A tag value can be up to 128 characters in length and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
  // 
  // example:
  // 
  // TestValue
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExportDesktopGroupInfoRequestTag) String() string {
  return dara.Prettify(s)
}

func (s ExportDesktopGroupInfoRequestTag) GoString() string {
  return s.String()
}

func (s *ExportDesktopGroupInfoRequestTag) GetKey() *string  {
  return s.Key
}

func (s *ExportDesktopGroupInfoRequestTag) GetValue() *string  {
  return s.Value
}

func (s *ExportDesktopGroupInfoRequestTag) SetKey(v string) *ExportDesktopGroupInfoRequestTag {
  s.Key = &v
  return s
}

func (s *ExportDesktopGroupInfoRequestTag) SetValue(v string) *ExportDesktopGroupInfoRequestTag {
  s.Value = &v
  return s
}

func (s *ExportDesktopGroupInfoRequestTag) Validate() error {
  return dara.Validate(s)
}

