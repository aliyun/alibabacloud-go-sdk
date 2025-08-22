// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeSrcTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeSrcTrafficDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2015-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
