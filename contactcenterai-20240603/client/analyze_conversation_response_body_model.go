// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AnalyzeConversationResponseBody
	GetErrorCode() *string
	SetErrorInfo(v string) *AnalyzeConversationResponseBody
	GetErrorInfo() *string
	SetFinishReason(v string) *AnalyzeConversationResponseBody
	GetFinishReason() *string
	SetInputTokens(v string) *AnalyzeConversationResponseBody
	GetInputTokens() *string
	SetOutputTokens(v string) *AnalyzeConversationResponseBody
	GetOutputTokens() *string
	SetRequestId(v string) *AnalyzeConversationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AnalyzeConversationResponseBody
	GetSuccess() *bool
	SetText(v string) *AnalyzeConversationResponseBody
	GetText() *string
	SetTotalTokens(v string) *AnalyzeConversationResponseBody
	GetTotalTokens() *string
}

type AnalyzeConversationResponseBody struct {
	// example:
	//
	// InvalidUser.NotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorInfo *string `json:"errorInfo,omitempty" xml:"errorInfo,omitempty"`
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	InputTokens  *string `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *string `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 968A8634-FA2C-5381-9B3E-C552DED7E8BF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success     *bool   `json:"success,omitempty" xml:"success,omitempty"`
	Text        *string `json:"text,omitempty" xml:"text,omitempty"`
	TotalTokens *string `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s AnalyzeConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AnalyzeConversationResponseBody) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *AnalyzeConversationResponseBody) GetFinishReason() *string {
	return s.FinishReason
}

func (s *AnalyzeConversationResponseBody) GetInputTokens() *string {
	return s.InputTokens
}

func (s *AnalyzeConversationResponseBody) GetOutputTokens() *string {
	return s.OutputTokens
}

func (s *AnalyzeConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnalyzeConversationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AnalyzeConversationResponseBody) GetText() *string {
	return s.Text
}

func (s *AnalyzeConversationResponseBody) GetTotalTokens() *string {
	return s.TotalTokens
}

func (s *AnalyzeConversationResponseBody) SetErrorCode(v string) *AnalyzeConversationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetErrorInfo(v string) *AnalyzeConversationResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetFinishReason(v string) *AnalyzeConversationResponseBody {
	s.FinishReason = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetInputTokens(v string) *AnalyzeConversationResponseBody {
	s.InputTokens = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetOutputTokens(v string) *AnalyzeConversationResponseBody {
	s.OutputTokens = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetRequestId(v string) *AnalyzeConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetSuccess(v bool) *AnalyzeConversationResponseBody {
	s.Success = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetText(v string) *AnalyzeConversationResponseBody {
	s.Text = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetTotalTokens(v string) *AnalyzeConversationResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *AnalyzeConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
