// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrateTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeMigrateTasksRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeMigrateTasksRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeMigrateTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeMigrateTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMigrateTasksRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeMigrateTasksRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeMigrateTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMigrateTasksRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeMigrateTasksRequest
	GetStartTime() *string
}

type DescribeMigrateTasksRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-25T01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: any non-zero positive integer.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30*	- to **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-20T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMigrateTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrateTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrateTasksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeMigrateTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMigrateTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMigrateTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMigrateTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMigrateTasksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMigrateTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMigrateTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMigrateTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMigrateTasksRequest) SetDBInstanceId(v string) *DescribeMigrateTasksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeMigrateTasksRequest) SetEndTime(v string) *DescribeMigrateTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMigrateTasksRequest) SetOwnerId(v int64) *DescribeMigrateTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrateTasksRequest) SetPageNumber(v int32) *DescribeMigrateTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMigrateTasksRequest) SetPageSize(v int32) *DescribeMigrateTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMigrateTasksRequest) SetResourceGroupId(v string) *DescribeMigrateTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMigrateTasksRequest) SetResourceOwnerAccount(v string) *DescribeMigrateTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMigrateTasksRequest) SetResourceOwnerId(v int64) *DescribeMigrateTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMigrateTasksRequest) SetStartTime(v string) *DescribeMigrateTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMigrateTasksRequest) Validate() error {
	return dara.Validate(s)
}
