// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappVerifyAndRegisterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChatappVerifyAndRegisterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChatappVerifyAndRegisterResponseBody
	GetCode() *string
	SetMessage(v string) *ChatappVerifyAndRegisterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChatappVerifyAndRegisterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatappVerifyAndRegisterResponseBody
	GetSuccess() *bool
}

type ChatappVerifyAndRegisterResponseBody struct {
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
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**: The call was successful.
	//
	// 	- **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatappVerifyAndRegisterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatappVerifyAndRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappVerifyAndRegisterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChatappVerifyAndRegisterResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatappVerifyAndRegisterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatappVerifyAndRegisterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatappVerifyAndRegisterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatappVerifyAndRegisterResponseBody) SetAccessDeniedDetail(v string) *ChatappVerifyAndRegisterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponseBody) SetCode(v string) *ChatappVerifyAndRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponseBody) SetMessage(v string) *ChatappVerifyAndRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponseBody) SetRequestId(v string) *ChatappVerifyAndRegisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponseBody) SetSuccess(v bool) *ChatappVerifyAndRegisterResponseBody {
	s.Success = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponseBody) Validate() error {
	return dara.Validate(s)
}
