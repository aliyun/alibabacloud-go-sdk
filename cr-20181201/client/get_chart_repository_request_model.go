// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChartRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetChartRepositoryRequest
	GetInstanceId() *string
	SetRepoName(v string) *GetChartRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *GetChartRepositoryRequest
	GetRepoNamespaceName() *string
}

type GetChartRepositoryRequest struct {
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetChartRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *GetChartRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetChartRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *GetChartRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *GetChartRepositoryRequest) SetInstanceId(v string) *GetChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetChartRepositoryRequest) SetRepoName(v string) *GetChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *GetChartRepositoryRequest) SetRepoNamespaceName(v string) *GetChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *GetChartRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
