// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *UpdateRepositoryRequest
	GetDetail() *string
	SetInstanceId(v string) *UpdateRepositoryRequest
	GetInstanceId() *string
	SetRepoId(v string) *UpdateRepositoryRequest
	GetRepoId() *string
	SetRepoName(v string) *UpdateRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *UpdateRepositoryRequest
	GetRepoNamespaceName() *string
	SetRepoType(v string) *UpdateRepositoryRequest
	GetRepoType() *string
	SetSummary(v string) *UpdateRepositoryRequest
	GetSummary() *string
	SetTagImmutability(v bool) *UpdateRepositoryRequest
	GetTagImmutability() *bool
}

type UpdateRepositoryRequest struct {
	// example:
	//
	// repo-for-test
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// example:
	//
	// dsp/domain-microapp
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// example:
	//
	// ejiayou-other
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test repo
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// true
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s UpdateRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryRequest) GetDetail() *string {
	return s.Detail
}

func (s *UpdateRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRepositoryRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *UpdateRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *UpdateRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *UpdateRepositoryRequest) GetRepoType() *string {
	return s.RepoType
}

func (s *UpdateRepositoryRequest) GetSummary() *string {
	return s.Summary
}

func (s *UpdateRepositoryRequest) GetTagImmutability() *bool {
	return s.TagImmutability
}

func (s *UpdateRepositoryRequest) SetDetail(v string) *UpdateRepositoryRequest {
	s.Detail = &v
	return s
}

func (s *UpdateRepositoryRequest) SetInstanceId(v string) *UpdateRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRepositoryRequest) SetRepoId(v string) *UpdateRepositoryRequest {
	s.RepoId = &v
	return s
}

func (s *UpdateRepositoryRequest) SetRepoName(v string) *UpdateRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *UpdateRepositoryRequest) SetRepoNamespaceName(v string) *UpdateRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *UpdateRepositoryRequest) SetRepoType(v string) *UpdateRepositoryRequest {
	s.RepoType = &v
	return s
}

func (s *UpdateRepositoryRequest) SetSummary(v string) *UpdateRepositoryRequest {
	s.Summary = &v
	return s
}

func (s *UpdateRepositoryRequest) SetTagImmutability(v bool) *UpdateRepositoryRequest {
	s.TagImmutability = &v
	return s
}

func (s *UpdateRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
