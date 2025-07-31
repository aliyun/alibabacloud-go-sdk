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
	SetResourceGroupId(v string) *DescribeActiveOperationTasksRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeActiveOperationTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationTasksRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *DescribeActiveOperationTasksRequest
	GetStatus() *int32
	SetTaskType(v string) *DescribeActiveOperationTasksRequest
	GetTaskType() *string
}

type DescribeActiveOperationTasksRequest struct {
	// Specifies whether to allow the cancellation operation. Valid values:
	//
	// 	- **0**: The cancellation operation is not allowed.
	//
	// 	- **1**: The cancellation operation is allowed.
	//
	// example:
	//
	// ***
	AllowCancel *int32 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// Specifies whether to allow the modification operation. Valid values:
	//
	// 	- **0**: The modification operation is not allowed.
	//
	// 	- **1**: The modification operation is allowed.
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
	// ***
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// ***
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// dds-bp16aaccfe10e3e4
	InsName      *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Specify the parameter to a positive integer that is greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// ***
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region ID of the instance.
	//
	// >  If you set the Region parameter to **all**, all tasks created within your Alibaba Cloud account are queried. In this case, you must set the **taskType*	- parameter to **all**.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **0**: waiting for execution
	//
	// 	- **1**: being executed
	//
	// 	- **2**: successful
	//
	// 	- **3**: failed
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task type.
	//
	// example:
	//
	// ***
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

func (s *DescribeActiveOperationTasksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeActiveOperationTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *DescribeActiveOperationTasksRequest) SetResourceGroupId(v string) *DescribeActiveOperationTasksRequest {
	s.ResourceGroupId = &v
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
