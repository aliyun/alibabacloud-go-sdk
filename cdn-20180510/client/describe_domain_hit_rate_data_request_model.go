// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainHitRateDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainHitRateDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDomainHitRateDataRequest
	GetStartTime() *string
}

type DescribeDomainHitRateDataRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries the byte hit ratios for all accelerated domain names that belong to your Alibaba Cloud account.
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
	// 2019-12-30T08:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-30T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainHitRateDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainHitRateDataRequest) SetDomainName(v string) *DescribeDomainHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetEndTime(v string) *DescribeDomainHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetInterval(v string) *DescribeDomainHitRateDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetStartTime(v string) *DescribeDomainHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
