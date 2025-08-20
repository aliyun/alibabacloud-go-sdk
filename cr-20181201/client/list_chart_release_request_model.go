// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChartReleaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChart(v string) *ListChartReleaseRequest
	GetChart() *string
	SetInstanceId(v string) *ListChartReleaseRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListChartReleaseRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListChartReleaseRequest
	GetPageSize() *int32
	SetRepoName(v string) *ListChartReleaseRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *ListChartReleaseRequest
	GetRepoNamespaceName() *string
}

type ListChartReleaseRequest struct {
	// The chart whose versions you want to query.
	//
	// example:
	//
	// null
	Chart *string `json:"Chart,omitempty" xml:"Chart,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListChartReleaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChartReleaseRequest) GoString() string {
	return s.String()
}

func (s *ListChartReleaseRequest) GetChart() *string {
	return s.Chart
}

func (s *ListChartReleaseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChartReleaseRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChartReleaseRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChartReleaseRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *ListChartReleaseRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListChartReleaseRequest) SetChart(v string) *ListChartReleaseRequest {
	s.Chart = &v
	return s
}

func (s *ListChartReleaseRequest) SetInstanceId(v string) *ListChartReleaseRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChartReleaseRequest) SetPageNo(v int32) *ListChartReleaseRequest {
	s.PageNo = &v
	return s
}

func (s *ListChartReleaseRequest) SetPageSize(v int32) *ListChartReleaseRequest {
	s.PageSize = &v
	return s
}

func (s *ListChartReleaseRequest) SetRepoName(v string) *ListChartReleaseRequest {
	s.RepoName = &v
	return s
}

func (s *ListChartReleaseRequest) SetRepoNamespaceName(v string) *ListChartReleaseRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListChartReleaseRequest) Validate() error {
	return dara.Validate(s)
}
