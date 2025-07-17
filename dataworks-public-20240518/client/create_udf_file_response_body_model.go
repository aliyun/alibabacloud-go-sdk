// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateUdfFileResponseBody
	GetData() *int64
	SetErrorCode(v string) *CreateUdfFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateUdfFileResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateUdfFileResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateUdfFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUdfFileResponseBody
	GetSuccess() *bool
}

type CreateUdfFileResponseBody struct {
	// example:
	//
	// 100000002
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateUdfFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUdfFileResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateUdfFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateUdfFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateUdfFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateUdfFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUdfFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUdfFileResponseBody) SetData(v int64) *CreateUdfFileResponseBody {
	s.Data = &v
	return s
}

func (s *CreateUdfFileResponseBody) SetErrorCode(v string) *CreateUdfFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUdfFileResponseBody) SetErrorMessage(v string) *CreateUdfFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUdfFileResponseBody) SetHttpStatusCode(v int32) *CreateUdfFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateUdfFileResponseBody) SetRequestId(v string) *CreateUdfFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUdfFileResponseBody) SetSuccess(v bool) *CreateUdfFileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUdfFileResponseBody) Validate() error {
	return dara.Validate(s)
}
