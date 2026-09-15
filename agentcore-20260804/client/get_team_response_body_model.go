// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTeamResponseBody
	GetCode() *string
	SetData(v *GetTeamResponseBodyData) *GetTeamResponseBody
	GetData() *GetTeamResponseBodyData
	SetHttpStatusCode(v int32) *GetTeamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTeamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTeamResponseBody
	GetSuccess() *bool
}

type GetTeamResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The team details.
	Data *GetTeamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBody) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTeamResponseBody) GetData() *GetTeamResponseBodyData {
	return s.Data
}

func (s *GetTeamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTeamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTeamResponseBody) SetCode(v string) *GetTeamResponseBody {
	s.Code = &v
	return s
}

func (s *GetTeamResponseBody) SetData(v *GetTeamResponseBodyData) *GetTeamResponseBody {
	s.Data = v
	return s
}

func (s *GetTeamResponseBody) SetHttpStatusCode(v int32) *GetTeamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTeamResponseBody) SetMessage(v string) *GetTeamResponseBody {
	s.Message = &v
	return s
}

func (s *GetTeamResponseBody) SetRequestId(v string) *GetTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTeamResponseBody) SetSuccess(v bool) *GetTeamResponseBody {
	s.Success = &v
	return s
}

func (s *GetTeamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTeamResponseBodyData struct {
	// The list of agent members in the team.
	Agents []*GetTeamResponseBodyDataAgents `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
	// The creation time in UTC, formatted in RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The team description.
	//
	// example:
	//
	// A team responsible for intelligent customer service operations
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The team name. The name can contain only lowercase letters, digits, and hyphens (-). It must start and end with a lowercase letter or digit. The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// team-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The team status. Valid values: Creating, Active, Updating, Deleting, Failed, and Deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The team ID.
	//
	// example:
	//
	// tm-123456
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The time of the last modification in UTC, formatted in RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The list of user members in the team.
	Users []*GetTeamResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetTeamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyData) GetAgents() []*GetTeamResponseBodyDataAgents {
	return s.Agents
}

func (s *GetTeamResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetTeamResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetTeamResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTeamResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTeamResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTeamResponseBodyData) GetTeamId() *string {
	return s.TeamId
}

func (s *GetTeamResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetTeamResponseBodyData) GetUsers() []*GetTeamResponseBodyDataUsers {
	return s.Users
}

func (s *GetTeamResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTeamResponseBodyData) SetAgents(v []*GetTeamResponseBodyDataAgents) *GetTeamResponseBodyData {
	s.Agents = v
	return s
}

func (s *GetTeamResponseBodyData) SetCreatedAt(v string) *GetTeamResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetTeamResponseBodyData) SetDescription(v string) *GetTeamResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTeamResponseBodyData) SetName(v string) *GetTeamResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTeamResponseBodyData) SetRegionId(v string) *GetTeamResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetTeamResponseBodyData) SetStatus(v string) *GetTeamResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTeamResponseBodyData) SetTeamId(v string) *GetTeamResponseBodyData {
	s.TeamId = &v
	return s
}

func (s *GetTeamResponseBodyData) SetUpdatedAt(v string) *GetTeamResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetTeamResponseBodyData) SetUsers(v []*GetTeamResponseBodyDataUsers) *GetTeamResponseBodyData {
	s.Users = v
	return s
}

func (s *GetTeamResponseBodyData) SetWorkspaceId(v string) *GetTeamResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetTeamResponseBodyData) Validate() error {
	if s.Agents != nil {
		for _, item := range s.Agents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTeamResponseBodyDataAgents struct {
	// The agent ID.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The creation mode of the agent. CUSTOM indicates custom creation. TEMPLATE indicates creation from a template.
	//
	// example:
	//
	// CUSTOM
	CreateMode *string `json:"createMode,omitempty" xml:"createMode,omitempty"`
	// The creation time in UTC, formatted in RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The deployment type of the agent. MANAGED indicates platform-managed deployment. SELF_HOSTED indicates self-hosted deployment.
	//
	// example:
	//
	// MANAGED
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// The agent description.
	//
	// example:
	//
	// An agent that handles after-sales inquiries
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The latest configuration version number of the agent.
	//
	// example:
	//
	// 2
	LatestSpecVersion *int64 `json:"latestSpecVersion,omitempty" xml:"latestSpecVersion,omitempty"`
	// The agent name.
	//
	// example:
	//
	// agent-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The runtime type of the agent.
	//
	// example:
	//
	// qwenpaw
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The agent status. Valid values: Creating, Running, Failed, Updating, Deleting, and Deleted.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The role of the agent in the team. Valid values: LEADER and WORKER.
	//
	// example:
	//
	// WORKER
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
	// The time of the last modification in UTC, formatted in RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetTeamResponseBodyDataAgents) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBodyDataAgents) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyDataAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *GetTeamResponseBodyDataAgents) GetCreateMode() *string {
	return s.CreateMode
}

func (s *GetTeamResponseBodyDataAgents) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetTeamResponseBodyDataAgents) GetDeployType() *string {
	return s.DeployType
}

func (s *GetTeamResponseBodyDataAgents) GetDescription() *string {
	return s.Description
}

func (s *GetTeamResponseBodyDataAgents) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *GetTeamResponseBodyDataAgents) GetName() *string {
	return s.Name
}

func (s *GetTeamResponseBodyDataAgents) GetRuntime() *string {
	return s.Runtime
}

func (s *GetTeamResponseBodyDataAgents) GetStatus() *string {
	return s.Status
}

func (s *GetTeamResponseBodyDataAgents) GetTeamRole() *string {
	return s.TeamRole
}

func (s *GetTeamResponseBodyDataAgents) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetTeamResponseBodyDataAgents) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTeamResponseBodyDataAgents) SetAgentId(v string) *GetTeamResponseBodyDataAgents {
	s.AgentId = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetCreateMode(v string) *GetTeamResponseBodyDataAgents {
	s.CreateMode = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetCreatedAt(v string) *GetTeamResponseBodyDataAgents {
	s.CreatedAt = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetDeployType(v string) *GetTeamResponseBodyDataAgents {
	s.DeployType = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetDescription(v string) *GetTeamResponseBodyDataAgents {
	s.Description = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetLatestSpecVersion(v int64) *GetTeamResponseBodyDataAgents {
	s.LatestSpecVersion = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetName(v string) *GetTeamResponseBodyDataAgents {
	s.Name = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetRuntime(v string) *GetTeamResponseBodyDataAgents {
	s.Runtime = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetStatus(v string) *GetTeamResponseBodyDataAgents {
	s.Status = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetTeamRole(v string) *GetTeamResponseBodyDataAgents {
	s.TeamRole = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetUpdatedAt(v string) *GetTeamResponseBodyDataAgents {
	s.UpdatedAt = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetWorkspaceId(v string) *GetTeamResponseBodyDataAgents {
	s.WorkspaceId = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) Validate() error {
	return dara.Validate(s)
}

type GetTeamResponseBodyDataUsers struct {
	// The authentication method of the user. password indicates local password authentication in the workspace. dingtalk and feishu indicate synchronization and authentication by the corresponding external identity provider.
	//
	// example:
	//
	// password
	AuthMethod *string `json:"authMethod,omitempty" xml:"authMethod,omitempty"`
	// The creation time in UTC, formatted in RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The display name of the user. The name must be 1 to 32 characters in length.
	//
	// example:
	//
	// John Doe
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address of the user. The address can be up to 256 characters in length.
	//
	// example:
	//
	// user-01@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The initial password of the user. If a password was specified during creation, that password is returned. If no password was specified, a random password generated by the server is returned.
	//
	// example:
	//
	// Example@2026
	InitialPassword *string `json:"initialPassword,omitempty" xml:"initialPassword,omitempty"`
	// The username. The name must be unique within the workspace. It can contain only lowercase letters, digits, and hyphens (-). It must start and end with a lowercase letter or digit. The name must be 1 to 32 characters in length.
	//
	// example:
	//
	// user-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The user note. The note can be up to 1024 characters in length.
	//
	// example:
	//
	// Member of the agent operations group
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// The user status. Valid values: Creating, Active, Updating, Deleting, Failed, and DeleteFailed.
	//
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The role of the user in the team. Valid values: ADMIN and MEMBER. Each team must have exactly one ADMIN.
	//
	// example:
	//
	// ADMIN
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
	// The time of the last modification in UTC, formatted in RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The user ID.
	//
	// example:
	//
	// usr-123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetTeamResponseBodyDataUsers) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyDataUsers) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *GetTeamResponseBodyDataUsers) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetTeamResponseBodyDataUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTeamResponseBodyDataUsers) GetEmail() *string {
	return s.Email
}

func (s *GetTeamResponseBodyDataUsers) GetInitialPassword() *string {
	return s.InitialPassword
}

func (s *GetTeamResponseBodyDataUsers) GetName() *string {
	return s.Name
}

func (s *GetTeamResponseBodyDataUsers) GetNote() *string {
	return s.Note
}

func (s *GetTeamResponseBodyDataUsers) GetStatus() *string {
	return s.Status
}

func (s *GetTeamResponseBodyDataUsers) GetTeamRole() *string {
	return s.TeamRole
}

func (s *GetTeamResponseBodyDataUsers) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetTeamResponseBodyDataUsers) GetUserId() *string {
	return s.UserId
}

func (s *GetTeamResponseBodyDataUsers) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTeamResponseBodyDataUsers) SetAuthMethod(v string) *GetTeamResponseBodyDataUsers {
	s.AuthMethod = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetCreatedAt(v string) *GetTeamResponseBodyDataUsers {
	s.CreatedAt = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetDisplayName(v string) *GetTeamResponseBodyDataUsers {
	s.DisplayName = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetEmail(v string) *GetTeamResponseBodyDataUsers {
	s.Email = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetInitialPassword(v string) *GetTeamResponseBodyDataUsers {
	s.InitialPassword = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetName(v string) *GetTeamResponseBodyDataUsers {
	s.Name = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetNote(v string) *GetTeamResponseBodyDataUsers {
	s.Note = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetStatus(v string) *GetTeamResponseBodyDataUsers {
	s.Status = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetTeamRole(v string) *GetTeamResponseBodyDataUsers {
	s.TeamRole = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetUpdatedAt(v string) *GetTeamResponseBodyDataUsers {
	s.UpdatedAt = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetUserId(v string) *GetTeamResponseBodyDataUsers {
	s.UserId = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetWorkspaceId(v string) *GetTeamResponseBodyDataUsers {
	s.WorkspaceId = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) Validate() error {
	return dara.Validate(s)
}
