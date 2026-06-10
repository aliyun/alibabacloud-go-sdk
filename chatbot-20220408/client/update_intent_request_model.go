// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateIntentRequest
	GetAgentKey() *string
	SetInstanceId(v string) *UpdateIntentRequest
	GetInstanceId() *string
	SetIntentDefinition(v *UpdateIntentRequestIntentDefinition) *UpdateIntentRequest
	GetIntentDefinition() *UpdateIntentRequestIntentDefinition
	SetIntentId(v int64) *UpdateIntentRequest
	GetIntentId() *int64
}

type UpdateIntentRequest struct {
	// The key for the business space. If you omit this parameter, the system uses the default business space. You can find this key on the Business Management page of your primary account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The ID of the chatbot instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The intent definition object.
	IntentDefinition *UpdateIntentRequestIntentDefinition `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty" type:"Struct"`
	// The ID of the intent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 234234234534
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s UpdateIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntentRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateIntentRequest) GetIntentDefinition() *UpdateIntentRequestIntentDefinition {
	return s.IntentDefinition
}

func (s *UpdateIntentRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *UpdateIntentRequest) SetAgentKey(v string) *UpdateIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateIntentRequest) SetInstanceId(v string) *UpdateIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIntentRequest) SetIntentDefinition(v *UpdateIntentRequestIntentDefinition) *UpdateIntentRequest {
	s.IntentDefinition = v
	return s
}

func (s *UpdateIntentRequest) SetIntentId(v int64) *UpdateIntentRequest {
	s.IntentId = &v
	return s
}

func (s *UpdateIntentRequest) Validate() error {
	if s.IntentDefinition != nil {
		if err := s.IntentDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIntentRequestIntentDefinition struct {
	// The alias for the intent.
	//
	// example:
	//
	// 查天气意图
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The name of the intent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 查天气意图
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// A list of slot objects.
	SlotInfos []*UpdateIntentRequestIntentDefinitionSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s UpdateIntentRequestIntentDefinition) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntentRequestIntentDefinition) GoString() string {
	return s.String()
}

func (s *UpdateIntentRequestIntentDefinition) GetAliasName() *string {
	return s.AliasName
}

func (s *UpdateIntentRequestIntentDefinition) GetIntentName() *string {
	return s.IntentName
}

func (s *UpdateIntentRequestIntentDefinition) GetSlotInfos() []*UpdateIntentRequestIntentDefinitionSlotInfos {
	return s.SlotInfos
}

func (s *UpdateIntentRequestIntentDefinition) SetAliasName(v string) *UpdateIntentRequestIntentDefinition {
	s.AliasName = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinition) SetIntentName(v string) *UpdateIntentRequestIntentDefinition {
	s.IntentName = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinition) SetSlotInfos(v []*UpdateIntentRequestIntentDefinitionSlotInfos) *UpdateIntentRequestIntentDefinition {
	s.SlotInfos = v
	return s
}

func (s *UpdateIntentRequestIntentDefinition) Validate() error {
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

type UpdateIntentRequestIntentDefinitionSlotInfos struct {
	// Specifies whether the slot can contain an array of values.
	//
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// Specifies whether to enable data masking for the slot value.
	//
	// example:
	//
	// false
	Encrypt *bool `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// Specifies whether the slot is interactive. If true, the chatbot prompts the user for a value if the user does not provide one.
	//
	// example:
	//
	// false
	Interactive *bool `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	// The name of the slot.
	//
	// This parameter is required.
	//
	// example:
	//
	// 天气
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the slot.
	//
	// This parameter is required.
	//
	// example:
	//
	// dgadf23dfg2f
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

func (s UpdateIntentRequestIntentDefinitionSlotInfos) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntentRequestIntentDefinitionSlotInfos) GoString() string {
	return s.String()
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) GetArray() *bool {
	return s.Array
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) GetEncrypt() *bool {
	return s.Encrypt
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) GetInteractive() *bool {
	return s.Interactive
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) GetName() *string {
	return s.Name
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) GetSlotId() *string {
	return s.SlotId
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) GetValue() *string {
	return s.Value
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetArray(v bool) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Array = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetEncrypt(v bool) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Encrypt = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetInteractive(v bool) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Interactive = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetName(v string) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Name = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetSlotId(v string) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.SlotId = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetValue(v string) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Value = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) Validate() error {
	return dara.Validate(s)
}
