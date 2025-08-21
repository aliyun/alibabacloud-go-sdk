// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateUserSayShrinkRequest
	GetAgentKey() *string
	SetInstanceId(v string) *CreateUserSayShrinkRequest
	GetInstanceId() *string
	SetUserSayDefinitionShrink(v string) *CreateUserSayShrinkRequest
	GetUserSayDefinitionShrink() *string
}

type CreateUserSayShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserSayDefinitionShrink *string `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty"`
}

func (s CreateUserSayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSayShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserSayShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateUserSayShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserSayShrinkRequest) GetUserSayDefinitionShrink() *string {
	return s.UserSayDefinitionShrink
}

func (s *CreateUserSayShrinkRequest) SetAgentKey(v string) *CreateUserSayShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateUserSayShrinkRequest) SetInstanceId(v string) *CreateUserSayShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserSayShrinkRequest) SetUserSayDefinitionShrink(v string) *CreateUserSayShrinkRequest {
	s.UserSayDefinitionShrink = &v
	return s
}

func (s *CreateUserSayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
