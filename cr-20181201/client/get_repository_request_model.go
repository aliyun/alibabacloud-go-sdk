// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRepositoryRequest
	GetInstanceId() *string
	SetRepoId(v string) *GetRepositoryRequest
	GetRepoId() *string
	SetRepoName(v string) *GetRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *GetRepositoryRequest
	GetRepoNamespaceName() *string
}

type GetRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crr-03cuozrsqhkw****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryRequest) GoString() string {
	return s.String()
}

func (s *GetRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepositoryRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *GetRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *GetRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *GetRepositoryRequest) SetInstanceId(v string) *GetRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepositoryRequest) SetRepoId(v string) *GetRepositoryRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepositoryRequest) SetRepoName(v string) *GetRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *GetRepositoryRequest) SetRepoNamespaceName(v string) *GetRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *GetRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
