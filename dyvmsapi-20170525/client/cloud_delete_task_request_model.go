// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudDeleteTaskRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudDeleteTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudDeleteTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CloudDeleteTaskRequest
	GetTaskId() *int64
}

type CloudDeleteTaskRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 外呼任务Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloudDeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteTaskRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudDeleteTaskRequest) SetEnterpriseId(v int64) *CloudDeleteTaskRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteTaskRequest) SetOwnerId(v int64) *CloudDeleteTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteTaskRequest) SetResourceOwnerAccount(v string) *CloudDeleteTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteTaskRequest) SetResourceOwnerId(v int64) *CloudDeleteTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteTaskRequest) SetTaskId(v int64) *CloudDeleteTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CloudDeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
