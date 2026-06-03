// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentLgfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateBeebotIntentLgfRequest
	GetInstanceId() *string
	SetLgfDefinition(v *CreateBeebotIntentLgfRequestLgfDefinition) *CreateBeebotIntentLgfRequest
	GetLgfDefinition() *CreateBeebotIntentLgfRequestLgfDefinition
	SetScriptId(v string) *CreateBeebotIntentLgfRequest
	GetScriptId() *string
}

type CreateBeebotIntentLgfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	LgfDefinition *CreateBeebotIntentLgfRequestLgfDefinition `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s CreateBeebotIntentLgfRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentLgfRequest) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentLgfRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBeebotIntentLgfRequest) GetLgfDefinition() *CreateBeebotIntentLgfRequestLgfDefinition {
	return s.LgfDefinition
}

func (s *CreateBeebotIntentLgfRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateBeebotIntentLgfRequest) SetInstanceId(v string) *CreateBeebotIntentLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBeebotIntentLgfRequest) SetLgfDefinition(v *CreateBeebotIntentLgfRequestLgfDefinition) *CreateBeebotIntentLgfRequest {
	s.LgfDefinition = v
	return s
}

func (s *CreateBeebotIntentLgfRequest) SetScriptId(v string) *CreateBeebotIntentLgfRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateBeebotIntentLgfRequest) Validate() error {
	if s.LgfDefinition != nil {
		if err := s.LgfDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBeebotIntentLgfRequestLgfDefinition struct {
	// This parameter is required.
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	RuleText *string `json:"RuleText,omitempty" xml:"RuleText,omitempty"`
}

func (s CreateBeebotIntentLgfRequestLgfDefinition) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentLgfRequestLgfDefinition) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentLgfRequestLgfDefinition) GetIntentId() *int64 {
	return s.IntentId
}

func (s *CreateBeebotIntentLgfRequestLgfDefinition) GetRuleText() *string {
	return s.RuleText
}

func (s *CreateBeebotIntentLgfRequestLgfDefinition) SetIntentId(v int64) *CreateBeebotIntentLgfRequestLgfDefinition {
	s.IntentId = &v
	return s
}

func (s *CreateBeebotIntentLgfRequestLgfDefinition) SetRuleText(v string) *CreateBeebotIntentLgfRequestLgfDefinition {
	s.RuleText = &v
	return s
}

func (s *CreateBeebotIntentLgfRequestLgfDefinition) Validate() error {
	return dara.Validate(s)
}
