// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeEventsRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeEventsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEventsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeEventsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeEventsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEventsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeEventsRequest
	GetStartTime() *string
}

type DescribeEventsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEventsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEventsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEventsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEventsRequest) SetEndTime(v string) *DescribeEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsRequest) SetOwnerId(v int64) *DescribeEventsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEventsRequest) SetPageNumber(v int32) *DescribeEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsRequest) SetPageSize(v int32) *DescribeEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsRequest) SetRegionId(v string) *DescribeEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventsRequest) SetResourceOwnerAccount(v string) *DescribeEventsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEventsRequest) SetResourceOwnerId(v int64) *DescribeEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEventsRequest) SetStartTime(v string) *DescribeEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsRequest) Validate() error {
	return dara.Validate(s)
}
