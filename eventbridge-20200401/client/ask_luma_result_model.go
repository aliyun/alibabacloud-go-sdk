// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAskLumaResult interface {
	dara.Model
	String() string
	GoString() string
	SetClarificationNeeded(v bool) *AskLumaResult
	GetClarificationNeeded() *bool
	SetClarificationQuestion(v string) *AskLumaResult
	GetClarificationQuestion() *string
	SetConstraints(v *Constraints) *AskLumaResult
	GetConstraints() *Constraints
	SetContent(v *Content) *AskLumaResult
	GetContent() *Content
	SetConversationId(v string) *AskLumaResult
	GetConversationId() *string
	SetErrorCode(v string) *AskLumaResult
	GetErrorCode() *string
	SetErrorMessage(v string) *AskLumaResult
	GetErrorMessage() *string
	SetIsError(v bool) *AskLumaResult
	GetIsError() *bool
	SetMessageId(v string) *AskLumaResult
	GetMessageId() *string
	SetStatus(v string) *AskLumaResult
	GetStatus() *string
	SetStorageTruncated(v bool) *AskLumaResult
	GetStorageTruncated() *bool
	SetWikiVersion(v string) *AskLumaResult
	GetWikiVersion() *string
}

type AskLumaResult struct {
	// Indicates whether clarification is needed.
	//
	// example:
	//
	// false
	ClarificationNeeded *bool `json:"ClarificationNeeded,omitempty" xml:"ClarificationNeeded,omitempty"`
	// The clarification question text.
	//
	// example:
	//
	// Which database does the employee table you are referring to belong to?
	ClarificationQuestion *string `json:"ClarificationQuestion,omitempty" xml:"ClarificationQuestion,omitempty"`
	// The query constraints.
	Constraints *Constraints `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The structured result body.
	Content *Content `json:"Content,omitempty" xml:"Content,omitempty"`
	// The conversation ID, used for multi-turn follow-up questions.
	//
	// example:
	//
	// conv_xxx
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The error code.
	//
	// example:
	//
	// ExecutionFailed, Timeout, RateLimited, InternalError, ConversationExpired
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error details.
	//
	// example:
	//
	// Agent with name \\"xxx\\" not found for account 1186xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether an error occurred. A value of false indicates that the query succeeded or a clarification is needed (including empty result sets). A value of true indicates that the execution failed due to a timeout, throttling, or internal error.
	//
	// example:
	//
	// false
	IsError *bool `json:"IsError,omitempty" xml:"IsError,omitempty"`
	// The message ID, used for polling with PollAskResult.
	//
	// example:
	//
	// msg_xxx
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The submit status.
	//
	// example:
	//
	// RUNNING, SUCCEEDED, FAILED, TIMEOUT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the result was truncated because it exceeded the storage limit. This field is returned only for large result sets.
	//
	// example:
	//
	// true
	StorageTruncated *bool `json:"StorageTruncated,omitempty" xml:"StorageTruncated,omitempty"`
	// The business Wiki version that was actually used for this response. This field is not returned if the agent does not have a Wiki configured.
	//
	// example:
	//
	// eventhouse-multisource-demo-v1
	WikiVersion *string `json:"WikiVersion,omitempty" xml:"WikiVersion,omitempty"`
}

func (s AskLumaResult) String() string {
	return dara.Prettify(s)
}

func (s AskLumaResult) GoString() string {
	return s.String()
}

func (s *AskLumaResult) GetClarificationNeeded() *bool {
	return s.ClarificationNeeded
}

func (s *AskLumaResult) GetClarificationQuestion() *string {
	return s.ClarificationQuestion
}

func (s *AskLumaResult) GetConstraints() *Constraints {
	return s.Constraints
}

func (s *AskLumaResult) GetContent() *Content {
	return s.Content
}

func (s *AskLumaResult) GetConversationId() *string {
	return s.ConversationId
}

func (s *AskLumaResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AskLumaResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AskLumaResult) GetIsError() *bool {
	return s.IsError
}

func (s *AskLumaResult) GetMessageId() *string {
	return s.MessageId
}

func (s *AskLumaResult) GetStatus() *string {
	return s.Status
}

func (s *AskLumaResult) GetStorageTruncated() *bool {
	return s.StorageTruncated
}

func (s *AskLumaResult) GetWikiVersion() *string {
	return s.WikiVersion
}

func (s *AskLumaResult) SetClarificationNeeded(v bool) *AskLumaResult {
	s.ClarificationNeeded = &v
	return s
}

func (s *AskLumaResult) SetClarificationQuestion(v string) *AskLumaResult {
	s.ClarificationQuestion = &v
	return s
}

func (s *AskLumaResult) SetConstraints(v *Constraints) *AskLumaResult {
	s.Constraints = v
	return s
}

func (s *AskLumaResult) SetContent(v *Content) *AskLumaResult {
	s.Content = v
	return s
}

func (s *AskLumaResult) SetConversationId(v string) *AskLumaResult {
	s.ConversationId = &v
	return s
}

func (s *AskLumaResult) SetErrorCode(v string) *AskLumaResult {
	s.ErrorCode = &v
	return s
}

func (s *AskLumaResult) SetErrorMessage(v string) *AskLumaResult {
	s.ErrorMessage = &v
	return s
}

func (s *AskLumaResult) SetIsError(v bool) *AskLumaResult {
	s.IsError = &v
	return s
}

func (s *AskLumaResult) SetMessageId(v string) *AskLumaResult {
	s.MessageId = &v
	return s
}

func (s *AskLumaResult) SetStatus(v string) *AskLumaResult {
	s.Status = &v
	return s
}

func (s *AskLumaResult) SetStorageTruncated(v bool) *AskLumaResult {
	s.StorageTruncated = &v
	return s
}

func (s *AskLumaResult) SetWikiVersion(v string) *AskLumaResult {
	s.WikiVersion = &v
	return s
}

func (s *AskLumaResult) Validate() error {
	if s.Constraints != nil {
		if err := s.Constraints.Validate(); err != nil {
			return err
		}
	}
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}
