// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRangeDataByLocateAndIspServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *DescribeRangeDataByLocateAndIspServiceRequest
	GetDomainNames() *string
	SetEndTime(v string) *DescribeRangeDataByLocateAndIspServiceRequest
	GetEndTime() *string
	SetIspNames(v string) *DescribeRangeDataByLocateAndIspServiceRequest
	GetIspNames() *string
	SetLocationNames(v string) *DescribeRangeDataByLocateAndIspServiceRequest
	GetLocationNames() *string
	SetStartTime(v string) *DescribeRangeDataByLocateAndIspServiceRequest
	GetStartTime() *string
}

type DescribeRangeDataByLocateAndIspServiceRequest struct {
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time. The maximum time range that can be specified is 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the ISP. You can specify only one ISP name in each call.
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query ISPs.
	//
	// example:
	//
	// unicom
	IspNames *string `json:"IspNames,omitempty" xml:"IspNames,omitempty"`
	// The names of the regions. Separate multiple region names with commas (,).
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query the most recent region list.
	//
	// example:
	//
	// liaoning,guangxi
	LocationNames *string `json:"LocationNames,omitempty" xml:"LocationNames,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRangeDataByLocateAndIspServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRangeDataByLocateAndIspServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) GetIspNames() *string {
	return s.IspNames
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) GetLocationNames() *string {
	return s.LocationNames
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetDomainNames(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.DomainNames = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetEndTime(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetIspNames(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.IspNames = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetLocationNames(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.LocationNames = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetStartTime(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) Validate() error {
	return dara.Validate(s)
}
