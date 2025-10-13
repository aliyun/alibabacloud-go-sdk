// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateInstanceNameResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateInstanceNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateInstanceNameResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UpdateInstanceNameResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateInstanceNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceNameResponseBody
	GetSuccess() *bool
}

type UpdateInstanceNameResponseBody struct {
	// The returned result, which indicates whether the operation was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6B55032-D41A-5FE0-9C07-8BD81C88422E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateInstanceNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateInstanceNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateInstanceNameResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceNameResponseBody) SetData(v bool) *UpdateInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrorCode(v string) *UpdateInstanceNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrorMessage(v string) *UpdateInstanceNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetHttpStatusCode(v string) *UpdateInstanceNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetRequestId(v string) *UpdateInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetSuccess(v bool) *UpdateInstanceNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}
