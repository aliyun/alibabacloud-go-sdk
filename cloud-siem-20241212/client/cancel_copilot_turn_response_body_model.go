// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCopilotTurnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiVersion(v string) *CancelCopilotTurnResponseBody
	GetApiVersion() *string
	SetRequestId(v string) *CancelCopilotTurnResponseBody
	GetRequestId() *string
	SetTurn(v *CancelCopilotTurnResponseBodyTurn) *CancelCopilotTurnResponseBody
	GetTurn() *CancelCopilotTurnResponseBodyTurn
}

type CancelCopilotTurnResponseBody struct {
	// The API version number. The value is 1.0.0.
	//
	// example:
	//
	// v1.2
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security operations Agent conversation turn.
	Turn *CancelCopilotTurnResponseBodyTurn `json:"Turn,omitempty" xml:"Turn,omitempty" type:"Struct"`
}

func (s CancelCopilotTurnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCopilotTurnResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCopilotTurnResponseBody) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *CancelCopilotTurnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCopilotTurnResponseBody) GetTurn() *CancelCopilotTurnResponseBodyTurn {
	return s.Turn
}

func (s *CancelCopilotTurnResponseBody) SetApiVersion(v string) *CancelCopilotTurnResponseBody {
	s.ApiVersion = &v
	return s
}

func (s *CancelCopilotTurnResponseBody) SetRequestId(v string) *CancelCopilotTurnResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCopilotTurnResponseBody) SetTurn(v *CancelCopilotTurnResponseBodyTurn) *CancelCopilotTurnResponseBody {
	s.Turn = v
	return s
}

func (s *CancelCopilotTurnResponseBody) Validate() error {
	if s.Turn != nil {
		if err := s.Turn.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelCopilotTurnResponseBodyTurn struct {
	// The client message ID.
	//
	// example:
	//
	// 366932fb-9a83-4a2d-94dd-c925b35f****
	ClientMessageId *string `json:"ClientMessageId,omitempty" xml:"ClientMessageId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The creation time. The time follows the ISO 8601 standard in UTC.
	//
	// example:
	//
	// 2026-08-23T02:13:49Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The conversation feedback information.
	//
	// example:
	//
	// 1
	Feedback *int64 `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// The number of tokens used for input.
	//
	// example:
	//
	// 1280
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// The model ID.
	//
	// example:
	//
	// model-****
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The number of tokens used for output.
	//
	// example:
	//
	// 856
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The number of reasoning tokens used on the current day.
	//
	// example:
	//
	// 420
	ReasoningTokens *int64 `json:"ReasoningTokens,omitempty" xml:"ReasoningTokens,omitempty"`
	// Indicates whether a retry is allowed.
	RetryAllowed *bool `json:"RetryAllowed,omitempty" xml:"RetryAllowed,omitempty"`
	// The status of the conversation turn.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// dc18b894-19a3-45ee-8768-a59478e0****
	TurnId *string `json:"TurnId,omitempty" xml:"TurnId,omitempty"`
}

func (s CancelCopilotTurnResponseBodyTurn) String() string {
	return dara.Prettify(s)
}

func (s CancelCopilotTurnResponseBodyTurn) GoString() string {
	return s.String()
}

func (s *CancelCopilotTurnResponseBodyTurn) GetClientMessageId() *string {
	return s.ClientMessageId
}

func (s *CancelCopilotTurnResponseBodyTurn) GetConversationId() *string {
	return s.ConversationId
}

func (s *CancelCopilotTurnResponseBodyTurn) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CancelCopilotTurnResponseBodyTurn) GetFeedback() *int64 {
	return s.Feedback
}

func (s *CancelCopilotTurnResponseBodyTurn) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *CancelCopilotTurnResponseBodyTurn) GetModelId() *string {
	return s.ModelId
}

func (s *CancelCopilotTurnResponseBodyTurn) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *CancelCopilotTurnResponseBodyTurn) GetReasoningTokens() *int64 {
	return s.ReasoningTokens
}

func (s *CancelCopilotTurnResponseBodyTurn) GetRetryAllowed() *bool {
	return s.RetryAllowed
}

func (s *CancelCopilotTurnResponseBodyTurn) GetStatus() *string {
	return s.Status
}

func (s *CancelCopilotTurnResponseBodyTurn) GetTurnId() *string {
	return s.TurnId
}

func (s *CancelCopilotTurnResponseBodyTurn) SetClientMessageId(v string) *CancelCopilotTurnResponseBodyTurn {
	s.ClientMessageId = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetConversationId(v string) *CancelCopilotTurnResponseBodyTurn {
	s.ConversationId = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetCreatedAt(v string) *CancelCopilotTurnResponseBodyTurn {
	s.CreatedAt = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetFeedback(v int64) *CancelCopilotTurnResponseBodyTurn {
	s.Feedback = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetInputTokens(v int64) *CancelCopilotTurnResponseBodyTurn {
	s.InputTokens = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetModelId(v string) *CancelCopilotTurnResponseBodyTurn {
	s.ModelId = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetOutputTokens(v int64) *CancelCopilotTurnResponseBodyTurn {
	s.OutputTokens = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetReasoningTokens(v int64) *CancelCopilotTurnResponseBodyTurn {
	s.ReasoningTokens = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetRetryAllowed(v bool) *CancelCopilotTurnResponseBodyTurn {
	s.RetryAllowed = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetStatus(v string) *CancelCopilotTurnResponseBodyTurn {
	s.Status = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) SetTurnId(v string) *CancelCopilotTurnResponseBodyTurn {
	s.TurnId = &v
	return s
}

func (s *CancelCopilotTurnResponseBodyTurn) Validate() error {
	return dara.Validate(s)
}
