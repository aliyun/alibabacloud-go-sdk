// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappVerifyCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetChatappVerifyCodeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetChatappVerifyCodeResponseBody
	GetCode() *string
	SetMessage(v string) *GetChatappVerifyCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChatappVerifyCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetChatappVerifyCodeResponseBody
	GetSuccess() *bool
}

type GetChatappVerifyCodeResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the request.
	//
	// - A value of OK indicates that the request is successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1612C226-E271-4CFE-9F18-4066D550F91B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// - **true**: successful.
	//
	// - **false**: failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetChatappVerifyCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatappVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappVerifyCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetChatappVerifyCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatappVerifyCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatappVerifyCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatappVerifyCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetChatappVerifyCodeResponseBody) SetAccessDeniedDetail(v string) *GetChatappVerifyCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappVerifyCodeResponseBody) SetCode(v string) *GetChatappVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappVerifyCodeResponseBody) SetMessage(v string) *GetChatappVerifyCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappVerifyCodeResponseBody) SetRequestId(v string) *GetChatappVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatappVerifyCodeResponseBody) SetSuccess(v bool) *GetChatappVerifyCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetChatappVerifyCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
