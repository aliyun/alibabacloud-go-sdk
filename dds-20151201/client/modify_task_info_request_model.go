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
	SetStepName(v string) *ModifyTaskInfoRequest
	GetStepName() *string
	SetTaskAction(v string) *ModifyTaskInfoRequest
	GetTaskAction() *string
	SetTaskId(v string) *ModifyTaskInfoRequest
	GetTaskId() *string
}

type ModifyTaskInfoRequest struct {
	// A action-related parameter. This parameter can be extended based on your business requirements. This parameter value varies with the value of the TaskAction parameter.
	//
	// example:
	//
	// {\\"recoverMode\\":\\"immediate\\"}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the step visible to the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec_task
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The action name that corresponds to the state described in the actionInfo parameter of the [DescribeHistoryTasks](https://help.aliyun.com/document_detail/2639186.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// modifySwitchTime
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The task ID. Separate multiple IDs with commas (,). You can specify up to 10 task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-83br18hlpdrw3uxxxx
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
