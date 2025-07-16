// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeSrcTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRealTimeSrcTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeSrcTrafficDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainRealTimeSrcTrafficDataRequest
	GetStartTime() *string
}

type DescribeDomainRealTimeSrcTrafficDataRequest struct {
	// The accelerated domain name. You can specify up to 100 domain names in each call. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
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
	// 2019-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeSrcTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) SetDomainName(v string) *DescribeDomainRealTimeSrcTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) SetEndTime(v string) *DescribeDomainRealTimeSrcTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) SetStartTime(v string) *DescribeDomainRealTimeSrcTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
