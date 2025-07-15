// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSessionStatisticRequest
	GetEndTime() *string
	SetOfficeSiteId(v string) *DescribeSessionStatisticRequest
	GetOfficeSiteId() *string
	SetPeriod(v int32) *DescribeSessionStatisticRequest
	GetPeriod() *int32
	SetRegionId(v string) *DescribeSessionStatisticRequest
	GetRegionId() *string
	SetSearchRegionId(v string) *DescribeSessionStatisticRequest
	GetSearchRegionId() *string
	SetStartTime(v string) *DescribeSessionStatisticRequest
	GetStartTime() *string
}

type DescribeSessionStatisticRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 1677808889806
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// cn-shanghai+dir-259382****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The query interval. Unit: seconds. Valid values:
	//
	// 	- 60
	//
	// 	- 120
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies to search for session information by region ID. This parameter is used to filter desktop information of a specific region.
	//
	// example:
	//
	// cn-hangzhou
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 1679449506572
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSessionStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeSessionStatisticRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSessionStatisticRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeSessionStatisticRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeSessionStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSessionStatisticRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeSessionStatisticRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSessionStatisticRequest) SetEndTime(v string) *DescribeSessionStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSessionStatisticRequest) SetOfficeSiteId(v string) *DescribeSessionStatisticRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSessionStatisticRequest) SetPeriod(v int32) *DescribeSessionStatisticRequest {
	s.Period = &v
	return s
}

func (s *DescribeSessionStatisticRequest) SetRegionId(v string) *DescribeSessionStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSessionStatisticRequest) SetSearchRegionId(v string) *DescribeSessionStatisticRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeSessionStatisticRequest) SetStartTime(v string) *DescribeSessionStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSessionStatisticRequest) Validate() error {
	return dara.Validate(s)
}
