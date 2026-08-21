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
	// The accelerated domain name to query.
	//
	// - If you do not specify this parameter, the pooled data of all accelerated domain names is returned by default.
	//
	// - Batch queries are supported. Separate multiple domain names with commas (,). You can specify up to 500 domain names at a time.
	//
	// - You can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com), and choose **Configuration Management > CDN Configuration > Domain Names*	- in the left-side navigation pane to view the accelerated domain names that you have added to ApsaraVideo VOD. You can also invoke the [DescribeVodUserDomains](~~DescribeVodUserDomains~~) operation to query the list of accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2015-12-10T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data. Unit: seconds. Valid values: **300**, **3600**, and **86400**. If you do not specify this parameter or specify an unsupported value, the default value is used. The supported time granularity varies based on the time span specified by `StartTime` and `EndTime`:
	//
	// - Less than 3 days (excluding exactly 3 days): **300*	- (default), **3600**, and **86400**.
	//
	// - 3 to 31 days (excluding exactly 31 days): **3600*	- (default) and **86400**.
	//
	// - 31 days or more: **86400*	- (default).
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP) in English. If you do not specify this parameter, data of all ISPs is queried by default.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region in English. If you do not specify this parameter, data of all regions is queried by default. Currently, only the Shanghai region is supported.
	//
	// example:
	//
	// shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
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
