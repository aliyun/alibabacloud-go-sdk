// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTeamResponseBody
	GetCode() *string
	SetData(v *CreateTeamResponseBodyData) *CreateTeamResponseBody
	GetData() *CreateTeamResponseBodyData
	SetHttpStatusCode(v int32) *CreateTeamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTeamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTeamResponseBody
	GetSuccess() *bool
}

type CreateTeamResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The information about the created team.
	Data *CreateTeamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTeamResponseBody) GetData() *CreateTeamResponseBodyData {
	return s.Data
}

func (s *CreateTeamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTeamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTeamResponseBody) SetCode(v string) *CreateTeamResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTeamResponseBody) SetData(v *CreateTeamResponseBodyData) *CreateTeamResponseBody {
	s.Data = v
	return s
}

func (s *CreateTeamResponseBody) SetHttpStatusCode(v int32) *CreateTeamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTeamResponseBody) SetMessage(v string) *CreateTeamResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTeamResponseBody) SetRequestId(v string) *CreateTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTeamResponseBody) SetSuccess(v bool) *CreateTeamResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTeamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTeamResponseBodyData struct {
	// The list of agent members in the team.
	Agents []*CreateTeamResponseBodyDataAgents `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
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
	// A team responsible for intelligent customer service
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The team name. The name can contain only lowercase letters, digits, and hyphens (-). It must start and end with a lowercase letter or digit. The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// team-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The team status. Valid values: Creating, Active, Updating, Deleting, Failed, Deleted.
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
	Users []*CreateTeamResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateTeamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTeamResponseBodyData) GetAgents() []*CreateTeamResponseBodyDataAgents {
	return s.Agents
}

func (s *CreateTeamResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateTeamResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateTeamResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateTeamResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateTeamResponseBodyData) GetTeamId() *string {
	return s.TeamId
}

func (s *CreateTeamResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateTeamResponseBodyData) GetUsers() []*CreateTeamResponseBodyDataUsers {
	return s.Users
}

func (s *CreateTeamResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateTeamResponseBodyData) SetAgents(v []*CreateTeamResponseBodyDataAgents) *CreateTeamResponseBodyData {
	s.Agents = v
	return s
}

func (s *CreateTeamResponseBodyData) SetCreatedAt(v string) *CreateTeamResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetDescription(v string) *CreateTeamResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetName(v string) *CreateTeamResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetStatus(v string) *CreateTeamResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetTeamId(v string) *CreateTeamResponseBodyData {
	s.TeamId = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetUpdatedAt(v string) *CreateTeamResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetUsers(v []*CreateTeamResponseBodyDataUsers) *CreateTeamResponseBodyData {
	s.Users = v
	return s
}

func (s *CreateTeamResponseBodyData) SetWorkspaceId(v string) *CreateTeamResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateTeamResponseBodyData) Validate() error {
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

type CreateTeamResponseBodyDataAgents struct {
	// The agent ID.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The role of the agent in the team. Valid values: LEADER, WORKER. Each team must include exactly one LEADER.
	//
	// example:
	//
	// WORKER
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
}

func (s CreateTeamResponseBodyDataAgents) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamResponseBodyDataAgents) GoString() string {
	return s.String()
}

func (s *CreateTeamResponseBodyDataAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateTeamResponseBodyDataAgents) GetTeamRole() *string {
	return s.TeamRole
}

func (s *CreateTeamResponseBodyDataAgents) SetAgentId(v string) *CreateTeamResponseBodyDataAgents {
	s.AgentId = &v
	return s
}

func (s *CreateTeamResponseBodyDataAgents) SetTeamRole(v string) *CreateTeamResponseBodyDataAgents {
	s.TeamRole = &v
	return s
}

func (s *CreateTeamResponseBodyDataAgents) Validate() error {
	return dara.Validate(s)
}

type CreateTeamResponseBodyDataUsers struct {
	// The role of the user in the team. Valid values: ADMIN, MEMBER. Each team must include exactly one ADMIN.
	//
	// example:
	//
	// ADMIN
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
	// The user ID.
	//
	// example:
	//
	// usr-123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateTeamResponseBodyDataUsers) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *CreateTeamResponseBodyDataUsers) GetTeamRole() *string {
	return s.TeamRole
}

func (s *CreateTeamResponseBodyDataUsers) GetUserId() *string {
	return s.UserId
}

func (s *CreateTeamResponseBodyDataUsers) SetTeamRole(v string) *CreateTeamResponseBodyDataUsers {
	s.TeamRole = &v
	return s
}

func (s *CreateTeamResponseBodyDataUsers) SetUserId(v string) *CreateTeamResponseBodyDataUsers {
	s.UserId = &v
	return s
}

func (s *CreateTeamResponseBodyDataUsers) Validate() error {
	return dara.Validate(s)
}
