// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeSrcBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRealTimeSrcBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeSrcBpsDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainRealTimeSrcBpsDataRequest
	GetStartTime() *string
}

type DescribeDomainRealTimeSrcBpsDataRequest struct {
	// The accelerated domain name. You can specify up to 100 domain names in each request. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeSrcBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) SetDomainName(v string) *DescribeDomainRealTimeSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) SetEndTime(v string) *DescribeDomainRealTimeSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) SetStartTime(v string) *DescribeDomainRealTimeSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
