// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainSrcTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainSrcTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainSrcTrafficDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDomainSrcTrafficDataRequest
	GetStartTime() *string
}

type DescribeDomainSrcTrafficDataRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries the origin traffic for all accelerated domain names that belong to your Alibaba Cloud account.
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
	// 2015-12-10T21:00:00Z
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
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainSrcTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainSrcTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcTrafficDataRequest) SetDomainName(v string) *DescribeDomainSrcTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataRequest) SetEndTime(v string) *DescribeDomainSrcTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataRequest) SetInterval(v string) *DescribeDomainSrcTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataRequest) SetStartTime(v string) *DescribeDomainSrcTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
