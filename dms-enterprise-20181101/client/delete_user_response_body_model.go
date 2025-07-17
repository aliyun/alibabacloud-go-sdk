// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteUserResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteUserResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUserResponseBody
	GetSuccess() *bool
}

type DeleteUserResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserResponseBody) SetErrorCode(v string) *DeleteUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetErrorMessage(v string) *DeleteUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserResponseBody) Validate() error {
	return dara.Validate(s)
}
