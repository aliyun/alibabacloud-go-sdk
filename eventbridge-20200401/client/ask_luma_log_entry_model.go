// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAskLumaLogEntry interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *AskLumaLogEntry
	GetAgentName() *string
	SetClarificationNeeded(v bool) *AskLumaLogEntry
	GetClarificationNeeded() *bool
	SetClarificationQuestion(v string) *AskLumaLogEntry
	GetClarificationQuestion() *string
	SetContent(v *Content) *AskLumaLogEntry
	GetContent() *Content
	SetConversationId(v string) *AskLumaLogEntry
	GetConversationId() *string
	SetCreatedAt(v string) *AskLumaLogEntry
	GetCreatedAt() *string
	SetDurationMs(v int64) *AskLumaLogEntry
	GetDurationMs() *int64
	SetErrorCode(v string) *AskLumaLogEntry
	GetErrorCode() *string
	SetErrorMessage(v string) *AskLumaLogEntry
	GetErrorMessage() *string
	SetIsError(v bool) *AskLumaLogEntry
	GetIsError() *bool
	SetMessageId(v string) *AskLumaLogEntry
	GetMessageId() *string
	SetQuestion(v string) *AskLumaLogEntry
	GetQuestion() *string
	SetSource(v string) *AskLumaLogEntry
	GetSource() *string
	SetStatus(v string) *AskLumaLogEntry
	GetStatus() *string
}

type AskLumaLogEntry struct {
	// example:
	//
	// my-agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// false
	ClarificationNeeded   *bool    `json:"ClarificationNeeded,omitempty" xml:"ClarificationNeeded,omitempty"`
	ClarificationQuestion *string  `json:"ClarificationQuestion,omitempty" xml:"ClarificationQuestion,omitempty"`
	Content               *Content `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// conv_xxx
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 1717200000000
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 3200
	DurationMs *int64 `json:"DurationMs,omitempty" xml:"DurationMs,omitempty"`
	// example:
	//
	// ExecutionFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Agent with name \\"xxx\\" not found for account 1186xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// false
	IsError *bool `json:"IsError,omitempty" xml:"IsError,omitempty"`
	// example:
	//
	// msg_xxx
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Question  *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// MCP, CHAT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// RUNNING, SUCCEEDED, FAILED, TIMEOUT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AskLumaLogEntry) String() string {
	return dara.Prettify(s)
}

func (s AskLumaLogEntry) GoString() string {
	return s.String()
}

func (s *AskLumaLogEntry) GetAgentName() *string {
	return s.AgentName
}

func (s *AskLumaLogEntry) GetClarificationNeeded() *bool {
	return s.ClarificationNeeded
}

func (s *AskLumaLogEntry) GetClarificationQuestion() *string {
	return s.ClarificationQuestion
}

func (s *AskLumaLogEntry) GetContent() *Content {
	return s.Content
}

func (s *AskLumaLogEntry) GetConversationId() *string {
	return s.ConversationId
}

func (s *AskLumaLogEntry) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *AskLumaLogEntry) GetDurationMs() *int64 {
	return s.DurationMs
}

func (s *AskLumaLogEntry) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AskLumaLogEntry) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AskLumaLogEntry) GetIsError() *bool {
	return s.IsError
}

func (s *AskLumaLogEntry) GetMessageId() *string {
	return s.MessageId
}

func (s *AskLumaLogEntry) GetQuestion() *string {
	return s.Question
}

func (s *AskLumaLogEntry) GetSource() *string {
	return s.Source
}

func (s *AskLumaLogEntry) GetStatus() *string {
	return s.Status
}

func (s *AskLumaLogEntry) SetAgentName(v string) *AskLumaLogEntry {
	s.AgentName = &v
	return s
}

func (s *AskLumaLogEntry) SetClarificationNeeded(v bool) *AskLumaLogEntry {
	s.ClarificationNeeded = &v
	return s
}

func (s *AskLumaLogEntry) SetClarificationQuestion(v string) *AskLumaLogEntry {
	s.ClarificationQuestion = &v
	return s
}

func (s *AskLumaLogEntry) SetContent(v *Content) *AskLumaLogEntry {
	s.Content = v
	return s
}

func (s *AskLumaLogEntry) SetConversationId(v string) *AskLumaLogEntry {
	s.ConversationId = &v
	return s
}

func (s *AskLumaLogEntry) SetCreatedAt(v string) *AskLumaLogEntry {
	s.CreatedAt = &v
	return s
}

func (s *AskLumaLogEntry) SetDurationMs(v int64) *AskLumaLogEntry {
	s.DurationMs = &v
	return s
}

func (s *AskLumaLogEntry) SetErrorCode(v string) *AskLumaLogEntry {
	s.ErrorCode = &v
	return s
}

func (s *AskLumaLogEntry) SetErrorMessage(v string) *AskLumaLogEntry {
	s.ErrorMessage = &v
	return s
}

func (s *AskLumaLogEntry) SetIsError(v bool) *AskLumaLogEntry {
	s.IsError = &v
	return s
}

func (s *AskLumaLogEntry) SetMessageId(v string) *AskLumaLogEntry {
	s.MessageId = &v
	return s
}

func (s *AskLumaLogEntry) SetQuestion(v string) *AskLumaLogEntry {
	s.Question = &v
	return s
}

func (s *AskLumaLogEntry) SetSource(v string) *AskLumaLogEntry {
	s.Source = &v
	return s
}

func (s *AskLumaLogEntry) SetStatus(v string) *AskLumaLogEntry {
	s.Status = &v
	return s
}

func (s *AskLumaLogEntry) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}
