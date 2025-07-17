// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateUdfFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateUdfFileResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateUdfFileResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateUdfFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUdfFileResponseBody
	GetSuccess() *bool
}

type UpdateUdfFileResponseBody struct {
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUdfFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUdfFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateUdfFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateUdfFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateUdfFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUdfFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUdfFileResponseBody) SetErrorCode(v string) *UpdateUdfFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateUdfFileResponseBody) SetErrorMessage(v string) *UpdateUdfFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateUdfFileResponseBody) SetHttpStatusCode(v int32) *UpdateUdfFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateUdfFileResponseBody) SetRequestId(v string) *UpdateUdfFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUdfFileResponseBody) SetSuccess(v bool) *UpdateUdfFileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUdfFileResponseBody) Validate() error {
	return dara.Validate(s)
}
