// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddedMetaData(v string) *UpdateAppChatMessageRequest
	GetAddedMetaData() *string
	SetContent(v string) *UpdateAppChatMessageRequest
	GetContent() *string
	SetConversationId(v string) *UpdateAppChatMessageRequest
	GetConversationId() *string
	SetMessageId(v string) *UpdateAppChatMessageRequest
	GetMessageId() *string
}

type UpdateAppChatMessageRequest struct {
	// Appended message metadata (JSON format)
	//
	// example:
	//
	// 123
	AddedMetaData *string `json:"AddedMetaData,omitempty" xml:"AddedMetaData,omitempty"`
	// Message content
	//
	// example:
	//
	// FormatVersion: OOS-2019-06-01nTasks:n  - Name: runCommandn    Action: \\"ACS::ECS::RunCommand\\"n    Properties:n      commandContent: \\" echo Hksqj@@883289
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Session ID
	//
	// example:
	//
	// 81bc5a34-1d8d-4ef7-a208-7401c51b054b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// Message ID
	//
	// example:
	//
	// 7baf7d67-1897-42ed-a380-f6ae825d6907
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s UpdateAppChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppChatMessageRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppChatMessageRequest) GetAddedMetaData() *string {
	return s.AddedMetaData
}

func (s *UpdateAppChatMessageRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateAppChatMessageRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *UpdateAppChatMessageRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *UpdateAppChatMessageRequest) SetAddedMetaData(v string) *UpdateAppChatMessageRequest {
	s.AddedMetaData = &v
	return s
}

func (s *UpdateAppChatMessageRequest) SetContent(v string) *UpdateAppChatMessageRequest {
	s.Content = &v
	return s
}

func (s *UpdateAppChatMessageRequest) SetConversationId(v string) *UpdateAppChatMessageRequest {
	s.ConversationId = &v
	return s
}

func (s *UpdateAppChatMessageRequest) SetMessageId(v string) *UpdateAppChatMessageRequest {
	s.MessageId = &v
	return s
}

func (s *UpdateAppChatMessageRequest) Validate() error {
	return dara.Validate(s)
}
