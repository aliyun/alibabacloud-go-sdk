// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRegionIdResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeEnsRegionIdResourceRequest
	GetEndTime() *string
	SetIsp(v string) *DescribeEnsRegionIdResourceRequest
	GetIsp() *string
	SetOrderByParams(v string) *DescribeEnsRegionIdResourceRequest
	GetOrderByParams() *string
	SetPageNumber(v int32) *DescribeEnsRegionIdResourceRequest
	GetPageNumber() *int32
	SetPageSize(v string) *DescribeEnsRegionIdResourceRequest
	GetPageSize() *string
	SetStartTime(v string) *DescribeEnsRegionIdResourceRequest
	GetStartTime() *string
}

type DescribeEnsRegionIdResourceRequest struct {
	// The end time of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-06-16T06:33:15
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The Internet service provider (ISP). Valid values:
	//
	// 	- cmcc: China Mobile
	//
	// 	- telecom: China Telecom
	//
	// 	- unicom: China Unicom
	//
	// 	- multiCarrier: multi-line ISP
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The order in which the resources to return are sorted. Valid values:
	//
	// 	- InstanceCount: desc
	//
	// 	- Area: asc
	//
	// 	- InternetBandwidth: asc
	//
	// example:
	//
	// InstanceCount: desc
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-06-16T06:33:15Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEnsRegionIdResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionIdResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEnsRegionIdResourceRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeEnsRegionIdResourceRequest) GetOrderByParams() *string {
	return s.OrderByParams
}

func (s *DescribeEnsRegionIdResourceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnsRegionIdResourceRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeEnsRegionIdResourceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEnsRegionIdResourceRequest) SetEndTime(v string) *DescribeEnsRegionIdResourceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetIsp(v string) *DescribeEnsRegionIdResourceRequest {
	s.Isp = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetOrderByParams(v string) *DescribeEnsRegionIdResourceRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetPageNumber(v int32) *DescribeEnsRegionIdResourceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetPageSize(v string) *DescribeEnsRegionIdResourceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetStartTime(v string) *DescribeEnsRegionIdResourceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) Validate() error {
	return dara.Validate(s)
}
