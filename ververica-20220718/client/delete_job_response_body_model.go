// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteJobResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteJobResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteJobResponseBody
	GetSuccess() *bool
}

type DeleteJobResponseBody struct {
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

func (s DeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteJobResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteJobResponseBody) SetErrorCode(v string) *DeleteJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteJobResponseBody) SetErrorMessage(v string) *DeleteJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteJobResponseBody) SetHttpCode(v int32) *DeleteJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobResponseBody) SetSuccess(v bool) *DeleteJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}
