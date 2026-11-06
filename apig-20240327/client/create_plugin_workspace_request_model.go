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
	// The gateway type. If this parameter is not specified, the server defaults to API.
	//
	// example:
	//
	// AI
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The Codeup organization ID for the Yunxiao DevOps organization. The frontend obtains this ID by parsing the ListPluginRepositories response or after creating an organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// 664f1e2xxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// The repository name, which also serves as the English ID of the plugin. The backend does not validate the format. The frontend restricts the value to the regular expression /^(?![-.])[a-z0-9.-]{1,20}(?<![-.])$/, which allows 1 to 20 lowercase letters, digits, periods, or hyphens, and must not start or end with a period or a hyphen.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-custom-plugin
	RepoName *string `json:"repoName,omitempty" xml:"repoName,omitempty"`
	// The workspace name, which is the Chinese name of the plugin. Note: This parameter is required. If it is omitted, the API returns a 400 InvalidParameter.IsEmpty error.
	//
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
