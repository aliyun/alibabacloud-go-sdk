// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIOrderApprovalCommentSSEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetAIOrderApprovalCommentSSEResponseBody
	GetData() *string
	SetErrorCode(v string) *GetAIOrderApprovalCommentSSEResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAIOrderApprovalCommentSSEResponseBody
	GetErrorMessage() *string
	SetOutput(v *GetAIOrderApprovalCommentSSEResponseBodyOutput) *GetAIOrderApprovalCommentSSEResponseBody
	GetOutput() *GetAIOrderApprovalCommentSSEResponseBodyOutput
	SetRequestId(v string) *GetAIOrderApprovalCommentSSEResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAIOrderApprovalCommentSSEResponseBody
	GetSuccess() *bool
}

type GetAIOrderApprovalCommentSSEResponseBody struct {
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Output       *GetAIOrderApprovalCommentSSEResponseBodyOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAIOrderApprovalCommentSSEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIOrderApprovalCommentSSEResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetOutput() *GetAIOrderApprovalCommentSSEResponseBodyOutput {
	return s.Output
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetData(v string) *GetAIOrderApprovalCommentSSEResponseBody {
	s.Data = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetErrorCode(v string) *GetAIOrderApprovalCommentSSEResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetErrorMessage(v string) *GetAIOrderApprovalCommentSSEResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetOutput(v *GetAIOrderApprovalCommentSSEResponseBodyOutput) *GetAIOrderApprovalCommentSSEResponseBody {
	s.Output = v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetRequestId(v string) *GetAIOrderApprovalCommentSSEResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetSuccess(v bool) *GetAIOrderApprovalCommentSSEResponseBody {
	s.Success = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAIOrderApprovalCommentSSEResponseBodyOutput struct {
	// example:
	//
	// {"approvalStatus":"建议拒绝","approvalSuggestion":"xxx","sessionId":"xxx"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetAIOrderApprovalCommentSSEResponseBodyOutput) String() string {
	return dara.Prettify(s)
}

func (s GetAIOrderApprovalCommentSSEResponseBodyOutput) GoString() string {
	return s.String()
}

func (s *GetAIOrderApprovalCommentSSEResponseBodyOutput) GetContent() *string {
	return s.Content
}

func (s *GetAIOrderApprovalCommentSSEResponseBodyOutput) SetContent(v string) *GetAIOrderApprovalCommentSSEResponseBodyOutput {
	s.Content = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBodyOutput) Validate() error {
	return dara.Validate(s)
}
