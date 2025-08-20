// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRepoSyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CancelRepoSyncTaskRequest
	GetInstanceId() *string
	SetSyncTaskId(v string) *CancelRepoSyncTaskRequest
	GetSyncTaskId() *string
}

type CancelRepoSyncTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the replication task.
	//
	// This parameter is required.
	//
	// example:
	//
	// rst-biu4u4pm4it5****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
}

func (s CancelRepoSyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelRepoSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelRepoSyncTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelRepoSyncTaskRequest) GetSyncTaskId() *string {
	return s.SyncTaskId
}

func (s *CancelRepoSyncTaskRequest) SetInstanceId(v string) *CancelRepoSyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelRepoSyncTaskRequest) SetSyncTaskId(v string) *CancelRepoSyncTaskRequest {
	s.SyncTaskId = &v
	return s
}

func (s *CancelRepoSyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
