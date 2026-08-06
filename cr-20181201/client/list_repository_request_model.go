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
	SetMaxResults(v int32) *ListRepositoryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRepositoryRequest
	GetNextToken() *string
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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The maximum value is 100. If the specified value exceeds 100, the system returns a parameter error or uses 100 as the actual maximum number of entries returned.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The repository name.
	//
	// example:
	//
	// repo-test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The repository namespace name.
	//
	// example:
	//
	// repo-namespace-test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The repository status. Valid values:
	//
	// - `NORMAL`: Normal.
	//
	// - `DELETING`: Being deleted.
	//
	// - `DELETED`: Deleted.
	//
	// - `ALL`: All repository statuses.
	//
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

func (s *ListRepositoryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRepositoryRequest) GetNextToken() *string {
	return s.NextToken
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

func (s *ListRepositoryRequest) SetMaxResults(v int32) *ListRepositoryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRepositoryRequest) SetNextToken(v string) *ListRepositoryRequest {
	s.NextToken = &v
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
