// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DisableUserResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DisableUserResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DisableUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableUserResponseBody
	GetSuccess() *bool
}

type DisableUserResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The specified user not exists.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34E01EDD-6A16-4CF0-9541-C644D1BE01AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableUserResponseBody) GoString() string {
	return s.String()
}

func (s *DisableUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DisableUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DisableUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableUserResponseBody) SetErrorCode(v string) *DisableUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DisableUserResponseBody) SetErrorMessage(v string) *DisableUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DisableUserResponseBody) SetRequestId(v string) *DisableUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableUserResponseBody) SetSuccess(v bool) *DisableUserResponseBody {
	s.Success = &v
	return s
}

func (s *DisableUserResponseBody) Validate() error {
	return dara.Validate(s)
}
