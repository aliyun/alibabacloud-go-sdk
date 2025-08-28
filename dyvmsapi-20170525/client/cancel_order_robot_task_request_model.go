// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderRobotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CancelOrderRobotTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CancelOrderRobotTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelOrderRobotTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CancelOrderRobotTaskRequest
	GetTaskId() *int64
}

type CancelOrderRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](https://help.aliyun.com/document_detail/393531.html) operation to obtain the ID of the robocall task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1045001
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelOrderRobotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRobotTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelOrderRobotTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelOrderRobotTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelOrderRobotTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CancelOrderRobotTaskRequest) SetOwnerId(v int64) *CancelOrderRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelOrderRobotTaskRequest) SetResourceOwnerAccount(v string) *CancelOrderRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelOrderRobotTaskRequest) SetResourceOwnerId(v int64) *CancelOrderRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelOrderRobotTaskRequest) SetTaskId(v int64) *CancelOrderRobotTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelOrderRobotTaskRequest) Validate() error {
	return dara.Validate(s)
}
