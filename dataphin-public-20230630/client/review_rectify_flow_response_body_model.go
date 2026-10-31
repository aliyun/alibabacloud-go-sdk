// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewRectifyFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReviewRectifyFlowResponseBody
	GetCode() *string
	SetData(v bool) *ReviewRectifyFlowResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *ReviewRectifyFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ReviewRectifyFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReviewRectifyFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReviewRectifyFlowResponseBody
	GetSuccess() *bool
}

type ReviewRectifyFlowResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation is complete. A value of true indicates that the operation is complete.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
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

func (s ReviewRectifyFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReviewRectifyFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ReviewRectifyFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReviewRectifyFlowResponseBody) GetData() *bool {
	return s.Data
}

func (s *ReviewRectifyFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReviewRectifyFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReviewRectifyFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReviewRectifyFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReviewRectifyFlowResponseBody) SetCode(v string) *ReviewRectifyFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ReviewRectifyFlowResponseBody) SetData(v bool) *ReviewRectifyFlowResponseBody {
	s.Data = &v
	return s
}

func (s *ReviewRectifyFlowResponseBody) SetHttpStatusCode(v int32) *ReviewRectifyFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReviewRectifyFlowResponseBody) SetMessage(v string) *ReviewRectifyFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ReviewRectifyFlowResponseBody) SetRequestId(v string) *ReviewRectifyFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReviewRectifyFlowResponseBody) SetSuccess(v bool) *ReviewRectifyFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ReviewRectifyFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
