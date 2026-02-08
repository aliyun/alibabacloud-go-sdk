// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentUserSayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyBeebotIntentUserSayRequest
	GetInstanceId() *string
	SetScriptId(v string) *ModifyBeebotIntentUserSayRequest
	GetScriptId() *string
	SetUserSayDefinition(v *ModifyBeebotIntentUserSayRequestUserSayDefinition) *ModifyBeebotIntentUserSayRequest
	GetUserSayDefinition() *ModifyBeebotIntentUserSayRequestUserSayDefinition
	SetUserSayId(v string) *ModifyBeebotIntentUserSayRequest
	GetUserSayId() *string
}

type ModifyBeebotIntentUserSayRequest struct {
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
	// User utterance definition
	//
	// This parameter is required.
	UserSayDefinition *ModifyBeebotIntentUserSayRequestUserSayDefinition `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty" type:"Struct"`
	// User utterance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 17448458
	UserSayId *string `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s ModifyBeebotIntentUserSayRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentUserSayRequest) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentUserSayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBeebotIntentUserSayRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyBeebotIntentUserSayRequest) GetUserSayDefinition() *ModifyBeebotIntentUserSayRequestUserSayDefinition {
	return s.UserSayDefinition
}

func (s *ModifyBeebotIntentUserSayRequest) GetUserSayId() *string {
	return s.UserSayId
}

func (s *ModifyBeebotIntentUserSayRequest) SetInstanceId(v string) *ModifyBeebotIntentUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayRequest) SetScriptId(v string) *ModifyBeebotIntentUserSayRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayRequest) SetUserSayDefinition(v *ModifyBeebotIntentUserSayRequestUserSayDefinition) *ModifyBeebotIntentUserSayRequest {
	s.UserSayDefinition = v
	return s
}

func (s *ModifyBeebotIntentUserSayRequest) SetUserSayId(v string) *ModifyBeebotIntentUserSayRequest {
	s.UserSayId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayRequest) Validate() error {
	if s.UserSayDefinition != nil {
		if err := s.UserSayDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyBeebotIntentUserSayRequestUserSayDefinition struct {
	// User utterance content
	//
	// This parameter is required.
	//
	// example:
	//
	// 你知道xxxxx么？
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Intent ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s ModifyBeebotIntentUserSayRequestUserSayDefinition) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentUserSayRequestUserSayDefinition) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentUserSayRequestUserSayDefinition) GetContent() *string {
	return s.Content
}

func (s *ModifyBeebotIntentUserSayRequestUserSayDefinition) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ModifyBeebotIntentUserSayRequestUserSayDefinition) SetContent(v string) *ModifyBeebotIntentUserSayRequestUserSayDefinition {
	s.Content = &v
	return s
}

func (s *ModifyBeebotIntentUserSayRequestUserSayDefinition) SetIntentId(v int64) *ModifyBeebotIntentUserSayRequestUserSayDefinition {
	s.IntentId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayRequestUserSayDefinition) Validate() error {
	return dara.Validate(s)
}
