// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainQpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainQpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainQpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVodDomainQpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainQpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainQpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainQpsDataRequest
	GetStartTime() *string
}

type DescribeVodDomainQpsDataRequest struct {
	// You can specify multiple domain names and separate them with commas (,). You can specify a maximum of 500 domain names in each call.
	//
	// By default, this operation queries the number and proportions of HTTP status codes for all accelerated domain names that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd**THH:mm:ssZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2024-05-02T15:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The interval at which to return data. Unit: seconds. Valid values: **300**, **3600**, and **86400**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). If you leave this parameter empty, data of all ISPs is queried.
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
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd**THH:mm:ssZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2024-05-02T15:50:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainQpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainQpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainQpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainQpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainQpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainQpsDataRequest) SetDomainName(v string) *DescribeVodDomainQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetEndTime(v string) *DescribeVodDomainQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetInterval(v string) *DescribeVodDomainQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetIspNameEn(v string) *DescribeVodDomainQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetLocationNameEn(v string) *DescribeVodDomainQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainQpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetStartTime(v string) *DescribeVodDomainQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
