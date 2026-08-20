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
	SetFiles(v []*ChatMessagesRequestFiles) *ChatMessagesRequest
	GetFiles() []*ChatMessagesRequestFiles
	SetInputs(v *ChatMessagesRequestInputs) *ChatMessagesRequest
	GetInputs() *ChatMessagesRequestInputs
	SetParentMessageId(v string) *ChatMessagesRequest
	GetParentMessageId() *string
	SetQuery(v string) *ChatMessagesRequest
	GetQuery() *string
}

type ChatMessagesRequest struct {
	// The conversation ID.
	//
	// example:
	//
	// fea7bdca-e848-44dd-b1ae-852472b8****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The event output type. Valid values: inline and separate. Default value: inline. When set to inline, tool invocation events, sub-node events, and document events are included in the answer field of event = message. When set to separate, tool invocation events, sub-node events, and document events each have their own event.
	//
	// example:
	//
	// inline
	EventMode *string                     `json:"EventMode,omitempty" xml:"EventMode,omitempty"`
	Files     []*ChatMessagesRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The task input.
	Inputs *ChatMessagesRequestInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The parent message ID.
	//
	// example:
	//
	// 84dc9f9b-424a-404d-9c36-35e9d000****
	ParentMessageId *string `json:"ParentMessageId,omitempty" xml:"ParentMessageId,omitempty"`
	// The query content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Disk usage of instance rm-bp14as9914vd3****, is capacity expansion needed
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

func (s *ChatMessagesRequest) GetFiles() []*ChatMessagesRequestFiles {
	return s.Files
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

func (s *ChatMessagesRequest) SetFiles(v []*ChatMessagesRequestFiles) *ChatMessagesRequest {
	s.Files = v
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
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Inputs != nil {
		if err := s.Inputs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatMessagesRequestFiles struct {
	TransferMethod *string `json:"TransferMethod,omitempty" xml:"TransferMethod,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UploadFileId   *string `json:"UploadFileId,omitempty" xml:"UploadFileId,omitempty"`
}

func (s ChatMessagesRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s ChatMessagesRequestFiles) GoString() string {
	return s.String()
}

func (s *ChatMessagesRequestFiles) GetTransferMethod() *string {
	return s.TransferMethod
}

func (s *ChatMessagesRequestFiles) GetType() *string {
	return s.Type
}

func (s *ChatMessagesRequestFiles) GetUploadFileId() *string {
	return s.UploadFileId
}

func (s *ChatMessagesRequestFiles) SetTransferMethod(v string) *ChatMessagesRequestFiles {
	s.TransferMethod = &v
	return s
}

func (s *ChatMessagesRequestFiles) SetType(v string) *ChatMessagesRequestFiles {
	s.Type = &v
	return s
}

func (s *ChatMessagesRequestFiles) SetUploadFileId(v string) *ChatMessagesRequestFiles {
	s.UploadFileId = &v
	return s
}

func (s *ChatMessagesRequestFiles) Validate() error {
	return dara.Validate(s)
}

type ChatMessagesRequestInputs struct {
	// The custom agent ID for the user.
	//
	// example:
	//
	// d1b7d639-f34e-44c7-8231-987da14d****
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// Specifies whether to enable deep thinking mode.
	//
	// example:
	//
	// true
	EnableThinking *string `json:"EnableThinking,omitempty" xml:"EnableThinking,omitempty"`
	// The conversation language.
	//
	// example:
	//
	// zh-cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The model ID.
	//
	// example:
	//
	// qwen3.7-max
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The thinking depth.
	//
	// example:
	//
	// default
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

func (s *ChatMessagesRequestInputs) GetModelId() *string {
	return s.ModelId
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

func (s *ChatMessagesRequestInputs) SetModelId(v string) *ChatMessagesRequestInputs {
	s.ModelId = &v
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
