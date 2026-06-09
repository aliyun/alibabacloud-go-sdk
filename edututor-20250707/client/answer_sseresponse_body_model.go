// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerSSEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AnswerSSEResponseBody
	GetCode() *string
	SetData(v string) *AnswerSSEResponseBody
	GetData() *string
	SetFinishReason(v string) *AnswerSSEResponseBody
	GetFinishReason() *string
	SetHttpStatusCode(v int32) *AnswerSSEResponseBody
	GetHttpStatusCode() *int32
	SetInputTokens(v int32) *AnswerSSEResponseBody
	GetInputTokens() *int32
	SetMessage(v string) *AnswerSSEResponseBody
	GetMessage() *string
	SetOutputTokens(v int32) *AnswerSSEResponseBody
	GetOutputTokens() *int32
	SetRequestId(v string) *AnswerSSEResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AnswerSSEResponseBody
	GetSuccess() *bool
}

type AnswerSSEResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	FinishReason   *string `json:"finish_reason,omitempty" xml:"finish_reason,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	InputTokens    *int32  `json:"input_tokens,omitempty" xml:"input_tokens,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	OutputTokens   *int32  `json:"output_tokens,omitempty" xml:"output_tokens,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AnswerSSEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnswerSSEResponseBody) GoString() string {
	return s.String()
}

func (s *AnswerSSEResponseBody) GetCode() *string {
	return s.Code
}

func (s *AnswerSSEResponseBody) GetData() *string {
	return s.Data
}

func (s *AnswerSSEResponseBody) GetFinishReason() *string {
	return s.FinishReason
}

func (s *AnswerSSEResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AnswerSSEResponseBody) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *AnswerSSEResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AnswerSSEResponseBody) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *AnswerSSEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnswerSSEResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AnswerSSEResponseBody) SetCode(v string) *AnswerSSEResponseBody {
	s.Code = &v
	return s
}

func (s *AnswerSSEResponseBody) SetData(v string) *AnswerSSEResponseBody {
	s.Data = &v
	return s
}

func (s *AnswerSSEResponseBody) SetFinishReason(v string) *AnswerSSEResponseBody {
	s.FinishReason = &v
	return s
}

func (s *AnswerSSEResponseBody) SetHttpStatusCode(v int32) *AnswerSSEResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AnswerSSEResponseBody) SetInputTokens(v int32) *AnswerSSEResponseBody {
	s.InputTokens = &v
	return s
}

func (s *AnswerSSEResponseBody) SetMessage(v string) *AnswerSSEResponseBody {
	s.Message = &v
	return s
}

func (s *AnswerSSEResponseBody) SetOutputTokens(v int32) *AnswerSSEResponseBody {
	s.OutputTokens = &v
	return s
}

func (s *AnswerSSEResponseBody) SetRequestId(v string) *AnswerSSEResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnswerSSEResponseBody) SetSuccess(v bool) *AnswerSSEResponseBody {
	s.Success = &v
	return s
}

func (s *AnswerSSEResponseBody) Validate() error {
	return dara.Validate(s)
}
