// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *ChatMessagesRequest
	GetConversationId() *string
	SetEventMode(v string) *ChatMessagesRequest
	GetEventMode() *string
	SetInputs(v *ChatMessagesRequestInputs) *ChatMessagesRequest
	GetInputs() *ChatMessagesRequestInputs
	SetParentMessageId(v string) *ChatMessagesRequest
	GetParentMessageId() *string
	SetQuery(v string) *ChatMessagesRequest
	GetQuery() *string
}

type ChatMessagesRequest struct {
	// The ID of the conversation.
	//
	// example:
	//
	// fea7bdca-e848-44dd-b1ae-852472b8****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	EventMode      *string `json:"EventMode,omitempty" xml:"EventMode,omitempty"`
	// The inputs for the task.
	Inputs *ChatMessagesRequestInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The ID of the parent message.
	//
	// example:
	//
	// 84dc9f9b-424a-404d-9c36-35e9d000****
	ParentMessageId *string `json:"ParentMessageId,omitempty" xml:"ParentMessageId,omitempty"`
	// The content of the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 实例rm-bp14as9914vd3***	- 磁盘使用率，是否需要进行扩容
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ChatMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatMessagesRequest) GoString() string {
	return s.String()
}

func (s *ChatMessagesRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ChatMessagesRequest) GetEventMode() *string {
	return s.EventMode
}

func (s *ChatMessagesRequest) GetInputs() *ChatMessagesRequestInputs {
	return s.Inputs
}

func (s *ChatMessagesRequest) GetParentMessageId() *string {
	return s.ParentMessageId
}

func (s *ChatMessagesRequest) GetQuery() *string {
	return s.Query
}

func (s *ChatMessagesRequest) SetConversationId(v string) *ChatMessagesRequest {
	s.ConversationId = &v
	return s
}

func (s *ChatMessagesRequest) SetEventMode(v string) *ChatMessagesRequest {
	s.EventMode = &v
	return s
}

func (s *ChatMessagesRequest) SetInputs(v *ChatMessagesRequestInputs) *ChatMessagesRequest {
	s.Inputs = v
	return s
}

func (s *ChatMessagesRequest) SetParentMessageId(v string) *ChatMessagesRequest {
	s.ParentMessageId = &v
	return s
}

func (s *ChatMessagesRequest) SetQuery(v string) *ChatMessagesRequest {
	s.Query = &v
	return s
}

func (s *ChatMessagesRequest) Validate() error {
	if s.Inputs != nil {
		if err := s.Inputs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatMessagesRequestInputs struct {
	// The custom agent ID.
	//
	// example:
	//
	// d1b7d639-f34e-44c7-8231-987da14d****
	CustomAgentId  *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	EnableThinking *string `json:"EnableThinking,omitempty" xml:"EnableThinking,omitempty"`
	// The language of the conversation.
	//
	// example:
	//
	// zh-cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ThinkEffort *string `json:"ThinkEffort,omitempty" xml:"ThinkEffort,omitempty"`
	// The time zone. Default value: **Asia/Shanghai**.
	//
	// example:
	//
	// UTC
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s ChatMessagesRequestInputs) String() string {
	return dara.Prettify(s)
}

func (s ChatMessagesRequestInputs) GoString() string {
	return s.String()
}

func (s *ChatMessagesRequestInputs) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ChatMessagesRequestInputs) GetEnableThinking() *string {
	return s.EnableThinking
}

func (s *ChatMessagesRequestInputs) GetLanguage() *string {
	return s.Language
}

func (s *ChatMessagesRequestInputs) GetRegionId() *string {
	return s.RegionId
}

func (s *ChatMessagesRequestInputs) GetThinkEffort() *string {
	return s.ThinkEffort
}

func (s *ChatMessagesRequestInputs) GetTimezone() *string {
	return s.Timezone
}

func (s *ChatMessagesRequestInputs) SetCustomAgentId(v string) *ChatMessagesRequestInputs {
	s.CustomAgentId = &v
	return s
}

func (s *ChatMessagesRequestInputs) SetEnableThinking(v string) *ChatMessagesRequestInputs {
	s.EnableThinking = &v
	return s
}

func (s *ChatMessagesRequestInputs) SetLanguage(v string) *ChatMessagesRequestInputs {
	s.Language = &v
	return s
}

func (s *ChatMessagesRequestInputs) SetRegionId(v string) *ChatMessagesRequestInputs {
	s.RegionId = &v
	return s
}

func (s *ChatMessagesRequestInputs) SetThinkEffort(v string) *ChatMessagesRequestInputs {
	s.ThinkEffort = &v
	return s
}

func (s *ChatMessagesRequestInputs) SetTimezone(v string) *ChatMessagesRequestInputs {
	s.Timezone = &v
	return s
}

func (s *ChatMessagesRequestInputs) Validate() error {
	return dara.Validate(s)
}
