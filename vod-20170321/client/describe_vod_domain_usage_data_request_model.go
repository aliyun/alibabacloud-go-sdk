// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeVodDomainUsageDataRequest
	GetArea() *string
	SetDomainName(v string) *DescribeVodDomainUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainUsageDataRequest
	GetEndTime() *string
	SetField(v string) *DescribeVodDomainUsageDataRequest
	GetField() *string
	SetInterval(v string) *DescribeVodDomainUsageDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeVodDomainUsageDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainUsageDataRequest
	GetStartTime() *string
	SetType(v string) *DescribeVodDomainUsageDataRequest
	GetType() *string
}

type DescribeVodDomainUsageDataRequest struct {
	// The region code. Default value: CN (the Chinese mainland). Valid values:
	//
	// - **CN**: the Chinese mainland.
	//
	// - **OverSeas**: outside the Chinese mainland.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The accelerated domain name. If this parameter is left empty, the merged data of all accelerated domain names is returned by default. Batch queries are supported. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T12:20:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data type. Valid values:
	//
	// - **bps**: bandwidth.
	//
	// - **traf**: traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// Forces the retrieval of data at the specified time granularity, in seconds. Valid values: **300*	- (5 minutes), **3600*	- (1 hour), and **86400*	- (1 day).
	//
	// - **Interval**=**300**: You can query data for up to the last half year. The maximum time span for a single query is 3 days.
	//
	// - **Interval**=**3600*	- or **86400**: You can query data for up to the last year.
	//
	// - If **Interval*	- is not specified: The maximum time span for a single query is 1 month. If the query time range is 1 to 3 days, data is returned at hourly granularity. If the query time range is 4 days or more, data is returned at daily granularity.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T10:20:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of usage data to retrieve. Valid values:
	//
	//  - **static**: static content.
	//
	// - **dynamic**: dynamic content.
	//
	// - **all**: all content.
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVodDomainUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataRequest) GetArea() *string {
	return s.Area
}

func (s *DescribeVodDomainUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainUsageDataRequest) GetField() *string {
	return s.Field
}

func (s *DescribeVodDomainUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainUsageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainUsageDataRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVodDomainUsageDataRequest) SetArea(v string) *DescribeVodDomainUsageDataRequest {
	s.Area = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetDomainName(v string) *DescribeVodDomainUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetEndTime(v string) *DescribeVodDomainUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetField(v string) *DescribeVodDomainUsageDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetInterval(v string) *DescribeVodDomainUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetOwnerId(v int64) *DescribeVodDomainUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetStartTime(v string) *DescribeVodDomainUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetType(v string) *DescribeVodDomainUsageDataRequest {
	s.Type = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
