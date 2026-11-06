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
	SetSecurityToken(v string) *DescribeActiveOperationTaskRequest
	GetSecurityToken() *string
	SetTaskType(v string) *DescribeActiveOperationTaskRequest
	GetTaskType() *string
}

type DescribeActiveOperationTaskRequest struct {
	// Specifies whether to return historical tasks. Valid values:
	//
	// 	- **0**: Returns current tasks.
	//
	// 	- **1**: Returns historical tasks.
	//
	// Default value: **0**.
	//
	// example:
	//
	// 1
	IsHistory    *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be greater than 0 and cannot exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The commodity ID.
	//
	// example:
	//
	// polardb
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the regions and zones supported by PolarDB.
	//
	// > 	- You can enter **all*	- to view all pending tasks under your account.
	//
	// 	- When `Region` is set to **all**, `TaskType` must also be set to **all**.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task type. You can call the [DescribeActiveOperationTaskType](ID not yet generated) operation to query the task types of pending tasks.
	//
	// > 	- You can also enter **all*	- to view all pending tasks under your account.
	//
	// 	- When `TaskType` is set to **all**, `Region` must also be set to **all**.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
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

func (s *DescribeActiveOperationTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
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

func (s *DescribeActiveOperationTaskRequest) SetSecurityToken(v string) *DescribeActiveOperationTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetTaskType(v string) *DescribeActiveOperationTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) Validate() error {
	return dara.Validate(s)
}
