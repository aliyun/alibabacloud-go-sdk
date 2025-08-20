// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFinishReason(v string) *RunCompletionResponseBody
	GetFinishReason() *string
	SetRequestId(v string) *RunCompletionResponseBody
	GetRequestId() *string
	SetText(v string) *RunCompletionResponseBody
	GetText() *string
	SetInputTokens(v string) *RunCompletionResponseBody
	GetInputTokens() *string
	SetOutputTokens(v string) *RunCompletionResponseBody
	GetOutputTokens() *string
	SetTotalTokens(v string) *RunCompletionResponseBody
	GetTotalTokens() *string
}

type RunCompletionResponseBody struct {
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

func (s RunCompletionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionResponseBody) GoString() string {
	return s.String()
}

func (s *RunCompletionResponseBody) GetFinishReason() *string {
	return s.FinishReason
}

func (s *RunCompletionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCompletionResponseBody) GetText() *string {
	return s.Text
}

func (s *RunCompletionResponseBody) GetInputTokens() *string {
	return s.InputTokens
}

func (s *RunCompletionResponseBody) GetOutputTokens() *string {
	return s.OutputTokens
}

func (s *RunCompletionResponseBody) GetTotalTokens() *string {
	return s.TotalTokens
}

func (s *RunCompletionResponseBody) SetFinishReason(v string) *RunCompletionResponseBody {
	s.FinishReason = &v
	return s
}

func (s *RunCompletionResponseBody) SetRequestId(v string) *RunCompletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCompletionResponseBody) SetText(v string) *RunCompletionResponseBody {
	s.Text = &v
	return s
}

func (s *RunCompletionResponseBody) SetInputTokens(v string) *RunCompletionResponseBody {
	s.InputTokens = &v
	return s
}

func (s *RunCompletionResponseBody) SetOutputTokens(v string) *RunCompletionResponseBody {
	s.OutputTokens = &v
	return s
}

func (s *RunCompletionResponseBody) SetTotalTokens(v string) *RunCompletionResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *RunCompletionResponseBody) Validate() error {
	return dara.Validate(s)
}
