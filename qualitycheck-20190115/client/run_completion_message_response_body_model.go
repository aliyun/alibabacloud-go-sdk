// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunCompletionMessageResponseBody
	GetCode() *string
	SetData(v *RunCompletionMessageResponseBodyData) *RunCompletionMessageResponseBody
	GetData() *RunCompletionMessageResponseBodyData
	SetMessage(v string) *RunCompletionMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunCompletionMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunCompletionMessageResponseBody
	GetSuccess() *bool
}

type RunCompletionMessageResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RunCompletionMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunCompletionMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunCompletionMessageResponseBody) GetData() *RunCompletionMessageResponseBodyData {
	return s.Data
}

func (s *RunCompletionMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunCompletionMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCompletionMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunCompletionMessageResponseBody) SetCode(v string) *RunCompletionMessageResponseBody {
	s.Code = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetData(v *RunCompletionMessageResponseBodyData) *RunCompletionMessageResponseBody {
	s.Data = v
	return s
}

func (s *RunCompletionMessageResponseBody) SetMessage(v string) *RunCompletionMessageResponseBody {
	s.Message = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetRequestId(v string) *RunCompletionMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetSuccess(v bool) *RunCompletionMessageResponseBody {
	s.Success = &v
	return s
}

func (s *RunCompletionMessageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunCompletionMessageResponseBodyData struct {
	// example:
	//
	// stop
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	LlmRequestId *string `json:"LlmRequestId,omitempty" xml:"LlmRequestId,omitempty"`
	// example:
	//
	// 200
	OutputTokens *int64  `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	Text         *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 300
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunCompletionMessageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageResponseBodyData) GetFinishReason() *string {
	return s.FinishReason
}

func (s *RunCompletionMessageResponseBodyData) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunCompletionMessageResponseBodyData) GetLlmRequestId() *string {
	return s.LlmRequestId
}

func (s *RunCompletionMessageResponseBodyData) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunCompletionMessageResponseBodyData) GetText() *string {
	return s.Text
}

func (s *RunCompletionMessageResponseBodyData) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunCompletionMessageResponseBodyData) SetFinishReason(v string) *RunCompletionMessageResponseBodyData {
	s.FinishReason = &v
	return s
}

func (s *RunCompletionMessageResponseBodyData) SetInputTokens(v int64) *RunCompletionMessageResponseBodyData {
	s.InputTokens = &v
	return s
}

func (s *RunCompletionMessageResponseBodyData) SetLlmRequestId(v string) *RunCompletionMessageResponseBodyData {
	s.LlmRequestId = &v
	return s
}

func (s *RunCompletionMessageResponseBodyData) SetOutputTokens(v int64) *RunCompletionMessageResponseBodyData {
	s.OutputTokens = &v
	return s
}

func (s *RunCompletionMessageResponseBodyData) SetText(v string) *RunCompletionMessageResponseBodyData {
	s.Text = &v
	return s
}

func (s *RunCompletionMessageResponseBodyData) SetTotalTokens(v int64) *RunCompletionMessageResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *RunCompletionMessageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
