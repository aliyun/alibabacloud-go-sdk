// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSwitchLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceSwitchLogRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBInstanceSwitchLogRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeDBInstanceSwitchLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceSwitchLogRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstanceSwitchLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstanceSwitchLogRequest
	GetPageSize() *int32
	SetRegionId(v []byte) *DescribeDBInstanceSwitchLogRequest
	GetRegionId() []byte
	SetResourceOwnerAccount(v string) *DescribeDBInstanceSwitchLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceSwitchLogRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeDBInstanceSwitchLogRequest
	GetStartTime() *string
}

type DescribeDBInstanceSwitchLogRequest struct {
	// The instance ID. You can call DescribeDBInstances to obtain the instance ID.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end time of the query. The end time must be later than the start time. Format: yyyy-MM-ddTHH:mmZ (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-04-02T00:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: values greater than 0 and not exceeding the maximum value of Integer.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call DescribeRegions to obtain the region ID.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// cn-hangzhou
	RegionId             []byte  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time of the query. Format: yyyy-MM-ddTHH:mmZ (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-04-01T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstanceSwitchLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSwitchLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSwitchLogRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceSwitchLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstanceSwitchLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceSwitchLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceSwitchLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceSwitchLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstanceSwitchLogRequest) GetRegionId() []byte {
	return s.RegionId
}

func (s *DescribeDBInstanceSwitchLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceSwitchLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceSwitchLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstanceSwitchLogRequest) SetDBInstanceId(v string) *DescribeDBInstanceSwitchLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetEndTime(v string) *DescribeDBInstanceSwitchLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetOwnerAccount(v string) *DescribeDBInstanceSwitchLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetOwnerId(v int64) *DescribeDBInstanceSwitchLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetPageNumber(v int32) *DescribeDBInstanceSwitchLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetPageSize(v int32) *DescribeDBInstanceSwitchLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetRegionId(v []byte) *DescribeDBInstanceSwitchLogRequest {
	s.RegionId = v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceSwitchLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceSwitchLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetStartTime(v string) *DescribeDBInstanceSwitchLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) Validate() error {
	return dara.Validate(s)
}
