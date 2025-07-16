// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainSrcBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainSrcBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainSrcBpsDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDomainSrcBpsDataRequest
	GetStartTime() *string
}

type DescribeDomainSrcBpsDataRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each call. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries the geographic distribution of users for all accelerated domain names.
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
	// 2019-12-10T20:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Description**.
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
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainSrcBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainSrcBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcBpsDataRequest) SetDomainName(v string) *DescribeDomainSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetEndTime(v string) *DescribeDomainSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetInterval(v string) *DescribeDomainSrcBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetStartTime(v string) *DescribeDomainSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
