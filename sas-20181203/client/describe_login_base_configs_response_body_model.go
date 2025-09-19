// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoginBaseConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaseConfigs(v []*DescribeLoginBaseConfigsResponseBodyBaseConfigs) *DescribeLoginBaseConfigsResponseBody
	GetBaseConfigs() []*DescribeLoginBaseConfigsResponseBodyBaseConfigs
	SetCurrentPage(v int32) *DescribeLoginBaseConfigsResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeLoginBaseConfigsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLoginBaseConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLoginBaseConfigsResponseBody
	GetTotalCount() *int32
}

type DescribeLoginBaseConfigsResponseBody struct {
	// The description of the configuration.
	BaseConfigs []*DescribeLoginBaseConfigsResponseBodyBaseConfigs `json:"BaseConfigs,omitempty" xml:"BaseConfigs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2C2D4B3C-0524-17B1-93D2-DA50119F4E1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoginBaseConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginBaseConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoginBaseConfigsResponseBody) GetBaseConfigs() []*DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	return s.BaseConfigs
}

func (s *DescribeLoginBaseConfigsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeLoginBaseConfigsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoginBaseConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoginBaseConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLoginBaseConfigsResponseBody) SetBaseConfigs(v []*DescribeLoginBaseConfigsResponseBodyBaseConfigs) *DescribeLoginBaseConfigsResponseBody {
	s.BaseConfigs = v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBody) SetCurrentPage(v int32) *DescribeLoginBaseConfigsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBody) SetPageSize(v int32) *DescribeLoginBaseConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBody) SetRequestId(v string) *DescribeLoginBaseConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBody) SetTotalCount(v int32) *DescribeLoginBaseConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLoginBaseConfigsResponseBodyBaseConfigs struct {
	// The common logon account.
	//
	// example:
	//
	// 1582318****
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The end time of the common logon time range.
	//
	// example:
	//
	// 07:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The common logon IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The common logon location.
	//
	// example:
	//
	// Montenegro
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// Corresponding configuration remark information.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The start time of the common logon time range.
	//
	// example:
	//
	// 08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The details of the servers to which the configuration is applied.
	TargetList []*DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The total number of servers.
	//
	// example:
	//
	// 172
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of servers to which the configuration is applied.
	//
	// example:
	//
	// 13
	UuidCount *int32 `json:"UuidCount,omitempty" xml:"UuidCount,omitempty"`
}

func (s DescribeLoginBaseConfigsResponseBodyBaseConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginBaseConfigsResponseBodyBaseConfigs) GoString() string {
	return s.String()
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) GetAccount() *string {
	return s.Account
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) GetIp() *string {
	return s.Ip
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) GetLocation() *string {
	return s.Location
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) GetRemark() *string {
	return s.Remark
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) GetTargetList() []*DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList {
	return s.TargetList
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) GetUuidCount() *int32 {
	return s.UuidCount
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) SetAccount(v string) *DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	s.Account = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) SetEndTime(v string) *DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	s.EndTime = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) SetIp(v string) *DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	s.Ip = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) SetLocation(v string) *DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	s.Location = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) SetRemark(v string) *DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	s.Remark = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) SetStartTime(v string) *DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	s.StartTime = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) SetTargetList(v []*DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList) *DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	s.TargetList = v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) SetTotalCount(v int32) *DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	s.TotalCount = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) SetUuidCount(v int32) *DescribeLoginBaseConfigsResponseBodyBaseConfigs {
	s.UuidCount = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList struct {
	// The UUID or group ID of the server.
	//
	// example:
	//
	// 0011ea53-738c-4bff-93be-ce6a1cc9****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the server to which the configuration is applied. Valid values:
	//
	// 	- **uuid**: a server
	//
	// 	- **groupId**: a server group
	//
	// 	- **global**: all servers
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList) GoString() string {
	return s.String()
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList) GetTarget() *string {
	return s.Target
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList) SetTarget(v string) *DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList {
	s.Target = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList) SetTargetType(v string) *DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList {
	s.TargetType = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponseBodyBaseConfigsTargetList) Validate() error {
	return dara.Validate(s)
}
