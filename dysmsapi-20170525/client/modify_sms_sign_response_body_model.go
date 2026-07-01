// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmsSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySmsSignResponseBody
	GetCode() *string
	SetMessage(v string) *ModifySmsSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySmsSignResponseBody
	GetRequestId() *string
	SetSignName(v string) *ModifySmsSignResponseBody
	GetSignName() *string
}

type ModifySmsSignResponseBody struct {
	// 请求状态码。
	//
	// 	- 返回OK代表请求成功。
	//
	// 	- 其他错误码，请参见[API错误码](https://help.aliyun.com/document_detail/101346.html)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 状态码的描述。
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 签名名称。
	//
	// example:
	//
	// 阿里云
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s ModifySmsSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmsSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySmsSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySmsSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySmsSignResponseBody) GetSignName() *string {
	return s.SignName
}

func (s *ModifySmsSignResponseBody) SetCode(v string) *ModifySmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySmsSignResponseBody) SetMessage(v string) *ModifySmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySmsSignResponseBody) SetRequestId(v string) *ModifySmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySmsSignResponseBody) SetSignName(v string) *ModifySmsSignResponseBody {
	s.SignName = &v
	return s
}

func (s *ModifySmsSignResponseBody) Validate() error {
	return dara.Validate(s)
}
