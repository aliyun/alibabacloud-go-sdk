// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWorkspaceAcrRamAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *VerifyWorkspaceAcrRamAuthorizationRequest
	GetAcrInstanceId() *string
	SetNamespace(v string) *VerifyWorkspaceAcrRamAuthorizationRequest
	GetNamespace() *string
	SetRepository(v string) *VerifyWorkspaceAcrRamAuthorizationRequest
	GetRepository() *string
}

type VerifyWorkspaceAcrRamAuthorizationRequest struct {
	// The ID of the ACR Enterprise instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-1234567890abcdef
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The target ACR namespace, which corresponds to artifact.container.namespace of the agent. This is not a Kubernetes namespace. It determines the authorization scope together with the instance and repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The target repository name, which corresponds to artifact.container.repo of the agent. It does not include tags, namespaces, or path separators. Wildcards are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-image
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
}

func (s VerifyWorkspaceAcrRamAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceAcrRamAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceAcrRamAuthorizationRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *VerifyWorkspaceAcrRamAuthorizationRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *VerifyWorkspaceAcrRamAuthorizationRequest) GetRepository() *string {
	return s.Repository
}

func (s *VerifyWorkspaceAcrRamAuthorizationRequest) SetAcrInstanceId(v string) *VerifyWorkspaceAcrRamAuthorizationRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationRequest) SetNamespace(v string) *VerifyWorkspaceAcrRamAuthorizationRequest {
	s.Namespace = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationRequest) SetRepository(v string) *VerifyWorkspaceAcrRamAuthorizationRequest {
	s.Repository = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
