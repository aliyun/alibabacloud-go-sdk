// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainSrcTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainSrcTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainSrcTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainSrcTrafficDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeVodDomainSrcTrafficDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainSrcTrafficDataRequest
	GetStartTime() *string
}

type DescribeVodDomainSrcTrafficDataRequest struct {
	// The accelerated domain name. You can specify a maximum of 500 domain names in a request. Separate multiple domain names with commas (,). If you specify multiple domain names in a request, aggregation results are returned.
	//
	// If you leave this parameter empty, the origin traffic data for all accelerated domain names is queried by default.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-09-24T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data entries to return. Unit: seconds. Valid values:
	//
	// 	- **300**: 5 minutes
	//
	// 	- **3600**: 1 hour
	//
	// 	- **86400**: 1 day
	//
	// > The time granularity supported by the Interval parameter varies based on the time range per query specified by using `StartTime` and `EndTime`. For more information, see the **Time granularity*	- section of this topic.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// If you leave this parameter empty, the origin traffic data that is generated in the last 24 hours is queried by default.
	//
	// example:
	//
	// 2022-03-22T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainSrcTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainSrcTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainSrcTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainSrcTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainSrcTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainSrcTrafficDataRequest) SetDomainName(v string) *DescribeVodDomainSrcTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataRequest) SetEndTime(v string) *DescribeVodDomainSrcTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataRequest) SetInterval(v string) *DescribeVodDomainSrcTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataRequest) SetOwnerId(v int64) *DescribeVodDomainSrcTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataRequest) SetStartTime(v string) *DescribeVodDomainSrcTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
