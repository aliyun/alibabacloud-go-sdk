// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTeamResponseBody
	GetCode() *string
	SetData(v *UpdateTeamResponseBodyData) *UpdateTeamResponseBody
	GetData() *UpdateTeamResponseBodyData
	SetHttpStatusCode(v int32) *UpdateTeamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTeamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTeamResponseBody
	GetSuccess() *bool
}

type UpdateTeamResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateTeamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTeamResponseBody) GetData() *UpdateTeamResponseBodyData {
	return s.Data
}

func (s *UpdateTeamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTeamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTeamResponseBody) SetCode(v string) *UpdateTeamResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTeamResponseBody) SetData(v *UpdateTeamResponseBodyData) *UpdateTeamResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTeamResponseBody) SetHttpStatusCode(v int32) *UpdateTeamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTeamResponseBody) SetMessage(v string) *UpdateTeamResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTeamResponseBody) SetRequestId(v string) *UpdateTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTeamResponseBody) SetSuccess(v bool) *UpdateTeamResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTeamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTeamResponseBodyData struct {
	Agents []*UpdateTeamResponseBodyDataAgents `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 负责智能客服业务的团队
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// team-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// tm-123456
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string                            `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Users     []*UpdateTeamResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateTeamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponseBodyData) GetAgents() []*UpdateTeamResponseBodyDataAgents {
	return s.Agents
}

func (s *UpdateTeamResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateTeamResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateTeamResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateTeamResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateTeamResponseBodyData) GetTeamId() *string {
	return s.TeamId
}

func (s *UpdateTeamResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateTeamResponseBodyData) GetUsers() []*UpdateTeamResponseBodyDataUsers {
	return s.Users
}

func (s *UpdateTeamResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateTeamResponseBodyData) SetAgents(v []*UpdateTeamResponseBodyDataAgents) *UpdateTeamResponseBodyData {
	s.Agents = v
	return s
}

func (s *UpdateTeamResponseBodyData) SetCreatedAt(v string) *UpdateTeamResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetDescription(v string) *UpdateTeamResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetName(v string) *UpdateTeamResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetStatus(v string) *UpdateTeamResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetTeamId(v string) *UpdateTeamResponseBodyData {
	s.TeamId = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetUpdatedAt(v string) *UpdateTeamResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetUsers(v []*UpdateTeamResponseBodyDataUsers) *UpdateTeamResponseBodyData {
	s.Users = v
	return s
}

func (s *UpdateTeamResponseBodyData) SetWorkspaceId(v string) *UpdateTeamResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateTeamResponseBodyData) Validate() error {
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

type UpdateTeamResponseBodyDataAgents struct {
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// WORKER
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
}

func (s UpdateTeamResponseBodyDataAgents) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamResponseBodyDataAgents) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponseBodyDataAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateTeamResponseBodyDataAgents) GetTeamRole() *string {
	return s.TeamRole
}

func (s *UpdateTeamResponseBodyDataAgents) SetAgentId(v string) *UpdateTeamResponseBodyDataAgents {
	s.AgentId = &v
	return s
}

func (s *UpdateTeamResponseBodyDataAgents) SetTeamRole(v string) *UpdateTeamResponseBodyDataAgents {
	s.TeamRole = &v
	return s
}

func (s *UpdateTeamResponseBodyDataAgents) Validate() error {
	return dara.Validate(s)
}

type UpdateTeamResponseBodyDataUsers struct {
	// example:
	//
	// ADMIN
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
	// example:
	//
	// usr-123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateTeamResponseBodyDataUsers) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponseBodyDataUsers) GetTeamRole() *string {
	return s.TeamRole
}

func (s *UpdateTeamResponseBodyDataUsers) GetUserId() *string {
	return s.UserId
}

func (s *UpdateTeamResponseBodyDataUsers) SetTeamRole(v string) *UpdateTeamResponseBodyDataUsers {
	s.TeamRole = &v
	return s
}

func (s *UpdateTeamResponseBodyDataUsers) SetUserId(v string) *UpdateTeamResponseBodyDataUsers {
	s.UserId = &v
	return s
}

func (s *UpdateTeamResponseBodyDataUsers) Validate() error {
	return dara.Validate(s)
}
