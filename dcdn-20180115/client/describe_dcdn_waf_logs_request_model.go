// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnWafLogsRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnWafLogsRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeDcdnWafLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnWafLogsRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeDcdnWafLogsRequest
	GetStartTime() *string
}

type DescribeDcdnWafLogsRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
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
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. Valid values: an integer greater than 0.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **300**. Valid values: **1 to 1000**.
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
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnWafLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafLogsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnWafLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnWafLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnWafLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnWafLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnWafLogsRequest) SetDomainName(v string) *DescribeDcdnWafLogsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafLogsRequest) SetEndTime(v string) *DescribeDcdnWafLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnWafLogsRequest) SetPageNumber(v int64) *DescribeDcdnWafLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafLogsRequest) SetPageSize(v int64) *DescribeDcdnWafLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafLogsRequest) SetStartTime(v string) *DescribeDcdnWafLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnWafLogsRequest) Validate() error {
	return dara.Validate(s)
}
