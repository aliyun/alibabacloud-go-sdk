// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudStartTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudStartTaskRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudStartTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudStartTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudStartTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CloudStartTaskRequest
	GetTaskId() *int64
}

type CloudStartTaskRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8001654
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 外呼任务id
	//
	// This parameter is required.
	//
	// example:
	//
	// 32
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloudStartTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudStartTaskRequest) GoString() string {
	return s.String()
}

func (s *CloudStartTaskRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudStartTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudStartTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudStartTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudStartTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudStartTaskRequest) SetEnterpriseId(v int64) *CloudStartTaskRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudStartTaskRequest) SetOwnerId(v int64) *CloudStartTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudStartTaskRequest) SetResourceOwnerAccount(v string) *CloudStartTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudStartTaskRequest) SetResourceOwnerId(v int64) *CloudStartTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudStartTaskRequest) SetTaskId(v int64) *CloudStartTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CloudStartTaskRequest) Validate() error {
	return dara.Validate(s)
}
