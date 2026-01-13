// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDesktopRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *DescribeGlobalDesktopRecordsRequest
	GetDesktopId() []*string
	SetDesktopName(v string) *DescribeGlobalDesktopRecordsRequest
	GetDesktopName() *string
	SetDesktopStatusList(v []*string) *DescribeGlobalDesktopRecordsRequest
	GetDesktopStatusList() []*string
	SetDesktopType(v string) *DescribeGlobalDesktopRecordsRequest
	GetDesktopType() *string
	SetEndTime(v string) *DescribeGlobalDesktopRecordsRequest
	GetEndTime() *string
	SetEndUserId(v string) *DescribeGlobalDesktopRecordsRequest
	GetEndUserId() *string
	SetExcludeDesktopStatusList(v []*string) *DescribeGlobalDesktopRecordsRequest
	GetExcludeDesktopStatusList() []*string
	SetOfficeSiteId(v string) *DescribeGlobalDesktopRecordsRequest
	GetOfficeSiteId() *string
	SetOrderBy(v string) *DescribeGlobalDesktopRecordsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeGlobalDesktopRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGlobalDesktopRecordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeGlobalDesktopRecordsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGlobalDesktopRecordsRequest
	GetResourceGroupId() *string
	SetScope(v string) *DescribeGlobalDesktopRecordsRequest
	GetScope() *string
	SetSortType(v string) *DescribeGlobalDesktopRecordsRequest
	GetSortType() *string
	SetStartTime(v string) *DescribeGlobalDesktopRecordsRequest
	GetStartTime() *string
	SetSubPayType(v string) *DescribeGlobalDesktopRecordsRequest
	GetSubPayType() *string
}

type DescribeGlobalDesktopRecordsRequest struct {
	// The cloud computer IDs. You can specify 1 to 100 office network IDs.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The name of the cloud computer.
	//
	// example:
	//
	// DemoComputer
	DesktopName       *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatusList []*string `json:"DesktopStatusList,omitempty" xml:"DesktopStatusList,omitempty" type:"Repeated"`
	// The cloud computer type. You can call the [DescribeDesktopTypes](https://help.aliyun.com/document_detail/188882.html) operation to query the IDs of the specifications supported by the cloud computer.
	//
	// example:
	//
	// eds.enterprise_office.2c4g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The end time. The interval between the start time and end time can be up to 30 days. Supported formats:
	//
	// 	- Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2022-08-31T06:56:45Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end user ID.
	//
	// example:
	//
	// TestUser
	EndUserId                *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ExcludeDesktopStatusList []*string `json:"ExcludeDesktopStatusList,omitempty" xml:"ExcludeDesktopStatusList,omitempty" type:"Repeated"`
	// The office network IDs.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The sorting field. If this parameter is not provided, results are sorted by creation time in descending order. Valid values:
	//
	// 	- uptime: indicates that the cloud computers are sorted by startup duration.
	//
	// example:
	//
	// uptime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number of the current page.\\
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// 	- China (Shanghai)
	//
	// 	- Singapore
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-3mtuc28rx95lx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The query range. This parameter is empty by default. Optional values are:
	//
	// 	- ADVANCED: indicates that statistics such as the connection duration are queried.
	//
	// example:
	//
	// ADVANCED
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The sorting method. Default value: ascending. Valid value:
	//
	// 	- Asc: ascending order
	//
	// 	- Desc: descending.
	//
	// example:
	//
	// Asc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The start time. Supported formats:
	//
	// 	- Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2022-03-23T04:10:21Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The way to purchase cloud computers. Valid values:
	//
	// 	- prePaid: The monthly purchase is unlimited.
	//
	// 	- postPaid: pay-as-you-go
	//
	// 	- monthPackage: monthly duration.
	//
	// example:
	//
	// monthPackage
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
}

func (s DescribeGlobalDesktopRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopRecordsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *DescribeGlobalDesktopRecordsRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeGlobalDesktopRecordsRequest) GetDesktopStatusList() []*string {
	return s.DesktopStatusList
}

func (s *DescribeGlobalDesktopRecordsRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeGlobalDesktopRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeGlobalDesktopRecordsRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeGlobalDesktopRecordsRequest) GetExcludeDesktopStatusList() []*string {
	return s.ExcludeDesktopStatusList
}

func (s *DescribeGlobalDesktopRecordsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeGlobalDesktopRecordsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeGlobalDesktopRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGlobalDesktopRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGlobalDesktopRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalDesktopRecordsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGlobalDesktopRecordsRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeGlobalDesktopRecordsRequest) GetSortType() *string {
	return s.SortType
}

func (s *DescribeGlobalDesktopRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeGlobalDesktopRecordsRequest) GetSubPayType() *string {
	return s.SubPayType
}

func (s *DescribeGlobalDesktopRecordsRequest) SetDesktopId(v []*string) *DescribeGlobalDesktopRecordsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetDesktopName(v string) *DescribeGlobalDesktopRecordsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetDesktopStatusList(v []*string) *DescribeGlobalDesktopRecordsRequest {
	s.DesktopStatusList = v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetDesktopType(v string) *DescribeGlobalDesktopRecordsRequest {
	s.DesktopType = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetEndTime(v string) *DescribeGlobalDesktopRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetEndUserId(v string) *DescribeGlobalDesktopRecordsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetExcludeDesktopStatusList(v []*string) *DescribeGlobalDesktopRecordsRequest {
	s.ExcludeDesktopStatusList = v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetOfficeSiteId(v string) *DescribeGlobalDesktopRecordsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetOrderBy(v string) *DescribeGlobalDesktopRecordsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetPageNumber(v int32) *DescribeGlobalDesktopRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetPageSize(v int32) *DescribeGlobalDesktopRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetRegionId(v string) *DescribeGlobalDesktopRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetResourceGroupId(v string) *DescribeGlobalDesktopRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetScope(v string) *DescribeGlobalDesktopRecordsRequest {
	s.Scope = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetSortType(v string) *DescribeGlobalDesktopRecordsRequest {
	s.SortType = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetStartTime(v string) *DescribeGlobalDesktopRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) SetSubPayType(v string) *DescribeGlobalDesktopRecordsRequest {
	s.SubPayType = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsRequest) Validate() error {
	return dara.Validate(s)
}
