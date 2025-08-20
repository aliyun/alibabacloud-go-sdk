// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeAudioSyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFinishReason(v string) *AnalyzeAudioSyncResponseBody
	GetFinishReason() *string
	SetInputTokens(v string) *AnalyzeAudioSyncResponseBody
	GetInputTokens() *string
	SetOutputTokens(v string) *AnalyzeAudioSyncResponseBody
	GetOutputTokens() *string
	SetRequestId(v string) *AnalyzeAudioSyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AnalyzeAudioSyncResponseBody
	GetSuccess() *bool
	SetText(v string) *AnalyzeAudioSyncResponseBody
	GetText() *string
	SetTotalTokens(v string) *AnalyzeAudioSyncResponseBody
	GetTotalTokens() *string
}

type AnalyzeAudioSyncResponseBody struct {
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
	// 968A8634-FA2C-5381-9B3E-*******F
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

func (s AnalyzeAudioSyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncResponseBody) GetFinishReason() *string {
	return s.FinishReason
}

func (s *AnalyzeAudioSyncResponseBody) GetInputTokens() *string {
	return s.InputTokens
}

func (s *AnalyzeAudioSyncResponseBody) GetOutputTokens() *string {
	return s.OutputTokens
}

func (s *AnalyzeAudioSyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnalyzeAudioSyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AnalyzeAudioSyncResponseBody) GetText() *string {
	return s.Text
}

func (s *AnalyzeAudioSyncResponseBody) GetTotalTokens() *string {
	return s.TotalTokens
}

func (s *AnalyzeAudioSyncResponseBody) SetFinishReason(v string) *AnalyzeAudioSyncResponseBody {
	s.FinishReason = &v
	return s
}

func (s *AnalyzeAudioSyncResponseBody) SetInputTokens(v string) *AnalyzeAudioSyncResponseBody {
	s.InputTokens = &v
	return s
}

func (s *AnalyzeAudioSyncResponseBody) SetOutputTokens(v string) *AnalyzeAudioSyncResponseBody {
	s.OutputTokens = &v
	return s
}

func (s *AnalyzeAudioSyncResponseBody) SetRequestId(v string) *AnalyzeAudioSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnalyzeAudioSyncResponseBody) SetSuccess(v bool) *AnalyzeAudioSyncResponseBody {
	s.Success = &v
	return s
}

func (s *AnalyzeAudioSyncResponseBody) SetText(v string) *AnalyzeAudioSyncResponseBody {
	s.Text = &v
	return s
}

func (s *AnalyzeAudioSyncResponseBody) SetTotalTokens(v string) *AnalyzeAudioSyncResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *AnalyzeAudioSyncResponseBody) Validate() error {
	return dara.Validate(s)
}
