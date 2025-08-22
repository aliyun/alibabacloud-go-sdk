// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainLogExTtlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainLogExTtlRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainLogExTtlRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeDcdnDomainLogExTtlRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnDomainLogExTtlRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeDcdnDomainLogExTtlRequest
	GetStartTime() *string
}

type DescribeDcdnDomainLogExTtlRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainLogExTtlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogExTtlRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogExTtlRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainLogExTtlRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainLogExTtlRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnDomainLogExTtlRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnDomainLogExTtlRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainLogExTtlRequest) SetDomainName(v string) *DescribeDcdnDomainLogExTtlRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlRequest) SetEndTime(v string) *DescribeDcdnDomainLogExTtlRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlRequest) SetPageNumber(v int64) *DescribeDcdnDomainLogExTtlRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlRequest) SetPageSize(v int64) *DescribeDcdnDomainLogExTtlRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlRequest) SetStartTime(v string) *DescribeDcdnDomainLogExTtlRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlRequest) Validate() error {
	return dara.Validate(s)
}
