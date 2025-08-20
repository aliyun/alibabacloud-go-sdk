// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoSourceCodeRepoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRepoSourceCodeRepoRequest
	GetInstanceId() *string
	SetRepoId(v string) *GetRepoSourceCodeRepoRequest
	GetRepoId() *string
}

type GetRepoSourceCodeRepoRequest struct {
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-gzsrlevmvoaq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s GetRepoSourceCodeRepoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepoSourceCodeRepoRequest) GoString() string {
	return s.String()
}

func (s *GetRepoSourceCodeRepoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepoSourceCodeRepoRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *GetRepoSourceCodeRepoRequest) SetInstanceId(v string) *GetRepoSourceCodeRepoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoSourceCodeRepoRequest) SetRepoId(v string) *GetRepoSourceCodeRepoRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoSourceCodeRepoRequest) Validate() error {
	return dara.Validate(s)
}
