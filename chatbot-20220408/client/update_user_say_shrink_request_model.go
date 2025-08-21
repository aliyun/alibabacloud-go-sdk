// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserSayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateUserSayShrinkRequest
	GetAgentKey() *string
	SetInstanceId(v string) *UpdateUserSayShrinkRequest
	GetInstanceId() *string
	SetUserSayDefinitionShrink(v string) *UpdateUserSayShrinkRequest
	GetUserSayDefinitionShrink() *string
	SetUserSayId(v int64) *UpdateUserSayShrinkRequest
	GetUserSayId() *int64
}

type UpdateUserSayShrinkRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 34512323
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s UpdateUserSayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserSayShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserSayShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateUserSayShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateUserSayShrinkRequest) GetUserSayDefinitionShrink() *string {
	return s.UserSayDefinitionShrink
}

func (s *UpdateUserSayShrinkRequest) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *UpdateUserSayShrinkRequest) SetAgentKey(v string) *UpdateUserSayShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateUserSayShrinkRequest) SetInstanceId(v string) *UpdateUserSayShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserSayShrinkRequest) SetUserSayDefinitionShrink(v string) *UpdateUserSayShrinkRequest {
	s.UserSayDefinitionShrink = &v
	return s
}

func (s *UpdateUserSayShrinkRequest) SetUserSayId(v int64) *UpdateUserSayShrinkRequest {
	s.UserSayId = &v
	return s
}

func (s *UpdateUserSayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
