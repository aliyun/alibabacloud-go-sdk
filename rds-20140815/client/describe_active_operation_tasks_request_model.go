// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowCancel(v int32) *DescribeActiveOperationTasksRequest
	GetAllowCancel() *int32
	SetAllowChange(v int32) *DescribeActiveOperationTasksRequest
	GetAllowChange() *int32
	SetChangeLevel(v string) *DescribeActiveOperationTasksRequest
	GetChangeLevel() *string
	SetDbType(v string) *DescribeActiveOperationTasksRequest
	GetDbType() *string
	SetInsName(v string) *DescribeActiveOperationTasksRequest
	GetInsName() *string
	SetOwnerAccount(v string) *DescribeActiveOperationTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActiveOperationTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeActiveOperationTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeActiveOperationTasksRequest
	GetPageSize() *int32
	SetProductId(v string) *DescribeActiveOperationTasksRequest
	GetProductId() *string
	SetRegion(v string) *DescribeActiveOperationTasksRequest
	GetRegion() *string
	SetResourceOwnerAccount(v string) *DescribeActiveOperationTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationTasksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeActiveOperationTasksRequest
	GetSecurityToken() *string
	SetStatus(v int32) *DescribeActiveOperationTasksRequest
	GetStatus() *int32
	SetTaskType(v string) *DescribeActiveOperationTasksRequest
	GetTaskType() *string
}

type DescribeActiveOperationTasksRequest struct {
	// Specifies whether the task can be canceled. Default value: -1. Valid values:
	//
	// - **-1**: all tasks.
	//
	// - **0**: Only tasks that cannot be canceled are returned.
	//
	// - **1**: Only tasks that can be canceled are returned.
	//
	// example:
	//
	// -1
	AllowCancel *int32 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// Specifies whether the task time can be modified. Default value: -1. Valid values:
	//
	// - **-1**: all tasks.
	//
	// - **0**: Only tasks whose time cannot be modified are returned.
	//
	// - **1**: Only tasks whose time can be modified are returned.
	//
	// example:
	//
	// -1
	AllowChange *int32 `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// The task level. Default value: all. Valid values:
	//
	// - **all**: all levels.
	//
	// - **S0**: Only tasks at the exception recovery level are returned.
	//
	// - **S1**: Only tasks at the system O&M level are returned.
	//
	// example:
	//
	// all
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// The database type. Default value: all. Valid values: mysql, pgsql, and mssql.
	//
	// example:
	//
	// all
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The instance name. This parameter is optional. You can specify at most one instance name.
	//
	// example:
	//
	// rm-bp191w771kd3****
	InsName      *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 25. Maximum value: 100.
	//
	// example:
	//
	// 25
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product name. Valid values: RDS, POLARDB, MongoDB, and Redis. For ApsaraDB RDS instances, set this parameter to RDS.
	//
	// example:
	//
	// RDS
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region ID of the pending event. You can call the DescribeRegions operation to query the most recent region list.
	//
	// > Set this parameter to **all*	- to specify all region IDs.
	//
	// example:
	//
	// cn-beijing
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task status. This parameter is used to filter the returned tasks. Valid values:
	//
	// 	- **-1**: all tasks.
	//
	// 	- **3**: pending tasks.
	//
	// 	- **4**: in-progress tasks.
	//
	// 	- **5**: succeeded tasks.
	//
	// 	- **6**: failed tasks.
	//
	// 	- **7**: canceled tasks.
	//
	// example:
	//
	// -1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task type. Valid values:
	//
	// 	- **rds_apsaradb_ha**: primary/secondary node switch.
	//
	// 	- **rds_apsaradb_transfer**: instance migration.
	//
	// 	- **rds_apsaradb_upgrade**: minor engine version update.
	//
	// 	- **rds_apsaradb_maxscale**: proxy minor version upgrade.
	//
	// 	- **all**: all task types.
	//
	// example:
	//
	// rds_apsaradb_upgrade
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeActiveOperationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksRequest) GetAllowCancel() *int32 {
	return s.AllowCancel
}

func (s *DescribeActiveOperationTasksRequest) GetAllowChange() *int32 {
	return s.AllowChange
}

func (s *DescribeActiveOperationTasksRequest) GetChangeLevel() *string {
	return s.ChangeLevel
}

func (s *DescribeActiveOperationTasksRequest) GetDbType() *string {
	return s.DbType
}

func (s *DescribeActiveOperationTasksRequest) GetInsName() *string {
	return s.InsName
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

func (s *DescribeActiveOperationTasksRequest) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeActiveOperationTasksRequest) GetRegion() *string {
	return s.Region
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

func (s *DescribeActiveOperationTasksRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeActiveOperationTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeActiveOperationTasksRequest) SetAllowCancel(v int32) *DescribeActiveOperationTasksRequest {
	s.AllowCancel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetAllowChange(v int32) *DescribeActiveOperationTasksRequest {
	s.AllowChange = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetChangeLevel(v string) *DescribeActiveOperationTasksRequest {
	s.ChangeLevel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetDbType(v string) *DescribeActiveOperationTasksRequest {
	s.DbType = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetInsName(v string) *DescribeActiveOperationTasksRequest {
	s.InsName = &v
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

func (s *DescribeActiveOperationTasksRequest) SetProductId(v string) *DescribeActiveOperationTasksRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetRegion(v string) *DescribeActiveOperationTasksRequest {
	s.Region = &v
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

func (s *DescribeActiveOperationTasksRequest) SetStatus(v int32) *DescribeActiveOperationTasksRequest {
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
