// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainEdgeLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainEdgeLogRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainEdgeLogRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveDomainEdgeLogRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeLiveDomainEdgeLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveDomainEdgeLogRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLiveDomainEdgeLogRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainEdgeLogRequest
	GetStartTime() *string
}

type DescribeLiveDomainEdgeLogRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainEdgeLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainEdgeLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainEdgeLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainEdgeLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainEdgeLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainEdgeLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveDomainEdgeLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveDomainEdgeLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainEdgeLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainEdgeLogRequest) SetDomainName(v string) *DescribeLiveDomainEdgeLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainEdgeLogRequest) SetEndTime(v string) *DescribeLiveDomainEdgeLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainEdgeLogRequest) SetOwnerId(v int64) *DescribeLiveDomainEdgeLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainEdgeLogRequest) SetPageNumber(v int64) *DescribeLiveDomainEdgeLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveDomainEdgeLogRequest) SetPageSize(v int64) *DescribeLiveDomainEdgeLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveDomainEdgeLogRequest) SetRegionId(v string) *DescribeLiveDomainEdgeLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainEdgeLogRequest) SetStartTime(v string) *DescribeLiveDomainEdgeLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainEdgeLogRequest) Validate() error {
	return dara.Validate(s)
}
