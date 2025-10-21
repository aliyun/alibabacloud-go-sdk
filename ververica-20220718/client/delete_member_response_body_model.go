// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteMemberResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteMemberResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMemberResponseBody
	GetSuccess() *bool
}

type DeleteMemberResponseBody struct {
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

func (s DeleteMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteMemberResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMemberResponseBody) SetErrorCode(v string) *DeleteMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMemberResponseBody) SetErrorMessage(v string) *DeleteMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMemberResponseBody) SetHttpCode(v int32) *DeleteMemberResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteMemberResponseBody) SetRequestId(v string) *DeleteMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemberResponseBody) SetSuccess(v bool) *DeleteMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
