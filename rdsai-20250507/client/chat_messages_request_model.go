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
	SetInputs(v *ChatMessagesRequestInputs) *ChatMessagesRequest
	GetInputs() *ChatMessagesRequestInputs
	SetParentMessageId(v string) *ChatMessagesRequest
	GetParentMessageId() *string
	SetQuery(v string) *ChatMessagesRequest
	GetQuery() *string
}

type ChatMessagesRequest struct {
	// example:
	//
	// fea7bdca-e848-44dd-b1ae-852472b8****
	ConversationId *string                    `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	Inputs         *ChatMessagesRequestInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// example:
	//
	// 84dc9f9b-424a-404d-9c36-35e9d000****
	ParentMessageId *string `json:"ParentMessageId,omitempty" xml:"ParentMessageId,omitempty"`
	// This parameter is required.
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
	// example:
	//
	// d1b7d639-f34e-44c7-8231-987da14d****
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// zh-cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ChatMessagesRequestInputs) GetLanguage() *string {
	return s.Language
}

func (s *ChatMessagesRequestInputs) GetRegionId() *string {
	return s.RegionId
}

func (s *ChatMessagesRequestInputs) GetTimezone() *string {
	return s.Timezone
}

func (s *ChatMessagesRequestInputs) SetCustomAgentId(v string) *ChatMessagesRequestInputs {
	s.CustomAgentId = &v
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

func (s *ChatMessagesRequestInputs) SetTimezone(v string) *ChatMessagesRequestInputs {
	s.Timezone = &v
	return s
}

func (s *ChatMessagesRequestInputs) Validate() error {
	return dara.Validate(s)
}
