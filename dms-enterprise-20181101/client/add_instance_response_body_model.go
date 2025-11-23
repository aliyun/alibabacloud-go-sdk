// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddInstanceResponseBody
	GetSuccess() *bool
}

type AddInstanceResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
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

func (s AddInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AddInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddInstanceResponseBody) SetErrorCode(v string) *AddInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddInstanceResponseBody) SetErrorMessage(v string) *AddInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddInstanceResponseBody) SetRequestId(v string) *AddInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddInstanceResponseBody) SetSuccess(v bool) *AddInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *AddInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
