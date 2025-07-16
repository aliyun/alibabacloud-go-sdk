// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcTopUrlVisitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainSrcTopUrlVisitRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainSrcTopUrlVisitRequest
	GetEndTime() *string
	SetSortBy(v string) *DescribeDomainSrcTopUrlVisitRequest
	GetSortBy() *string
	SetStartTime(v string) *DescribeDomainSrcTopUrlVisitRequest
	GetStartTime() *string
}

type DescribeDomainSrcTopUrlVisitRequest struct {
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
	// > The end time must be later than the start time. The duration between the end time and the start time cannot exceed seven days.
	//
	// example:
	//
	// 2018-10-03T20:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The method that is used to sort the returned URLs. Default value: **pv**. Valid values:
	//
	// 	- **traf**: by network traffic
	//
	// 	- **pv**: by the number of visits
	//
	// example:
	//
	// pv
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > If you leave this parameter empty, data within the previous day is queried.
	//
	// example:
	//
	// 2018-10-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcTopUrlVisitRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainSrcTopUrlVisitRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeDomainSrcTopUrlVisitRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcTopUrlVisitRequest) SetDomainName(v string) *DescribeDomainSrcTopUrlVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitRequest) SetEndTime(v string) *DescribeDomainSrcTopUrlVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitRequest) SetSortBy(v string) *DescribeDomainSrcTopUrlVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitRequest) SetStartTime(v string) *DescribeDomainSrcTopUrlVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitRequest) Validate() error {
	return dara.Validate(s)
}
