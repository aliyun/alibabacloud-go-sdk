// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRobotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteRobotTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteRobotTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRobotTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *DeleteRobotTaskRequest
	GetTaskId() *int64
}

type DeleteRobotTaskRequest struct {
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

func (s DeleteRobotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteRobotTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRobotTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRobotTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRobotTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DeleteRobotTaskRequest) SetOwnerId(v int64) *DeleteRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRobotTaskRequest) SetResourceOwnerAccount(v string) *DeleteRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRobotTaskRequest) SetResourceOwnerId(v int64) *DeleteRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRobotTaskRequest) SetTaskId(v int64) *DeleteRobotTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteRobotTaskRequest) Validate() error {
	return dara.Validate(s)
}
