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
	// The cycle to query the 95th percentile bandwidth data. Valid values:
	//
	// 	- day (default)
	//
	// 	- month
	//
	// example:
	//
	// month
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The domain name to be queried for acceleration. If the parameter is empty, the data merged from all accelerated domain names will be returned by default.
	//
	// > Batch domain name queries are not supported.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// End time point. The date format follows the ISO8601 representation and uses UTC time, in the format yyyy-MM-dd\\"T\\"HH:mm:ssZ.
	//
	// example:
	//
	// 2017-01-12T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Start time point. The date format follows the ISO8601 representation and uses UTC time, in the format yyyy-MM-dd\\"T\\"HH:mm:ssZ.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The start time point for getting the data. The date format follows the ISO8601 representation and uses UTC time, in the format yyyy-MM-dd\\"T\\"HH:mm:ssZ.
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
