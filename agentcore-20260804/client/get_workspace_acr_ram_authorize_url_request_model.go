// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceAcrRamAuthorizeUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *GetWorkspaceAcrRamAuthorizeUrlRequest
	GetAcrInstanceId() *string
	SetNamespace(v string) *GetWorkspaceAcrRamAuthorizeUrlRequest
	GetNamespace() *string
	SetRepository(v string) *GetWorkspaceAcrRamAuthorizeUrlRequest
	GetRepository() *string
}

type GetWorkspaceAcrRamAuthorizeUrlRequest struct {
	// The ACR Enterprise instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-1234567890abcdef
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The target ACR namespace, which corresponds to Agent artifact.container.namespace. This is not a Kubernetes namespace. Together with the instance and repository, this parameter determines the authorization scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The target repository name, which corresponds to Agent artifact.container.repo. Do not include a tag, namespace, or path separator. Wildcards are not accepted.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-image
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
}

func (s GetWorkspaceAcrRamAuthorizeUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceAcrRamAuthorizeUrlRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceAcrRamAuthorizeUrlRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *GetWorkspaceAcrRamAuthorizeUrlRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetWorkspaceAcrRamAuthorizeUrlRequest) GetRepository() *string {
	return s.Repository
}

func (s *GetWorkspaceAcrRamAuthorizeUrlRequest) SetAcrInstanceId(v string) *GetWorkspaceAcrRamAuthorizeUrlRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlRequest) SetNamespace(v string) *GetWorkspaceAcrRamAuthorizeUrlRequest {
	s.Namespace = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlRequest) SetRepository(v string) *GetWorkspaceAcrRamAuthorizeUrlRequest {
	s.Repository = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlRequest) Validate() error {
	return dara.Validate(s)
}
