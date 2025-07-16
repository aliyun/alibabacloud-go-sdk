// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopUrlVisitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainTopUrlVisitRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainTopUrlVisitRequest
	GetEndTime() *string
	SetSortBy(v string) *DescribeDomainTopUrlVisitRequest
	GetSortBy() *string
	SetStartTime(v string) *DescribeDomainTopUrlVisitRequest
	GetStartTime() *string
}

type DescribeDomainTopUrlVisitRequest struct {
	// The accelerated domain name that you want to query.
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
	// > The end time must be later than the start time. The maximum time range that can be specified is seven days.
	//
	// example:
	//
	// 2019-10-04T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The method that is used to sort the returned URLs. Default value: **pv**. Valid values:
	//
	// 	- **traf**: by network traffic
	//
	// 	- **pv**: by the number of page views
	//
	// example:
	//
	// pv
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// If you want to query data of a specific day, we recommend that you set the value in the yyyy-MM-ddT16:00:00Z format.
	//
	// example:
	//
	// 2019-10-04T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainTopUrlVisitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainTopUrlVisitRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainTopUrlVisitRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeDomainTopUrlVisitRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainTopUrlVisitRequest) SetDomainName(v string) *DescribeDomainTopUrlVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopUrlVisitRequest) SetEndTime(v string) *DescribeDomainTopUrlVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopUrlVisitRequest) SetSortBy(v string) *DescribeDomainTopUrlVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDomainTopUrlVisitRequest) SetStartTime(v string) *DescribeDomainTopUrlVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopUrlVisitRequest) Validate() error {
	return dara.Validate(s)
}
