// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainSrcBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainSrcBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainSrcBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainSrcBpsDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeVodDomainSrcBpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainSrcBpsDataRequest
	GetStartTime() *string
}

type DescribeVodDomainSrcBpsDataRequest struct {
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
	// 2022-04-26T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity. Unit: seconds. Valid values: **300**, **3600**, and **86400**. If you leave this parameter empty or specify an invalid value, the default value is used. The supported time granularity varies based on the time range specified by `EndTime` and `StartTime`. The following content describes the supported time granularity.
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
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-25T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainSrcBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetDomainName(v string) *DescribeVodDomainSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetEndTime(v string) *DescribeVodDomainSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetInterval(v string) *DescribeVodDomainSrcBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainSrcBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetStartTime(v string) *DescribeVodDomainSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
