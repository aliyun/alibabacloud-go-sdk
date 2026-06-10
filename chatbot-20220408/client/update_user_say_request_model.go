// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserSayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateUserSayRequest
	GetAgentKey() *string
	SetInstanceId(v string) *UpdateUserSayRequest
	GetInstanceId() *string
	SetUserSayDefinition(v *UpdateUserSayRequestUserSayDefinition) *UpdateUserSayRequest
	GetUserSayDefinition() *UpdateUserSayRequestUserSayDefinition
	SetUserSayId(v int64) *UpdateUserSayRequest
	GetUserSayId() *int64
}

type UpdateUserSayRequest struct {
	// The workspace key. If you omit this parameter, the default workspace is used. You can find this key on the Workspace Management page in your main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The ID of the bot.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The definition of the user say.
	UserSayDefinition *UpdateUserSayRequestUserSayDefinition `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty" type:"Struct"`
	// The ID of the user say.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34512323
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s UpdateUserSayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserSayRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserSayRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateUserSayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateUserSayRequest) GetUserSayDefinition() *UpdateUserSayRequestUserSayDefinition {
	return s.UserSayDefinition
}

func (s *UpdateUserSayRequest) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *UpdateUserSayRequest) SetAgentKey(v string) *UpdateUserSayRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateUserSayRequest) SetInstanceId(v string) *UpdateUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserSayRequest) SetUserSayDefinition(v *UpdateUserSayRequestUserSayDefinition) *UpdateUserSayRequest {
	s.UserSayDefinition = v
	return s
}

func (s *UpdateUserSayRequest) SetUserSayId(v int64) *UpdateUserSayRequest {
	s.UserSayId = &v
	return s
}

func (s *UpdateUserSayRequest) Validate() error {
	if s.UserSayDefinition != nil {
		if err := s.UserSayDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserSayRequestUserSayDefinition struct {
	// The content of the user say.
	//
	// This parameter is required.
	//
	// example:
	//
	// 做核酸了吗您
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the intent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// An array that contains information for each slot.
	SlotInfos []*UpdateUserSayRequestUserSayDefinitionSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s UpdateUserSayRequestUserSayDefinition) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserSayRequestUserSayDefinition) GoString() string {
	return s.String()
}

func (s *UpdateUserSayRequestUserSayDefinition) GetContent() *string {
	return s.Content
}

func (s *UpdateUserSayRequestUserSayDefinition) GetIntentId() *int64 {
	return s.IntentId
}

func (s *UpdateUserSayRequestUserSayDefinition) GetSlotInfos() []*UpdateUserSayRequestUserSayDefinitionSlotInfos {
	return s.SlotInfos
}

func (s *UpdateUserSayRequestUserSayDefinition) SetContent(v string) *UpdateUserSayRequestUserSayDefinition {
	s.Content = &v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinition) SetIntentId(v int64) *UpdateUserSayRequestUserSayDefinition {
	s.IntentId = &v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinition) SetSlotInfos(v []*UpdateUserSayRequestUserSayDefinitionSlotInfos) *UpdateUserSayRequestUserSayDefinition {
	s.SlotInfos = v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinition) Validate() error {
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

type UpdateUserSayRequestUserSayDefinitionSlotInfos struct {
	// The zero-based end index (exclusive) of the slot within the content of the user say.
	//
	// example:
	//
	// 3
	EndIndex *int32 `json:"EndIndex,omitempty" xml:"EndIndex,omitempty"`
	// The ID of the slot.
	//
	// example:
	//
	// 346ffg3q23dv
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// The zero-based start index of the slot within the content of the user say.
	//
	// example:
	//
	// 1
	StartIndex *int32 `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
}

func (s UpdateUserSayRequestUserSayDefinitionSlotInfos) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserSayRequestUserSayDefinitionSlotInfos) GoString() string {
	return s.String()
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) GetEndIndex() *int32 {
	return s.EndIndex
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) GetSlotId() *string {
	return s.SlotId
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) GetStartIndex() *int32 {
	return s.StartIndex
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) SetEndIndex(v int32) *UpdateUserSayRequestUserSayDefinitionSlotInfos {
	s.EndIndex = &v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) SetSlotId(v string) *UpdateUserSayRequestUserSayDefinitionSlotInfos {
	s.SlotId = &v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) SetStartIndex(v int32) *UpdateUserSayRequestUserSayDefinitionSlotInfos {
	s.StartIndex = &v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) Validate() error {
	return dara.Validate(s)
}
