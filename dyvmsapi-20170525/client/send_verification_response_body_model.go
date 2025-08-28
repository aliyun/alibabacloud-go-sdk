// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendVerificationResponseBody
	GetCode() *string
	SetData(v bool) *SendVerificationResponseBody
	GetData() *bool
	SetMessage(v string) *SendVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendVerificationResponseBody
	GetRequestId() *string
}

type SendVerificationResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the verification code was sent successfully.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendVerificationResponseBody) GetData() *bool {
	return s.Data
}

func (s *SendVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendVerificationResponseBody) SetCode(v string) *SendVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *SendVerificationResponseBody) SetData(v bool) *SendVerificationResponseBody {
	s.Data = &v
	return s
}

func (s *SendVerificationResponseBody) SetMessage(v string) *SendVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *SendVerificationResponseBody) SetRequestId(v string) *SendVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendVerificationResponseBody) Validate() error {
	return dara.Validate(s)
}
