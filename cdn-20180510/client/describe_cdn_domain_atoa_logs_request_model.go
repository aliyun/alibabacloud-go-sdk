// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainAtoaLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnDomainAtoaLogsRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeCdnDomainAtoaLogsRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeCdnDomainAtoaLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCdnDomainAtoaLogsRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeCdnDomainAtoaLogsRequest
	GetStartTime() *string
}

type DescribeCdnDomainAtoaLogsRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnDomainAtoaLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainAtoaLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainAtoaLogsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainAtoaLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnDomainAtoaLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCdnDomainAtoaLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnDomainAtoaLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnDomainAtoaLogsRequest) SetDomainName(v string) *DescribeCdnDomainAtoaLogsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainAtoaLogsRequest) SetEndTime(v string) *DescribeCdnDomainAtoaLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainAtoaLogsRequest) SetPageNumber(v int64) *DescribeCdnDomainAtoaLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnDomainAtoaLogsRequest) SetPageSize(v int64) *DescribeCdnDomainAtoaLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainAtoaLogsRequest) SetStartTime(v string) *DescribeCdnDomainAtoaLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnDomainAtoaLogsRequest) Validate() error {
	return dara.Validate(s)
}
