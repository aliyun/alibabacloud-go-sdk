// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainMultiUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainMultiUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainMultiUsageDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainMultiUsageDataRequest
	GetStartTime() *string
}

type DescribeDomainMultiUsageDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// > 	- You can specify a maximum of 30 domain names at a time.
	//
	// >	- If this parameter is not set, data of all your accelerated domain names is queried.
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
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainMultiUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMultiUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainMultiUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainMultiUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainMultiUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainMultiUsageDataRequest) SetDomainName(v string) *DescribeDomainMultiUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainMultiUsageDataRequest) SetEndTime(v string) *DescribeDomainMultiUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainMultiUsageDataRequest) SetStartTime(v string) *DescribeDomainMultiUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainMultiUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
