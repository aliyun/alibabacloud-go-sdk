// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAudienceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddCustomAudienceUserResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddCustomAudienceUserResponseBody
	GetCode() *string
	SetMessage(v string) *AddCustomAudienceUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCustomAudienceUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddCustomAudienceUserResponseBody
	GetSuccess() *bool
}

type AddCustomAudienceUserResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - A value of OK indicates that the request was successful.
	//
	// - For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCustomAudienceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAudienceUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomAudienceUserResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddCustomAudienceUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddCustomAudienceUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCustomAudienceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomAudienceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddCustomAudienceUserResponseBody) SetAccessDeniedDetail(v string) *AddCustomAudienceUserResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddCustomAudienceUserResponseBody) SetCode(v string) *AddCustomAudienceUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddCustomAudienceUserResponseBody) SetMessage(v string) *AddCustomAudienceUserResponseBody {
	s.Message = &v
	return s
}

func (s *AddCustomAudienceUserResponseBody) SetRequestId(v string) *AddCustomAudienceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomAudienceUserResponseBody) SetSuccess(v bool) *AddCustomAudienceUserResponseBody {
	s.Success = &v
	return s
}

func (s *AddCustomAudienceUserResponseBody) Validate() error {
	return dara.Validate(s)
}
