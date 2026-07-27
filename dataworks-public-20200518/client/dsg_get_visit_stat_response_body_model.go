// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgGetVisitStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DsgGetVisitStatResponseBody
	GetAccessDeniedDetail() *string
	SetData(v interface{}) *DsgGetVisitStatResponseBody
	GetData() interface{}
	SetDynamicErrorCode(v string) *DsgGetVisitStatResponseBody
	GetDynamicErrorCode() *string
	SetDynamicErrorMessage(v string) *DsgGetVisitStatResponseBody
	GetDynamicErrorMessage() *string
	SetErrorCode(v string) *DsgGetVisitStatResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgGetVisitStatResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgGetVisitStatResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgGetVisitStatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgGetVisitStatResponseBody
	GetSuccess() *bool
}

type DsgGetVisitStatResponseBody struct {
	// The authentication error details.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The details of the file.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code. The value is the same as ErrorCode.
	//
	// example:
	//
	// 400
	DynamicErrorCode *string `json:"DynamicErrorCode,omitempty" xml:"DynamicErrorCode,omitempty"`
	// The error message. The value is the same as ErrorMessage.
	//
	// example:
	//
	// Missing parameter
	DynamicErrorMessage *string `json:"DynamicErrorMessage,omitempty" xml:"DynamicErrorMessage,omitempty"`
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
	// success
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
	// 952795279527ab****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgGetVisitStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgGetVisitStatResponseBody) GoString() string {
	return s.String()
}

func (s *DsgGetVisitStatResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DsgGetVisitStatResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DsgGetVisitStatResponseBody) GetDynamicErrorCode() *string {
	return s.DynamicErrorCode
}

func (s *DsgGetVisitStatResponseBody) GetDynamicErrorMessage() *string {
	return s.DynamicErrorMessage
}

func (s *DsgGetVisitStatResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgGetVisitStatResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgGetVisitStatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgGetVisitStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgGetVisitStatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgGetVisitStatResponseBody) SetAccessDeniedDetail(v string) *DsgGetVisitStatResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DsgGetVisitStatResponseBody) SetData(v interface{}) *DsgGetVisitStatResponseBody {
	s.Data = v
	return s
}

func (s *DsgGetVisitStatResponseBody) SetDynamicErrorCode(v string) *DsgGetVisitStatResponseBody {
	s.DynamicErrorCode = &v
	return s
}

func (s *DsgGetVisitStatResponseBody) SetDynamicErrorMessage(v string) *DsgGetVisitStatResponseBody {
	s.DynamicErrorMessage = &v
	return s
}

func (s *DsgGetVisitStatResponseBody) SetErrorCode(v string) *DsgGetVisitStatResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgGetVisitStatResponseBody) SetErrorMessage(v string) *DsgGetVisitStatResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgGetVisitStatResponseBody) SetHttpStatusCode(v int32) *DsgGetVisitStatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgGetVisitStatResponseBody) SetRequestId(v string) *DsgGetVisitStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgGetVisitStatResponseBody) SetSuccess(v bool) *DsgGetVisitStatResponseBody {
	s.Success = &v
	return s
}

func (s *DsgGetVisitStatResponseBody) Validate() error {
	return dara.Validate(s)
}
