// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChartRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListChartRepositoryRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListChartRepositoryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListChartRepositoryRequest
	GetPageSize() *int32
	SetRepoName(v string) *ListChartRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *ListChartRepositoryRequest
	GetRepoNamespaceName() *string
	SetRepoStatus(v string) *ListChartRepositoryRequest
	GetRepoStatus() *string
}

type ListChartRepositoryRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The repository name.
	//
	// example:
	//
	// ns1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace.
	//
	// example:
	//
	// repo1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the chart repository. Valid values:
	//
	// - `NORMAL`: The chart repository is in the Normal state.
	//
	//
	//
	// - `DELETING`: The chart repository is being deleted.
	//
	// example:
	//
	// NORMAL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
}

func (s ListChartRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *ListChartRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChartRepositoryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChartRepositoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChartRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *ListChartRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListChartRepositoryRequest) GetRepoStatus() *string {
	return s.RepoStatus
}

func (s *ListChartRepositoryRequest) SetInstanceId(v string) *ListChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChartRepositoryRequest) SetPageNo(v int32) *ListChartRepositoryRequest {
	s.PageNo = &v
	return s
}

func (s *ListChartRepositoryRequest) SetPageSize(v int32) *ListChartRepositoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListChartRepositoryRequest) SetRepoName(v string) *ListChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *ListChartRepositoryRequest) SetRepoNamespaceName(v string) *ListChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListChartRepositoryRequest) SetRepoStatus(v string) *ListChartRepositoryRequest {
	s.RepoStatus = &v
	return s
}

func (s *ListChartRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
