// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLgfShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateLgfShrinkRequest
	GetAgentKey() *string
	SetInstanceId(v string) *CreateLgfShrinkRequest
	GetInstanceId() *string
	SetLgfDefinitionShrink(v string) *CreateLgfShrinkRequest
	GetLgfDefinitionShrink() *string
}

type CreateLgfShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LgfDefinitionShrink *string `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty"`
}

func (s CreateLgfShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLgfShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLgfShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateLgfShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLgfShrinkRequest) GetLgfDefinitionShrink() *string {
	return s.LgfDefinitionShrink
}

func (s *CreateLgfShrinkRequest) SetAgentKey(v string) *CreateLgfShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateLgfShrinkRequest) SetInstanceId(v string) *CreateLgfShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLgfShrinkRequest) SetLgfDefinitionShrink(v string) *CreateLgfShrinkRequest {
	s.LgfDefinitionShrink = &v
	return s
}

func (s *CreateLgfShrinkRequest) Validate() error {
	return dara.Validate(s)
}
