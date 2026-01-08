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
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// cn
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
