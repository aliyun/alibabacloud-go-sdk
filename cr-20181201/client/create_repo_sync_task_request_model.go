// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateRepoSyncTaskRequest
	GetInstanceId() *string
	SetOverride(v bool) *CreateRepoSyncTaskRequest
	GetOverride() *bool
	SetRepoId(v string) *CreateRepoSyncTaskRequest
	GetRepoId() *string
	SetTag(v string) *CreateRepoSyncTaskRequest
	GetTag() *string
	SetTargetInstanceId(v string) *CreateRepoSyncTaskRequest
	GetTargetInstanceId() *string
	SetTargetNamespace(v string) *CreateRepoSyncTaskRequest
	GetTargetNamespace() *string
	SetTargetRegionId(v string) *CreateRepoSyncTaskRequest
	GetTargetRegionId() *string
	SetTargetRepoName(v string) *CreateRepoSyncTaskRequest
	GetTargetRepoName() *string
	SetTargetTag(v string) *CreateRepoSyncTaskRequest
	GetTargetTag() *string
	SetTargetUserId(v string) *CreateRepoSyncTaskRequest
	GetTargetUserId() *string
}

type CreateRepoSyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Override *bool `json:"Override,omitempty" xml:"Override,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crr-iql7jalx4g0****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cri-ibxs3piklys3****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ns1
	TargetNamespace *string `json:"TargetNamespace,omitempty" xml:"TargetNamespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// repo1
	TargetRepoName *string `json:"TargetRepoName,omitempty" xml:"TargetRepoName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tag1
	TargetTag *string `json:"TargetTag,omitempty" xml:"TargetTag,omitempty"`
	// example:
	//
	// 12345***
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s CreateRepoSyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRepoSyncTaskRequest) GetOverride() *bool {
	return s.Override
}

func (s *CreateRepoSyncTaskRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateRepoSyncTaskRequest) GetTag() *string {
	return s.Tag
}

func (s *CreateRepoSyncTaskRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *CreateRepoSyncTaskRequest) GetTargetNamespace() *string {
	return s.TargetNamespace
}

func (s *CreateRepoSyncTaskRequest) GetTargetRegionId() *string {
	return s.TargetRegionId
}

func (s *CreateRepoSyncTaskRequest) GetTargetRepoName() *string {
	return s.TargetRepoName
}

func (s *CreateRepoSyncTaskRequest) GetTargetTag() *string {
	return s.TargetTag
}

func (s *CreateRepoSyncTaskRequest) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *CreateRepoSyncTaskRequest) SetInstanceId(v string) *CreateRepoSyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetOverride(v bool) *CreateRepoSyncTaskRequest {
	s.Override = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetRepoId(v string) *CreateRepoSyncTaskRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTag(v string) *CreateRepoSyncTaskRequest {
	s.Tag = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetInstanceId(v string) *CreateRepoSyncTaskRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetNamespace(v string) *CreateRepoSyncTaskRequest {
	s.TargetNamespace = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetRegionId(v string) *CreateRepoSyncTaskRequest {
	s.TargetRegionId = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetRepoName(v string) *CreateRepoSyncTaskRequest {
	s.TargetRepoName = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetTag(v string) *CreateRepoSyncTaskRequest {
	s.TargetTag = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetUserId(v string) *CreateRepoSyncTaskRequest {
	s.TargetUserId = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
