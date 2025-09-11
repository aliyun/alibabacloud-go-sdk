// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddSmsTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *AddSmsTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSmsTemplateResponseBody
	GetRequestId() *string
	SetTemplateCode(v string) *AddSmsTemplateResponseBody
	GetTemplateCode() *string
}

type AddSmsTemplateResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
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

func (s AddSmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddSmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddSmsTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSmsTemplateResponseBody) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *AddSmsTemplateResponseBody) SetCode(v string) *AddSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *AddSmsTemplateResponseBody) SetMessage(v string) *AddSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *AddSmsTemplateResponseBody) SetRequestId(v string) *AddSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSmsTemplateResponseBody) SetTemplateCode(v string) *AddSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *AddSmsTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
