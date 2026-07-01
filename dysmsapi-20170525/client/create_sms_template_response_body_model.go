// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSmsTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *CreateSmsTemplateResponseBody
	GetMessage() *string
	SetOrderId(v string) *CreateSmsTemplateResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateSmsTemplateResponseBody
	GetRequestId() *string
	SetTemplateCode(v string) *CreateSmsTemplateResponseBody
	GetTemplateCode() *string
	SetTemplateName(v string) *CreateSmsTemplateResponseBody
	GetTemplateName() *string
}

type CreateSmsTemplateResponseBody struct {
	// The status code of the request. Valid values:
	//
	// - OK: The request was successful.
	//
	// - For other error codes, see the **Error codes*	- list in this topic or [API Error Codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ticket ID.
	//
	// Reviewers use this ID to check the review status. You must provide this ID when you request an expedited review.
	//
	// example:
	//
	// 2005020****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request. Alibaba Cloud generates a unique ID for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE991
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code of the message template.
	//
	// After you submit a template application, you can use the template code to call the [GetSmsTemplate](https://help.aliyun.com/document_detail/2807433.html) operation and query the review details of the template. You can also [configure delivery receipts](https://help.aliyun.com/document_detail/101508.html) to receive the review status of the template in the [TemplateSmsReport](https://help.aliyun.com/document_detail/120999.html) message.
	//
	// example:
	//
	// SMS_10000****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// 验证码模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateSmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSmsTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSmsTemplateResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateSmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmsTemplateResponseBody) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateSmsTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateSmsTemplateResponseBody) SetCode(v string) *CreateSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetMessage(v string) *CreateSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetOrderId(v string) *CreateSmsTemplateResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetRequestId(v string) *CreateSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetTemplateCode(v string) *CreateSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetTemplateName(v string) *CreateSmsTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
