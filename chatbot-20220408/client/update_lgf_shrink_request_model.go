// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLgfShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateLgfShrinkRequest
	GetAgentKey() *string
	SetInstanceId(v string) *UpdateLgfShrinkRequest
	GetInstanceId() *string
	SetLgfDefinitionShrink(v string) *UpdateLgfShrinkRequest
	GetLgfDefinitionShrink() *string
	SetLgfId(v int64) *UpdateLgfShrinkRequest
	GetLgfId() *int64
}

type UpdateLgfShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	LgfDefinitionShrink *string `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty"`
	// LGF ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 12121
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
}

func (s UpdateLgfShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLgfShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLgfShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateLgfShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLgfShrinkRequest) GetLgfDefinitionShrink() *string {
	return s.LgfDefinitionShrink
}

func (s *UpdateLgfShrinkRequest) GetLgfId() *int64 {
	return s.LgfId
}

func (s *UpdateLgfShrinkRequest) SetAgentKey(v string) *UpdateLgfShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateLgfShrinkRequest) SetInstanceId(v string) *UpdateLgfShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLgfShrinkRequest) SetLgfDefinitionShrink(v string) *UpdateLgfShrinkRequest {
	s.LgfDefinitionShrink = &v
	return s
}

func (s *UpdateLgfShrinkRequest) SetLgfId(v int64) *UpdateLgfShrinkRequest {
	s.LgfId = &v
	return s
}

func (s *UpdateLgfShrinkRequest) Validate() error {
	return dara.Validate(s)
}
