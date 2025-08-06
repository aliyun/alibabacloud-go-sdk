// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVodDomainBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainBpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainBpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainBpsDataRequest
	GetStartTime() *string
}

type DescribeVodDomainBpsDataRequest struct {
	// The domain name to be queried. If you do not specify this parameter, the merged data of all your domain names for CDN is returned. You can specify multiple domain names. Separate them with commas (,).
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2015-12-10T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The query interval. Unit: seconds. Valid values: **300**, **3600**, and **86400**.
	//
	// 	- If the time range to query is less than 3 days, valid values are **300**, **3600**, and **86400**. The default value is 300.
	//
	// 	- If the time range to query is from 3 to less than 31 days, valid values are **3600*	- and **86400**. The default value is 3600.
	//
	// 	- If the time range to query is from 31 to 90 days, the valid value is **86400**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). If you do not specify this parameter, the data of all ISPs is returned.
	//
	// example:
	//
	// Alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. If you do not specify this parameter, the data in all regions is returned. Only data in the China (Shanghai) region can be queried.
	//
	// example:
	//
	// cn-shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > The minimum query interval is 5 minutes. If you do not specify this parameter, the data in the last 24 hours is queried.
	//
	// example:
	//
	// 2015-12-10T13:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainBpsDataRequest) SetDomainName(v string) *DescribeVodDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetEndTime(v string) *DescribeVodDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetInterval(v string) *DescribeVodDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetIspNameEn(v string) *DescribeVodDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeVodDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetStartTime(v string) *DescribeVodDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
