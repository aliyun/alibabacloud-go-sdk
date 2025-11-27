// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeMajorVersionTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeUpgradeMajorVersionTasksRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeUpgradeMajorVersionTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUpgradeMajorVersionTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeUpgradeMajorVersionTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeUpgradeMajorVersionTasksRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeUpgradeMajorVersionTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUpgradeMajorVersionTasksRequest
	GetResourceOwnerId() *int64
	SetTargetMajorVersion(v string) *DescribeUpgradeMajorVersionTasksRequest
	GetTargetMajorVersion() *string
	SetTaskId(v int32) *DescribeUpgradeMajorVersionTasksRequest
	GetTaskId() *int32
}

type DescribeUpgradeMajorVersionTasksRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1gm3yh0ht1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Pages start from 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: **30*	- to **100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The major engine version of the new instance. Valid values:
	//
	// 	- **10.0**
	//
	// 	- **11.0**
	//
	// 	- **12.0**
	//
	// 	- **13.0**
	//
	// 	- **14.0**
	//
	// 	- **15.0**
	//
	// example:
	//
	// 12.0
	TargetMajorVersion *string `json:"TargetMajorVersion,omitempty" xml:"TargetMajorVersion,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// 417450000
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeUpgradeMajorVersionTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeMajorVersionTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeMajorVersionTasksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeUpgradeMajorVersionTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUpgradeMajorVersionTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUpgradeMajorVersionTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUpgradeMajorVersionTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUpgradeMajorVersionTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUpgradeMajorVersionTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUpgradeMajorVersionTasksRequest) GetTargetMajorVersion() *string {
	return s.TargetMajorVersion
}

func (s *DescribeUpgradeMajorVersionTasksRequest) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeUpgradeMajorVersionTasksRequest) SetDBInstanceId(v string) *DescribeUpgradeMajorVersionTasksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksRequest) SetOwnerAccount(v string) *DescribeUpgradeMajorVersionTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksRequest) SetOwnerId(v int64) *DescribeUpgradeMajorVersionTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksRequest) SetPageNumber(v int32) *DescribeUpgradeMajorVersionTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksRequest) SetPageSize(v int32) *DescribeUpgradeMajorVersionTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksRequest) SetResourceOwnerAccount(v string) *DescribeUpgradeMajorVersionTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksRequest) SetResourceOwnerId(v int64) *DescribeUpgradeMajorVersionTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksRequest) SetTargetMajorVersion(v string) *DescribeUpgradeMajorVersionTasksRequest {
	s.TargetMajorVersion = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksRequest) SetTaskId(v int32) *DescribeUpgradeMajorVersionTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksRequest) Validate() error {
	return dara.Validate(s)
}
