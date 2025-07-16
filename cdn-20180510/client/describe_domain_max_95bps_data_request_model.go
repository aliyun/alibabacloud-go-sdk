// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainMax95BpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCycle(v string) *DescribeDomainMax95BpsDataRequest
	GetCycle() *string
	SetDomainName(v string) *DescribeDomainMax95BpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainMax95BpsDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainMax95BpsDataRequest
	GetStartTime() *string
	SetTimePoint(v string) *DescribeDomainMax95BpsDataRequest
	GetTimePoint() *string
}

type DescribeDomainMax95BpsDataRequest struct {
	// The cycle to query the 95th percentile bandwidth data. Default value: **day**. Valid values:
	//
	// 	- **day**: queries the 95th percentile bandwidth data by day.
	//
	// 	- **month**: queries the 95th percentile bandwidth data by month.
	//
	// example:
	//
	// month
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The accelerated domain name. If you do not specify a domain name, data of all domain names is queried.
	//
	// > You cannot specify multiple domain names in a DescribeDomainMax95BpsData request.
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
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-21T10:00:00Z
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s DescribeDomainMax95BpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMax95BpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainMax95BpsDataRequest) GetCycle() *string {
	return s.Cycle
}

func (s *DescribeDomainMax95BpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainMax95BpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainMax95BpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainMax95BpsDataRequest) GetTimePoint() *string {
	return s.TimePoint
}

func (s *DescribeDomainMax95BpsDataRequest) SetCycle(v string) *DescribeDomainMax95BpsDataRequest {
	s.Cycle = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) SetDomainName(v string) *DescribeDomainMax95BpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) SetEndTime(v string) *DescribeDomainMax95BpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) SetStartTime(v string) *DescribeDomainMax95BpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) SetTimePoint(v string) *DescribeDomainMax95BpsDataRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) Validate() error {
	return dara.Validate(s)
}
