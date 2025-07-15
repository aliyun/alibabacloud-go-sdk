// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveTrafficDomainLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveTrafficDomainLogRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveTrafficDomainLogRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveTrafficDomainLogRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeLiveTrafficDomainLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveTrafficDomainLogRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLiveTrafficDomainLogRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveTrafficDomainLogRequest
	GetStartTime() *string
}

type DescribeLiveTrafficDomainLogRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveTrafficDomainLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTrafficDomainLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveTrafficDomainLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveTrafficDomainLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveTrafficDomainLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveTrafficDomainLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveTrafficDomainLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveTrafficDomainLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveTrafficDomainLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveTrafficDomainLogRequest) SetDomainName(v string) *DescribeLiveTrafficDomainLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogRequest) SetEndTime(v string) *DescribeLiveTrafficDomainLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogRequest) SetOwnerId(v int64) *DescribeLiveTrafficDomainLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogRequest) SetPageNumber(v int64) *DescribeLiveTrafficDomainLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogRequest) SetPageSize(v int64) *DescribeLiveTrafficDomainLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogRequest) SetRegionId(v string) *DescribeLiveTrafficDomainLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogRequest) SetStartTime(v string) *DescribeLiveTrafficDomainLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogRequest) Validate() error {
	return dara.Validate(s)
}
