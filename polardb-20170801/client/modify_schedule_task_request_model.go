// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyScheduleTaskRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyScheduleTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyScheduleTaskRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *ModifyScheduleTaskRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *ModifyScheduleTaskRequest
	GetPlannedStartTime() *string
	SetResourceGroupId(v string) *ModifyScheduleTaskRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyScheduleTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyScheduleTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *ModifyScheduleTaskRequest
	GetTaskId() *string
}

type ModifyScheduleTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-18T19:20:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-15T20:00:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 53879cdb-9a00-428e-acaf-ff4cff******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduleTaskRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyScheduleTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyScheduleTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyScheduleTaskRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *ModifyScheduleTaskRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *ModifyScheduleTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyScheduleTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyScheduleTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyScheduleTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyScheduleTaskRequest) SetDBClusterId(v string) *ModifyScheduleTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyScheduleTaskRequest) SetOwnerAccount(v string) *ModifyScheduleTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScheduleTaskRequest) SetOwnerId(v int64) *ModifyScheduleTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScheduleTaskRequest) SetPlannedEndTime(v string) *ModifyScheduleTaskRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyScheduleTaskRequest) SetPlannedStartTime(v string) *ModifyScheduleTaskRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyScheduleTaskRequest) SetResourceGroupId(v string) *ModifyScheduleTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyScheduleTaskRequest) SetResourceOwnerAccount(v string) *ModifyScheduleTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScheduleTaskRequest) SetResourceOwnerId(v int64) *ModifyScheduleTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScheduleTaskRequest) SetTaskId(v string) *ModifyScheduleTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}
