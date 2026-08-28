// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayType(v string) *CreatePluginWorkspaceRequest
	GetGatewayType() *string
	SetOrganizationId(v string) *CreatePluginWorkspaceRequest
	GetOrganizationId() *string
	SetRepoName(v string) *CreatePluginWorkspaceRequest
	GetRepoName() *string
	SetWorkspaceName(v string) *CreatePluginWorkspaceRequest
	GetWorkspaceName() *string
}

type CreatePluginWorkspaceRequest struct {
	// example:
	//
	// AI
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 664f1e2xxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-custom-plugin
	RepoName *string `json:"repoName,omitempty" xml:"repoName,omitempty"`
	// example:
	//
	// my-plugin
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s CreatePluginWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreatePluginWorkspaceRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreatePluginWorkspaceRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreatePluginWorkspaceRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *CreatePluginWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreatePluginWorkspaceRequest) SetGatewayType(v string) *CreatePluginWorkspaceRequest {
	s.GatewayType = &v
	return s
}

func (s *CreatePluginWorkspaceRequest) SetOrganizationId(v string) *CreatePluginWorkspaceRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreatePluginWorkspaceRequest) SetRepoName(v string) *CreatePluginWorkspaceRequest {
	s.RepoName = &v
	return s
}

func (s *CreatePluginWorkspaceRequest) SetWorkspaceName(v string) *CreatePluginWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreatePluginWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
