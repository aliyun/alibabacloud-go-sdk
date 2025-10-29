// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateFileResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateFileResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFileResponseBody
	GetSuccess() *bool
}

type UpdateFileResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFGH-IJKLMNOPQ
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded. Valid values:
	//
	// 	- true: The call succeeded.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFileResponseBody) SetErrorCode(v string) *UpdateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFileResponseBody) SetErrorMessage(v string) *UpdateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFileResponseBody) SetHttpStatusCode(v int32) *UpdateFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateFileResponseBody) SetRequestId(v string) *UpdateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileResponseBody) SetSuccess(v bool) *UpdateFileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFileResponseBody) Validate() error {
	return dara.Validate(s)
}
