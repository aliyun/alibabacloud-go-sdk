// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteTaskResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTaskResponseBody
	GetSuccess() *bool
}

type DeleteTaskResponseBody struct {
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
	// Unknown server error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3B460770-40D9-5F07-B68A-173D1D708B72
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

func (s DeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTaskResponseBody) SetErrorCode(v string) *DeleteTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTaskResponseBody) SetErrorMessage(v string) *DeleteTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTaskResponseBody) SetRequestId(v string) *DeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskResponseBody) SetSuccess(v bool) *DeleteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
