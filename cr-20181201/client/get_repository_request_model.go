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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The repository ID. This API supports querying by RepoId, or by RepoNamespaceName and RepoName together.
	//
	// example:
	//
	// crr-03cuozrsqhkw****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The repository name. This API supports querying by RepoId, or by RepoNamespaceName and RepoName together.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The repository namespace name. This API supports querying by RepoId, or by RepoNamespaceName and RepoName together.
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
