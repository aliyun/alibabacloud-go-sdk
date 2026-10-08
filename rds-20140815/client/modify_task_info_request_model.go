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
	// The action-related parameters, which can be extended as needed. When taskAction is set to modifySwitchTime, set ActionParams to `{"recoverMode": "xxx", "recoverTime": "xxx"}`.
	//
	// recoverMode specifies the task recovery pattern. Valid values:
	//
	// - **timePoint**: Execute at a specified point in time.
	//
	// - **immediate**: Execute immediately.
	//
	// recoverTime specifies the recovery time in UTC+0. Format: yyyy-MM-ddTHH:mm:ssZ. This parameter is required when recoverMode is set to timePoint.
	//
	// example:
	//
	// {"recoverTime":"2023-04-12T18:30:00Z","recoverMode":"timePoint"}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) operation to query available region IDs.
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
	// The name of the execution step.
	//
	// example:
	//
	// ha_switch
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The task action. Set the value to modifySwitchTime, which indicates modifying the switchover time or recovery time.
	//
	// example:
	//
	// modifySwitchTime
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The task ID. You can call the DescribeTasks operation to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-83br18hloum8u3948s
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
