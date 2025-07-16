// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeTasksRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeTasksRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTasksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTasksRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeTasksRequest
	GetStartTime() *string
}

type DescribeTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-04
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 30
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTasksRequest) SetDBInstanceId(v string) *DescribeTasksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeTasksRequest) SetEndTime(v string) *DescribeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerAccount(v string) *DescribeTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerId(v int64) *DescribeTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetRegionId(v string) *DescribeTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerAccount(v string) *DescribeTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerId(v int64) *DescribeTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetStartTime(v string) *DescribeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksRequest) Validate() error {
	return dara.Validate(s)
}
