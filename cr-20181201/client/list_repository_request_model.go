// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRepositoryRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListRepositoryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepositoryRequest
	GetPageSize() *int32
	SetRepoName(v string) *ListRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *ListRepositoryRequest
	GetRepoNamespaceName() *string
	SetRepoStatus(v string) *ListRepositoryRequest
	GetRepoStatus() *string
}

type ListRepositoryRequest struct {
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100. If you specify a value larger than 100 for this parameter, the system reports a parameter error or uses 100 as the maximum value.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// repo-test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// repo-namespace-test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// example:
	//
	// ALL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
}

func (s ListRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepositoryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepositoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *ListRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListRepositoryRequest) GetRepoStatus() *string {
	return s.RepoStatus
}

func (s *ListRepositoryRequest) SetInstanceId(v string) *ListRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepositoryRequest) SetPageNo(v int32) *ListRepositoryRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepositoryRequest) SetPageSize(v int32) *ListRepositoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryRequest) SetRepoName(v string) *ListRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *ListRepositoryRequest) SetRepoNamespaceName(v string) *ListRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListRepositoryRequest) SetRepoStatus(v string) *ListRepositoryRequest {
	s.RepoStatus = &v
	return s
}

func (s *ListRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
