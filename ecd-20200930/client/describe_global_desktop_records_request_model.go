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
	SetDesktopType(v string) *DescribeGlobalDesktopRecordsRequest
	GetDesktopType() *string
	SetEndTime(v string) *DescribeGlobalDesktopRecordsRequest
	GetEndTime() *string
	SetEndUserId(v string) *DescribeGlobalDesktopRecordsRequest
	GetEndUserId() *string
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
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// example:
	//
	// DemoComputer
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// eds.enterprise_office.2c4g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// example:
	//
	// 2022-08-31T06:56:45Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TestUser
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// uptime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-3mtuc28rx95lx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ADVANCED
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// Asc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// example:
	//
	// 2022-03-23T04:10:21Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeGlobalDesktopRecordsRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeGlobalDesktopRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeGlobalDesktopRecordsRequest) GetEndUserId() *string {
	return s.EndUserId
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
