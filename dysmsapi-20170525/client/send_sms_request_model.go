// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutId(v string) *SendSmsRequest
	GetOutId() *string
	SetOwnerId(v int64) *SendSmsRequest
	GetOwnerId() *int64
	SetPhoneNumbers(v string) *SendSmsRequest
	GetPhoneNumbers() *string
	SetResourceOwnerAccount(v string) *SendSmsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendSmsRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *SendSmsRequest
	GetSignName() *string
	SetSmsUpExtendCode(v string) *SendSmsRequest
	GetSmsUpExtendCode() *string
	SetTemplateCode(v string) *SendSmsRequest
	GetTemplateCode() *string
	SetTemplateParam(v string) *SendSmsRequest
	GetTemplateParam() *string
}

type SendSmsRequest struct {
	// The external transaction ID.
	//
	// > You can ignore this parameter if you do not have special requirements.
	//
	// example:
	//
	// abcdefgh
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recipient\\"s mobile number. The format varies based on the destination region:
	//
	// - For messages to the Chinese mainland: A mobile number, with or without a country code. Valid prefixes are +, +86, 0086, and 86. Example: 1390000\\*\\*\\*\\*.
	//
	// - For international messages or messages to Hong Kong, Macao, or Taiwan: Use the format [Country code][Mobile number]. Example: 852000012\\*\\*\\*\\*.
	//
	// - To send a test message, bind a test mobile number in the [console](https://dysms.console.aliyun.com/quickstart).
	//
	// > To send a message to multiple numbers, separate them with commas (,). You can specify up to 1,000 mobile numbers per request. Bulk sending may have higher latency than sending single messages. For time-sensitive messages such as verification codes, we recommend sending them individually.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNumbers         *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature name.
	//
	// Call the [QuerySmsSignList](https://help.aliyun.com/document_detail/419282.html) API or view your list of signatures in the [Short Message Service console](https://dysms.console.aliyun.com/domestic/text/sign). You must use a signature that has been **approved**.
	//
	// > - 1\\. If a verification code signature and a general-purpose signature share the same name, the general-purpose signature is used by default.
	//
	// >
	//
	// > - 2\\. The system uses the specified signature to send the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The upstream SMS extension code. An upstream SMS message is a message sent from a user to a service provider to subscribe to a service, make a query, or perform other tasks. Such messages are charged by the carrier at standard rates.
	//
	// > The system assigns a default extension code when a signature is created. Use this parameter to specify a different code. You can ignore this parameter if you do not have special requirements.
	//
	// example:
	//
	// 90999
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	// The code of the template.
	//
	// Call the [QuerySmsTemplateList](https://help.aliyun.com/document_detail/419288.html) API or view your list of templates in the [Short Message Service console](https://dysms.console.aliyun.com/domestic/text/template). You must use the code of a template that has been **approved**.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_15305****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The values for the template variables, specified as a **JSON string**. This parameter is required if the template contains variables. The JSON string must provide a value for each variable.
	//
	// > - If the JSON string needs to include line breaks, format it according to standard JSON specifications.
	//
	// >
	//
	// > - For more information about template variable formatting, see [SMS template specifications](https://help.aliyun.com/document_detail/463161.html).
	//
	// example:
	//
	// {"name":"张三","number":"1390000****"}
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
}

func (s SendSmsRequest) String() string {
	return dara.Prettify(s)
}

func (s SendSmsRequest) GoString() string {
	return s.String()
}

func (s *SendSmsRequest) GetOutId() *string {
	return s.OutId
}

func (s *SendSmsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendSmsRequest) GetPhoneNumbers() *string {
	return s.PhoneNumbers
}

func (s *SendSmsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendSmsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendSmsRequest) GetSignName() *string {
	return s.SignName
}

func (s *SendSmsRequest) GetSmsUpExtendCode() *string {
	return s.SmsUpExtendCode
}

func (s *SendSmsRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendSmsRequest) GetTemplateParam() *string {
	return s.TemplateParam
}

func (s *SendSmsRequest) SetOutId(v string) *SendSmsRequest {
	s.OutId = &v
	return s
}

func (s *SendSmsRequest) SetOwnerId(v int64) *SendSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *SendSmsRequest) SetPhoneNumbers(v string) *SendSmsRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *SendSmsRequest) SetResourceOwnerAccount(v string) *SendSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendSmsRequest) SetResourceOwnerId(v int64) *SendSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendSmsRequest) SetSignName(v string) *SendSmsRequest {
	s.SignName = &v
	return s
}

func (s *SendSmsRequest) SetSmsUpExtendCode(v string) *SendSmsRequest {
	s.SmsUpExtendCode = &v
	return s
}

func (s *SendSmsRequest) SetTemplateCode(v string) *SendSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendSmsRequest) SetTemplateParam(v string) *SendSmsRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendSmsRequest) Validate() error {
	return dara.Validate(s)
}
