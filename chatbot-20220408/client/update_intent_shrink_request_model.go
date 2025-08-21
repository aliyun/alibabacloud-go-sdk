// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateIntentShrinkRequest
	GetAgentKey() *string
	SetInstanceId(v string) *UpdateIntentShrinkRequest
	GetInstanceId() *string
	SetIntentDefinitionShrink(v string) *UpdateIntentShrinkRequest
	GetIntentDefinitionShrink() *string
	SetIntentId(v int64) *UpdateIntentShrinkRequest
	GetIntentId() *int64
}

type UpdateIntentShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentDefinitionShrink *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 234234234534
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s UpdateIntentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntentShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateIntentShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateIntentShrinkRequest) GetIntentDefinitionShrink() *string {
	return s.IntentDefinitionShrink
}

func (s *UpdateIntentShrinkRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *UpdateIntentShrinkRequest) SetAgentKey(v string) *UpdateIntentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateIntentShrinkRequest) SetInstanceId(v string) *UpdateIntentShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIntentShrinkRequest) SetIntentDefinitionShrink(v string) *UpdateIntentShrinkRequest {
	s.IntentDefinitionShrink = &v
	return s
}

func (s *UpdateIntentShrinkRequest) SetIntentId(v int64) *UpdateIntentShrinkRequest {
	s.IntentId = &v
	return s
}

func (s *UpdateIntentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
