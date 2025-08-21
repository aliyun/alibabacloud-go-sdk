// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateUserSayRequest
	GetAgentKey() *string
	SetInstanceId(v string) *CreateUserSayRequest
	GetInstanceId() *string
	SetUserSayDefinition(v *CreateUserSayRequestUserSayDefinition) *CreateUserSayRequest
	GetUserSayDefinition() *CreateUserSayRequestUserSayDefinition
}

type CreateUserSayRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId        *string                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserSayDefinition *CreateUserSayRequestUserSayDefinition `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty" type:"Struct"`
}

func (s CreateUserSayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSayRequest) GoString() string {
	return s.String()
}

func (s *CreateUserSayRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateUserSayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserSayRequest) GetUserSayDefinition() *CreateUserSayRequestUserSayDefinition {
	return s.UserSayDefinition
}

func (s *CreateUserSayRequest) SetAgentKey(v string) *CreateUserSayRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateUserSayRequest) SetInstanceId(v string) *CreateUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserSayRequest) SetUserSayDefinition(v *CreateUserSayRequestUserSayDefinition) *CreateUserSayRequest {
	s.UserSayDefinition = v
	return s
}

func (s *CreateUserSayRequest) Validate() error {
	return dara.Validate(s)
}

type CreateUserSayRequestUserSayDefinition struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123232
	IntentId  *int64                                            `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	SlotInfos []*CreateUserSayRequestUserSayDefinitionSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s CreateUserSayRequestUserSayDefinition) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSayRequestUserSayDefinition) GoString() string {
	return s.String()
}

func (s *CreateUserSayRequestUserSayDefinition) GetContent() *string {
	return s.Content
}

func (s *CreateUserSayRequestUserSayDefinition) GetIntentId() *int64 {
	return s.IntentId
}

func (s *CreateUserSayRequestUserSayDefinition) GetSlotInfos() []*CreateUserSayRequestUserSayDefinitionSlotInfos {
	return s.SlotInfos
}

func (s *CreateUserSayRequestUserSayDefinition) SetContent(v string) *CreateUserSayRequestUserSayDefinition {
	s.Content = &v
	return s
}

func (s *CreateUserSayRequestUserSayDefinition) SetIntentId(v int64) *CreateUserSayRequestUserSayDefinition {
	s.IntentId = &v
	return s
}

func (s *CreateUserSayRequestUserSayDefinition) SetSlotInfos(v []*CreateUserSayRequestUserSayDefinitionSlotInfos) *CreateUserSayRequestUserSayDefinition {
	s.SlotInfos = v
	return s
}

func (s *CreateUserSayRequestUserSayDefinition) Validate() error {
	return dara.Validate(s)
}

type CreateUserSayRequestUserSayDefinitionSlotInfos struct {
	// example:
	//
	// 6
	EndIndex *int32 `json:"EndIndex,omitempty" xml:"EndIndex,omitempty"`
	// example:
	//
	// fb34adf2fv43f2
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// example:
	//
	// 4
	StartIndex *int32 `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
}

func (s CreateUserSayRequestUserSayDefinitionSlotInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSayRequestUserSayDefinitionSlotInfos) GoString() string {
	return s.String()
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) GetEndIndex() *int32 {
	return s.EndIndex
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) GetSlotId() *string {
	return s.SlotId
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) GetStartIndex() *int32 {
	return s.StartIndex
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) SetEndIndex(v int32) *CreateUserSayRequestUserSayDefinitionSlotInfos {
	s.EndIndex = &v
	return s
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) SetSlotId(v string) *CreateUserSayRequestUserSayDefinitionSlotInfos {
	s.SlotId = &v
	return s
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) SetStartIndex(v int32) *CreateUserSayRequestUserSayDefinitionSlotInfos {
	s.StartIndex = &v
	return s
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) Validate() error {
	return dara.Validate(s)
}
