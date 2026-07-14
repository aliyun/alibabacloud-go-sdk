// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetCustomerSiteResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetCustomerSiteResponseBody
	GetCode() *string
	SetData(v string) *GetCustomerSiteResponseBody
	GetData() *string
	SetMessage(v string) *GetCustomerSiteResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomerSiteResponseBody
	GetRequestId() *string
}

type GetCustomerSiteResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK: The request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data object.
	//
	// example:
	//
	// cn
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 3R938***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCustomerSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerSiteResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerSiteResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetCustomerSiteResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomerSiteResponseBody) GetData() *string {
	return s.Data
}

func (s *GetCustomerSiteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomerSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerSiteResponseBody) SetAccessDeniedDetail(v string) *GetCustomerSiteResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCustomerSiteResponseBody) SetCode(v string) *GetCustomerSiteResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerSiteResponseBody) SetData(v string) *GetCustomerSiteResponseBody {
	s.Data = &v
	return s
}

func (s *GetCustomerSiteResponseBody) SetMessage(v string) *GetCustomerSiteResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerSiteResponseBody) SetRequestId(v string) *GetCustomerSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
