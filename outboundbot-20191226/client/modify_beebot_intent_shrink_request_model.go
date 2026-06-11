// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyBeebotIntentShrinkRequest
	GetInstanceId() *string
	SetIntentDefinitionShrink(v string) *ModifyBeebotIntentShrinkRequest
	GetIntentDefinitionShrink() *string
	SetIntentId(v int64) *ModifyBeebotIntentShrinkRequest
	GetIntentId() *int64
	SetScriptId(v string) *ModifyBeebotIntentShrinkRequest
	GetScriptId() *string
}

type ModifyBeebotIntentShrinkRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Intent definition
	//
	// This parameter is required.
	IntentDefinitionShrink *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
	// Intent ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ModifyBeebotIntentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBeebotIntentShrinkRequest) GetIntentDefinitionShrink() *string {
	return s.IntentDefinitionShrink
}

func (s *ModifyBeebotIntentShrinkRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ModifyBeebotIntentShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyBeebotIntentShrinkRequest) SetInstanceId(v string) *ModifyBeebotIntentShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBeebotIntentShrinkRequest) SetIntentDefinitionShrink(v string) *ModifyBeebotIntentShrinkRequest {
	s.IntentDefinitionShrink = &v
	return s
}

func (s *ModifyBeebotIntentShrinkRequest) SetIntentId(v int64) *ModifyBeebotIntentShrinkRequest {
	s.IntentId = &v
	return s
}

func (s *ModifyBeebotIntentShrinkRequest) SetScriptId(v string) *ModifyBeebotIntentShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyBeebotIntentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
