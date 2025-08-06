// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealTimeQpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeQpsDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeVodDomainRealTimeQpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainRealTimeQpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainRealTimeQpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRealTimeQpsDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRealTimeQpsDataRequest struct {
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
	// 2019-12-02T11:26:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the Internet service provider (ISP). If you do not set this parameter, all ISPs are queried.
	//
	// example:
	//
	// unicom
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
	// 2019-12-02T11:25:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) SetDomainName(v string) *DescribeVodDomainRealTimeQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) SetEndTime(v string) *DescribeVodDomainRealTimeQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) SetIspNameEn(v string) *DescribeVodDomainRealTimeQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) SetLocationNameEn(v string) *DescribeVodDomainRealTimeQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainRealTimeQpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) SetStartTime(v string) *DescribeVodDomainRealTimeQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
