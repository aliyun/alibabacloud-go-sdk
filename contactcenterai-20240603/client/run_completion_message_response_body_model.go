// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFinishReason(v string) *RunCompletionMessageResponseBody
	GetFinishReason() *string
	SetRequestId(v string) *RunCompletionMessageResponseBody
	GetRequestId() *string
	SetText(v string) *RunCompletionMessageResponseBody
	GetText() *string
	SetInputTokens(v string) *RunCompletionMessageResponseBody
	GetInputTokens() *string
	SetOutputTokens(v string) *RunCompletionMessageResponseBody
	GetOutputTokens() *string
	SetTotalTokens(v string) *RunCompletionMessageResponseBody
	GetTotalTokens() *string
}

type RunCompletionMessageResponseBody struct {
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Text         *string `json:"Text,omitempty" xml:"Text,omitempty"`
	InputTokens  *string `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *string `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	TotalTokens  *string `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunCompletionMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageResponseBody) GetFinishReason() *string {
	return s.FinishReason
}

func (s *RunCompletionMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCompletionMessageResponseBody) GetText() *string {
	return s.Text
}

func (s *RunCompletionMessageResponseBody) GetInputTokens() *string {
	return s.InputTokens
}

func (s *RunCompletionMessageResponseBody) GetOutputTokens() *string {
	return s.OutputTokens
}

func (s *RunCompletionMessageResponseBody) GetTotalTokens() *string {
	return s.TotalTokens
}

func (s *RunCompletionMessageResponseBody) SetFinishReason(v string) *RunCompletionMessageResponseBody {
	s.FinishReason = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetRequestId(v string) *RunCompletionMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetText(v string) *RunCompletionMessageResponseBody {
	s.Text = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetInputTokens(v string) *RunCompletionMessageResponseBody {
	s.InputTokens = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetOutputTokens(v string) *RunCompletionMessageResponseBody {
	s.OutputTokens = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetTotalTokens(v string) *RunCompletionMessageResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *RunCompletionMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
