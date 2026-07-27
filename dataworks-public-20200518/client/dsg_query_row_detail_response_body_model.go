// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryRowDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DsgQueryRowDetailResponseBody
	GetAccessDeniedDetail() *string
	SetData(v interface{}) *DsgQueryRowDetailResponseBody
	GetData() interface{}
	SetDynamicErrorCode(v string) *DsgQueryRowDetailResponseBody
	GetDynamicErrorCode() *string
	SetDynamicErrorMessage(v string) *DsgQueryRowDetailResponseBody
	GetDynamicErrorMessage() *string
	SetErrorCode(v string) *DsgQueryRowDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgQueryRowDetailResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgQueryRowDetailResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgQueryRowDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgQueryRowDetailResponseBody
	GetSuccess() *bool
}

type DsgQueryRowDetailResponseBody struct {
	// The details of the access denial. This field is returned only when RAM authentication fails.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether the deletion is successful.
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
	// 401
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
	// 10000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the workspace information is queried successfully.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgQueryRowDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryRowDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DsgQueryRowDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DsgQueryRowDetailResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DsgQueryRowDetailResponseBody) GetDynamicErrorCode() *string {
	return s.DynamicErrorCode
}

func (s *DsgQueryRowDetailResponseBody) GetDynamicErrorMessage() *string {
	return s.DynamicErrorMessage
}

func (s *DsgQueryRowDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgQueryRowDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgQueryRowDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgQueryRowDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgQueryRowDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgQueryRowDetailResponseBody) SetAccessDeniedDetail(v string) *DsgQueryRowDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DsgQueryRowDetailResponseBody) SetData(v interface{}) *DsgQueryRowDetailResponseBody {
	s.Data = v
	return s
}

func (s *DsgQueryRowDetailResponseBody) SetDynamicErrorCode(v string) *DsgQueryRowDetailResponseBody {
	s.DynamicErrorCode = &v
	return s
}

func (s *DsgQueryRowDetailResponseBody) SetDynamicErrorMessage(v string) *DsgQueryRowDetailResponseBody {
	s.DynamicErrorMessage = &v
	return s
}

func (s *DsgQueryRowDetailResponseBody) SetErrorCode(v string) *DsgQueryRowDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgQueryRowDetailResponseBody) SetErrorMessage(v string) *DsgQueryRowDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgQueryRowDetailResponseBody) SetHttpStatusCode(v int32) *DsgQueryRowDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgQueryRowDetailResponseBody) SetRequestId(v string) *DsgQueryRowDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgQueryRowDetailResponseBody) SetSuccess(v bool) *DsgQueryRowDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DsgQueryRowDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
