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
	// The filter condition that is used to return events based on the settings of event cancellation. Default value: -1. Valid values:
	//
	// 	- **-1**: returns all events.
	//
	// 	- **0**: returns only O\\&M events that cannot be canceled.
	//
	// 	- **1**: returns only O\\&M events that can be canceled.
	//
	// example:
	//
	// 1
	AllowCancel *int32 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// The filter condition that is used to return events based on the settings of the switching time. Default value: -1. Valid values:
	//
	// 	- **-1**: returns all events.
	//
	// 	- **0**: returns only O\\&M events for which the switching time cannot be changed.
	//
	// 	- **-1**: returns only O\\&M events for which the switching time can be changed.
	//
	// example:
	//
	// -1
	AllowChange *int32 `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// The type of task configuration change. Valid values:
	//
	// 	- **all*	- (default): The configurations of all O\\&M tasks are changed.
	//
	// 	- **S0**: The configurations of tasks initiated to fix exceptions are changed.
	//
	// 	- **S1**: The configurations of system O\\&M tasks are changed.
	//
	// example:
	//
	// all
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// The database type. Valid values: **redis**
	//
	// example:
	//
	// redis
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The name of the instance. You can leave this parameter empty. If you configure this parameter, you can specify the name only of one instance.
	//
	// example:
	//
	// r-wz96fzmpvpr2qnqnlb
	InsName      *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
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
	// The name of the service. Valid values: RDS, POLARDB, MongoDB, and Redis. For Redis instances, set the value to Redis.
	//
	// example:
	//
	// Redis
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region ID of the O&M task. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// > A value of **all*	- indicates all region IDs.
	//
	// example:
	//
	// cn-shanghai
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The status of an O\\&M event. This parameter is used to filter returned tasks. Valid values:
	//
	// 	- **-1**: filters all events.
	//
	// 	- **3**: filters pending events.
	//
	// 	- **4**: filters in-progress events.
	//
	// 	- **5**: filters successful events.
	//
	// 	- **6**: filters failed events.
	//
	// 	- **7**: filters canceled events.
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the O\\&M event. If this parameter is not specified, all types of O\\&M events are queried.
	//
	// Valid values:
	//
	// 	- rds_apsradb_upgrade: minor version update
	//
	// 	- rds_apsaradb_ha: primary/secondary switchover
	//
	// 	- rds_apsaradb_ssl_update: SSL certificate update
	//
	// 	- rds_apsaradb_major_upgrade: major version upgrade
	//
	// 	- rds_apsradb_transfer: instance migration
	//
	// 	- rds_apsaradb_modify_config: network upgrade
	//
	// 	- rds_apsaradb_modify_config: instance parameter adjustment
	//
	// 	- rds_apsaradb_maxscale: proxy minor version update
	//
	// example:
	//
	// all
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
