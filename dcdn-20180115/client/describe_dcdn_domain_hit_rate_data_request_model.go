// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainHitRateDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainHitRateDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDcdnDomainHitRateDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainHitRateDataRequest struct {
	// The accelerated domain name. You can specify only one domain name.
	//
	// If you do not specify a value for this parameter, all domain names are queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time needs to be in UTC.
	//
	// The end time needs to be later than the start time.
	//
	// example:
	//
	// 2018-03-02T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for a query. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time needs to be in UTC.
	//
	// example:
	//
	// 2018-03-02T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainHitRateDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainHitRateDataRequest) SetDomainName(v string) *DescribeDcdnDomainHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataRequest) SetEndTime(v string) *DescribeDcdnDomainHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataRequest) SetInterval(v string) *DescribeDcdnDomainHitRateDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataRequest) SetStartTime(v string) *DescribeDcdnDomainHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
