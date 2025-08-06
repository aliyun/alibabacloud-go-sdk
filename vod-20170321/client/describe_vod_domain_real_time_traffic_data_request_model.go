// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealTimeTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeTrafficDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeVodDomainRealTimeTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainRealTimeTrafficDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainRealTimeTrafficDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRealTimeTrafficDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRealTimeTrafficDataRequest struct {
	// You can specify multiple domain names and separate them with commas (,). You can specify a maximum of 500 domain names in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the ISP. If you do not set this parameter, all ISPs are queried.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. If you do not set this parameter, data in all regions is queried.
	//
	// example:
	//
	// shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) SetDomainName(v string) *DescribeVodDomainRealTimeTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) SetEndTime(v string) *DescribeVodDomainRealTimeTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) SetIspNameEn(v string) *DescribeVodDomainRealTimeTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) SetLocationNameEn(v string) *DescribeVodDomainRealTimeTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) SetOwnerId(v int64) *DescribeVodDomainRealTimeTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) SetStartTime(v string) *DescribeVodDomainRealTimeTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
