// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateIntentShrinkRequest
	GetAgentKey() *string
	SetInstanceId(v string) *CreateIntentShrinkRequest
	GetInstanceId() *string
	SetIntentDefinitionShrink(v string) *CreateIntentShrinkRequest
	GetIntentDefinitionShrink() *string
}

type CreateIntentShrinkRequest struct {
	// The key of the business space. If you do not specify this parameter, the default business space is used. You can get the key from the business management page of your main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The intent definition.
	IntentDefinitionShrink *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
}

func (s CreateIntentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIntentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateIntentShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateIntentShrinkRequest) GetIntentDefinitionShrink() *string {
	return s.IntentDefinitionShrink
}

func (s *CreateIntentShrinkRequest) SetAgentKey(v string) *CreateIntentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateIntentShrinkRequest) SetInstanceId(v string) *CreateIntentShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIntentShrinkRequest) SetIntentDefinitionShrink(v string) *CreateIntentShrinkRequest {
	s.IntentDefinitionShrink = &v
	return s
}

func (s *CreateIntentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
