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
  // The billing method of the shared cloud computer.
  // 
  // example:
  // 
  // PrePaid
  ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
  // The list of shared cloud computer IDs.
  DesktopGroupId []*string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" type:"Repeated"`
  // The name of the shared cloud computer.
  // 
  // example:
  // 
  // CloudComputerPool01
  DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
  // The list of authorized user IDs for the shared cloud computer.
  EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
  // The expiration time of the subscription shared cloud computer. The time is in the ISO 8601 standard (UTC).
  // 
  // example:
  // 
  // 2022-12-31T15:59Z
  ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
  // The language type.
  // 
  // example:
  // 
  // zh-CN
  LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
  // The number of entries per page for a paged query.    
  // 
  // - Maximum value: 100.
  // 
  // - Default value: 10.
  // 
  // example:
  // 
  // 10
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // The pagination token for the next query. An empty value indicates that no more results exist.
  // 
  // example:
  // 
  // caeba0bbb2be03f84eb48b699f0a4883
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  // The ID of the office network to which the shared cloud computer belongs.
  // 
  // example:
  // 
  // cn-hangzhou+dir-467671****
  OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
  // The ID of the policy associated with the shared cloud computer.
  // 
  // example:
  // 
  // pg-53iyi2aar0nd6****
  PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
  // The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by WUYING Workspace.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The list of tags. A maximum of 20 tags can be specified.
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
  if s.Tag != nil {
    for _, item := range s.Tag {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExportDesktopGroupInfoRequestTag struct {
  // The tag key. This parameter cannot be an empty string if specified. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
  // 
  // example:
  // 
  // TestKey
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // The tag value. The tag value can be an empty string and can be up to 128 characters in length. It cannot start with `acs:` or contain `http://` or `https://`.
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

