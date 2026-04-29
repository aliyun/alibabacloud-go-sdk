// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitorTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudMonitorTaskRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudMonitorTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudMonitorTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudMonitorTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CloudMonitorTaskRequest
	GetTaskId() *int64
}

type CloudMonitorTaskRequest struct {
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
	// 外呼任务id
	//
	// This parameter is required.
	//
	// example:
	//
	// 19
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloudMonitorTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskRequest) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudMonitorTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudMonitorTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudMonitorTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudMonitorTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudMonitorTaskRequest) SetEnterpriseId(v int64) *CloudMonitorTaskRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudMonitorTaskRequest) SetOwnerId(v int64) *CloudMonitorTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudMonitorTaskRequest) SetResourceOwnerAccount(v string) *CloudMonitorTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudMonitorTaskRequest) SetResourceOwnerId(v int64) *CloudMonitorTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudMonitorTaskRequest) SetTaskId(v int64) *CloudMonitorTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CloudMonitorTaskRequest) Validate() error {
	return dara.Validate(s)
}
