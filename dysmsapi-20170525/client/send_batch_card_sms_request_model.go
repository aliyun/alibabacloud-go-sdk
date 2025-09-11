// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBatchCardSmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCardTemplateCode(v string) *SendBatchCardSmsRequest
	GetCardTemplateCode() *string
	SetCardTemplateParamJson(v string) *SendBatchCardSmsRequest
	GetCardTemplateParamJson() *string
	SetDigitalTemplateCode(v string) *SendBatchCardSmsRequest
	GetDigitalTemplateCode() *string
	SetDigitalTemplateParamJson(v string) *SendBatchCardSmsRequest
	GetDigitalTemplateParamJson() *string
	SetFallbackType(v string) *SendBatchCardSmsRequest
	GetFallbackType() *string
	SetOutId(v string) *SendBatchCardSmsRequest
	GetOutId() *string
	SetPhoneNumberJson(v string) *SendBatchCardSmsRequest
	GetPhoneNumberJson() *string
	SetSignNameJson(v string) *SendBatchCardSmsRequest
	GetSignNameJson() *string
	SetSmsTemplateCode(v string) *SendBatchCardSmsRequest
	GetSmsTemplateCode() *string
	SetSmsTemplateParamJson(v string) *SendBatchCardSmsRequest
	GetSmsTemplateParamJson() *string
	SetSmsUpExtendCodeJson(v string) *SendBatchCardSmsRequest
	GetSmsUpExtendCodeJson() *string
	SetTemplateCode(v string) *SendBatchCardSmsRequest
	GetTemplateCode() *string
	SetTemplateParamJson(v string) *SendBatchCardSmsRequest
	GetTemplateParamJson() *string
}

type SendBatchCardSmsRequest struct {
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_3245
	CardTemplateCode *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	// The variables of the card message template.
	//
	// example:
	//
	// [{\\"customurl\\":\\"http://www.alibaba.com\\",\\"dyncParams\\":\\"{\\\\\\"a\\\\\\":\\\\\\"hello\\\\\\",\\\\\\"b\\\\\\":\\\\\\"world\\\\\\"}\\"}]
	CardTemplateParamJson *string `json:"CardTemplateParamJson,omitempty" xml:"CardTemplateParamJson,omitempty"`
	// The code of the digital message template that applies when the card message is rolled back. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// example:
	//
	// DIGITAL_SMS_234080176
	DigitalTemplateCode *string `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	// The variables of the digital message template.
	//
	// example:
	//
	// [{"a":1,"b":2},{"a":9,"b":8}]
	DigitalTemplateParamJson *string `json:"DigitalTemplateParamJson,omitempty" xml:"DigitalTemplateParamJson,omitempty"`
	// The rollback type. Valid values:
	//
	// 	- **SMS**: text message
	//
	// 	- **DIGITALSMS**: digital message
	//
	// 	- **NONE**: none
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS
	FallbackType *string `json:"FallbackType,omitempty" xml:"FallbackType,omitempty"`
	// The ID that is reserved for the caller of the operation.
	//
	// example:
	//
	// 16545681783595370
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The mobile numbers of the recipients.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"1390000****\\",\\"1370000****\\"]"
	PhoneNumberJson *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	// The signature. You can view the template code in the **Signature*	- column on the **Signaturess*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > The signatures must be approved and correspond to the mobile numbers in sequence.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"aliyun\\",\\"aliyuncode\\"]
	SignNameJson *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	// The code of the text message template that applies when the card message is rolled back. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// example:
	//
	// SMS_234251075
	SmsTemplateCode *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	// The variables of the text message template.
	//
	// example:
	//
	// [{"a":1,"b":2},{"a":9,"b":8}]
	SmsTemplateParamJson *string `json:"SmsTemplateParamJson,omitempty" xml:"SmsTemplateParamJson,omitempty"`
	// The extension code of the upstream message.
	//
	// example:
	//
	// [\\"6\\",\\"6\\"]
	SmsUpExtendCodeJson *string `json:"SmsUpExtendCodeJson,omitempty" xml:"SmsUpExtendCodeJson,omitempty"`
	// The code of the message template.
	//
	// You can log on to the [Alibaba Cloud console](https://dysms.console.aliyun.com/dysms.htm?spm=5176.12818093.categories-n-products.ddysms.3b2816d0xml2NA#/overview), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the **template code*	- on the **Templates*	- tab.
	//
	// > You must specify a message template that is created in the SMS console and approved by Alibaba Cloud. If you send messages to countries or regions outside the Chinese mainland, use the corresponding message templates.
	//
	// example:
	//
	// SMS_20375****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The value of the variable in the message template.
	//
	// > If you need to add line breaks to the JSON template, make sure that the format is valid. In addition, the sequence of variable values must be the same as that of the mobile numbers and signatures.
	//
	// example:
	//
	// [{"name":"TemplateParamJson"},{"name":"TemplateParamJson"}]
	TemplateParamJson *string `json:"TemplateParamJson,omitempty" xml:"TemplateParamJson,omitempty"`
}

func (s SendBatchCardSmsRequest) String() string {
	return dara.Prettify(s)
}

func (s SendBatchCardSmsRequest) GoString() string {
	return s.String()
}

func (s *SendBatchCardSmsRequest) GetCardTemplateCode() *string {
	return s.CardTemplateCode
}

func (s *SendBatchCardSmsRequest) GetCardTemplateParamJson() *string {
	return s.CardTemplateParamJson
}

func (s *SendBatchCardSmsRequest) GetDigitalTemplateCode() *string {
	return s.DigitalTemplateCode
}

func (s *SendBatchCardSmsRequest) GetDigitalTemplateParamJson() *string {
	return s.DigitalTemplateParamJson
}

func (s *SendBatchCardSmsRequest) GetFallbackType() *string {
	return s.FallbackType
}

func (s *SendBatchCardSmsRequest) GetOutId() *string {
	return s.OutId
}

func (s *SendBatchCardSmsRequest) GetPhoneNumberJson() *string {
	return s.PhoneNumberJson
}

func (s *SendBatchCardSmsRequest) GetSignNameJson() *string {
	return s.SignNameJson
}

func (s *SendBatchCardSmsRequest) GetSmsTemplateCode() *string {
	return s.SmsTemplateCode
}

func (s *SendBatchCardSmsRequest) GetSmsTemplateParamJson() *string {
	return s.SmsTemplateParamJson
}

func (s *SendBatchCardSmsRequest) GetSmsUpExtendCodeJson() *string {
	return s.SmsUpExtendCodeJson
}

func (s *SendBatchCardSmsRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendBatchCardSmsRequest) GetTemplateParamJson() *string {
	return s.TemplateParamJson
}

func (s *SendBatchCardSmsRequest) SetCardTemplateCode(v string) *SendBatchCardSmsRequest {
	s.CardTemplateCode = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetCardTemplateParamJson(v string) *SendBatchCardSmsRequest {
	s.CardTemplateParamJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetDigitalTemplateCode(v string) *SendBatchCardSmsRequest {
	s.DigitalTemplateCode = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetDigitalTemplateParamJson(v string) *SendBatchCardSmsRequest {
	s.DigitalTemplateParamJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetFallbackType(v string) *SendBatchCardSmsRequest {
	s.FallbackType = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetOutId(v string) *SendBatchCardSmsRequest {
	s.OutId = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetPhoneNumberJson(v string) *SendBatchCardSmsRequest {
	s.PhoneNumberJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetSignNameJson(v string) *SendBatchCardSmsRequest {
	s.SignNameJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetSmsTemplateCode(v string) *SendBatchCardSmsRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetSmsTemplateParamJson(v string) *SendBatchCardSmsRequest {
	s.SmsTemplateParamJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetSmsUpExtendCodeJson(v string) *SendBatchCardSmsRequest {
	s.SmsUpExtendCodeJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetTemplateCode(v string) *SendBatchCardSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetTemplateParamJson(v string) *SendBatchCardSmsRequest {
	s.TemplateParamJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) Validate() error {
	return dara.Validate(s)
}
