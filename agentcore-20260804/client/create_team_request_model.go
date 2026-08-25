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
	Body *CreateTeamRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// example:
	//
	// 暂不支持
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
	Agents []*CreateTeamRequestBodyAgents `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
	// example:
	//
	// 负责智能客服业务的团队
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// team-01
	Name  *string                       `json:"name,omitempty" xml:"name,omitempty"`
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
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
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
	// example:
	//
	// ADMIN
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
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
