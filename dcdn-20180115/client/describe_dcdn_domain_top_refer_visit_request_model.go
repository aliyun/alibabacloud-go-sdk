// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainTopReferVisitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainTopReferVisitRequest
	GetDomainName() *string
	SetSortBy(v string) *DescribeDcdnDomainTopReferVisitRequest
	GetSortBy() *string
	SetStartTime(v string) *DescribeDcdnDomainTopReferVisitRequest
	GetStartTime() *string
}

type DescribeDcdnDomainTopReferVisitRequest struct {
	// The accelerated domain name. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- **traf**: by network traffic
	//
	// 	- **pv**: by the number of visits
	//
	// Default value: **pv**.
	//
	// example:
	//
	// pv
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// To query the data on a specified day, use the yyyy-MM-ddT16:00:00Z format.
	//
	// If you do not set this parameter, data collected within the last 24 hours is queried.
	//
	// example:
	//
	// 2018-10-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainTopReferVisitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainTopReferVisitRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeDcdnDomainTopReferVisitRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainTopReferVisitRequest) SetDomainName(v string) *DescribeDcdnDomainTopReferVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitRequest) SetSortBy(v string) *DescribeDcdnDomainTopReferVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitRequest) SetStartTime(v string) *DescribeDcdnDomainTopReferVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitRequest) Validate() error {
	return dara.Validate(s)
}
