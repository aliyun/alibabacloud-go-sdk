// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRobotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *StartRobotTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StartRobotTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartRobotTaskRequest
	GetResourceOwnerId() *int64
	SetScheduleTime(v string) *StartRobotTaskRequest
	GetScheduleTime() *string
	SetTaskId(v int64) *StartRobotTaskRequest
	GetTaskId() *int64
}

type StartRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time scheduled for starting the robocall task, in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2019-06-14 14:55:23
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~CreateRobotTask~~) operation to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 104500****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartRobotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *StartRobotTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartRobotTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartRobotTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartRobotTaskRequest) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *StartRobotTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StartRobotTaskRequest) SetOwnerId(v int64) *StartRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartRobotTaskRequest) SetResourceOwnerAccount(v string) *StartRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartRobotTaskRequest) SetResourceOwnerId(v int64) *StartRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartRobotTaskRequest) SetScheduleTime(v string) *StartRobotTaskRequest {
	s.ScheduleTime = &v
	return s
}

func (s *StartRobotTaskRequest) SetTaskId(v int64) *StartRobotTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartRobotTaskRequest) Validate() error {
	return dara.Validate(s)
}
