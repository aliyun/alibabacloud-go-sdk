// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudGetTaskRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudGetTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudGetTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudGetTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CloudGetTaskRequest
	GetTaskId() *int64
}

type CloudGetTaskRequest struct {
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
	// 31
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloudGetTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudGetTaskRequest) GoString() string {
	return s.String()
}

func (s *CloudGetTaskRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudGetTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudGetTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudGetTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudGetTaskRequest) SetEnterpriseId(v int64) *CloudGetTaskRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetTaskRequest) SetOwnerId(v int64) *CloudGetTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudGetTaskRequest) SetResourceOwnerAccount(v string) *CloudGetTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudGetTaskRequest) SetResourceOwnerId(v int64) *CloudGetTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudGetTaskRequest) SetTaskId(v int64) *CloudGetTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CloudGetTaskRequest) Validate() error {
	return dara.Validate(s)
}
