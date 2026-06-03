// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentLgfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyBeebotIntentLgfRequest
	GetInstanceId() *string
	SetLgfDefinition(v *ModifyBeebotIntentLgfRequestLgfDefinition) *ModifyBeebotIntentLgfRequest
	GetLgfDefinition() *ModifyBeebotIntentLgfRequestLgfDefinition
	SetLgfId(v int64) *ModifyBeebotIntentLgfRequest
	GetLgfId() *int64
	SetScriptId(v string) *ModifyBeebotIntentLgfRequest
	GetScriptId() *string
}

type ModifyBeebotIntentLgfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	LgfDefinition *ModifyBeebotIntentLgfRequestLgfDefinition `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty" type:"Struct"`
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

func (s ModifyBeebotIntentLgfRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentLgfRequest) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentLgfRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBeebotIntentLgfRequest) GetLgfDefinition() *ModifyBeebotIntentLgfRequestLgfDefinition {
	return s.LgfDefinition
}

func (s *ModifyBeebotIntentLgfRequest) GetLgfId() *int64 {
	return s.LgfId
}

func (s *ModifyBeebotIntentLgfRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyBeebotIntentLgfRequest) SetInstanceId(v string) *ModifyBeebotIntentLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBeebotIntentLgfRequest) SetLgfDefinition(v *ModifyBeebotIntentLgfRequestLgfDefinition) *ModifyBeebotIntentLgfRequest {
	s.LgfDefinition = v
	return s
}

func (s *ModifyBeebotIntentLgfRequest) SetLgfId(v int64) *ModifyBeebotIntentLgfRequest {
	s.LgfId = &v
	return s
}

func (s *ModifyBeebotIntentLgfRequest) SetScriptId(v string) *ModifyBeebotIntentLgfRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyBeebotIntentLgfRequest) Validate() error {
	if s.LgfDefinition != nil {
		if err := s.LgfDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyBeebotIntentLgfRequestLgfDefinition struct {
	// This parameter is required.
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	RuleText *string `json:"RuleText,omitempty" xml:"RuleText,omitempty"`
}

func (s ModifyBeebotIntentLgfRequestLgfDefinition) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentLgfRequestLgfDefinition) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentLgfRequestLgfDefinition) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ModifyBeebotIntentLgfRequestLgfDefinition) GetRuleText() *string {
	return s.RuleText
}

func (s *ModifyBeebotIntentLgfRequestLgfDefinition) SetIntentId(v int64) *ModifyBeebotIntentLgfRequestLgfDefinition {
	s.IntentId = &v
	return s
}

func (s *ModifyBeebotIntentLgfRequestLgfDefinition) SetRuleText(v string) *ModifyBeebotIntentLgfRequestLgfDefinition {
	s.RuleText = &v
	return s
}

func (s *ModifyBeebotIntentLgfRequestLgfDefinition) Validate() error {
	return dara.Validate(s)
}
