// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainOriginTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainOriginTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainOriginTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainOriginTrafficDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDcdnDomainOriginTrafficDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainOriginTrafficDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,). If you do not specify a value for this parameter, all accelerated domain names are queried.
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
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data entries to return. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Description**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainOriginTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainOriginTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainOriginTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) SetInterval(v string) *DescribeDcdnDomainOriginTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainOriginTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
