// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgent(v *CreatePolarClawAgentResponseBodyAgent) *CreatePolarClawAgentResponseBody
	GetAgent() *CreatePolarClawAgentResponseBodyAgent
	SetAgentId(v string) *CreatePolarClawAgentResponseBody
	GetAgentId() *string
	SetApplicationId(v string) *CreatePolarClawAgentResponseBody
	GetApplicationId() *string
	SetCode(v int32) *CreatePolarClawAgentResponseBody
	GetCode() *int32
	SetMessage(v string) *CreatePolarClawAgentResponseBody
	GetMessage() *string
	SetName(v string) *CreatePolarClawAgentResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePolarClawAgentResponseBody
	GetRequestId() *string
	SetWorkspace(v string) *CreatePolarClawAgentResponseBody
	GetWorkspace() *string
}

type CreatePolarClawAgentResponseBody struct {
	// Details of the new agent.
	Agent *CreatePolarClawAgentResponseBodyAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// The agent ID.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The display name of the agent.
	//
	// example:
	//
	// work
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The workspace path.
	//
	// example:
	//
	// /home/node/.openclaw/workspace-work
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreatePolarClawAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarClawAgentResponseBody) GetAgent() *CreatePolarClawAgentResponseBodyAgent {
	return s.Agent
}

func (s *CreatePolarClawAgentResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *CreatePolarClawAgentResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePolarClawAgentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePolarClawAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePolarClawAgentResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePolarClawAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolarClawAgentResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreatePolarClawAgentResponseBody) SetAgent(v *CreatePolarClawAgentResponseBodyAgent) *CreatePolarClawAgentResponseBody {
	s.Agent = v
	return s
}

func (s *CreatePolarClawAgentResponseBody) SetAgentId(v string) *CreatePolarClawAgentResponseBody {
	s.AgentId = &v
	return s
}

func (s *CreatePolarClawAgentResponseBody) SetApplicationId(v string) *CreatePolarClawAgentResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreatePolarClawAgentResponseBody) SetCode(v int32) *CreatePolarClawAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolarClawAgentResponseBody) SetMessage(v string) *CreatePolarClawAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolarClawAgentResponseBody) SetName(v string) *CreatePolarClawAgentResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePolarClawAgentResponseBody) SetRequestId(v string) *CreatePolarClawAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolarClawAgentResponseBody) SetWorkspace(v string) *CreatePolarClawAgentResponseBody {
	s.Workspace = &v
	return s
}

func (s *CreatePolarClawAgentResponseBody) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolarClawAgentResponseBodyAgent struct {
	// The agent ID.
	//
	// example:
	//
	// work
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The agent\\"s identity.
	Identity *CreatePolarClawAgentResponseBodyAgentIdentity `json:"Identity,omitempty" xml:"Identity,omitempty" type:"Struct"`
	// The display name of the agent.
	//
	// example:
	//
	// work
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The workspace path.
	//
	// example:
	//
	// /home/node/.openclaw/workspace-work
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreatePolarClawAgentResponseBodyAgent) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawAgentResponseBodyAgent) GoString() string {
	return s.String()
}

func (s *CreatePolarClawAgentResponseBodyAgent) GetId() *string {
	return s.Id
}

func (s *CreatePolarClawAgentResponseBodyAgent) GetIdentity() *CreatePolarClawAgentResponseBodyAgentIdentity {
	return s.Identity
}

func (s *CreatePolarClawAgentResponseBodyAgent) GetName() *string {
	return s.Name
}

func (s *CreatePolarClawAgentResponseBodyAgent) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreatePolarClawAgentResponseBodyAgent) SetId(v string) *CreatePolarClawAgentResponseBodyAgent {
	s.Id = &v
	return s
}

func (s *CreatePolarClawAgentResponseBodyAgent) SetIdentity(v *CreatePolarClawAgentResponseBodyAgentIdentity) *CreatePolarClawAgentResponseBodyAgent {
	s.Identity = v
	return s
}

func (s *CreatePolarClawAgentResponseBodyAgent) SetName(v string) *CreatePolarClawAgentResponseBodyAgent {
	s.Name = &v
	return s
}

func (s *CreatePolarClawAgentResponseBodyAgent) SetWorkspace(v string) *CreatePolarClawAgentResponseBodyAgent {
	s.Workspace = &v
	return s
}

func (s *CreatePolarClawAgentResponseBodyAgent) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolarClawAgentResponseBodyAgentIdentity struct {
	// The avatar\\"s path or content.
	//
	// example:
	//
	// test
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// The avatar URL.
	//
	// example:
	//
	// test
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// The emoji identifier.
	//
	// example:
	//
	// U+1F99E
	Emoji *string `json:"Emoji,omitempty" xml:"Emoji,omitempty"`
	// The identity name.
	//
	// example:
	//
	// work
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The theme.
	//
	// example:
	//
	// work
	Theme *string `json:"Theme,omitempty" xml:"Theme,omitempty"`
}

func (s CreatePolarClawAgentResponseBodyAgentIdentity) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawAgentResponseBodyAgentIdentity) GoString() string {
	return s.String()
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) GetAvatar() *string {
	return s.Avatar
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) GetEmoji() *string {
	return s.Emoji
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) GetName() *string {
	return s.Name
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) GetTheme() *string {
	return s.Theme
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) SetAvatar(v string) *CreatePolarClawAgentResponseBodyAgentIdentity {
	s.Avatar = &v
	return s
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) SetAvatarUrl(v string) *CreatePolarClawAgentResponseBodyAgentIdentity {
	s.AvatarUrl = &v
	return s
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) SetEmoji(v string) *CreatePolarClawAgentResponseBodyAgentIdentity {
	s.Emoji = &v
	return s
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) SetName(v string) *CreatePolarClawAgentResponseBodyAgentIdentity {
	s.Name = &v
	return s
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) SetTheme(v string) *CreatePolarClawAgentResponseBodyAgentIdentity {
	s.Theme = &v
	return s
}

func (s *CreatePolarClawAgentResponseBodyAgentIdentity) Validate() error {
	return dara.Validate(s)
}
