// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgGetVisitDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DsgGetVisitDetailResponseBody
	GetAccessDeniedDetail() *string
	SetData(v interface{}) *DsgGetVisitDetailResponseBody
	GetData() interface{}
	SetDynamicErrorCode(v string) *DsgGetVisitDetailResponseBody
	GetDynamicErrorCode() *string
	SetDynamicErrorMessage(v string) *DsgGetVisitDetailResponseBody
	GetDynamicErrorMessage() *string
	SetErrorCode(v string) *DsgGetVisitDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgGetVisitDetailResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgGetVisitDetailResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgGetVisitDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgGetVisitDetailResponseBody
	GetSuccess() *bool
}

type DsgGetVisitDetailResponseBody struct {
	// The details about the access denial. This field is returned only when RAM verification fails.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The details returned after the DeleteQualityEntity operation is called.
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
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. Used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 58D5334A-B013-430E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgGetVisitDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgGetVisitDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DsgGetVisitDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DsgGetVisitDetailResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DsgGetVisitDetailResponseBody) GetDynamicErrorCode() *string {
	return s.DynamicErrorCode
}

func (s *DsgGetVisitDetailResponseBody) GetDynamicErrorMessage() *string {
	return s.DynamicErrorMessage
}

func (s *DsgGetVisitDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgGetVisitDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgGetVisitDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgGetVisitDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgGetVisitDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgGetVisitDetailResponseBody) SetAccessDeniedDetail(v string) *DsgGetVisitDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DsgGetVisitDetailResponseBody) SetData(v interface{}) *DsgGetVisitDetailResponseBody {
	s.Data = v
	return s
}

func (s *DsgGetVisitDetailResponseBody) SetDynamicErrorCode(v string) *DsgGetVisitDetailResponseBody {
	s.DynamicErrorCode = &v
	return s
}

func (s *DsgGetVisitDetailResponseBody) SetDynamicErrorMessage(v string) *DsgGetVisitDetailResponseBody {
	s.DynamicErrorMessage = &v
	return s
}

func (s *DsgGetVisitDetailResponseBody) SetErrorCode(v string) *DsgGetVisitDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgGetVisitDetailResponseBody) SetErrorMessage(v string) *DsgGetVisitDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgGetVisitDetailResponseBody) SetHttpStatusCode(v int32) *DsgGetVisitDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgGetVisitDetailResponseBody) SetRequestId(v string) *DsgGetVisitDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgGetVisitDetailResponseBody) SetSuccess(v bool) *DsgGetVisitDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DsgGetVisitDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
