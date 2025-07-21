// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIsvCustomerTermsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SubmitIsvCustomerTermsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SubmitIsvCustomerTermsResponseBody
	GetCode() *string
	SetMessage(v string) *SubmitIsvCustomerTermsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitIsvCustomerTermsResponseBody
	GetRequestId() *string
}

type SubmitIsvCustomerTermsResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// /
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitIsvCustomerTermsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitIsvCustomerTermsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIsvCustomerTermsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SubmitIsvCustomerTermsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitIsvCustomerTermsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitIsvCustomerTermsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitIsvCustomerTermsResponseBody) SetAccessDeniedDetail(v string) *SubmitIsvCustomerTermsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SubmitIsvCustomerTermsResponseBody) SetCode(v string) *SubmitIsvCustomerTermsResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitIsvCustomerTermsResponseBody) SetMessage(v string) *SubmitIsvCustomerTermsResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitIsvCustomerTermsResponseBody) SetRequestId(v string) *SubmitIsvCustomerTermsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIsvCustomerTermsResponseBody) Validate() error {
	return dara.Validate(s)
}
