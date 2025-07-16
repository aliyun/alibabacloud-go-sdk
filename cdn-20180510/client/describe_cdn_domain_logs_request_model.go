// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnDomainLogsRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeCdnDomainLogsRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeCdnDomainLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCdnDomainLogsRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeCdnDomainLogsRequest
	GetStartTime() *string
}

type DescribeCdnDomainLogsRequest struct {
	// The domain name. You can specify only one domain name.
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
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. Pages start from page **1**.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **300**. Maximum value: **1000**. Valid values: **1*	- to **1000**.
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

func (s DescribeCdnDomainLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnDomainLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCdnDomainLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnDomainLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnDomainLogsRequest) SetDomainName(v string) *DescribeCdnDomainLogsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetEndTime(v string) *DescribeCdnDomainLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetPageNumber(v int64) *DescribeCdnDomainLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetPageSize(v int64) *DescribeCdnDomainLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetStartTime(v string) *DescribeCdnDomainLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) Validate() error {
	return dara.Validate(s)
}
