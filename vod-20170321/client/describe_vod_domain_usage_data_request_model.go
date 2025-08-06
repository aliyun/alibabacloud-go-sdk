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
	// The region in which you want to query data. Valid values:
	//
	// 	- **CN**: Chinese mainland
	//
	// 	- **OverSeas**: outside the Chinese mainland
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The accelerated domain name. If you leave this parameter empty, the merged data of all your accelerated domain names is returned. Separate multiple accelerated domain names with commas (,).
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T12:20:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the data to return. Valid values:
	//
	// 	- **bps**: bandwidth
	//
	// 	- **traf**: traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The time interval between the data entries to return. Unit: seconds. Valid values: **300*	- (5 minutes), **3600*	- (1 hour), and **86400*	- (1 day).
	//
	// 	- If **Interval*	- is set to **300**, you can query usage data in the last six months. The maximum time range per query that can be specified is three days.
	//
	// 	- If **Interval*	- is set to **3600*	- or **86400**, you can query usage data of the previous year.
	//
	// 	- If you do not set the **Interval*	- parameter, the maximum time range that you can query is one month. If you specify a time range of 1 to 3 days, the time interval between the entries that are returned is 1 hour. If you specify a time range of at least 4 days, the time interval between the entries that are returned is 1 day.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T10:20:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of content that you want to query. Valid values:
	//
	// 	- **static**: static content
	//
	// 	- **dynamic**: dynamic requests
	//
	// 	- **all**: all content
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
