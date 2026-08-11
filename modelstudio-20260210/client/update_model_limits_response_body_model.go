// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelLimitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateModelLimitsResponseBody
	GetCode() *string
	SetErrorMessage(v string) *UpdateModelLimitsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateModelLimitsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateModelLimitsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateModelLimitsResponseBody
	GetSuccess() *bool
}

type UpdateModelLimitsResponseBody struct {
	// The error code. This parameter is empty when the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// A workspace with ID beb173d2361941 does not exist.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// E4C14AE6-E987-5C2F-9230-9960AB48F4F2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateModelLimitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelLimitsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelLimitsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelLimitsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateModelLimitsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateModelLimitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelLimitsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModelLimitsResponseBody) SetCode(v string) *UpdateModelLimitsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelLimitsResponseBody) SetErrorMessage(v string) *UpdateModelLimitsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateModelLimitsResponseBody) SetHttpStatusCode(v int32) *UpdateModelLimitsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateModelLimitsResponseBody) SetRequestId(v string) *UpdateModelLimitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelLimitsResponseBody) SetSuccess(v bool) *UpdateModelLimitsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModelLimitsResponseBody) Validate() error {
	return dara.Validate(s)
}
