// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRangeDataByLocateAndIspServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetStartTime() *string
}

type DescribeVodRangeDataByLocateAndIspServiceRequest struct {
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-29T03:55:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the ISP. If you leave this parameter empty, all ISPs are queried.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. If you leave this parameter empty, data in all regions is queried.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-29T03:50:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodRangeDataByLocateAndIspServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRangeDataByLocateAndIspServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetDomainName(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetEndTime(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetIspNameEn(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetLocationNameEn(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetOwnerId(v int64) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetStartTime(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) Validate() error {
	return dara.Validate(s)
}
