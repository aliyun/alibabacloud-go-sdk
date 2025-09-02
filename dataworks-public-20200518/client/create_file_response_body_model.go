// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateFileResponseBody
	GetData() *int64
	SetErrorCode(v string) *CreateFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateFileResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateFileResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFileResponseBody
	GetSuccess() *bool
}

type CreateFileResponseBody struct {
	// The ID of the file that was created.
	//
	// example:
	//
	// 1000001
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFileResponseBody) SetData(v int64) *CreateFileResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFileResponseBody) SetErrorCode(v string) *CreateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFileResponseBody) SetErrorMessage(v string) *CreateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFileResponseBody) SetHttpStatusCode(v int32) *CreateFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFileResponseBody) SetRequestId(v string) *CreateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileResponseBody) SetSuccess(v bool) *CreateFileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFileResponseBody) Validate() error {
	return dara.Validate(s)
}
