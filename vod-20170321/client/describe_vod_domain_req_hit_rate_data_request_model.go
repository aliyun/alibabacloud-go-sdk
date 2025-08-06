// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainReqHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainReqHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainReqHitRateDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainReqHitRateDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeVodDomainReqHitRateDataRequest
	GetStartTime() *string
}

type DescribeVodDomainReqHitRateDataRequest struct {
	// The accelerated domain name.
	//
	// 	- If you leave this parameter empty, the merged data of all your accelerated domain names is returned.
	//
	// 	- You can specify a maximum of 500 accelerated domain names. Separate multiple domain names with commas (,).
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
	// 2023-12-22T08:00:00Z
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
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainReqHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainReqHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainReqHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainReqHitRateDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainReqHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainReqHitRateDataRequest) SetDomainName(v string) *DescribeVodDomainReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataRequest) SetEndTime(v string) *DescribeVodDomainReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataRequest) SetInterval(v string) *DescribeVodDomainReqHitRateDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataRequest) SetStartTime(v string) *DescribeVodDomainReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
