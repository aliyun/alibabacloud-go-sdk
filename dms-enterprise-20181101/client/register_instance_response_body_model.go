// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RegisterInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RegisterInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RegisterInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RegisterInstanceResponseBody
	GetSuccess() *bool
}

type RegisterInstanceResponseBody struct {
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
	// F4E2A94B-604F-43FF-93E7-F4EE3DCF412E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true:*	- The request was successful.
	//
	// 	- **false:*	- The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RegisterInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RegisterInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RegisterInstanceResponseBody) SetErrorCode(v string) *RegisterInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterInstanceResponseBody) SetErrorMessage(v string) *RegisterInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterInstanceResponseBody) SetRequestId(v string) *RegisterInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterInstanceResponseBody) SetSuccess(v bool) *RegisterInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RegisterInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
