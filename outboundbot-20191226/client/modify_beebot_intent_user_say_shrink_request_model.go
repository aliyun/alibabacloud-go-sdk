// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentUserSayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyBeebotIntentUserSayShrinkRequest
	GetInstanceId() *string
	SetScriptId(v string) *ModifyBeebotIntentUserSayShrinkRequest
	GetScriptId() *string
	SetUserSayDefinitionShrink(v string) *ModifyBeebotIntentUserSayShrinkRequest
	GetUserSayDefinitionShrink() *string
	SetUserSayId(v string) *ModifyBeebotIntentUserSayShrinkRequest
	GetUserSayId() *string
}

type ModifyBeebotIntentUserSayShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
	UserSayDefinitionShrink *string `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 17448458
	UserSayId *string `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s ModifyBeebotIntentUserSayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentUserSayShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentUserSayShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBeebotIntentUserSayShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyBeebotIntentUserSayShrinkRequest) GetUserSayDefinitionShrink() *string {
	return s.UserSayDefinitionShrink
}

func (s *ModifyBeebotIntentUserSayShrinkRequest) GetUserSayId() *string {
	return s.UserSayId
}

func (s *ModifyBeebotIntentUserSayShrinkRequest) SetInstanceId(v string) *ModifyBeebotIntentUserSayShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayShrinkRequest) SetScriptId(v string) *ModifyBeebotIntentUserSayShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayShrinkRequest) SetUserSayDefinitionShrink(v string) *ModifyBeebotIntentUserSayShrinkRequest {
	s.UserSayDefinitionShrink = &v
	return s
}

func (s *ModifyBeebotIntentUserSayShrinkRequest) SetUserSayId(v string) *ModifyBeebotIntentUserSayShrinkRequest {
	s.UserSayId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
