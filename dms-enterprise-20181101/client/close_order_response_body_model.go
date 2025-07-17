// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CloseOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CloseOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CloseOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CloseOrderResponseBody
	GetSuccess() *bool
}

type CloseOrderResponseBody struct {
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
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CloseOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CloseOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CloseOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloseOrderResponseBody) SetErrorCode(v string) *CloseOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CloseOrderResponseBody) SetErrorMessage(v string) *CloseOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CloseOrderResponseBody) SetRequestId(v string) *CloseOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseOrderResponseBody) SetSuccess(v bool) *CloseOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CloseOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
