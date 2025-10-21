// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteUdfFunctionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteUdfFunctionResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteUdfFunctionResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteUdfFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUdfFunctionResponseBody
	GetSuccess() *bool
}

type DeleteUdfFunctionResponseBody struct {
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteUdfFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUdfFunctionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteUdfFunctionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteUdfFunctionResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteUdfFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUdfFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUdfFunctionResponseBody) SetErrorCode(v string) *DeleteUdfFunctionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUdfFunctionResponseBody) SetErrorMessage(v string) *DeleteUdfFunctionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUdfFunctionResponseBody) SetHttpCode(v int32) *DeleteUdfFunctionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteUdfFunctionResponseBody) SetRequestId(v string) *DeleteUdfFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUdfFunctionResponseBody) SetSuccess(v bool) *DeleteUdfFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUdfFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
