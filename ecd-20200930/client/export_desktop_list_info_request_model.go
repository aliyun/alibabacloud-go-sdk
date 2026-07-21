// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDesktopListInfoRequest interface {
  dara.Model
  String() string
  GoString() string
  SetChargeType(v string) *ExportDesktopListInfoRequest
  GetChargeType() *string 
  SetDesktopId(v []*string) *ExportDesktopListInfoRequest
  GetDesktopId() []*string 
  SetDesktopName(v string) *ExportDesktopListInfoRequest
  GetDesktopName() *string 
  SetDesktopStatus(v string) *ExportDesktopListInfoRequest
  GetDesktopStatus() *string 
  SetEndUserId(v []*string) *ExportDesktopListInfoRequest
  GetEndUserId() []*string 
  SetExpiredTime(v string) *ExportDesktopListInfoRequest
  GetExpiredTime() *string 
  SetGroupId(v string) *ExportDesktopListInfoRequest
  GetGroupId() *string 
  SetLangType(v string) *ExportDesktopListInfoRequest
  GetLangType() *string 
  SetMaxResults(v int32) *ExportDesktopListInfoRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *ExportDesktopListInfoRequest
  GetNextToken() *string 
  SetOfficeSiteId(v string) *ExportDesktopListInfoRequest
  GetOfficeSiteId() *string 
  SetPolicyGroupId(v string) *ExportDesktopListInfoRequest
  GetPolicyGroupId() *string 
  SetRegionId(v string) *ExportDesktopListInfoRequest
  GetRegionId() *string 
  SetTag(v []*ExportDesktopListInfoRequestTag) *ExportDesktopListInfoRequest
  GetTag() []*ExportDesktopListInfoRequestTag 
  SetUserName(v string) *ExportDesktopListInfoRequest
  GetUserName() *string 
}

type ExportDesktopListInfoRequest struct {
  // The billing method of the cloud desktop.
  // 
  // example:
  // 
  // PostPaid
  ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
  // The cloud desktop IDs. You can specify 1 to 100 IDs.
  DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
  // The name of the cloud desktop.
  // 
  // example:
  // 
  // DemoComputer01
  DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
  // The status of the cloud desktop.
  // 
  // example:
  // 
  // Running
  DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
  // The list of authorized usernames for the cloud desktop. You can specify 1 to 100 usernames.
  // 
  // > Only one user can connect to and use the cloud desktop at a time.
  EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
  // The expiration time of the subscription cloud desktop.
  // 
  // example:
  // 
  // 2022-12-31T15:59Z
  ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
  // The ID of the cloud desktop pool to which the cloud desktop belongs.
  // 
  // example:
  // 
  // dg-boyczi8enfyc5***
  GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The language type displayed on the frontend page. The backend sets the language type of the exported file based on this value.
  // 
  // example:
  // 
  // `zh-CN`
  LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
  // The number of entries per page for a paged query.
  // 
  // Maximum value: 100.
  // 
  // Default value: 10.
  // 
  // example:
  // 
  // 10
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // The pagination token for the next query. An empty value indicates that there are no more results.
  // 
  // example:
  // 
  // caeba0bbb2be03f84eb48b699f0a4883
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  // The office network ID.
  // 
  // example:
  // 
  // cn-hangzhou+dir-363353****
  OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
  // The ID of the policy associated with the cloud desktop.
  // 
  // example:
  // 
  // system-all-enabled-policy
  PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
  // The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by WUYING Workspace.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The list of tags. A tag consists of a key-value pair and is used to mark resources. You can use tags to group and manage cloud desktops for easier searching and batch operations. For more information, see [Use tags to manage cloud desktops](https://help.aliyun.com/document_detail/203781.html).
  Tag []*ExportDesktopListInfoRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
  // The username of the user who is currently using the cloud desktop.
  // 
  // example:
  // 
  // alice
  UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ExportDesktopListInfoRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportDesktopListInfoRequest) GoString() string {
  return s.String()
}

func (s *ExportDesktopListInfoRequest) GetChargeType() *string  {
  return s.ChargeType
}

func (s *ExportDesktopListInfoRequest) GetDesktopId() []*string  {
  return s.DesktopId
}

func (s *ExportDesktopListInfoRequest) GetDesktopName() *string  {
  return s.DesktopName
}

func (s *ExportDesktopListInfoRequest) GetDesktopStatus() *string  {
  return s.DesktopStatus
}

func (s *ExportDesktopListInfoRequest) GetEndUserId() []*string  {
  return s.EndUserId
}

func (s *ExportDesktopListInfoRequest) GetExpiredTime() *string  {
  return s.ExpiredTime
}

func (s *ExportDesktopListInfoRequest) GetGroupId() *string  {
  return s.GroupId
}

func (s *ExportDesktopListInfoRequest) GetLangType() *string  {
  return s.LangType
}

func (s *ExportDesktopListInfoRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *ExportDesktopListInfoRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportDesktopListInfoRequest) GetOfficeSiteId() *string  {
  return s.OfficeSiteId
}

func (s *ExportDesktopListInfoRequest) GetPolicyGroupId() *string  {
  return s.PolicyGroupId
}

func (s *ExportDesktopListInfoRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportDesktopListInfoRequest) GetTag() []*ExportDesktopListInfoRequestTag  {
  return s.Tag
}

func (s *ExportDesktopListInfoRequest) GetUserName() *string  {
  return s.UserName
}

func (s *ExportDesktopListInfoRequest) SetChargeType(v string) *ExportDesktopListInfoRequest {
  s.ChargeType = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetDesktopId(v []*string) *ExportDesktopListInfoRequest {
  s.DesktopId = v
  return s
}

func (s *ExportDesktopListInfoRequest) SetDesktopName(v string) *ExportDesktopListInfoRequest {
  s.DesktopName = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetDesktopStatus(v string) *ExportDesktopListInfoRequest {
  s.DesktopStatus = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetEndUserId(v []*string) *ExportDesktopListInfoRequest {
  s.EndUserId = v
  return s
}

func (s *ExportDesktopListInfoRequest) SetExpiredTime(v string) *ExportDesktopListInfoRequest {
  s.ExpiredTime = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetGroupId(v string) *ExportDesktopListInfoRequest {
  s.GroupId = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetLangType(v string) *ExportDesktopListInfoRequest {
  s.LangType = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetMaxResults(v int32) *ExportDesktopListInfoRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetNextToken(v string) *ExportDesktopListInfoRequest {
  s.NextToken = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetOfficeSiteId(v string) *ExportDesktopListInfoRequest {
  s.OfficeSiteId = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetPolicyGroupId(v string) *ExportDesktopListInfoRequest {
  s.PolicyGroupId = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetRegionId(v string) *ExportDesktopListInfoRequest {
  s.RegionId = &v
  return s
}

func (s *ExportDesktopListInfoRequest) SetTag(v []*ExportDesktopListInfoRequestTag) *ExportDesktopListInfoRequest {
  s.Tag = v
  return s
}

func (s *ExportDesktopListInfoRequest) SetUserName(v string) *ExportDesktopListInfoRequest {
  s.UserName = &v
  return s
}

func (s *ExportDesktopListInfoRequest) Validate() error {
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

type ExportDesktopListInfoRequestTag struct {
  // The tag key. If you specify `Tag`, `Key` is required. The tag key cannot exceed 128 characters, cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`. It also cannot consist of only spaces.
  // 
  // example:
  // 
  // TestKey
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // The tag value. The tag value cannot exceed 128 characters, cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
  // 
  // example:
  // 
  // TestValue
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExportDesktopListInfoRequestTag) String() string {
  return dara.Prettify(s)
}

func (s ExportDesktopListInfoRequestTag) GoString() string {
  return s.String()
}

func (s *ExportDesktopListInfoRequestTag) GetKey() *string  {
  return s.Key
}

func (s *ExportDesktopListInfoRequestTag) GetValue() *string  {
  return s.Value
}

func (s *ExportDesktopListInfoRequestTag) SetKey(v string) *ExportDesktopListInfoRequestTag {
  s.Key = &v
  return s
}

func (s *ExportDesktopListInfoRequestTag) SetValue(v string) *ExportDesktopListInfoRequestTag {
  s.Value = &v
  return s
}

func (s *ExportDesktopListInfoRequestTag) Validate() error {
  return dara.Validate(s)
}

