// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentUserSayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateBeebotIntentUserSayRequest
	GetInstanceId() *string
	SetScriptId(v string) *CreateBeebotIntentUserSayRequest
	GetScriptId() *string
	SetUserSayDefinition(v *CreateBeebotIntentUserSayRequestUserSayDefinition) *CreateBeebotIntentUserSayRequest
	GetUserSayDefinition() *CreateBeebotIntentUserSayRequestUserSayDefinition
}

type CreateBeebotIntentUserSayRequest struct {
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
	UserSayDefinition *CreateBeebotIntentUserSayRequestUserSayDefinition `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty" type:"Struct"`
}

func (s CreateBeebotIntentUserSayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentUserSayRequest) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentUserSayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBeebotIntentUserSayRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateBeebotIntentUserSayRequest) GetUserSayDefinition() *CreateBeebotIntentUserSayRequestUserSayDefinition {
	return s.UserSayDefinition
}

func (s *CreateBeebotIntentUserSayRequest) SetInstanceId(v string) *CreateBeebotIntentUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBeebotIntentUserSayRequest) SetScriptId(v string) *CreateBeebotIntentUserSayRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateBeebotIntentUserSayRequest) SetUserSayDefinition(v *CreateBeebotIntentUserSayRequestUserSayDefinition) *CreateBeebotIntentUserSayRequest {
	s.UserSayDefinition = v
	return s
}

func (s *CreateBeebotIntentUserSayRequest) Validate() error {
	if s.UserSayDefinition != nil {
		if err := s.UserSayDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBeebotIntentUserSayRequestUserSayDefinition struct {
	// User utterance
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

func (s CreateBeebotIntentUserSayRequestUserSayDefinition) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentUserSayRequestUserSayDefinition) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentUserSayRequestUserSayDefinition) GetContent() *string {
	return s.Content
}

func (s *CreateBeebotIntentUserSayRequestUserSayDefinition) GetIntentId() *int64 {
	return s.IntentId
}

func (s *CreateBeebotIntentUserSayRequestUserSayDefinition) SetContent(v string) *CreateBeebotIntentUserSayRequestUserSayDefinition {
	s.Content = &v
	return s
}

func (s *CreateBeebotIntentUserSayRequestUserSayDefinition) SetIntentId(v int64) *CreateBeebotIntentUserSayRequestUserSayDefinition {
	s.IntentId = &v
	return s
}

func (s *CreateBeebotIntentUserSayRequestUserSayDefinition) Validate() error {
	return dara.Validate(s)
}
