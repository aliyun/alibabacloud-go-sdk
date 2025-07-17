// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RegisterUserResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RegisterUserResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RegisterUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RegisterUserResponseBody
	GetSuccess() *bool
}

type RegisterUserResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The specified user already exists.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34E01EDD-6A16-4CF0-9541-C644D1BE01AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterUserResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RegisterUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RegisterUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RegisterUserResponseBody) SetErrorCode(v string) *RegisterUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterUserResponseBody) SetErrorMessage(v string) *RegisterUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterUserResponseBody) SetRequestId(v string) *RegisterUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterUserResponseBody) SetSuccess(v bool) *RegisterUserResponseBody {
	s.Success = &v
	return s
}

func (s *RegisterUserResponseBody) Validate() error {
	return dara.Validate(s)
}
