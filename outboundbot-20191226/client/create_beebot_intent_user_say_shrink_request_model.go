// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentUserSayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateBeebotIntentUserSayShrinkRequest
	GetInstanceId() *string
	SetScriptId(v string) *CreateBeebotIntentUserSayShrinkRequest
	GetScriptId() *string
	SetUserSayDefinitionShrink(v string) *CreateBeebotIntentUserSayShrinkRequest
	GetUserSayDefinitionShrink() *string
}

type CreateBeebotIntentUserSayShrinkRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// User utterance definition description
	//
	// This parameter is required.
	UserSayDefinitionShrink *string `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty"`
}

func (s CreateBeebotIntentUserSayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentUserSayShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentUserSayShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBeebotIntentUserSayShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateBeebotIntentUserSayShrinkRequest) GetUserSayDefinitionShrink() *string {
	return s.UserSayDefinitionShrink
}

func (s *CreateBeebotIntentUserSayShrinkRequest) SetInstanceId(v string) *CreateBeebotIntentUserSayShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBeebotIntentUserSayShrinkRequest) SetScriptId(v string) *CreateBeebotIntentUserSayShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateBeebotIntentUserSayShrinkRequest) SetUserSayDefinitionShrink(v string) *CreateBeebotIntentUserSayShrinkRequest {
	s.UserSayDefinitionShrink = &v
	return s
}

func (s *CreateBeebotIntentUserSayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
