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
	// Specifies whether to return historical tasks. Valid values:
	//
	// - 0: returns current tasks. This is the default value.
	//
	// - 1: returns historical tasks.
	//
	// example:
	//
	// 0
	IsHistory    *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be greater than 0 and not exceed the maximum value of the Integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page. The value must be greater than 10. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product ID.
	//
	// example:
	//
	// MongoDB
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the region to which the pending events belong. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query region IDs. If you set this parameter to all, all tasks under the current account are queried, and the TaskType parameter must also be set to all.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task type. Valid values:
	//
	// - rds_apsaradb_ha: primary/secondary node switch.
	//
	// - rds_apsaradb_transfer: instance migration.
	//
	// - rds_apsaradb_upgrade: minor version update.
	//
	// You can obtain the valid values from the [DescribeHistoryTasks](https://help.aliyun.com/document_detail/2639186.html) operation.
	//
	// If you set this parameter to all, all tasks under the current account are queried, and the Region parameter must also be set to all.
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
