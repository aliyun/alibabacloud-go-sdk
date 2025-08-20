// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoSyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRepoSyncTaskRequest
	GetInstanceId() *string
	SetSyncTaskId(v string) *GetRepoSyncTaskRequest
	GetSyncTaskId() *string
}

type GetRepoSyncTaskRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-sgedpenzw80e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// rst-zxjkiv5oil6f****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
}

func (s GetRepoSyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepoSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepoSyncTaskRequest) GetSyncTaskId() *string {
	return s.SyncTaskId
}

func (s *GetRepoSyncTaskRequest) SetInstanceId(v string) *GetRepoSyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoSyncTaskRequest) SetSyncTaskId(v string) *GetRepoSyncTaskRequest {
	s.SyncTaskId = &v
	return s
}

func (s *GetRepoSyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
