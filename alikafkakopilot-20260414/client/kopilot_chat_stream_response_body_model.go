// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotChatStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *KopilotChatStreamResponseBody
	GetContent() *string
	SetDelta(v string) *KopilotChatStreamResponseBody
	GetDelta() *string
	SetMessage(v string) *KopilotChatStreamResponseBody
	GetMessage() *string
	SetMessageId(v string) *KopilotChatStreamResponseBody
	GetMessageId() *string
	SetRequestId(v string) *KopilotChatStreamResponseBody
	GetRequestId() *string
	SetRole(v string) *KopilotChatStreamResponseBody
	GetRole() *string
	SetRunId(v string) *KopilotChatStreamResponseBody
	GetRunId() *string
	SetThreadId(v string) *KopilotChatStreamResponseBody
	GetThreadId() *string
	SetToolCallId(v string) *KopilotChatStreamResponseBody
	GetToolCallId() *string
	SetToolCallName(v string) *KopilotChatStreamResponseBody
	GetToolCallName() *string
	SetType(v string) *KopilotChatStreamResponseBody
	GetType() *string
}

type KopilotChatStreamResponseBody struct {
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Delta        *string `json:"Delta,omitempty" xml:"Delta,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageId    *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	RunId        *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	ThreadId     *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	ToolCallId   *string `json:"ToolCallId,omitempty" xml:"ToolCallId,omitempty"`
	ToolCallName *string `json:"ToolCallName,omitempty" xml:"ToolCallName,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s KopilotChatStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KopilotChatStreamResponseBody) GoString() string {
	return s.String()
}

func (s *KopilotChatStreamResponseBody) GetContent() *string {
	return s.Content
}

func (s *KopilotChatStreamResponseBody) GetDelta() *string {
	return s.Delta
}

func (s *KopilotChatStreamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *KopilotChatStreamResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *KopilotChatStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KopilotChatStreamResponseBody) GetRole() *string {
	return s.Role
}

func (s *KopilotChatStreamResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *KopilotChatStreamResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *KopilotChatStreamResponseBody) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *KopilotChatStreamResponseBody) GetToolCallName() *string {
	return s.ToolCallName
}

func (s *KopilotChatStreamResponseBody) GetType() *string {
	return s.Type
}

func (s *KopilotChatStreamResponseBody) SetContent(v string) *KopilotChatStreamResponseBody {
	s.Content = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetDelta(v string) *KopilotChatStreamResponseBody {
	s.Delta = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetMessage(v string) *KopilotChatStreamResponseBody {
	s.Message = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetMessageId(v string) *KopilotChatStreamResponseBody {
	s.MessageId = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetRequestId(v string) *KopilotChatStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetRole(v string) *KopilotChatStreamResponseBody {
	s.Role = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetRunId(v string) *KopilotChatStreamResponseBody {
	s.RunId = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetThreadId(v string) *KopilotChatStreamResponseBody {
	s.ThreadId = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetToolCallId(v string) *KopilotChatStreamResponseBody {
	s.ToolCallId = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetToolCallName(v string) *KopilotChatStreamResponseBody {
	s.ToolCallName = &v
	return s
}

func (s *KopilotChatStreamResponseBody) SetType(v string) *KopilotChatStreamResponseBody {
	s.Type = &v
	return s
}

func (s *KopilotChatStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
