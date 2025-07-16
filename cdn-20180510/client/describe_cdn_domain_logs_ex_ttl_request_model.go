// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainLogsExTtlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnDomainLogsExTtlRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeCdnDomainLogsExTtlRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeCdnDomainLogsExTtlRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCdnDomainLogsExTtlRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeCdnDomainLogsExTtlRequest
	GetStartTime() *string
}

type DescribeCdnDomainLogsExTtlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2019-12-22T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 300
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2019-12-21T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnDomainLogsExTtlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsExTtlRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsExTtlRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainLogsExTtlRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnDomainLogsExTtlRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCdnDomainLogsExTtlRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnDomainLogsExTtlRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnDomainLogsExTtlRequest) SetDomainName(v string) *DescribeCdnDomainLogsExTtlRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlRequest) SetEndTime(v string) *DescribeCdnDomainLogsExTtlRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlRequest) SetPageNumber(v int64) *DescribeCdnDomainLogsExTtlRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlRequest) SetPageSize(v int64) *DescribeCdnDomainLogsExTtlRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlRequest) SetStartTime(v string) *DescribeCdnDomainLogsExTtlRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlRequest) Validate() error {
	return dara.Validate(s)
}
