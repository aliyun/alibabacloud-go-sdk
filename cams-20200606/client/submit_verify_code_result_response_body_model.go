// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVerifyCodeResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SubmitVerifyCodeResultResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SubmitVerifyCodeResultResponseBody
	GetCode() *string
	SetData(v string) *SubmitVerifyCodeResultResponseBody
	GetData() *string
	SetMessage(v string) *SubmitVerifyCodeResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitVerifyCodeResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitVerifyCodeResultResponseBody
	GetSuccess() *bool
}

type SubmitVerifyCodeResultResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - A value of `OK` indicates that the request was successful.
	//
	// - See the [error code list](https://help.aliyun.com/document_detail/196974.html) for other error codes.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// example
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message.
	//
	// example:
	//
	// example
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// example
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitVerifyCodeResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVerifyCodeResultResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVerifyCodeResultResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SubmitVerifyCodeResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitVerifyCodeResultResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitVerifyCodeResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitVerifyCodeResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVerifyCodeResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitVerifyCodeResultResponseBody) SetAccessDeniedDetail(v string) *SubmitVerifyCodeResultResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SubmitVerifyCodeResultResponseBody) SetCode(v string) *SubmitVerifyCodeResultResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitVerifyCodeResultResponseBody) SetData(v string) *SubmitVerifyCodeResultResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitVerifyCodeResultResponseBody) SetMessage(v string) *SubmitVerifyCodeResultResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitVerifyCodeResultResponseBody) SetRequestId(v string) *SubmitVerifyCodeResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVerifyCodeResultResponseBody) SetSuccess(v bool) *SubmitVerifyCodeResultResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitVerifyCodeResultResponseBody) Validate() error {
	return dara.Validate(s)
}
