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
  // The billing method of the cloud computer.
  // 
  // Default value: Postpaid. Valid values:
  // 
  // 	- Postpaid: pay-as-you-go
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- PrePaid: subscription
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // example:
  // 
  // PostPaid
  ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
  // The IDs of the cloud computers. You can specify 1 to 100 IDs.
  DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
  // The name of the cloud computer.
  // 
  // example:
  // 
  // testName
  DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
  // The status of the cloud computers.
  // 
  // Valid values:
  // 
  // 	- Stopped
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- Starting
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- Rebuilding
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- Running
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- Stopping
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- Expired
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- Deleted
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- Pending
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // example:
  // 
  // Running
  DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
  // The IDs of the end users of the cloud computer. You can specify 1 to 100 IDs.
  // 
  // >  During a specific period of time, only one user can connect to and use the cloud computer.
  EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
  // The time when a subscription cloud computer expires.
  // 
  // example:
  // 
  // 2022-12-31T15:59Z
  ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
  // The ID of the cloud computer pool to which the cloud computers belong.
  // 
  // example:
  // 
  // dg-boyczi8enfyc5***
  GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The language in which the cloud computer is displayed in the console UI. You can export the list of cloud computers in the specified language.
  // 
  // Default value: zh-CN. Valid values:
  // 
  // 	- zh-CN: Simplified Chinese
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- en-GB: British English
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // example:
  // 
  // `zh-CN`
  LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
  // The number of entries per page.
  // 
  // Maximum value: 100.
  // 
  // Default value: 10.
  // 
  // example:
  // 
  // 10
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // The token that is used for the next query. If this parameter is empty, all results are returned.
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
  // The ID of the policy that is attached to the cloud computer.
  // 
  // example:
  // 
  // system-all-enabled-policy
  PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
  // The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The tags that are added to the cloud computer. A tag is a key-value pair that consists of a tag key and a tag value. Tags are used to identify resources. You can use tags to manage cloud computers by group. This facilitates search and batch operations. For more information, see [Use tags to manage cloud computers](https://help.aliyun.com/document_detail/203781.html).
  Tag []*ExportDesktopListInfoRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
  // The username of the end user who is using the cloud computer.
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
  return dara.Validate(s)
}

type ExportDesktopListInfoRequestTag struct {
  // The tag key. If you specify the `Tag` parameter, you must also specify the `Key` parameter. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun` and cannot contain only spaces.
  // 
  // example:
  // 
  // TestKey
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // The tag value. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
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

