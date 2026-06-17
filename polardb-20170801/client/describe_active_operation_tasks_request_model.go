// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowCancel(v int64) *DescribeActiveOperationTasksRequest
	GetAllowCancel() *int64
	SetAllowChange(v int64) *DescribeActiveOperationTasksRequest
	GetAllowChange() *int64
	SetChangeLevel(v string) *DescribeActiveOperationTasksRequest
	GetChangeLevel() *string
	SetDBClusterId(v string) *DescribeActiveOperationTasksRequest
	GetDBClusterId() *string
	SetDBType(v string) *DescribeActiveOperationTasksRequest
	GetDBType() *string
	SetOwnerAccount(v string) *DescribeActiveOperationTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActiveOperationTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeActiveOperationTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeActiveOperationTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeActiveOperationTasksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeActiveOperationTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationTasksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeActiveOperationTasksRequest
	GetSecurityToken() *string
	SetStatus(v int64) *DescribeActiveOperationTasksRequest
	GetStatus() *int64
	SetTaskType(v string) *DescribeActiveOperationTasksRequest
	GetTaskType() *string
}

type DescribeActiveOperationTasksRequest struct {
	// Specifies whether to allow cancellation. Valid values:
	//
	// - **-1*	- (default): all.
	//
	// - **0**: returns only tasks that do not allow cancellation.
	//
	// - **1**: returns only tasks that allow cancellation.
	//
	// example:
	//
	// -1
	AllowCancel *int64 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// Specifies whether to allow time modification. Valid values:
	//
	// - **-1*	- (default): all.
	//
	// - **0**: returns only tasks that do not allow time modification.
	//
	// - **1**: returns only tasks that allow time modification.
	//
	// example:
	//
	// -1
	AllowChange *int64 `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// The task level. Valid values:
	//
	// - **all*	- (default): all.
	//
	// - **S0**: returns tasks at the abnormal repair level.
	//
	// - **S1**: returns tasks at the system maintenance level.
	//
	// example:
	//
	// all
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query detailed information about all clusters under your account, including cluster IDs.
	//
	// example:
	//
	// pc-3ns***********d5d
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine type. Valid values:
	//
	// - **MySQL**
	//
	// - **PostgreSQL**
	//
	// - **Oracle**
	//
	// example:
	//
	// MySQL
	DBType       *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be greater than 0 and cannot exceed the maximum value of the Integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// - **30*	- (default)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the pending event.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task status. Valid values:
	//
	// - -1: all tasks.
	//
	// - 3: pending tasks.
	//
	// - 4: tasks in progress.
	//
	// - 5: successfully completed tasks.
	//
	// - 6: failed tasks.
	//
	// - 7: canceled tasks.
	//
	// example:
	//
	// -1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the pending event task. Valid values:
	//
	// - **DatabaseSoftwareUpgrading**: database software upgrade
	//
	// - **DatabaseHardwareMaintenance**: hardware maintenance and upgrade
	//
	// - **DatabaseStorageUpgrading**: database storage upgrade
	//
	// - **DatabaseProxyUpgrading**: proxy minor version upgrade
	//
	// - **all**: returns all types of pending events
	//
	// > When `Region` is set to **all**, `TaskType` must also be set to **all**.
	//
	// example:
	//
	// DatabaseProxyUpgrading
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeActiveOperationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksRequest) GetAllowCancel() *int64 {
	return s.AllowCancel
}

func (s *DescribeActiveOperationTasksRequest) GetAllowChange() *int64 {
	return s.AllowChange
}

func (s *DescribeActiveOperationTasksRequest) GetChangeLevel() *string {
	return s.ChangeLevel
}

func (s *DescribeActiveOperationTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeActiveOperationTasksRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeActiveOperationTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActiveOperationTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActiveOperationTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeActiveOperationTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeActiveOperationTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeActiveOperationTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActiveOperationTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeActiveOperationTasksRequest) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeActiveOperationTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeActiveOperationTasksRequest) SetAllowCancel(v int64) *DescribeActiveOperationTasksRequest {
	s.AllowCancel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetAllowChange(v int64) *DescribeActiveOperationTasksRequest {
	s.AllowChange = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetChangeLevel(v string) *DescribeActiveOperationTasksRequest {
	s.ChangeLevel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetDBClusterId(v string) *DescribeActiveOperationTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetDBType(v string) *DescribeActiveOperationTasksRequest {
	s.DBType = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetOwnerAccount(v string) *DescribeActiveOperationTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetOwnerId(v int64) *DescribeActiveOperationTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageNumber(v int32) *DescribeActiveOperationTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageSize(v int32) *DescribeActiveOperationTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetRegionId(v string) *DescribeActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetSecurityToken(v string) *DescribeActiveOperationTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetStatus(v int64) *DescribeActiveOperationTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetTaskType(v string) *DescribeActiveOperationTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) Validate() error {
	return dara.Validate(s)
}
