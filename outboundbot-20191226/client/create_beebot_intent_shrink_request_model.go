// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateBeebotIntentShrinkRequest
	GetInstanceId() *string
	SetIntentDefinitionShrink(v string) *CreateBeebotIntentShrinkRequest
	GetIntentDefinitionShrink() *string
	SetScriptId(v string) *CreateBeebotIntentShrinkRequest
	GetScriptId() *string
}

type CreateBeebotIntentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	IntentDefinitionShrink *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s CreateBeebotIntentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBeebotIntentShrinkRequest) GetIntentDefinitionShrink() *string {
	return s.IntentDefinitionShrink
}

func (s *CreateBeebotIntentShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateBeebotIntentShrinkRequest) SetInstanceId(v string) *CreateBeebotIntentShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBeebotIntentShrinkRequest) SetIntentDefinitionShrink(v string) *CreateBeebotIntentShrinkRequest {
	s.IntentDefinitionShrink = &v
	return s
}

func (s *CreateBeebotIntentShrinkRequest) SetScriptId(v string) *CreateBeebotIntentShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateBeebotIntentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
