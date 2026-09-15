// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateTeamRequestBody) *UpdateTeamRequest
	GetBody() *UpdateTeamRequestBody
	SetClientToken(v string) *UpdateTeamRequest
	GetClientToken() *string
}

type UpdateTeamRequest struct {
	// The request body for updating the team.
	Body *UpdateTeamRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// Not supported.
	//
	// example:
	//
	// Not supported
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamRequest) GoString() string {
	return s.String()
}

func (s *UpdateTeamRequest) GetBody() *UpdateTeamRequestBody {
	return s.Body
}

func (s *UpdateTeamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTeamRequest) SetBody(v *UpdateTeamRequestBody) *UpdateTeamRequest {
	s.Body = v
	return s
}

func (s *UpdateTeamRequest) SetClientToken(v string) *UpdateTeamRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTeamRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTeamRequestBody struct {
	// The new agent member list. Replaces the existing agent members using full overwrite semantics. If not specified, the existing agent members remain unchanged.
	Agents []*UpdateTeamRequestBodyAgents `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
	// The new team description. If not specified, the existing description remains unchanged.
	//
	// example:
	//
	// A team responsible for intelligent customer service
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The new user member list. Replaces the existing user members using full overwrite semantics. When specified, the list must contain exactly one member with the ADMIN role. If not specified, the existing user members remain unchanged.
	Users []*UpdateTeamRequestBodyUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s UpdateTeamRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateTeamRequestBody) GetAgents() []*UpdateTeamRequestBodyAgents {
	return s.Agents
}

func (s *UpdateTeamRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateTeamRequestBody) GetUsers() []*UpdateTeamRequestBodyUsers {
	return s.Users
}

func (s *UpdateTeamRequestBody) SetAgents(v []*UpdateTeamRequestBodyAgents) *UpdateTeamRequestBody {
	s.Agents = v
	return s
}

func (s *UpdateTeamRequestBody) SetDescription(v string) *UpdateTeamRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateTeamRequestBody) SetUsers(v []*UpdateTeamRequestBodyUsers) *UpdateTeamRequestBody {
	s.Users = v
	return s
}

func (s *UpdateTeamRequestBody) Validate() error {
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

type UpdateTeamRequestBodyAgents struct {
	// The agent ID.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The role of the agent in the team. Valid values:
	//
	// - LEADER
	//
	// - WORKER
	//
	// example:
	//
	// WORKER
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
}

func (s UpdateTeamRequestBodyAgents) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamRequestBodyAgents) GoString() string {
	return s.String()
}

func (s *UpdateTeamRequestBodyAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateTeamRequestBodyAgents) GetTeamRole() *string {
	return s.TeamRole
}

func (s *UpdateTeamRequestBodyAgents) SetAgentId(v string) *UpdateTeamRequestBodyAgents {
	s.AgentId = &v
	return s
}

func (s *UpdateTeamRequestBodyAgents) SetTeamRole(v string) *UpdateTeamRequestBodyAgents {
	s.TeamRole = &v
	return s
}

func (s *UpdateTeamRequestBodyAgents) Validate() error {
	return dara.Validate(s)
}

type UpdateTeamRequestBodyUsers struct {
	// The role of the user in the team. Valid values:
	//
	// - ADMIN
	//
	// - MEMBER
	//
	// Each team must have exactly one ADMIN.
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

func (s UpdateTeamRequestBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamRequestBodyUsers) GoString() string {
	return s.String()
}

func (s *UpdateTeamRequestBodyUsers) GetTeamRole() *string {
	return s.TeamRole
}

func (s *UpdateTeamRequestBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *UpdateTeamRequestBodyUsers) SetTeamRole(v string) *UpdateTeamRequestBodyUsers {
	s.TeamRole = &v
	return s
}

func (s *UpdateTeamRequestBodyUsers) SetUserId(v string) *UpdateTeamRequestBodyUsers {
	s.UserId = &v
	return s
}

func (s *UpdateTeamRequestBodyUsers) Validate() error {
	return dara.Validate(s)
}
