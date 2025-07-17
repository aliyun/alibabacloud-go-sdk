// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstanceResponseBody
	GetSuccess() *bool
}

type DeleteInstanceResponseBody struct {
	// The error code that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B4B07137-F6AE-4756-8474-7F92BB6C4E04
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

func (s DeleteInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstanceResponseBody) SetErrorCode(v string) *DeleteInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetErrorMessage(v string) *DeleteInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
