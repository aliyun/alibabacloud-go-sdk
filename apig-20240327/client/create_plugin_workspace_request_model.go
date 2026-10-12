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
	// The gateway type. If this parameter is not specified, the server uses API as the default value.
	//
	// example:
	//
	// AI
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The Codeup organization ID. This is the Yunxiao DevOps organization ID, which can be obtained by parsing the response of the ListPluginRepositories operation or by creating an organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// 664f1e2xxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// The repository name, which also serves as the English ID of the plugin. The backend does not validate the format. The frontend restricts the value to the pattern /^(?![-.])[a-z0-9.-]{1,20}(?<![-.])$/, which means 1 to 20 lowercase letters, digits, periods, or hyphens that cannot start or end with a period or hyphen.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-custom-plugin
	RepoName *string `json:"repoName,omitempty" xml:"repoName,omitempty"`
	// The workspace name, which is the Chinese name of the plugin. This parameter is required. If you do not specify this parameter, the API returns a 400 InvalidParameter.IsEmpty error.
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
