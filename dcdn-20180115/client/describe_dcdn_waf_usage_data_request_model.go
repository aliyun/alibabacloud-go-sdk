// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnWafUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnWafUsageDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnWafUsageDataRequest
	GetInterval() *string
	SetSplitBy(v string) *DescribeDcdnWafUsageDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeDcdnWafUsageDataRequest
	GetStartTime() *string
}

type DescribeDcdnWafUsageDataRequest struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2018-10-01T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of used SeCUs.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The WAF information about the accelerated domain name.
	//
	// example:
	//
	// domain
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnWafUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnWafUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnWafUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnWafUsageDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeDcdnWafUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnWafUsageDataRequest) SetDomainName(v string) *DescribeDcdnWafUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafUsageDataRequest) SetEndTime(v string) *DescribeDcdnWafUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnWafUsageDataRequest) SetInterval(v string) *DescribeDcdnWafUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnWafUsageDataRequest) SetSplitBy(v string) *DescribeDcdnWafUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeDcdnWafUsageDataRequest) SetStartTime(v string) *DescribeDcdnWafUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnWafUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
