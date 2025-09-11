// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySmsTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *ModifySmsTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySmsTemplateResponseBody
	GetRequestId() *string
	SetTemplateCode(v string) *ModifySmsTemplateResponseBody
	GetTemplateCode() *string
}

type ModifySmsTemplateResponseBody struct {
	// The response code.
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
	// The code of the message template.
	//
	// example:
	//
	// SMS_15255****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s ModifySmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySmsTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySmsTemplateResponseBody) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ModifySmsTemplateResponseBody) SetCode(v string) *ModifySmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySmsTemplateResponseBody) SetMessage(v string) *ModifySmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySmsTemplateResponseBody) SetRequestId(v string) *ModifySmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySmsTemplateResponseBody) SetTemplateCode(v string) *ModifySmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *ModifySmsTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
