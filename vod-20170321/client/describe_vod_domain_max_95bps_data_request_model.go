// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainMax95BpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCycle(v string) *DescribeVodDomainMax95BpsDataRequest
	GetCycle() *string
	SetDomainName(v string) *DescribeVodDomainMax95BpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainMax95BpsDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainMax95BpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainMax95BpsDataRequest
	GetStartTime() *string
	SetTimePoint(v string) *DescribeVodDomainMax95BpsDataRequest
	GetTimePoint() *string
}

type DescribeVodDomainMax95BpsDataRequest struct {
	// The cycle for the 95th percentile bandwidth. Default value: day. Valid values:
	//
	// - day: queries the 95th percentile bandwidth by day.
	//
	// - month: queries the 95th percentile bandwidth by month.
	//
	// example:
	//
	// month
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The accelerated domain name to query. If this parameter is left empty, the merged data of all accelerated domain names is returned by default.
	//
	//
	// > Batch domain name queries are not supported.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2017-01-12T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The start time point for data retrieval. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-21T10:00:00Z
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s DescribeVodDomainMax95BpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainMax95BpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainMax95BpsDataRequest) GetCycle() *string {
	return s.Cycle
}

func (s *DescribeVodDomainMax95BpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainMax95BpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainMax95BpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainMax95BpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainMax95BpsDataRequest) GetTimePoint() *string {
	return s.TimePoint
}

func (s *DescribeVodDomainMax95BpsDataRequest) SetCycle(v string) *DescribeVodDomainMax95BpsDataRequest {
	s.Cycle = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataRequest) SetDomainName(v string) *DescribeVodDomainMax95BpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataRequest) SetEndTime(v string) *DescribeVodDomainMax95BpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainMax95BpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataRequest) SetStartTime(v string) *DescribeVodDomainMax95BpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataRequest) SetTimePoint(v string) *DescribeVodDomainMax95BpsDataRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataRequest) Validate() error {
	return dara.Validate(s)
}
