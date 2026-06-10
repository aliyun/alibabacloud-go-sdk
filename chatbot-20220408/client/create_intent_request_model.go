// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateIntentRequest
	GetAgentKey() *string
	SetInstanceId(v string) *CreateIntentRequest
	GetInstanceId() *string
	SetIntentDefinition(v *CreateIntentRequestIntentDefinition) *CreateIntentRequest
	GetIntentDefinition() *CreateIntentRequestIntentDefinition
}

type CreateIntentRequest struct {
	// The key of the business space. If you do not specify this parameter, the default business space is used. You can get the key from the business management page of your main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The intent definition.
	IntentDefinition *CreateIntentRequestIntentDefinition `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty" type:"Struct"`
}

func (s CreateIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIntentRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateIntentRequest) GetIntentDefinition() *CreateIntentRequestIntentDefinition {
	return s.IntentDefinition
}

func (s *CreateIntentRequest) SetAgentKey(v string) *CreateIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateIntentRequest) SetInstanceId(v string) *CreateIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIntentRequest) SetIntentDefinition(v *CreateIntentRequestIntentDefinition) *CreateIntentRequest {
	s.IntentDefinition = v
	return s
}

func (s *CreateIntentRequest) Validate() error {
	if s.IntentDefinition != nil {
		if err := s.IntentDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIntentRequestIntentDefinition struct {
	// The intent alias.
	//
	// example:
	//
	// 查天气意图
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The intent name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 查天气意图
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// An array of objects, each defining a slot for the intent.
	SlotInfos []*CreateIntentRequestIntentDefinitionSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s CreateIntentRequestIntentDefinition) String() string {
	return dara.Prettify(s)
}

func (s CreateIntentRequestIntentDefinition) GoString() string {
	return s.String()
}

func (s *CreateIntentRequestIntentDefinition) GetAliasName() *string {
	return s.AliasName
}

func (s *CreateIntentRequestIntentDefinition) GetIntentName() *string {
	return s.IntentName
}

func (s *CreateIntentRequestIntentDefinition) GetSlotInfos() []*CreateIntentRequestIntentDefinitionSlotInfos {
	return s.SlotInfos
}

func (s *CreateIntentRequestIntentDefinition) SetAliasName(v string) *CreateIntentRequestIntentDefinition {
	s.AliasName = &v
	return s
}

func (s *CreateIntentRequestIntentDefinition) SetIntentName(v string) *CreateIntentRequestIntentDefinition {
	s.IntentName = &v
	return s
}

func (s *CreateIntentRequestIntentDefinition) SetSlotInfos(v []*CreateIntentRequestIntentDefinitionSlotInfos) *CreateIntentRequestIntentDefinition {
	s.SlotInfos = v
	return s
}

func (s *CreateIntentRequestIntentDefinition) Validate() error {
	if s.SlotInfos != nil {
		for _, item := range s.SlotInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateIntentRequestIntentDefinitionSlotInfos struct {
	// Specifies whether the slot can have multiple values.
	//
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// Specifies whether to encrypt the slot value.
	//
	// example:
	//
	// false
	Encrypt *bool `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// Specifies whether the slot is interactive.
	//
	// example:
	//
	// false
	Interactive *bool `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	// The slot name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 天气
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The slot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fg452dfg3df23
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// The name of the associated entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 天气
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIntentRequestIntentDefinitionSlotInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateIntentRequestIntentDefinitionSlotInfos) GoString() string {
	return s.String()
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) GetArray() *bool {
	return s.Array
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) GetEncrypt() *bool {
	return s.Encrypt
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) GetInteractive() *bool {
	return s.Interactive
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) GetName() *string {
	return s.Name
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) GetSlotId() *string {
	return s.SlotId
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) GetValue() *string {
	return s.Value
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetArray(v bool) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Array = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetEncrypt(v bool) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Encrypt = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetInteractive(v bool) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Interactive = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetName(v string) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Name = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetSlotId(v string) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.SlotId = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetValue(v string) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Value = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) Validate() error {
	return dara.Validate(s)
}
