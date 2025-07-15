// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainLogExTtlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainLogExTtlRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainLogExTtlRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveDomainLogExTtlRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeLiveDomainLogExTtlRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveDomainLogExTtlRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLiveDomainLogExTtlRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainLogExTtlRequest
	GetStartTime() *string
}

type DescribeLiveDomainLogExTtlRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainLogExTtlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogExTtlRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogExTtlRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainLogExTtlRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainLogExTtlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainLogExTtlRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveDomainLogExTtlRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveDomainLogExTtlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainLogExTtlRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainLogExTtlRequest) SetDomainName(v string) *DescribeLiveDomainLogExTtlRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlRequest) SetEndTime(v string) *DescribeLiveDomainLogExTtlRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlRequest) SetOwnerId(v int64) *DescribeLiveDomainLogExTtlRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlRequest) SetPageNumber(v int64) *DescribeLiveDomainLogExTtlRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlRequest) SetPageSize(v int64) *DescribeLiveDomainLogExTtlRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlRequest) SetRegionId(v string) *DescribeLiveDomainLogExTtlRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlRequest) SetStartTime(v string) *DescribeLiveDomainLogExTtlRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlRequest) Validate() error {
	return dara.Validate(s)
}
