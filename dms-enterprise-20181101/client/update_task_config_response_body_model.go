// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskConfigResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskConfigResponseBody
	GetSuccess() *bool
}

type UpdateTaskConfigResponseBody struct {
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
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskConfigResponseBody) SetErrorCode(v string) *UpdateTaskConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskConfigResponseBody) SetErrorMessage(v string) *UpdateTaskConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskConfigResponseBody) SetRequestId(v string) *UpdateTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskConfigResponseBody) SetSuccess(v bool) *UpdateTaskConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
