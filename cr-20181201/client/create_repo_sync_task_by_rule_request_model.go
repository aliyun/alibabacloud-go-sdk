// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSyncTaskByRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateRepoSyncTaskByRuleRequest
	GetInstanceId() *string
	SetPriority(v int32) *CreateRepoSyncTaskByRuleRequest
	GetPriority() *int32
	SetRepoId(v string) *CreateRepoSyncTaskByRuleRequest
	GetRepoId() *string
	SetSyncRuleId(v string) *CreateRepoSyncTaskByRuleRequest
	GetSyncRuleId() *string
	SetTag(v string) *CreateRepoSyncTaskByRuleRequest
	GetTag() *string
}

type CreateRepoSyncTaskByRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The execution priority of the synchronization task. Synchronization tasks are executed in descending order of priority. Synchronization tasks with the same priority are executed in random order.
	//
	// Valid values: 1 to 5.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The image repository ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-hnoq7j93or3k****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The synchronization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crsr-o8n4dijbumgq****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The image version to be synchronized.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.24
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateRepoSyncTaskByRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSyncTaskByRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskByRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRepoSyncTaskByRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateRepoSyncTaskByRuleRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateRepoSyncTaskByRuleRequest) GetSyncRuleId() *string {
	return s.SyncRuleId
}

func (s *CreateRepoSyncTaskByRuleRequest) GetTag() *string {
	return s.Tag
}

func (s *CreateRepoSyncTaskByRuleRequest) SetInstanceId(v string) *CreateRepoSyncTaskByRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleRequest) SetPriority(v int32) *CreateRepoSyncTaskByRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleRequest) SetRepoId(v string) *CreateRepoSyncTaskByRuleRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleRequest) SetSyncRuleId(v string) *CreateRepoSyncTaskByRuleRequest {
	s.SyncRuleId = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleRequest) SetTag(v string) *CreateRepoSyncTaskByRuleRequest {
	s.Tag = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleRequest) Validate() error {
	return dara.Validate(s)
}
