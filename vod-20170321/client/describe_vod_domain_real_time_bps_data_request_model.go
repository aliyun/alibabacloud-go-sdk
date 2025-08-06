// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealTimeBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeBpsDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeVodDomainRealTimeBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainRealTimeBpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainRealTimeBpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRealTimeBpsDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRealTimeBpsDataRequest struct {
	// You can specify multiple domain names and separate them with commas (,). You can specify a maximum of 500 domain names in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
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
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) SetDomainName(v string) *DescribeVodDomainRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) SetEndTime(v string) *DescribeVodDomainRealTimeBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) SetIspNameEn(v string) *DescribeVodDomainRealTimeBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) SetLocationNameEn(v string) *DescribeVodDomainRealTimeBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainRealTimeBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) SetStartTime(v string) *DescribeVodDomainRealTimeBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
