// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeSrcBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeSrcBpsDataRequest struct {
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

func (s DescribeDcdnDomainRealTimeSrcBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
