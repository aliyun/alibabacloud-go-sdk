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
	SetRagStatus(v string) *RunCompletionResponseBody
	GetRagStatus() *string
	SetTotalTokens(v string) *RunCompletionResponseBody
	GetTotalTokens() *string
	SetUsage(v *RunCompletionResponseBodyUsage) *RunCompletionResponseBody
	GetUsage() *RunCompletionResponseBodyUsage
}

type RunCompletionResponseBody struct {
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Text         *string                         `json:"Text,omitempty" xml:"Text,omitempty"`
	InputTokens  *string                         `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *string                         `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	RagStatus    *string                         `json:"ragStatus,omitempty" xml:"ragStatus,omitempty"`
	TotalTokens  *string                         `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	Usage        *RunCompletionResponseBodyUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
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

func (s *RunCompletionResponseBody) GetRagStatus() *string {
	return s.RagStatus
}

func (s *RunCompletionResponseBody) GetTotalTokens() *string {
	return s.TotalTokens
}

func (s *RunCompletionResponseBody) GetUsage() *RunCompletionResponseBodyUsage {
	return s.Usage
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

func (s *RunCompletionResponseBody) SetRagStatus(v string) *RunCompletionResponseBody {
	s.RagStatus = &v
	return s
}

func (s *RunCompletionResponseBody) SetTotalTokens(v string) *RunCompletionResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *RunCompletionResponseBody) SetUsage(v *RunCompletionResponseBodyUsage) *RunCompletionResponseBody {
	s.Usage = v
	return s
}

func (s *RunCompletionResponseBody) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunCompletionResponseBodyUsage struct {
	Rag *RunCompletionResponseBodyUsageRag `json:"rag,omitempty" xml:"rag,omitempty" type:"Struct"`
}

func (s RunCompletionResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunCompletionResponseBodyUsage) GetRag() *RunCompletionResponseBodyUsageRag {
	return s.Rag
}

func (s *RunCompletionResponseBodyUsage) SetRag(v *RunCompletionResponseBodyUsageRag) *RunCompletionResponseBodyUsage {
	s.Rag = v
	return s
}

func (s *RunCompletionResponseBodyUsage) Validate() error {
	if s.Rag != nil {
		if err := s.Rag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunCompletionResponseBodyUsageRag struct {
	Adaptive      *RunCompletionResponseBodyUsageRagAdaptive      `json:"adaptive,omitempty" xml:"adaptive,omitempty" type:"Struct"`
	DialogSummary *RunCompletionResponseBodyUsageRagDialogSummary `json:"dialogSummary,omitempty" xml:"dialogSummary,omitempty" type:"Struct"`
}

func (s RunCompletionResponseBodyUsageRag) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionResponseBodyUsageRag) GoString() string {
	return s.String()
}

func (s *RunCompletionResponseBodyUsageRag) GetAdaptive() *RunCompletionResponseBodyUsageRagAdaptive {
	return s.Adaptive
}

func (s *RunCompletionResponseBodyUsageRag) GetDialogSummary() *RunCompletionResponseBodyUsageRagDialogSummary {
	return s.DialogSummary
}

func (s *RunCompletionResponseBodyUsageRag) SetAdaptive(v *RunCompletionResponseBodyUsageRagAdaptive) *RunCompletionResponseBodyUsageRag {
	s.Adaptive = v
	return s
}

func (s *RunCompletionResponseBodyUsageRag) SetDialogSummary(v *RunCompletionResponseBodyUsageRagDialogSummary) *RunCompletionResponseBodyUsageRag {
	s.DialogSummary = v
	return s
}

func (s *RunCompletionResponseBodyUsageRag) Validate() error {
	if s.Adaptive != nil {
		if err := s.Adaptive.Validate(); err != nil {
			return err
		}
	}
	if s.DialogSummary != nil {
		if err := s.DialogSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunCompletionResponseBodyUsageRagAdaptive struct {
	InputTokens  *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	InvokeCount  *int32 `json:"invokeCount,omitempty" xml:"invokeCount,omitempty"`
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
}

func (s RunCompletionResponseBodyUsageRagAdaptive) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionResponseBodyUsageRagAdaptive) GoString() string {
	return s.String()
}

func (s *RunCompletionResponseBodyUsageRagAdaptive) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *RunCompletionResponseBodyUsageRagAdaptive) GetInvokeCount() *int32 {
	return s.InvokeCount
}

func (s *RunCompletionResponseBodyUsageRagAdaptive) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *RunCompletionResponseBodyUsageRagAdaptive) SetInputTokens(v int32) *RunCompletionResponseBodyUsageRagAdaptive {
	s.InputTokens = &v
	return s
}

func (s *RunCompletionResponseBodyUsageRagAdaptive) SetInvokeCount(v int32) *RunCompletionResponseBodyUsageRagAdaptive {
	s.InvokeCount = &v
	return s
}

func (s *RunCompletionResponseBodyUsageRagAdaptive) SetOutputTokens(v int32) *RunCompletionResponseBodyUsageRagAdaptive {
	s.OutputTokens = &v
	return s
}

func (s *RunCompletionResponseBodyUsageRagAdaptive) Validate() error {
	return dara.Validate(s)
}

type RunCompletionResponseBodyUsageRagDialogSummary struct {
	InputTokens  *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	InvokeCount  *int32 `json:"invokeCount,omitempty" xml:"invokeCount,omitempty"`
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
}

func (s RunCompletionResponseBodyUsageRagDialogSummary) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionResponseBodyUsageRagDialogSummary) GoString() string {
	return s.String()
}

func (s *RunCompletionResponseBodyUsageRagDialogSummary) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *RunCompletionResponseBodyUsageRagDialogSummary) GetInvokeCount() *int32 {
	return s.InvokeCount
}

func (s *RunCompletionResponseBodyUsageRagDialogSummary) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *RunCompletionResponseBodyUsageRagDialogSummary) SetInputTokens(v int32) *RunCompletionResponseBodyUsageRagDialogSummary {
	s.InputTokens = &v
	return s
}

func (s *RunCompletionResponseBodyUsageRagDialogSummary) SetInvokeCount(v int32) *RunCompletionResponseBodyUsageRagDialogSummary {
	s.InvokeCount = &v
	return s
}

func (s *RunCompletionResponseBodyUsageRagDialogSummary) SetOutputTokens(v int32) *RunCompletionResponseBodyUsageRagDialogSummary {
	s.OutputTokens = &v
	return s
}

func (s *RunCompletionResponseBodyUsageRagDialogSummary) Validate() error {
	return dara.Validate(s)
}
