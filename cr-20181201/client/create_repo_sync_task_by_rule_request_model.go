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
	SetRepoId(v string) *CreateRepoSyncTaskByRuleRequest
	GetRepoId() *string
	SetSyncRuleId(v string) *CreateRepoSyncTaskByRuleRequest
	GetSyncRuleId() *string
	SetTag(v string) *CreateRepoSyncTaskByRuleRequest
	GetTag() *string
}

type CreateRepoSyncTaskByRuleRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-hnoq7j93or3k****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the synchronization rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crsr-o8n4dijbumgq****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The version of the image to be synchronized.
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
