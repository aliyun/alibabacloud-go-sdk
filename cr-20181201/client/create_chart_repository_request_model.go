// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChartRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateChartRepositoryRequest
	GetInstanceId() *string
	SetRepoName(v string) *CreateChartRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *CreateChartRepositoryRequest
	GetRepoNamespaceName() *string
	SetRepoType(v string) *CreateChartRepositoryRequest
	GetRepoType() *string
	SetSummary(v string) *CreateChartRepositoryRequest
	GetSummary() *string
}

type CreateChartRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
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
	// The default repository type. Valid values:
	//
	// 	- `PUBLIC`: a public repository.
	//
	// 	- `PRIVATE`: a private repository.
	//
	// You can specify the RepoType or Summary parameter. The RepoType parameter is optional.
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The summary of the repository.
	//
	// example:
	//
	// summary
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s CreateChartRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *CreateChartRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateChartRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateChartRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *CreateChartRepositoryRequest) GetRepoType() *string {
	return s.RepoType
}

func (s *CreateChartRepositoryRequest) GetSummary() *string {
	return s.Summary
}

func (s *CreateChartRepositoryRequest) SetInstanceId(v string) *CreateChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateChartRepositoryRequest) SetRepoName(v string) *CreateChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *CreateChartRepositoryRequest) SetRepoNamespaceName(v string) *CreateChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *CreateChartRepositoryRequest) SetRepoType(v string) *CreateChartRepositoryRequest {
	s.RepoType = &v
	return s
}

func (s *CreateChartRepositoryRequest) SetSummary(v string) *CreateChartRepositoryRequest {
	s.Summary = &v
	return s
}

func (s *CreateChartRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
