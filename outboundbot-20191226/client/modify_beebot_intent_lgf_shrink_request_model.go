// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentLgfShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyBeebotIntentLgfShrinkRequest
	GetInstanceId() *string
	SetLgfDefinitionShrink(v string) *ModifyBeebotIntentLgfShrinkRequest
	GetLgfDefinitionShrink() *string
	SetLgfId(v int64) *ModifyBeebotIntentLgfShrinkRequest
	GetLgfId() *int64
	SetScriptId(v string) *ModifyBeebotIntentLgfShrinkRequest
	GetScriptId() *string
}

type ModifyBeebotIntentLgfShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	LgfDefinitionShrink *string `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5666117
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ModifyBeebotIntentLgfShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentLgfShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentLgfShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBeebotIntentLgfShrinkRequest) GetLgfDefinitionShrink() *string {
	return s.LgfDefinitionShrink
}

func (s *ModifyBeebotIntentLgfShrinkRequest) GetLgfId() *int64 {
	return s.LgfId
}

func (s *ModifyBeebotIntentLgfShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyBeebotIntentLgfShrinkRequest) SetInstanceId(v string) *ModifyBeebotIntentLgfShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBeebotIntentLgfShrinkRequest) SetLgfDefinitionShrink(v string) *ModifyBeebotIntentLgfShrinkRequest {
	s.LgfDefinitionShrink = &v
	return s
}

func (s *ModifyBeebotIntentLgfShrinkRequest) SetLgfId(v int64) *ModifyBeebotIntentLgfShrinkRequest {
	s.LgfId = &v
	return s
}

func (s *ModifyBeebotIntentLgfShrinkRequest) SetScriptId(v string) *ModifyBeebotIntentLgfShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyBeebotIntentLgfShrinkRequest) Validate() error {
	return dara.Validate(s)
}
