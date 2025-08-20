// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteRepositoryRequest
	GetInstanceId() *string
	SetRepoId(v string) *DeleteRepositoryRequest
	GetRepoId() *string
	SetRepoName(v string) *DeleteRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *DeleteRepositoryRequest
	GetRepoNamespaceName() *string
}

type DeleteRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-l4933wbcmun2****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// test-namespace
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s DeleteRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRepositoryRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DeleteRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DeleteRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *DeleteRepositoryRequest) SetInstanceId(v string) *DeleteRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepositoryRequest) SetRepoId(v string) *DeleteRepositoryRequest {
	s.RepoId = &v
	return s
}

func (s *DeleteRepositoryRequest) SetRepoName(v string) *DeleteRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *DeleteRepositoryRequest) SetRepoNamespaceName(v string) *DeleteRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *DeleteRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
