// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSessionIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationSessionIdsRequest
	GetApplicationId() *string
	SetEndTime(v string) *DescribeApplicationSessionIdsRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeApplicationSessionIdsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeApplicationSessionIdsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeApplicationSessionIdsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApplicationSessionIdsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApplicationSessionIdsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeApplicationSessionIdsRequest
	GetResourceOwnerAccount() *string
	SetStartTime(v string) *DescribeApplicationSessionIdsRequest
	GetStartTime() *string
}

type DescribeApplicationSessionIdsRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-02-01T18:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The beginning of the time range to query. Specify the time in the `YYYY-MM-DDThh:mmZ` format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-15T16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApplicationSessionIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionIdsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionIdsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationSessionIdsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApplicationSessionIdsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeApplicationSessionIdsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeApplicationSessionIdsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationSessionIdsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationSessionIdsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationSessionIdsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeApplicationSessionIdsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApplicationSessionIdsRequest) SetApplicationId(v string) *DescribeApplicationSessionIdsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationSessionIdsRequest) SetEndTime(v string) *DescribeApplicationSessionIdsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationSessionIdsRequest) SetOwnerAccount(v string) *DescribeApplicationSessionIdsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeApplicationSessionIdsRequest) SetOwnerId(v int64) *DescribeApplicationSessionIdsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeApplicationSessionIdsRequest) SetPageNumber(v int32) *DescribeApplicationSessionIdsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationSessionIdsRequest) SetPageSize(v int32) *DescribeApplicationSessionIdsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationSessionIdsRequest) SetRegionId(v string) *DescribeApplicationSessionIdsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationSessionIdsRequest) SetResourceOwnerAccount(v string) *DescribeApplicationSessionIdsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeApplicationSessionIdsRequest) SetStartTime(v string) *DescribeApplicationSessionIdsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationSessionIdsRequest) Validate() error {
	return dara.Validate(s)
}
