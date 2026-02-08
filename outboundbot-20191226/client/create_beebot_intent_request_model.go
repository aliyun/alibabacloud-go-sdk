// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateBeebotIntentRequest
	GetInstanceId() *string
	SetIntentDefinition(v *CreateBeebotIntentRequestIntentDefinition) *CreateBeebotIntentRequest
	GetIntentDefinition() *CreateBeebotIntentRequestIntentDefinition
	SetScriptId(v string) *CreateBeebotIntentRequest
	GetScriptId() *string
}

type CreateBeebotIntentRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Intent definition description
	//
	// This parameter is required.
	IntentDefinition *CreateBeebotIntentRequestIntentDefinition `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty" type:"Struct"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s CreateBeebotIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentRequest) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBeebotIntentRequest) GetIntentDefinition() *CreateBeebotIntentRequestIntentDefinition {
	return s.IntentDefinition
}

func (s *CreateBeebotIntentRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateBeebotIntentRequest) SetInstanceId(v string) *CreateBeebotIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBeebotIntentRequest) SetIntentDefinition(v *CreateBeebotIntentRequestIntentDefinition) *CreateBeebotIntentRequest {
	s.IntentDefinition = v
	return s
}

func (s *CreateBeebotIntentRequest) SetScriptId(v string) *CreateBeebotIntentRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateBeebotIntentRequest) Validate() error {
	if s.IntentDefinition != nil {
		if err := s.IntentDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBeebotIntentRequestIntentDefinition struct {
	// Intent alias
	//
	// example:
	//
	// 嗯明白了
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Intent name
	//
	// > This serves as the intent code and is a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 知道了
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
}

func (s CreateBeebotIntentRequestIntentDefinition) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentRequestIntentDefinition) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentRequestIntentDefinition) GetAliasName() *string {
	return s.AliasName
}

func (s *CreateBeebotIntentRequestIntentDefinition) GetIntentName() *string {
	return s.IntentName
}

func (s *CreateBeebotIntentRequestIntentDefinition) SetAliasName(v string) *CreateBeebotIntentRequestIntentDefinition {
	s.AliasName = &v
	return s
}

func (s *CreateBeebotIntentRequestIntentDefinition) SetIntentName(v string) *CreateBeebotIntentRequestIntentDefinition {
	s.IntentName = &v
	return s
}

func (s *CreateBeebotIntentRequestIntentDefinition) Validate() error {
	return dara.Validate(s)
}
