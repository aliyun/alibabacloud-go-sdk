// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChatappPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddChatappPhoneNumberResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddChatappPhoneNumberResponseBody
	GetCode() *string
	SetMessage(v string) *AddChatappPhoneNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddChatappPhoneNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddChatappPhoneNumberResponseBody
	GetSuccess() *bool
}

type AddChatappPhoneNumberResponseBody struct {
	// com.alicom.access.oxs.client.channel.aliyun.flow.AyFlowExecuteService
	//
	// example:
	//
	// http://pop_access_slb_sgvpc/#vpc
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The phone number.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// com.alicom.access.oxs.client.channel.aliyun.flow.dto.AyCommonApiRequest
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// formData
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 13800000000
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddChatappPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddChatappPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *AddChatappPhoneNumberResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddChatappPhoneNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddChatappPhoneNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddChatappPhoneNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddChatappPhoneNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddChatappPhoneNumberResponseBody) SetAccessDeniedDetail(v string) *AddChatappPhoneNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddChatappPhoneNumberResponseBody) SetCode(v string) *AddChatappPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *AddChatappPhoneNumberResponseBody) SetMessage(v string) *AddChatappPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *AddChatappPhoneNumberResponseBody) SetRequestId(v string) *AddChatappPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddChatappPhoneNumberResponseBody) SetSuccess(v bool) *AddChatappPhoneNumberResponseBody {
	s.Success = &v
	return s
}

func (s *AddChatappPhoneNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
