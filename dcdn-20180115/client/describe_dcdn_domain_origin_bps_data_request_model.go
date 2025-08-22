// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainOriginBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainOriginBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainOriginBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainOriginBpsDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDcdnDomainOriginBpsDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainOriginBpsDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,). If you do not specify a value for this parameter, all accelerated domain names are queried.
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
	// 2019-12-11T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for a query. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-10T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainOriginBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainOriginBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainOriginBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) SetInterval(v string) *DescribeDcdnDomainOriginBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainOriginBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
