// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCutQuestionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CutQuestionsResponseBody
	GetCode() *string
	SetData(v string) *CutQuestionsResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CutQuestionsResponseBody
	GetHttpStatusCode() *int32
	SetInputTokens(v int32) *CutQuestionsResponseBody
	GetInputTokens() *int32
	SetMessage(v string) *CutQuestionsResponseBody
	GetMessage() *string
	SetOutputTokens(v int32) *CutQuestionsResponseBody
	GetOutputTokens() *int32
	SetRequestId(v string) *CutQuestionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CutQuestionsResponseBody
	GetSuccess() *bool
}

type CutQuestionsResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	InputTokens    *int32  `json:"input_tokens,omitempty" xml:"input_tokens,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	OutputTokens   *int32  `json:"output_tokens,omitempty" xml:"output_tokens,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CutQuestionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CutQuestionsResponseBody) GoString() string {
	return s.String()
}

func (s *CutQuestionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CutQuestionsResponseBody) GetData() *string {
	return s.Data
}

func (s *CutQuestionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CutQuestionsResponseBody) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *CutQuestionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CutQuestionsResponseBody) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *CutQuestionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CutQuestionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CutQuestionsResponseBody) SetCode(v string) *CutQuestionsResponseBody {
	s.Code = &v
	return s
}

func (s *CutQuestionsResponseBody) SetData(v string) *CutQuestionsResponseBody {
	s.Data = &v
	return s
}

func (s *CutQuestionsResponseBody) SetHttpStatusCode(v int32) *CutQuestionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CutQuestionsResponseBody) SetInputTokens(v int32) *CutQuestionsResponseBody {
	s.InputTokens = &v
	return s
}

func (s *CutQuestionsResponseBody) SetMessage(v string) *CutQuestionsResponseBody {
	s.Message = &v
	return s
}

func (s *CutQuestionsResponseBody) SetOutputTokens(v int32) *CutQuestionsResponseBody {
	s.OutputTokens = &v
	return s
}

func (s *CutQuestionsResponseBody) SetRequestId(v string) *CutQuestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CutQuestionsResponseBody) SetSuccess(v bool) *CutQuestionsResponseBody {
	s.Success = &v
	return s
}

func (s *CutQuestionsResponseBody) Validate() error {
	return dara.Validate(s)
}
