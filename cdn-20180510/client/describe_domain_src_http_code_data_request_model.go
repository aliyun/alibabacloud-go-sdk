// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainSrcHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainSrcHttpCodeDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainSrcHttpCodeDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDomainSrcHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeDomainSrcHttpCodeDataRequest struct {
	// The accelerated domain name. You can specify a maximum of 500 domain names in a request. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// example.com,example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format in the ISO 8601 standard. The time is displayed in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
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
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainSrcHttpCodeDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainSrcHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcHttpCodeDataRequest) SetDomainName(v string) *DescribeDomainSrcHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataRequest) SetEndTime(v string) *DescribeDomainSrcHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataRequest) SetInterval(v string) *DescribeDomainSrcHttpCodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataRequest) SetStartTime(v string) *DescribeDomainSrcHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
