// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommerceSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateCommerceSettingResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateCommerceSettingResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateCommerceSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCommerceSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCommerceSettingResponseBody
	GetSuccess() *bool
}

type UpdateCommerceSettingResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
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

func (s UpdateCommerceSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommerceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCommerceSettingResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateCommerceSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCommerceSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCommerceSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCommerceSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCommerceSettingResponseBody) SetAccessDeniedDetail(v string) *UpdateCommerceSettingResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateCommerceSettingResponseBody) SetCode(v string) *UpdateCommerceSettingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCommerceSettingResponseBody) SetMessage(v string) *UpdateCommerceSettingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCommerceSettingResponseBody) SetRequestId(v string) *UpdateCommerceSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCommerceSettingResponseBody) SetSuccess(v bool) *UpdateCommerceSettingResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCommerceSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
