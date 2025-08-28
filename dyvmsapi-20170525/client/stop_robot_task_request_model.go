// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRobotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *StopRobotTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StopRobotTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopRobotTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *StopRobotTaskRequest
	GetTaskId() *int64
}

type StopRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~CreateRobotTask~~) operation to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1045001
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopRobotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *StopRobotTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopRobotTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopRobotTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopRobotTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StopRobotTaskRequest) SetOwnerId(v int64) *StopRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopRobotTaskRequest) SetResourceOwnerAccount(v string) *StopRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopRobotTaskRequest) SetResourceOwnerId(v int64) *StopRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopRobotTaskRequest) SetTaskId(v int64) *StopRobotTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopRobotTaskRequest) Validate() error {
	return dara.Validate(s)
}
