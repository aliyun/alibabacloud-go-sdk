// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeMajorVersionPrecheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeUpgradeMajorVersionPrecheckTaskRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeUpgradeMajorVersionPrecheckTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUpgradeMajorVersionPrecheckTaskRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeUpgradeMajorVersionPrecheckTaskRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeUpgradeMajorVersionPrecheckTaskRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeUpgradeMajorVersionPrecheckTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUpgradeMajorVersionPrecheckTaskRequest
	GetResourceOwnerId() *int64
	SetTargetMajorVersion(v string) *DescribeUpgradeMajorVersionPrecheckTaskRequest
	GetTargetMajorVersion() *string
	SetTaskId(v int32) *DescribeUpgradeMajorVersionPrecheckTaskRequest
	GetTaskId() *int32
}

type DescribeUpgradeMajorVersionPrecheckTaskRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1c808s731l****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Valid values: any non-zero positive integer. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values:
	//
	// 	- 30 (default)
	//
	// 	- 50
	//
	// 	- 100
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new major engine version of the instance. The new major engine version must be later than the original major engine version.
	//
	// example:
	//
	// 12.0
	TargetMajorVersion *string `json:"TargetMajorVersion,omitempty" xml:"TargetMajorVersion,omitempty"`
	// The ID of the upgrade check task. You can obtain the ID of the upgrade check task from the **TaskId*	- parameter in the response to the UpgradeDBInstanceMajorVersionPrecheck operation.
	//
	// example:
	//
	// 416980000
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeUpgradeMajorVersionPrecheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeMajorVersionPrecheckTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) GetTargetMajorVersion() *string {
	return s.TargetMajorVersion
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) SetDBInstanceId(v string) *DescribeUpgradeMajorVersionPrecheckTaskRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) SetOwnerAccount(v string) *DescribeUpgradeMajorVersionPrecheckTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) SetOwnerId(v int64) *DescribeUpgradeMajorVersionPrecheckTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) SetPageNumber(v int32) *DescribeUpgradeMajorVersionPrecheckTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) SetPageSize(v int32) *DescribeUpgradeMajorVersionPrecheckTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) SetResourceOwnerAccount(v string) *DescribeUpgradeMajorVersionPrecheckTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) SetResourceOwnerId(v int64) *DescribeUpgradeMajorVersionPrecheckTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) SetTargetMajorVersion(v string) *DescribeUpgradeMajorVersionPrecheckTaskRequest {
	s.TargetMajorVersion = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) SetTaskId(v int32) *DescribeUpgradeMajorVersionPrecheckTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
