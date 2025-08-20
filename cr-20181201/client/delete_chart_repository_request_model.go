// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChartRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteChartRepositoryRequest
	GetInstanceId() *string
	SetRepoName(v string) *DeleteChartRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *DeleteChartRepositoryRequest
	GetRepoNamespaceName() *string
}

type DeleteChartRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo01
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace01
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s DeleteChartRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteChartRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteChartRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DeleteChartRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *DeleteChartRepositoryRequest) SetInstanceId(v string) *DeleteChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteChartRepositoryRequest) SetRepoName(v string) *DeleteChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *DeleteChartRepositoryRequest) SetRepoNamespaceName(v string) *DeleteChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *DeleteChartRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
