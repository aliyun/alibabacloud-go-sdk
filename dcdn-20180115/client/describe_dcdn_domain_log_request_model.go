// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainLogRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainLogRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeDcdnDomainLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnDomainLogRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeDcdnDomainLogRequest
	GetStartTime() *string
}

type DescribeDcdnDomainLogRequest struct {
	// The accelerated domain name that you want to query. You can specify only one domain name in each request.
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
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2021-11-07T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. Pages start from page **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: **1*	- to **1000**. Default value: **300**. Maximum value: **1000**.
	//
	// example:
	//
	// 300
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-11-07T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnDomainLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnDomainLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainLogRequest) SetDomainName(v string) *DescribeDcdnDomainLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) SetEndTime(v string) *DescribeDcdnDomainLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) SetPageNumber(v int64) *DescribeDcdnDomainLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) SetPageSize(v int64) *DescribeDcdnDomainLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) SetStartTime(v string) *DescribeDcdnDomainLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) Validate() error {
	return dara.Validate(s)
}
