// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVsDomainBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVsDomainBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVsDomainBpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVsDomainBpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainBpsDataRequest
	GetStartTime() *string
}

type DescribeVsDomainBpsDataRequest struct {
	// Domain Names. If this parameter is empty, the system returns merged data for all accelerated Domain Names. Enter the accelerated Domain Names to query. Separate multiple Domain Names with commas.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time must be later than the start time. The date format follows ISO8601 notation and uses UTC time. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2021-10-02T02:30:48Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for query data. Supports 300, 3600, and 86400 seconds. If this parameter is not specified or the specified value is not supported, the system uses 300 seconds by default.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The English name of the carrier (ISP). Obtain this from the DescribeCdnRegionAndIsp interface. If not specified, the system queries all carriers (ISPs).
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The English name of the region. Obtain this from the DescribeCdnRegionAndIsp interface. If not specified, the system queries all regions.
	//
	// example:
	//
	// guangdong
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time for data retrieval. The date format follows ISO8601 notation and uses UTC time. Format: YYYY-MM-DDThh:mm:ssZ. The minimum data granularity is 5 minutes. If not specified, the system reads data from the past 24 hours.
	//
	// example:
	//
	// 2021-12-26T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVsDomainBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVsDomainBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVsDomainBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainBpsDataRequest) SetDomainName(v string) *DescribeVsDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetEndTime(v string) *DescribeVsDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetInterval(v string) *DescribeVsDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetIspNameEn(v string) *DescribeVsDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeVsDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetOwnerId(v int64) *DescribeVsDomainBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetStartTime(v string) *DescribeVsDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
