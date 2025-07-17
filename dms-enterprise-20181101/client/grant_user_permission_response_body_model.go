// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantUserPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GrantUserPermissionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GrantUserPermissionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GrantUserPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GrantUserPermissionResponseBody
	GetSuccess() *bool
}

type GrantUserPermissionResponseBody struct {
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantUserPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantUserPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GrantUserPermissionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GrantUserPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantUserPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GrantUserPermissionResponseBody) SetErrorCode(v string) *GrantUserPermissionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GrantUserPermissionResponseBody) SetErrorMessage(v string) *GrantUserPermissionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GrantUserPermissionResponseBody) SetRequestId(v string) *GrantUserPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantUserPermissionResponseBody) SetSuccess(v bool) *GrantUserPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *GrantUserPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
