// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DsgQueryDetailsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v interface{}) *DsgQueryDetailsResponseBody
	GetData() interface{}
	SetDynamicErrorCode(v string) *DsgQueryDetailsResponseBody
	GetDynamicErrorCode() *string
	SetDynamicErrorMessage(v string) *DsgQueryDetailsResponseBody
	GetDynamicErrorMessage() *string
	SetErrorCode(v string) *DsgQueryDetailsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgQueryDetailsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgQueryDetailsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgQueryDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgQueryDetailsResponseBody
	GetSuccess() *bool
}

type DsgQueryDetailsResponseBody struct {
	// The details of the access denial. This field is returned only when RAM authentication fails.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The operation result. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// select a;
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
	// 1031203110005
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
	// The request ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADE490
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgQueryDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DsgQueryDetailsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DsgQueryDetailsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DsgQueryDetailsResponseBody) GetDynamicErrorCode() *string {
	return s.DynamicErrorCode
}

func (s *DsgQueryDetailsResponseBody) GetDynamicErrorMessage() *string {
	return s.DynamicErrorMessage
}

func (s *DsgQueryDetailsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgQueryDetailsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgQueryDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgQueryDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgQueryDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgQueryDetailsResponseBody) SetAccessDeniedDetail(v string) *DsgQueryDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DsgQueryDetailsResponseBody) SetData(v interface{}) *DsgQueryDetailsResponseBody {
	s.Data = v
	return s
}

func (s *DsgQueryDetailsResponseBody) SetDynamicErrorCode(v string) *DsgQueryDetailsResponseBody {
	s.DynamicErrorCode = &v
	return s
}

func (s *DsgQueryDetailsResponseBody) SetDynamicErrorMessage(v string) *DsgQueryDetailsResponseBody {
	s.DynamicErrorMessage = &v
	return s
}

func (s *DsgQueryDetailsResponseBody) SetErrorCode(v string) *DsgQueryDetailsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgQueryDetailsResponseBody) SetErrorMessage(v string) *DsgQueryDetailsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgQueryDetailsResponseBody) SetHttpStatusCode(v int32) *DsgQueryDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgQueryDetailsResponseBody) SetRequestId(v string) *DsgQueryDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgQueryDetailsResponseBody) SetSuccess(v bool) *DsgQueryDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *DsgQueryDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}
