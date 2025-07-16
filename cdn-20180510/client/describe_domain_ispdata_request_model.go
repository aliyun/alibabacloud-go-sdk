// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainISPDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainISPDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainISPDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainISPDataRequest
	GetStartTime() *string
}

type DescribeDomainISPDataRequest struct {
	// The accelerated domain name. You can specify only one domain name in each call.
	//
	// By default, this operation queries the proportions of data usage for all accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-11-29T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainISPDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainISPDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainISPDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainISPDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainISPDataRequest) SetDomainName(v string) *DescribeDomainISPDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainISPDataRequest) SetEndTime(v string) *DescribeDomainISPDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainISPDataRequest) SetStartTime(v string) *DescribeDomainISPDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainISPDataRequest) Validate() error {
	return dara.Validate(s)
}
