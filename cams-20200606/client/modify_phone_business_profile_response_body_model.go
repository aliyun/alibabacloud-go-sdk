// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhoneBusinessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyPhoneBusinessProfileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyPhoneBusinessProfileResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyPhoneBusinessProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyPhoneBusinessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyPhoneBusinessProfileResponseBody
	GetSuccess() *bool
}

type ModifyPhoneBusinessProfileResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The URL of the website.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The websites.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPhoneBusinessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhoneBusinessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPhoneBusinessProfileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyPhoneBusinessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyPhoneBusinessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyPhoneBusinessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPhoneBusinessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyPhoneBusinessProfileResponseBody) SetAccessDeniedDetail(v string) *ModifyPhoneBusinessProfileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponseBody) SetCode(v string) *ModifyPhoneBusinessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponseBody) SetMessage(v string) *ModifyPhoneBusinessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponseBody) SetRequestId(v string) *ModifyPhoneBusinessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponseBody) SetSuccess(v bool) *ModifyPhoneBusinessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
