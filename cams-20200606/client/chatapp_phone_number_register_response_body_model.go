// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappPhoneNumberRegisterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChatappPhoneNumberRegisterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChatappPhoneNumberRegisterResponseBody
	GetCode() *string
	SetMessage(v string) *ChatappPhoneNumberRegisterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChatappPhoneNumberRegisterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatappPhoneNumberRegisterResponseBody
	GetSuccess() *bool
}

type ChatappPhoneNumberRegisterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
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
	// None.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
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

func (s ChatappPhoneNumberRegisterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatappPhoneNumberRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberRegisterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChatappPhoneNumberRegisterResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatappPhoneNumberRegisterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatappPhoneNumberRegisterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatappPhoneNumberRegisterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatappPhoneNumberRegisterResponseBody) SetAccessDeniedDetail(v string) *ChatappPhoneNumberRegisterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponseBody) SetCode(v string) *ChatappPhoneNumberRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponseBody) SetMessage(v string) *ChatappPhoneNumberRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponseBody) SetRequestId(v string) *ChatappPhoneNumberRegisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponseBody) SetSuccess(v bool) *ChatappPhoneNumberRegisterResponseBody {
	s.Success = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponseBody) Validate() error {
	return dara.Validate(s)
}
