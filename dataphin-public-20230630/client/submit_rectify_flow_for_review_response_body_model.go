// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRectifyFlowForReviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitRectifyFlowForReviewResponseBody
	GetCode() *string
	SetData(v bool) *SubmitRectifyFlowForReviewResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *SubmitRectifyFlowForReviewResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitRectifyFlowForReviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitRectifyFlowForReviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitRectifyFlowForReviewResponseBody
	GetSuccess() *bool
}

type SubmitRectifyFlowForReviewResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation is completed. The value is true when the operation is completed.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitRectifyFlowForReviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitRectifyFlowForReviewResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitRectifyFlowForReviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitRectifyFlowForReviewResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitRectifyFlowForReviewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitRectifyFlowForReviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitRectifyFlowForReviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitRectifyFlowForReviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitRectifyFlowForReviewResponseBody) SetCode(v string) *SubmitRectifyFlowForReviewResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitRectifyFlowForReviewResponseBody) SetData(v bool) *SubmitRectifyFlowForReviewResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitRectifyFlowForReviewResponseBody) SetHttpStatusCode(v int32) *SubmitRectifyFlowForReviewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitRectifyFlowForReviewResponseBody) SetMessage(v string) *SubmitRectifyFlowForReviewResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitRectifyFlowForReviewResponseBody) SetRequestId(v string) *SubmitRectifyFlowForReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitRectifyFlowForReviewResponseBody) SetSuccess(v bool) *SubmitRectifyFlowForReviewResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitRectifyFlowForReviewResponseBody) Validate() error {
	return dara.Validate(s)
}
