// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalSmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDigitalSmsTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *CreateDigitalSmsTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDigitalSmsTemplateResponseBody
	GetRequestId() *string
	SetTemplateCode(v string) *CreateDigitalSmsTemplateResponseBody
	GetTemplateCode() *string
}

type CreateDigitalSmsTemplateResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// DIGITAL_SMS_0001
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CreateDigitalSmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDigitalSmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDigitalSmsTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDigitalSmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDigitalSmsTemplateResponseBody) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateDigitalSmsTemplateResponseBody) SetCode(v string) *CreateDigitalSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDigitalSmsTemplateResponseBody) SetMessage(v string) *CreateDigitalSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDigitalSmsTemplateResponseBody) SetRequestId(v string) *CreateDigitalSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDigitalSmsTemplateResponseBody) SetTemplateCode(v string) *CreateDigitalSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *CreateDigitalSmsTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
