// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgent(v *UpdatePolarClawAgentResponseBodyAgent) *UpdatePolarClawAgentResponseBody
	GetAgent() *UpdatePolarClawAgentResponseBodyAgent
	SetAgentId(v string) *UpdatePolarClawAgentResponseBody
	GetAgentId() *string
	SetApplicationId(v string) *UpdatePolarClawAgentResponseBody
	GetApplicationId() *string
	SetCode(v int32) *UpdatePolarClawAgentResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdatePolarClawAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePolarClawAgentResponseBody
	GetRequestId() *string
}

type UpdatePolarClawAgentResponseBody struct {
	Agent *UpdatePolarClawAgentResponseBodyAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// Agent ID
	//
	// example:
	//
	// main
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePolarClawAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentResponseBody) GetAgent() *UpdatePolarClawAgentResponseBodyAgent {
	return s.Agent
}

func (s *UpdatePolarClawAgentResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawAgentResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawAgentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePolarClawAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolarClawAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolarClawAgentResponseBody) SetAgent(v *UpdatePolarClawAgentResponseBodyAgent) *UpdatePolarClawAgentResponseBody {
	s.Agent = v
	return s
}

func (s *UpdatePolarClawAgentResponseBody) SetAgentId(v string) *UpdatePolarClawAgentResponseBody {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBody) SetApplicationId(v string) *UpdatePolarClawAgentResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBody) SetCode(v int32) *UpdatePolarClawAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBody) SetMessage(v string) *UpdatePolarClawAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBody) SetRequestId(v string) *UpdatePolarClawAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBody) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolarClawAgentResponseBodyAgent struct {
	// Agent ID
	//
	// example:
	//
	// main
	Id       *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Identity *UpdatePolarClawAgentResponseBodyAgentIdentity `json:"Identity,omitempty" xml:"Identity,omitempty" type:"Struct"`
	// example:
	//
	// PolarClaw
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /home/node/.openclaw/workspace-work-v2
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdatePolarClawAgentResponseBodyAgent) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentResponseBodyAgent) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentResponseBodyAgent) GetId() *string {
	return s.Id
}

func (s *UpdatePolarClawAgentResponseBodyAgent) GetIdentity() *UpdatePolarClawAgentResponseBodyAgentIdentity {
	return s.Identity
}

func (s *UpdatePolarClawAgentResponseBodyAgent) GetName() *string {
	return s.Name
}

func (s *UpdatePolarClawAgentResponseBodyAgent) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdatePolarClawAgentResponseBodyAgent) SetId(v string) *UpdatePolarClawAgentResponseBodyAgent {
	s.Id = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBodyAgent) SetIdentity(v *UpdatePolarClawAgentResponseBodyAgentIdentity) *UpdatePolarClawAgentResponseBodyAgent {
	s.Identity = v
	return s
}

func (s *UpdatePolarClawAgentResponseBodyAgent) SetName(v string) *UpdatePolarClawAgentResponseBodyAgent {
	s.Name = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBodyAgent) SetWorkspace(v string) *UpdatePolarClawAgentResponseBodyAgent {
	s.Workspace = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBodyAgent) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolarClawAgentResponseBodyAgentIdentity struct {
	// example:
	//
	// test
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// test
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// example:
	//
	// U+1F99E
	Emoji *string `json:"Emoji,omitempty" xml:"Emoji,omitempty"`
	// example:
	//
	// PolarClaw
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// soul lobster
	Theme *string `json:"Theme,omitempty" xml:"Theme,omitempty"`
}

func (s UpdatePolarClawAgentResponseBodyAgentIdentity) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentResponseBodyAgentIdentity) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) GetEmoji() *string {
	return s.Emoji
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) GetName() *string {
	return s.Name
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) GetTheme() *string {
	return s.Theme
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) SetAvatar(v string) *UpdatePolarClawAgentResponseBodyAgentIdentity {
	s.Avatar = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) SetAvatarUrl(v string) *UpdatePolarClawAgentResponseBodyAgentIdentity {
	s.AvatarUrl = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) SetEmoji(v string) *UpdatePolarClawAgentResponseBodyAgentIdentity {
	s.Emoji = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) SetName(v string) *UpdatePolarClawAgentResponseBodyAgentIdentity {
	s.Name = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) SetTheme(v string) *UpdatePolarClawAgentResponseBodyAgentIdentity {
	s.Theme = &v
	return s
}

func (s *UpdatePolarClawAgentResponseBodyAgentIdentity) Validate() error {
	return dara.Validate(s)
}
