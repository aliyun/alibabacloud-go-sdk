// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateTeamRequestBody) *CreateTeamRequest
	GetBody() *CreateTeamRequestBody
	SetClientToken(v string) *CreateTeamRequest
	GetClientToken() *string
}

type CreateTeamRequest struct {
	// The request body for creating a team.
	Body *CreateTeamRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// Not supported.
	//
	// example:
	//
	// Not supported
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamRequest) GoString() string {
	return s.String()
}

func (s *CreateTeamRequest) GetBody() *CreateTeamRequestBody {
	return s.Body
}

func (s *CreateTeamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTeamRequest) SetBody(v *CreateTeamRequestBody) *CreateTeamRequest {
	s.Body = v
	return s
}

func (s *CreateTeamRequest) SetClientToken(v string) *CreateTeamRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTeamRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTeamRequestBody struct {
	// The list of agent members in the team.
	Agents []*CreateTeamRequestBodyAgents `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
	// The team description.
	//
	// example:
	//
	// A team responsible for intelligent customer service
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The team name. The name can contain only lowercase letters, digits, and hyphens (-). It must start and end with a lowercase letter or digit. The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// team-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of user members in the team. The list must include exactly one member with the ADMIN role.
	Users []*CreateTeamRequestBodyUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s CreateTeamRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamRequestBody) GoString() string {
	return s.String()
}

func (s *CreateTeamRequestBody) GetAgents() []*CreateTeamRequestBodyAgents {
	return s.Agents
}

func (s *CreateTeamRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateTeamRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateTeamRequestBody) GetUsers() []*CreateTeamRequestBodyUsers {
	return s.Users
}

func (s *CreateTeamRequestBody) SetAgents(v []*CreateTeamRequestBodyAgents) *CreateTeamRequestBody {
	s.Agents = v
	return s
}

func (s *CreateTeamRequestBody) SetDescription(v string) *CreateTeamRequestBody {
	s.Description = &v
	return s
}

func (s *CreateTeamRequestBody) SetName(v string) *CreateTeamRequestBody {
	s.Name = &v
	return s
}

func (s *CreateTeamRequestBody) SetUsers(v []*CreateTeamRequestBodyUsers) *CreateTeamRequestBody {
	s.Users = v
	return s
}

func (s *CreateTeamRequestBody) Validate() error {
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

type CreateTeamRequestBodyAgents struct {
	// The agent ID.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The role of the agent in the team. Valid values: LEADER, WORKER.
	//
	// example:
	//
	// WORKER
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
}

func (s CreateTeamRequestBodyAgents) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamRequestBodyAgents) GoString() string {
	return s.String()
}

func (s *CreateTeamRequestBodyAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateTeamRequestBodyAgents) GetTeamRole() *string {
	return s.TeamRole
}

func (s *CreateTeamRequestBodyAgents) SetAgentId(v string) *CreateTeamRequestBodyAgents {
	s.AgentId = &v
	return s
}

func (s *CreateTeamRequestBodyAgents) SetTeamRole(v string) *CreateTeamRequestBodyAgents {
	s.TeamRole = &v
	return s
}

func (s *CreateTeamRequestBodyAgents) Validate() error {
	return dara.Validate(s)
}

type CreateTeamRequestBodyUsers struct {
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

func (s CreateTeamRequestBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamRequestBodyUsers) GoString() string {
	return s.String()
}

func (s *CreateTeamRequestBodyUsers) GetTeamRole() *string {
	return s.TeamRole
}

func (s *CreateTeamRequestBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *CreateTeamRequestBodyUsers) SetTeamRole(v string) *CreateTeamRequestBodyUsers {
	s.TeamRole = &v
	return s
}

func (s *CreateTeamRequestBodyUsers) SetUserId(v string) *CreateTeamRequestBodyUsers {
	s.UserId = &v
	return s
}

func (s *CreateTeamRequestBodyUsers) Validate() error {
	return dara.Validate(s)
}
