// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsHistory(v int32) *DescribeActiveOperationTaskRegionRequest
	GetIsHistory() *int32
	SetOwnerAccount(v string) *DescribeActiveOperationTaskRegionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActiveOperationTaskRegionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskRegionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationTaskRegionRequest
	GetResourceOwnerId() *int64
	SetTaskType(v string) *DescribeActiveOperationTaskRegionRequest
	GetTaskType() *string
}

type DescribeActiveOperationTaskRegionRequest struct {
	// Specifies whether to return the historical tasks. Default value: 0. Valid values:
	//
	// - 0: returns the current task.
	//
	// - 1: returns the historical tasks.
	//
	// example:
	//
	// 0
	IsHistory            *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// - all: all types
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_apsaradb_upgrade
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeActiveOperationTaskRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskRegionRequest) GetIsHistory() *int32 {
	return s.IsHistory
}

func (s *DescribeActiveOperationTaskRegionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActiveOperationTaskRegionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActiveOperationTaskRegionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationTaskRegionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActiveOperationTaskRegionRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeActiveOperationTaskRegionRequest) SetIsHistory(v int32) *DescribeActiveOperationTaskRegionRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionRequest) SetOwnerAccount(v string) *DescribeActiveOperationTaskRegionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionRequest) SetOwnerId(v int64) *DescribeActiveOperationTaskRegionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskRegionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTaskRegionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionRequest) SetTaskType(v string) *DescribeActiveOperationTaskRegionRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionRequest) Validate() error {
	return dara.Validate(s)
}
