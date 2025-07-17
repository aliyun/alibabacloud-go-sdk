// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOwnersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SetOwnersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SetOwnersResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SetOwnersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetOwnersResponseBody
	GetSuccess() *bool
}

type SetOwnersResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A99CD576-1E18-4E86-931E-C3CCE56DC030
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetOwnersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetOwnersResponseBody) GoString() string {
	return s.String()
}

func (s *SetOwnersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SetOwnersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SetOwnersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetOwnersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetOwnersResponseBody) SetErrorCode(v string) *SetOwnersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetOwnersResponseBody) SetErrorMessage(v string) *SetOwnersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetOwnersResponseBody) SetRequestId(v string) *SetOwnersResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetOwnersResponseBody) SetSuccess(v bool) *SetOwnersResponseBody {
	s.Success = &v
	return s
}

func (s *SetOwnersResponseBody) Validate() error {
	return dara.Validate(s)
}
