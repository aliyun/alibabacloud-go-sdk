// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *DescribeFlowStatisticRequest
	GetDesktopId() *string
	SetOfficeSiteId(v string) *DescribeFlowStatisticRequest
	GetOfficeSiteId() *string
	SetPageNumber(v int32) *DescribeFlowStatisticRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFlowStatisticRequest
	GetPageSize() *int32
	SetPeriod(v int32) *DescribeFlowStatisticRequest
	GetPeriod() *int32
	SetRegionId(v string) *DescribeFlowStatisticRequest
	GetRegionId() *string
}

type DescribeFlowStatisticRequest struct {
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The number of the page to return.\\
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The statistic collection interval. Unit: seconds.
	//
	// Valid values:
	//
	// 	- 3600: 1 hour
	//
	// 	- 10800: 3 hours
	//
	// 	- 86400: 24 hours
	//
	// This parameter is required.
	//
	// example:
	//
	// 3600
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowStatisticRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeFlowStatisticRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeFlowStatisticRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFlowStatisticRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFlowStatisticRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeFlowStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowStatisticRequest) SetDesktopId(v string) *DescribeFlowStatisticRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeFlowStatisticRequest) SetOfficeSiteId(v string) *DescribeFlowStatisticRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeFlowStatisticRequest) SetPageNumber(v int32) *DescribeFlowStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowStatisticRequest) SetPageSize(v int32) *DescribeFlowStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowStatisticRequest) SetPeriod(v int32) *DescribeFlowStatisticRequest {
	s.Period = &v
	return s
}

func (s *DescribeFlowStatisticRequest) SetRegionId(v string) *DescribeFlowStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowStatisticRequest) Validate() error {
	return dara.Validate(s)
}
