// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteVariableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteVariableResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteVariableResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteVariableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteVariableResponseBody
	GetSuccess() *bool
}

type DeleteVariableResponseBody struct {
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
	// The value was fixed to 200.
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

func (s DeleteVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVariableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteVariableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteVariableResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVariableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteVariableResponseBody) SetErrorCode(v string) *DeleteVariableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteVariableResponseBody) SetErrorMessage(v string) *DeleteVariableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteVariableResponseBody) SetHttpCode(v int32) *DeleteVariableResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteVariableResponseBody) SetRequestId(v string) *DeleteVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVariableResponseBody) SetSuccess(v bool) *DeleteVariableResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
