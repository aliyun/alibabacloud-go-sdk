// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVodDomainTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainTrafficDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainTrafficDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainTrafficDataRequest
	GetStartTime() *string
}

type DescribeVodDomainTrafficDataRequest struct {
	// The accelerated domain name.
	//
	// 	- If you leave this parameter empty, the merged data of all your accelerated domain names is returned.
	//
	// 	- You can specify multiple domain names and separate them with commas (,). You can specify a maximum of 500 domain names in each call.
	//
	// 	- To obtain the accelerated domain name, perform the following steps: Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management > CDN Configuration > Domain Names**. On the Domain Names page, view the accelerated domain names. Alternatively, you can call the [DescribeVodUserDomains](~~DescribeVodUserDomains~~) operation to query the accelerated domain names.
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
	// 2019-01-20T14:59:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values: **300**, **3600**, and **86400**. If you leave this parameter empty or specify an invalid value, the default value is used. The supported time granularity varies based on the time range specified by `EndTime` and `StartTime`. The following content describes the supported time granularity.
	//
	// 	- Time range per query < 3 days: **300*	- (default), **3600**, and **86400**
	//
	// 	- 3 days ≤ Time range per query < 31 days: **3600*	- (default) and **86400**
	//
	// 	- 31 days ≤ Time range per query ≤ 366 days: **86400*	- (default)
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). If you leave this parameter empty, all ISPs are queried.
	//
	// example:
	//
	// Alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. If you leave this parameter empty, all regions are queried. You can specify only the China (Shanghai) region.
	//
	// example:
	//
	// cn-shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-01-20T13:59:58Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainTrafficDataRequest) SetDomainName(v string) *DescribeVodDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetEndTime(v string) *DescribeVodDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetInterval(v string) *DescribeVodDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeVodDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeVodDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetOwnerId(v int64) *DescribeVodDomainTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetStartTime(v string) *DescribeVodDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
