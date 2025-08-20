// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *CreateRepositoryRequest
	GetDetail() *string
	SetInstanceId(v string) *CreateRepositoryRequest
	GetInstanceId() *string
	SetRepoName(v string) *CreateRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *CreateRepositoryRequest
	GetRepoNamespaceName() *string
	SetRepoType(v string) *CreateRepositoryRequest
	GetRepoType() *string
	SetSummary(v string) *CreateRepositoryRequest
	GetSummary() *string
	SetTagImmutability(v bool) *CreateRepositoryRequest
	GetTagImmutability() *bool
}

type CreateRepositoryRequest struct {
	// The description of the repository.
	//
	// example:
	//
	// repo1
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the image repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace01
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PUBLIC`: The repository is a public repository.
	//
	// 	- `PRIVATE`: The repository is a private repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// PRIVATE
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The summary about the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo1
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Specifies whether to enable the feature of image tag immutability. Valid values:
	//
	// 	- `true`: enables the feature of image tag immutability.
	//
	// 	- `false`: disables the feature of image tag immutability.
	//
	// example:
	//
	// true
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s CreateRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryRequest) GetDetail() *string {
	return s.Detail
}

func (s *CreateRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *CreateRepositoryRequest) GetRepoType() *string {
	return s.RepoType
}

func (s *CreateRepositoryRequest) GetSummary() *string {
	return s.Summary
}

func (s *CreateRepositoryRequest) GetTagImmutability() *bool {
	return s.TagImmutability
}

func (s *CreateRepositoryRequest) SetDetail(v string) *CreateRepositoryRequest {
	s.Detail = &v
	return s
}

func (s *CreateRepositoryRequest) SetInstanceId(v string) *CreateRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepositoryRequest) SetRepoName(v string) *CreateRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *CreateRepositoryRequest) SetRepoNamespaceName(v string) *CreateRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *CreateRepositoryRequest) SetRepoType(v string) *CreateRepositoryRequest {
	s.RepoType = &v
	return s
}

func (s *CreateRepositoryRequest) SetSummary(v string) *CreateRepositoryRequest {
	s.Summary = &v
	return s
}

func (s *CreateRepositoryRequest) SetTagImmutability(v bool) *CreateRepositoryRequest {
	s.TagImmutability = &v
	return s
}

func (s *CreateRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
