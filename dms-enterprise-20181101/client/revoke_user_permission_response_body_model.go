// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeUserPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RevokeUserPermissionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RevokeUserPermissionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RevokeUserPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevokeUserPermissionResponseBody
	GetSuccess() *bool
}

type RevokeUserPermissionResponseBody struct {
	// The error code that is returned.
	//
	// example:
	//
	// MissingUserId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// UserId is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A99CD576-1E18-4E86-931E-C3CCE56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeUserPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeUserPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeUserPermissionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RevokeUserPermissionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RevokeUserPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeUserPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeUserPermissionResponseBody) SetErrorCode(v string) *RevokeUserPermissionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RevokeUserPermissionResponseBody) SetErrorMessage(v string) *RevokeUserPermissionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RevokeUserPermissionResponseBody) SetRequestId(v string) *RevokeUserPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeUserPermissionResponseBody) SetSuccess(v bool) *RevokeUserPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeUserPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
