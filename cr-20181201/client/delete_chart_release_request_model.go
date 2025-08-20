// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChartReleaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChart(v string) *DeleteChartReleaseRequest
	GetChart() *string
	SetInstanceId(v string) *DeleteChartReleaseRequest
	GetInstanceId() *string
	SetRelease(v string) *DeleteChartReleaseRequest
	GetRelease() *string
	SetRepoName(v string) *DeleteChartReleaseRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *DeleteChartReleaseRequest
	GetRepoNamespaceName() *string
}

type DeleteChartReleaseRequest struct {
	// The name of the chart.
	//
	// This parameter is required.
	//
	// example:
	//
	// chart3
	Chart *string `json:"Chart,omitempty" xml:"Chart,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The version of the chart that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.1.0
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
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

func (s DeleteChartReleaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChartReleaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteChartReleaseRequest) GetChart() *string {
	return s.Chart
}

func (s *DeleteChartReleaseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteChartReleaseRequest) GetRelease() *string {
	return s.Release
}

func (s *DeleteChartReleaseRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DeleteChartReleaseRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *DeleteChartReleaseRequest) SetChart(v string) *DeleteChartReleaseRequest {
	s.Chart = &v
	return s
}

func (s *DeleteChartReleaseRequest) SetInstanceId(v string) *DeleteChartReleaseRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteChartReleaseRequest) SetRelease(v string) *DeleteChartReleaseRequest {
	s.Release = &v
	return s
}

func (s *DeleteChartReleaseRequest) SetRepoName(v string) *DeleteChartReleaseRequest {
	s.RepoName = &v
	return s
}

func (s *DeleteChartReleaseRequest) SetRepoNamespaceName(v string) *DeleteChartReleaseRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *DeleteChartReleaseRequest) Validate() error {
	return dara.Validate(s)
}
