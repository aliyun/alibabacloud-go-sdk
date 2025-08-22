// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeTrafficDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeTrafficDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeTrafficDataRequest struct {
	// The accelerated domain name. You can specify one or more domain names and separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The end time must be later than the start time.
	//
	// > If you do not specify StartTime or EndTime, data within the last hour is queried. If you specify both StartTime and EndTime, data within the specified time range is queried.
	//
	// example:
	//
	// 2015-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > If you do not specify StartTime or EndTime, data within the last hour is queried. If you specify both StartTime and EndTime, data within the specified time range is queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
