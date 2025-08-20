// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralAnalyzeImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFinishReason(v string) *GeneralAnalyzeImageResponseBody
	GetFinishReason() *string
	SetInputTokens(v int32) *GeneralAnalyzeImageResponseBody
	GetInputTokens() *int32
	SetOutputTokens(v int32) *GeneralAnalyzeImageResponseBody
	GetOutputTokens() *int32
	SetRequestId(v string) *GeneralAnalyzeImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GeneralAnalyzeImageResponseBody
	GetSuccess() *bool
	SetText(v string) *GeneralAnalyzeImageResponseBody
	GetText() *string
	SetTotalTokens(v int32) *GeneralAnalyzeImageResponseBody
	GetTotalTokens() *int32
}

type GeneralAnalyzeImageResponseBody struct {
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 1000
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 2000
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D718325-92F9-5588-803B-C4A69A5F0AE1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
	Text    *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// 3000
	TotalTokens *int32 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GeneralAnalyzeImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GeneralAnalyzeImageResponseBody) GoString() string {
	return s.String()
}

func (s *GeneralAnalyzeImageResponseBody) GetFinishReason() *string {
	return s.FinishReason
}

func (s *GeneralAnalyzeImageResponseBody) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *GeneralAnalyzeImageResponseBody) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *GeneralAnalyzeImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GeneralAnalyzeImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GeneralAnalyzeImageResponseBody) GetText() *string {
	return s.Text
}

func (s *GeneralAnalyzeImageResponseBody) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *GeneralAnalyzeImageResponseBody) SetFinishReason(v string) *GeneralAnalyzeImageResponseBody {
	s.FinishReason = &v
	return s
}

func (s *GeneralAnalyzeImageResponseBody) SetInputTokens(v int32) *GeneralAnalyzeImageResponseBody {
	s.InputTokens = &v
	return s
}

func (s *GeneralAnalyzeImageResponseBody) SetOutputTokens(v int32) *GeneralAnalyzeImageResponseBody {
	s.OutputTokens = &v
	return s
}

func (s *GeneralAnalyzeImageResponseBody) SetRequestId(v string) *GeneralAnalyzeImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GeneralAnalyzeImageResponseBody) SetSuccess(v bool) *GeneralAnalyzeImageResponseBody {
	s.Success = &v
	return s
}

func (s *GeneralAnalyzeImageResponseBody) SetText(v string) *GeneralAnalyzeImageResponseBody {
	s.Text = &v
	return s
}

func (s *GeneralAnalyzeImageResponseBody) SetTotalTokens(v int32) *GeneralAnalyzeImageResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *GeneralAnalyzeImageResponseBody) Validate() error {
	return dara.Validate(s)
}
