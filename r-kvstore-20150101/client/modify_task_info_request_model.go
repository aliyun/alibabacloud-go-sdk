// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionParams(v string) *ModifyTaskInfoRequest
	GetActionParams() *string
	SetRegionId(v string) *ModifyTaskInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyTaskInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTaskInfoRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyTaskInfoRequest
	GetSecurityToken() *string
	SetStepName(v string) *ModifyTaskInfoRequest
	GetStepName() *string
	SetTaskAction(v string) *ModifyTaskInfoRequest
	GetTaskAction() *string
	SetTaskId(v string) *ModifyTaskInfoRequest
	GetTaskId() *string
}

type ModifyTaskInfoRequest struct {
	// The parameters related to the action, in JSON format.
	//
	// If **TaskAction*	- is set to **modifySwitchTime**, set this parameter to `{"recoverMode": "xxx", "recoverTime": "xxx"}`.
	//
	// 	- **recoverMode**: the recovery mode of the task. Valid values:
	//
	//   	- **imePoint**: executes the task at a specified point in time.
	//
	//   	- **immediate**: executes the task immediately.
	//
	//   	- **maintainTime**: executes the task during the O&M window.
	//
	// 	- **recoverTime**: the specified point in time. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. This parameter is required when **recoverMode*	- is set to **timePoint**.
	//
	// example:
	//
	// {"recoverMode":"immediate"}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the current execution step. You can call the [DescribeHistoryTasks](https://help.aliyun.com/document_detail/2400077.html) operation to query the current execution step of a specified task. A common value is **do_pause**, which indicates that the task is waiting to be executed at a preset time.
	//
	// example:
	//
	// do_pause
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The operation name. Valid value: **modifySwitchTime**, which indicates modifying the switchover time or recovery time.
	//
	// example:
	//
	// modifySwitchTime
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The task ID. Separate multiple task IDs with commas (,). A maximum of 30 task IDs are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-0mq3kfhn8i1nlt****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyTaskInfoRequest) GetActionParams() *string {
	return s.ActionParams
}

func (s *ModifyTaskInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyTaskInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTaskInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTaskInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyTaskInfoRequest) GetStepName() *string {
	return s.StepName
}

func (s *ModifyTaskInfoRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *ModifyTaskInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyTaskInfoRequest) SetActionParams(v string) *ModifyTaskInfoRequest {
	s.ActionParams = &v
	return s
}

func (s *ModifyTaskInfoRequest) SetRegionId(v string) *ModifyTaskInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTaskInfoRequest) SetResourceOwnerAccount(v string) *ModifyTaskInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTaskInfoRequest) SetResourceOwnerId(v int64) *ModifyTaskInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTaskInfoRequest) SetSecurityToken(v string) *ModifyTaskInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyTaskInfoRequest) SetStepName(v string) *ModifyTaskInfoRequest {
	s.StepName = &v
	return s
}

func (s *ModifyTaskInfoRequest) SetTaskAction(v string) *ModifyTaskInfoRequest {
	s.TaskAction = &v
	return s
}

func (s *ModifyTaskInfoRequest) SetTaskId(v string) *ModifyTaskInfoRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
