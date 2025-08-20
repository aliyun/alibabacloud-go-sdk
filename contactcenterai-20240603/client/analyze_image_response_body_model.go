// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFinishReason(v string) *AnalyzeImageResponseBody
	GetFinishReason() *string
	SetInputTokens(v string) *AnalyzeImageResponseBody
	GetInputTokens() *string
	SetOutputTokens(v string) *AnalyzeImageResponseBody
	GetOutputTokens() *string
	SetRequestId(v string) *AnalyzeImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AnalyzeImageResponseBody
	GetSuccess() *bool
	SetText(v string) *AnalyzeImageResponseBody
	GetText() *string
	SetTotalTokens(v string) *AnalyzeImageResponseBody
	GetTotalTokens() *string
}

type AnalyzeImageResponseBody struct {
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 1000
	InputTokens *string `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 2000
	OutputTokens *string `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 9*****-AE0D-5EE3-B1AF-48632CB0831C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
	Text    *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// 3000
	TotalTokens *string `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s AnalyzeImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeImageResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeImageResponseBody) GetFinishReason() *string {
	return s.FinishReason
}

func (s *AnalyzeImageResponseBody) GetInputTokens() *string {
	return s.InputTokens
}

func (s *AnalyzeImageResponseBody) GetOutputTokens() *string {
	return s.OutputTokens
}

func (s *AnalyzeImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnalyzeImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AnalyzeImageResponseBody) GetText() *string {
	return s.Text
}

func (s *AnalyzeImageResponseBody) GetTotalTokens() *string {
	return s.TotalTokens
}

func (s *AnalyzeImageResponseBody) SetFinishReason(v string) *AnalyzeImageResponseBody {
	s.FinishReason = &v
	return s
}

func (s *AnalyzeImageResponseBody) SetInputTokens(v string) *AnalyzeImageResponseBody {
	s.InputTokens = &v
	return s
}

func (s *AnalyzeImageResponseBody) SetOutputTokens(v string) *AnalyzeImageResponseBody {
	s.OutputTokens = &v
	return s
}

func (s *AnalyzeImageResponseBody) SetRequestId(v string) *AnalyzeImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnalyzeImageResponseBody) SetSuccess(v bool) *AnalyzeImageResponseBody {
	s.Success = &v
	return s
}

func (s *AnalyzeImageResponseBody) SetText(v string) *AnalyzeImageResponseBody {
	s.Text = &v
	return s
}

func (s *AnalyzeImageResponseBody) SetTotalTokens(v string) *AnalyzeImageResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *AnalyzeImageResponseBody) Validate() error {
	return dara.Validate(s)
}
