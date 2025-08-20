// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChartRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateChartRepositoryRequest
	GetInstanceId() *string
	SetRepoName(v string) *UpdateChartRepositoryRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *UpdateChartRepositoryRequest
	GetRepoNamespaceName() *string
	SetRepoType(v string) *UpdateChartRepositoryRequest
	GetRepoType() *string
	SetSummary(v string) *UpdateChartRepositoryRequest
	GetSummary() *string
}

type UpdateChartRepositoryRequest struct {
	// The ID of the instance.
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
	// The name of the namespace to which the repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PUBLIC`: a public repository.
	//
	// 	- `PRIVATE`: a private repository.
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The summary of the repository.
	//
	// example:
	//
	// test
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s UpdateChartRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateChartRepositoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateChartRepositoryRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *UpdateChartRepositoryRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *UpdateChartRepositoryRequest) GetRepoType() *string {
	return s.RepoType
}

func (s *UpdateChartRepositoryRequest) GetSummary() *string {
	return s.Summary
}

func (s *UpdateChartRepositoryRequest) SetInstanceId(v string) *UpdateChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateChartRepositoryRequest) SetRepoName(v string) *UpdateChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *UpdateChartRepositoryRequest) SetRepoNamespaceName(v string) *UpdateChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *UpdateChartRepositoryRequest) SetRepoType(v string) *UpdateChartRepositoryRequest {
	s.RepoType = &v
	return s
}

func (s *UpdateChartRepositoryRequest) SetSummary(v string) *UpdateChartRepositoryRequest {
	s.Summary = &v
	return s
}

func (s *UpdateChartRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
