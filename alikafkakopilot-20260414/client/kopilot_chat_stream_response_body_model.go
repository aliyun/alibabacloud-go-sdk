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
	// The actual content.
	//
	// example:
	//
	// test
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The streaming incremental content.
	//
	// example:
	//
	// hello
	Delta *string `json:"Delta,omitempty" xml:"Delta,omitempty"`
	// The message body.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique message ID.
	//
	// example:
	//
	// 4b209618fd066c4354037b4b0634ffc9
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 76E1F1AA-6046-5074-96E2-79A37AFBD2FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role identifier.
	//
	// example:
	//
	// assistant
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The run task ID.
	//
	// example:
	//
	// 5737d000********
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The session thread ID.
	//
	// example:
	//
	// thread_abc123xyz
	ThreadId *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// The unique tool calling invoke ID.
	//
	// example:
	//
	// call_xyz789012
	ToolCallId *string `json:"ToolCallId,omitempty" xml:"ToolCallId,omitempty"`
	// The tool or function name.
	//
	// example:
	//
	// search_knowledge_base
	ToolCallName *string `json:"ToolCallName,omitempty" xml:"ToolCallName,omitempty"`
	// The event or message type.
	//
	// example:
	//
	// delta
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
