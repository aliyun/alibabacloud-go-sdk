// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPauseTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudPauseTaskRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudPauseTaskRequest
	GetOwnerId() *int64
	SetPauseDuration(v int64) *CloudPauseTaskRequest
	GetPauseDuration() *int64
	SetResourceOwnerAccount(v string) *CloudPauseTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudPauseTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CloudPauseTaskRequest
	GetTaskId() *int64
}

type CloudPauseTaskRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 暂停时长；默认0
	//
	// example:
	//
	// 30
	PauseDuration        *int64  `json:"PauseDuration,omitempty" xml:"PauseDuration,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 外呼任务id
	//
	// This parameter is required.
	//
	// example:
	//
	// 15
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloudPauseTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudPauseTaskRequest) GoString() string {
	return s.String()
}

func (s *CloudPauseTaskRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudPauseTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudPauseTaskRequest) GetPauseDuration() *int64 {
	return s.PauseDuration
}

func (s *CloudPauseTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudPauseTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudPauseTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudPauseTaskRequest) SetEnterpriseId(v int64) *CloudPauseTaskRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudPauseTaskRequest) SetOwnerId(v int64) *CloudPauseTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudPauseTaskRequest) SetPauseDuration(v int64) *CloudPauseTaskRequest {
	s.PauseDuration = &v
	return s
}

func (s *CloudPauseTaskRequest) SetResourceOwnerAccount(v string) *CloudPauseTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudPauseTaskRequest) SetResourceOwnerId(v int64) *CloudPauseTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudPauseTaskRequest) SetTaskId(v int64) *CloudPauseTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CloudPauseTaskRequest) Validate() error {
	return dara.Validate(s)
}
