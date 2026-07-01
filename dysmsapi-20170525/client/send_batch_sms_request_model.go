// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBatchSmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutId(v string) *SendBatchSmsRequest
	GetOutId() *string
	SetOwnerId(v int64) *SendBatchSmsRequest
	GetOwnerId() *int64
	SetPhoneNumberJson(v string) *SendBatchSmsRequest
	GetPhoneNumberJson() *string
	SetResourceOwnerAccount(v string) *SendBatchSmsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendBatchSmsRequest
	GetResourceOwnerId() *int64
	SetSignNameJson(v string) *SendBatchSmsRequest
	GetSignNameJson() *string
	SetSmsUpExtendCodeJson(v string) *SendBatchSmsRequest
	GetSmsUpExtendCodeJson() *string
	SetTemplateCode(v string) *SendBatchSmsRequest
	GetTemplateCode() *string
	SetTemplateParamJson(v string) *SendBatchSmsRequest
	GetTemplateParamJson() *string
}

type SendBatchSmsRequest struct {
	// An external business ID. It must be a string of fewer than 256 characters.
	//
	// > You can leave this parameter empty if you have no special requirements.
	//
	// example:
	//
	// abcdefg
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recipient phone numbers. Format:
	//
	// - For domestic SMS: Phone numbers with or without a country code such as `+`, `+86`, `0086`, or `86`. Example: `1590000****`.
	//
	// - For international SMS: The country code followed by the phone number. Example: `852000012****`.
	//
	// > For time-sensitive messages like verification codes, use the [SendSms](https://help.aliyun.com/document_detail/419273.html) operation to send messages individually.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1590000****","1350000****"]
	PhoneNumberJson      *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature names. The number of signatures must match the number of phone numbers.
	//
	// You can call the [QuerySmsSignList](https://help.aliyun.com/document_detail/419282.html) operation or check the [Short Message Service console](https://dysms.console.aliyun.com/domestic/text/sign) to find approved signatures. You must use an approved signature.
	//
	// > - The system uses the selected signature to send SMS messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["阿里云","阿里巴巴"]
	SignNameJson *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	// A JSON array of MO SMS extension codes.
	//
	// > You can leave this parameter empty if you have no special requirements.
	//
	// example:
	//
	// ["90999","90998"]
	SmsUpExtendCodeJson *string `json:"SmsUpExtendCodeJson,omitempty" xml:"SmsUpExtendCodeJson,omitempty"`
	// The message template code. You cannot use templates for domestic SMS and international SMS interchangeably.
	//
	// You can call the [QuerySmsTemplateList](https://help.aliyun.com/document_detail/419288.html) operation or check the [Short Message Service console](https://dysms.console.aliyun.com/domestic/text/template) to find approved template codes. You must use an approved template code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_15255****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The actual values for the template variables. This parameter is required if the template contains variables.
	//
	// > - The number of template variable sets must match the number of phone numbers and signatures. The elements in the PhoneNumberJson, SignNameJson, and TemplateParamJson arrays must correspond by index to ensure each message is sent with the correct signature and variable values.
	//
	// >
	//
	// > - If you need to include a line break in the JSON string, follow the standard JSON format.
	//
	// example:
	//
	// [{"name":"TemplateParamJson"},{"name":"TemplateParamJson"}]
	TemplateParamJson *string `json:"TemplateParamJson,omitempty" xml:"TemplateParamJson,omitempty"`
}

func (s SendBatchSmsRequest) String() string {
	return dara.Prettify(s)
}

func (s SendBatchSmsRequest) GoString() string {
	return s.String()
}

func (s *SendBatchSmsRequest) GetOutId() *string {
	return s.OutId
}

func (s *SendBatchSmsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendBatchSmsRequest) GetPhoneNumberJson() *string {
	return s.PhoneNumberJson
}

func (s *SendBatchSmsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendBatchSmsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendBatchSmsRequest) GetSignNameJson() *string {
	return s.SignNameJson
}

func (s *SendBatchSmsRequest) GetSmsUpExtendCodeJson() *string {
	return s.SmsUpExtendCodeJson
}

func (s *SendBatchSmsRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendBatchSmsRequest) GetTemplateParamJson() *string {
	return s.TemplateParamJson
}

func (s *SendBatchSmsRequest) SetOutId(v string) *SendBatchSmsRequest {
	s.OutId = &v
	return s
}

func (s *SendBatchSmsRequest) SetOwnerId(v int64) *SendBatchSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *SendBatchSmsRequest) SetPhoneNumberJson(v string) *SendBatchSmsRequest {
	s.PhoneNumberJson = &v
	return s
}

func (s *SendBatchSmsRequest) SetResourceOwnerAccount(v string) *SendBatchSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendBatchSmsRequest) SetResourceOwnerId(v int64) *SendBatchSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendBatchSmsRequest) SetSignNameJson(v string) *SendBatchSmsRequest {
	s.SignNameJson = &v
	return s
}

func (s *SendBatchSmsRequest) SetSmsUpExtendCodeJson(v string) *SendBatchSmsRequest {
	s.SmsUpExtendCodeJson = &v
	return s
}

func (s *SendBatchSmsRequest) SetTemplateCode(v string) *SendBatchSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendBatchSmsRequest) SetTemplateParamJson(v string) *SendBatchSmsRequest {
	s.TemplateParamJson = &v
	return s
}

func (s *SendBatchSmsRequest) Validate() error {
	return dara.Validate(s)
}
