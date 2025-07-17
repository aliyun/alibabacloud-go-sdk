// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateUserResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateUserResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUserResponseBody
	GetSuccess() *bool
}

type UpdateUserResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E9BEBF41-4F69-4605-A5D5-A67955173941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUserResponseBody) SetErrorCode(v string) *UpdateUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateUserResponseBody) SetErrorMessage(v string) *UpdateUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v bool) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserResponseBody) Validate() error {
	return dara.Validate(s)
}
