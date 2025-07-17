// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskContentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskContentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskContentResponseBody
	GetSuccess() *bool
}

type UpdateTaskContentResponseBody struct {
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
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 39557312-28D5-528F-9554-80C0700EB489
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

func (s UpdateTaskContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskContentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskContentResponseBody) SetErrorCode(v string) *UpdateTaskContentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskContentResponseBody) SetErrorMessage(v string) *UpdateTaskContentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskContentResponseBody) SetRequestId(v string) *UpdateTaskContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskContentResponseBody) SetSuccess(v bool) *UpdateTaskContentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskContentResponseBody) Validate() error {
	return dara.Validate(s)
}
