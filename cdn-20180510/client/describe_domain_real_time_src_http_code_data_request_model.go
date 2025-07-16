// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeSrcHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeDomainRealTimeSrcHttpCodeDataRequest struct {
	// The accelerated domain name. You can specify up to 100 domain names in each call. Separate multiple domain names with commas (,).
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
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-11-30T04:40:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) SetDomainName(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) SetEndTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) SetStartTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
