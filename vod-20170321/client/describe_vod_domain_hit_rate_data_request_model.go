// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainHitRateDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainHitRateDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeVodDomainHitRateDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainHitRateDataRequest
	GetStartTime() *string
}

type DescribeVodDomainHitRateDataRequest struct {
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
	// 2024-01-20T14:59:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity. Unit: seconds. Valid values: **300**, **3600**, and **86400**. If you leave this parameter empty or specify an invalid value, the default value is used. The supported time granularity varies based on the time range specified by `EndTime` and `StartTime`. The following content describes the supported time granularity.
	//
	// 	- Time range per query < 3 days: **300*	- (default), **3600**, and **86400**
	//
	// 	- 3 days ≤ Time range per query < 31 days: **3600*	- (default) and **86400**
	//
	// 	- 31 days ≤ Time range per query ≤ 90 days: **86400*	- (default)
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2024-01-20T13:59:58Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainHitRateDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainHitRateDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainHitRateDataRequest) SetDomainName(v string) *DescribeVodDomainHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainHitRateDataRequest) SetEndTime(v string) *DescribeVodDomainHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainHitRateDataRequest) SetInterval(v string) *DescribeVodDomainHitRateDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainHitRateDataRequest) SetOwnerId(v int64) *DescribeVodDomainHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainHitRateDataRequest) SetStartTime(v string) *DescribeVodDomainHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
