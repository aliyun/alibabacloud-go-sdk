// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRobotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CancelRobotTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CancelRobotTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelRobotTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CancelRobotTaskRequest
	GetTaskId() *int64
}

type CancelRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](https://help.aliyun.com/document_detail/393531.html) operation to obtain the task ID.
	//
	// example:
	//
	// 1045001
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelRobotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelRobotTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelRobotTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelRobotTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelRobotTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CancelRobotTaskRequest) SetOwnerId(v int64) *CancelRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelRobotTaskRequest) SetResourceOwnerAccount(v string) *CancelRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelRobotTaskRequest) SetResourceOwnerId(v int64) *CancelRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelRobotTaskRequest) SetTaskId(v int64) *CancelRobotTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelRobotTaskRequest) Validate() error {
	return dara.Validate(s)
}
