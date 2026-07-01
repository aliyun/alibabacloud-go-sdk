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
	// The code of the card SMS template. On the **Card SMS*	- [Template Management](https://dysms.console.aliyun.com/domestic/card) page in the console, select the code of a card SMS template that has been **approved**.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_3**5
	CardTemplateCode *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	// The actual values of the variables in the card SMS template. This parameter is required when the card SMS template specified by **CardTemplateCode*	- contains variables.
	//
	// >If the JSON contains line breaks, handle them based on the standard JSON protocol.
	//
	// example:
	//
	// [{"customurl":"http://www.alibaba.com","dyncParams":"{"a":"hello","b":"world"}"}]
	CardTemplateParamJson *string `json:"CardTemplateParamJson,omitempty" xml:"CardTemplateParamJson,omitempty"`
	// The code of the digital SMS template used for fallback. This parameter is required when **FallbackType*	- is set to **DIGITALSMS*	- (fallback to digital SMS).
	//
	// You can view the list of digital SMS templates on the **Domestic Digital SMS*	- [Template Management](https://dysms.console.aliyun.com/domestic/digit) page in the console.
	//
	// >The template must be added and approved.
	//
	// example:
	//
	// DIGITAL_SMS_23408****
	DigitalTemplateCode *string `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	// The actual values of the variables in the digital SMS template. This parameter is required when the fallback digital SMS template specified by **DigitalTemplateCode*	- contains variables.
	//
	// >If the JSON contains line breaks, handle them based on the standard JSON protocol.
	//
	// example:
	//
	// [{"a":1,"b":2},{"a":9,"b":8}]
	DigitalTemplateParamJson *string `json:"DigitalTemplateParamJson,omitempty" xml:"DigitalTemplateParamJson,omitempty"`
	// The fallback type. Valid values:
	//
	// - **SMS**: Phone numbers that do not support card SMS messages fall back to text SMS messages.
	//
	// - **DIGITALSMS**: Phone numbers that do not support card SMS messages fall back to digital SMS messages.
	//
	// - **NONE**: No fallback is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS
	FallbackType *string `json:"FallbackType,omitempty" xml:"FallbackType,omitempty"`
	// The ID reserved for the caller.
	//
	// example:
	//
	// 16545681783595370
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The mobile phone numbers that receive the SMS messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1390000****","1370000****"]
	PhoneNumberJson *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	// The name of the SMS signature.
	//
	// You can call the [QuerySmsSignList](https://help.aliyun.com/document_detail/419282.html) operation to query the signatures that have been submitted under the current account, or you can view the list of signatures in the [Short Message Service console](https://dysms.console.aliyun.com/domestic/text/sign).
	//
	// >The signature must be added and approved. The number of SMS signatures must be the same as the number of phone numbers, and the signatures must be in one-to-one correspondence with the phone numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["阿里云","阿里巴巴"]
	SignNameJson *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	// The code of the text SMS template used for fallback. This parameter is required when **FallbackType*	- is set to **SMS*	- (fallback to text SMS).
	//
	// You can call the [QuerySmsTemplateList](https://help.aliyun.com/document_detail/419288.html) operation to query the templates that have been submitted under the current account, or you can view the list of templates in the [Short Message Service console](https://dysms.console.aliyun.com/domestic/text/template).
	//
	// >The template must be added and approved.
	//
	// example:
	//
	// SMS_23425****
	SmsTemplateCode *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	// The actual values of the variables in the text SMS template. This parameter is required when the fallback text SMS template specified by **SmsTemplateCode*	- contains variables.
	//
	// >If the JSON contains line breaks, handle them based on the standard JSON protocol.
	//
	// example:
	//
	// [{"a":1,"b":2},{"a":9,"b":8}]
	SmsTemplateParamJson *string `json:"SmsTemplateParamJson,omitempty" xml:"SmsTemplateParamJson,omitempty"`
	// The extension code of the MO (mobile-originated) SMS message.
	//
	// example:
	//
	// [\\"6\\",\\"6\\"]
	SmsUpExtendCodeJson *string `json:"SmsUpExtendCodeJson,omitempty" xml:"SmsUpExtendCodeJson,omitempty"`
	// The code of the custom send content template.
	//
	// The custom content is sent to the terminal in the form of the selected text SMS template plus the card parsing link. You can log on to the [Short Message Service console](https://dysms.console.aliyun.com/overview), choose **Domestic Messages*	- or **International/Hong Kong, Macao, and Taiwan Messages**, and then view the **Template Code*	- on the **Template Management*	- page.
	//
	// > - The template must be added and approved. To send international or Hong Kong, Macao, and Taiwan messages, use an international or Hong Kong, Macao, and Taiwan SMS template.
	//
	// > - For example, the selected text SMS template content is: You have a message to check; the card parsing link is: `1*.cn/2**d`. The final delivered content is: `You have a message to check 1*.cn/2**d`. Perform testing and control the number of characters before sending.
	//
	// example:
	//
	// SMS_20375****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The actual values of the variables in the custom send content template. This parameter is required when the SMS template specified by **TemplateCode*	- contains variables.
	//
	// > - If the JSON contains line breaks, handle them based on the standard JSON protocol.
	//
	// > - The number of template variable values must be the same as the number of phone numbers and signatures, and they must be in one-to-one correspondence. This indicates that an SMS message with the corresponding signature is sent to the specified phone number, and the variable parameters in the SMS template are replaced with the corresponding values.
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
