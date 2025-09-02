// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuccessInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetSuccessInstanceResponseBody
	GetData() *bool
	SetErrorCode(v string) *SetSuccessInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SetSuccessInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *SetSuccessInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SetSuccessInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetSuccessInstanceResponseBody
	GetSuccess() *bool
}

type SetSuccessInstanceResponseBody struct {
	// Indicates whether result details are returned.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetSuccessInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSuccessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SetSuccessInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetSuccessInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SetSuccessInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SetSuccessInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SetSuccessInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSuccessInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetSuccessInstanceResponseBody) SetData(v bool) *SetSuccessInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *SetSuccessInstanceResponseBody) SetErrorCode(v string) *SetSuccessInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetSuccessInstanceResponseBody) SetErrorMessage(v string) *SetSuccessInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetSuccessInstanceResponseBody) SetHttpStatusCode(v int32) *SetSuccessInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SetSuccessInstanceResponseBody) SetRequestId(v string) *SetSuccessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSuccessInstanceResponseBody) SetSuccess(v bool) *SetSuccessInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *SetSuccessInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
