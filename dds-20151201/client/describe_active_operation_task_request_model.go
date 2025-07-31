// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsHistory(v int32) *DescribeActiveOperationTaskRequest
	GetIsHistory() *int32
	SetOwnerAccount(v string) *DescribeActiveOperationTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActiveOperationTaskRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeActiveOperationTaskRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeActiveOperationTaskRequest
	GetPageSize() *int32
	SetProductId(v string) *DescribeActiveOperationTaskRequest
	GetProductId() *string
	SetRegion(v string) *DescribeActiveOperationTaskRequest
	GetRegion() *string
	SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationTaskRequest
	GetResourceOwnerId() *int64
	SetTaskType(v string) *DescribeActiveOperationTaskRequest
	GetTaskType() *string
}

type DescribeActiveOperationTaskRequest struct {
	// Specifies whether to return the historical tasks.
	//
	// Default value: 0. Valid values:
	//
	// - 0: returns the current task.
	//
	// - 1: returns the historical tasks.
	//
	// example:
	//
	// 0
	IsHistory    *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Specify a value greater than 10. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// MongoDB
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region ID of the instance.
	//
	// >  If you set the Region parameter to **all**, all tasks created within your Alibaba Cloud account are queried. In this case, you must set the **taskType*	- parameter to **all**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the task. Valid values:
	//
	// - rds_apsaradb_ha: master-replica switchover
	//
	// - rds_apsaradb_transfer: instance migration
	//
	// - rds_apsaradb_upgrade: minor version update
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_apsaradb_ha
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeActiveOperationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskRequest) GetIsHistory() *int32 {
	return s.IsHistory
}

func (s *DescribeActiveOperationTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActiveOperationTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActiveOperationTaskRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeActiveOperationTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeActiveOperationTaskRequest) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeActiveOperationTaskRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeActiveOperationTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActiveOperationTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeActiveOperationTaskRequest) SetIsHistory(v int32) *DescribeActiveOperationTaskRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetOwnerAccount(v string) *DescribeActiveOperationTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetOwnerId(v int64) *DescribeActiveOperationTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetPageNumber(v int32) *DescribeActiveOperationTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetPageSize(v int32) *DescribeActiveOperationTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetProductId(v string) *DescribeActiveOperationTaskRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetRegion(v string) *DescribeActiveOperationTaskRequest {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetTaskType(v string) *DescribeActiveOperationTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) Validate() error {
	return dara.Validate(s)
}
