// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainTopUrlVisitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainTopUrlVisitRequest
	GetDomainName() *string
	SetSortBy(v string) *DescribeDcdnDomainTopUrlVisitRequest
	GetSortBy() *string
	SetStartTime(v string) *DescribeDcdnDomainTopUrlVisitRequest
	GetStartTime() *string
}

type DescribeDcdnDomainTopUrlVisitRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
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
	// Default value: **pv**
	//
	// example:
	//
	// pv
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// To query the data on a specified day, use the format: yyyy-MM-ddT16:00:00Z.
	//
	// > If you do not specify this parameter, the data in the last 24 hours is queried.
	//
	// example:
	//
	// 2018-10-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) SetDomainName(v string) *DescribeDcdnDomainTopUrlVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) SetSortBy(v string) *DescribeDcdnDomainTopUrlVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) SetStartTime(v string) *DescribeDcdnDomainTopUrlVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) Validate() error {
	return dara.Validate(s)
}
