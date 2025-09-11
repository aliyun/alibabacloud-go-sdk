// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SendSmsResponseBody
	GetBizId() *string
	SetCode(v string) *SendSmsResponseBody
	GetCode() *string
	SetMessage(v string) *SendSmsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendSmsResponseBody
	GetRequestId() *string
}

type SendSmsResponseBody struct {
	// The ID of the delivery receipt.
	//
	// You can call the [QuerySendDetails](~~QuerySendDetails~~) operation to query the delivery status based on the receipt ID.
	//
	// example:
	//
	// 9006197469364984****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendSmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendSmsResponseBody) GetBizId() *string {
	return s.BizId
}

func (s *SendSmsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendSmsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendSmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendSmsResponseBody) SetBizId(v string) *SendSmsResponseBody {
	s.BizId = &v
	return s
}

func (s *SendSmsResponseBody) SetCode(v string) *SendSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendSmsResponseBody) SetMessage(v string) *SendSmsResponseBody {
	s.Message = &v
	return s
}

func (s *SendSmsResponseBody) SetRequestId(v string) *SendSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendSmsResponseBody) Validate() error {
	return dara.Validate(s)
}
