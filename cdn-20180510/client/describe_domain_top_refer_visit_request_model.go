// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopReferVisitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainTopReferVisitRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainTopReferVisitRequest
	GetEndTime() *string
	SetSortBy(v string) *DescribeDomainTopReferVisitRequest
	GetSortBy() *string
	SetStartTime(v string) *DescribeDomainTopReferVisitRequest
	GetStartTime() *string
}

type DescribeDomainTopReferVisitRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2019-12-22T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The order in which you want to sort the queried information. Valid values:
	//
	// 	- **traf**: by network traffic.
	//
	// 	- **pv**: by the number of page views. This is the default value.
	//
	// example:
	//
	// pv
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-21T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainTopReferVisitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopReferVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainTopReferVisitRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainTopReferVisitRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeDomainTopReferVisitRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainTopReferVisitRequest) SetDomainName(v string) *DescribeDomainTopReferVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopReferVisitRequest) SetEndTime(v string) *DescribeDomainTopReferVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopReferVisitRequest) SetSortBy(v string) *DescribeDomainTopReferVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDomainTopReferVisitRequest) SetStartTime(v string) *DescribeDomainTopReferVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopReferVisitRequest) Validate() error {
	return dara.Validate(s)
}
