// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserTrafficLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveUserTrafficLogRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveUserTrafficLogRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveUserTrafficLogRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeLiveUserTrafficLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveUserTrafficLogRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLiveUserTrafficLogRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveUserTrafficLogRequest
	GetStartTime() *string
}

type DescribeLiveUserTrafficLogRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveUserTrafficLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTrafficLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTrafficLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveUserTrafficLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveUserTrafficLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveUserTrafficLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveUserTrafficLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveUserTrafficLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveUserTrafficLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveUserTrafficLogRequest) SetDomainName(v string) *DescribeLiveUserTrafficLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUserTrafficLogRequest) SetEndTime(v string) *DescribeLiveUserTrafficLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveUserTrafficLogRequest) SetOwnerId(v int64) *DescribeLiveUserTrafficLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveUserTrafficLogRequest) SetPageNumber(v int64) *DescribeLiveUserTrafficLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveUserTrafficLogRequest) SetPageSize(v int64) *DescribeLiveUserTrafficLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveUserTrafficLogRequest) SetRegionId(v string) *DescribeLiveUserTrafficLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveUserTrafficLogRequest) SetStartTime(v string) *DescribeLiveUserTrafficLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveUserTrafficLogRequest) Validate() error {
	return dara.Validate(s)
}
