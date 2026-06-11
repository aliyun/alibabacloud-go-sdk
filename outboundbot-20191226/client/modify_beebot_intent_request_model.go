// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyBeebotIntentRequest
	GetInstanceId() *string
	SetIntentDefinition(v *ModifyBeebotIntentRequestIntentDefinition) *ModifyBeebotIntentRequest
	GetIntentDefinition() *ModifyBeebotIntentRequestIntentDefinition
	SetIntentId(v int64) *ModifyBeebotIntentRequest
	GetIntentId() *int64
	SetScriptId(v string) *ModifyBeebotIntentRequest
	GetScriptId() *string
}

type ModifyBeebotIntentRequest struct {
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
	IntentDefinition *ModifyBeebotIntentRequestIntentDefinition `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty" type:"Struct"`
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

func (s ModifyBeebotIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentRequest) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBeebotIntentRequest) GetIntentDefinition() *ModifyBeebotIntentRequestIntentDefinition {
	return s.IntentDefinition
}

func (s *ModifyBeebotIntentRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ModifyBeebotIntentRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyBeebotIntentRequest) SetInstanceId(v string) *ModifyBeebotIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBeebotIntentRequest) SetIntentDefinition(v *ModifyBeebotIntentRequestIntentDefinition) *ModifyBeebotIntentRequest {
	s.IntentDefinition = v
	return s
}

func (s *ModifyBeebotIntentRequest) SetIntentId(v int64) *ModifyBeebotIntentRequest {
	s.IntentId = &v
	return s
}

func (s *ModifyBeebotIntentRequest) SetScriptId(v string) *ModifyBeebotIntentRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyBeebotIntentRequest) Validate() error {
	if s.IntentDefinition != nil {
		if err := s.IntentDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyBeebotIntentRequestIntentDefinition struct {
	// Intent alias
	//
	// example:
	//
	// 嗯明白了
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Intent name
	//
	// This parameter is required.
	//
	// example:
	//
	// 知道了么
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
}

func (s ModifyBeebotIntentRequestIntentDefinition) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentRequestIntentDefinition) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentRequestIntentDefinition) GetAliasName() *string {
	return s.AliasName
}

func (s *ModifyBeebotIntentRequestIntentDefinition) GetIntentName() *string {
	return s.IntentName
}

func (s *ModifyBeebotIntentRequestIntentDefinition) SetAliasName(v string) *ModifyBeebotIntentRequestIntentDefinition {
	s.AliasName = &v
	return s
}

func (s *ModifyBeebotIntentRequestIntentDefinition) SetIntentName(v string) *ModifyBeebotIntentRequestIntentDefinition {
	s.IntentName = &v
	return s
}

func (s *ModifyBeebotIntentRequestIntentDefinition) Validate() error {
	return dara.Validate(s)
}
