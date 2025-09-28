// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSmsVerifyCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRetry(v int64) *SendSmsVerifyCodeRequest
	GetAutoRetry() *int64
	SetCodeLength(v int64) *SendSmsVerifyCodeRequest
	GetCodeLength() *int64
	SetCodeType(v int64) *SendSmsVerifyCodeRequest
	GetCodeType() *int64
	SetCountryCode(v string) *SendSmsVerifyCodeRequest
	GetCountryCode() *string
	SetDuplicatePolicy(v int64) *SendSmsVerifyCodeRequest
	GetDuplicatePolicy() *int64
	SetInterval(v int64) *SendSmsVerifyCodeRequest
	GetInterval() *int64
	SetOutId(v string) *SendSmsVerifyCodeRequest
	GetOutId() *string
	SetOwnerId(v int64) *SendSmsVerifyCodeRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *SendSmsVerifyCodeRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *SendSmsVerifyCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendSmsVerifyCodeRequest
	GetResourceOwnerId() *int64
	SetReturnVerifyCode(v bool) *SendSmsVerifyCodeRequest
	GetReturnVerifyCode() *bool
	SetSchemeName(v string) *SendSmsVerifyCodeRequest
	GetSchemeName() *string
	SetSignName(v string) *SendSmsVerifyCodeRequest
	GetSignName() *string
	SetSmsUpExtendCode(v string) *SendSmsVerifyCodeRequest
	GetSmsUpExtendCode() *string
	SetTemplateCode(v string) *SendSmsVerifyCodeRequest
	GetTemplateCode() *string
	SetTemplateParam(v string) *SendSmsVerifyCodeRequest
	GetTemplateParam() *string
	SetValidTime(v int64) *SendSmsVerifyCodeRequest
	GetValidTime() *int64
}

type SendSmsVerifyCodeRequest struct {
	// example:
	//
	// 是否自动重试
	AutoRetry *int64 `json:"AutoRetry,omitempty" xml:"AutoRetry,omitempty"`
	// The length of the verification code. Default value: 4. Valid values: 4 to 8.
	//
	// example:
	//
	// 4
	CodeLength *int64 `json:"CodeLength,omitempty" xml:"CodeLength,omitempty"`
	// The type of the generated verification code. Default value: 1. Valid values:
	//
	// 	- 1: digits only
	//
	// 	- 2: uppercase letters only
	//
	// 	- 3: lowercase letters only
	//
	// 	- 4: uppercase and lowercase letters
	//
	// 	- 5: digits and uppercase letters
	//
	// 	- 6: digits and lowercase letters
	//
	// 	- 7: digits and uppercase and lowercase letters
	//
	// example:
	//
	// 1
	CodeType *int64 `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	// The country code of the phone number. SMS verification codes can be sent only by using phone numbers in the Chinese mainland. Default value: 86.
	//
	// example:
	//
	// 86
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// Specifies how to handle the verification codes received earlier in a case where verification codes are sent to the same phone number for the same scenario within the validity period.
	//
	// 	- 1 (default): The latest verification code overwrites the verification codes received earlier. In this case, verification codes received earlier expire.
	//
	// 	- 2: Verification codes within their validity period are valid and can be used for verification.
	//
	// example:
	//
	// 1
	DuplicatePolicy *int64 `json:"DuplicatePolicy,omitempty" xml:"DuplicatePolicy,omitempty"`
	// The time interval. Unit: seconds. Default value: 60. This parameter specifies how often you can send a verification code.
	//
	// example:
	//
	// 60
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 12358794Aqzaq
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86130****0000
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to return a verification code.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ReturnVerifyCode *bool `json:"ReturnVerifyCode,omitempty" xml:"ReturnVerifyCode,omitempty"`
	// The verification service name. If this parameter is not specified, the default service is used. The name can be up to 20 characters in length.
	//
	// example:
	//
	// Aliyun
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// The signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun Test
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The extension code of the upstream text message. Upstream text messages are text messages sent to the communication service provider. Upstream text messages are used to customize a service, complete an inquiry, or send a request. You are charged for sending upstream text messages based on the billing standards of the service provider.
	//
	// >  The extension code is automatically generated by the system when the signature is generated. You do not need to specify the extension code. You can skip this parameter based on your business requirements. If you want to use custom extension codes, contact your account manager.
	//
	// example:
	//
	// 1213123
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	// The code of the text message template.
	//
	// Log on to the [SMS console](https://dysms.console.aliyun.com/dysms.htm?spm=5176.12818093.categories-n-products.ddysms.3b2816d0xml2NA#/overview). In the left-side navigation pane, click **Go China*	- or **Go Globe**. You can view the text message template code in the **Template Code*	- column on the **Message Templates*	- tab.
	//
	// >  The text message templates must be created on the Go Globe page and approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// azsq_*****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The value of the variable in the text message template. The verification code is replaced with "##code##".
	//
	// Example 1: For a system-defined template that contains variables, if the template content is "Your verification code is ${code} and valid for 5 minutes. Do not disclose the verification code to others.", specify the value of this parameter as {"code":"##code##"}
	//
	// Example 2: For a custom template, if the template content is ${content}, specify the value of this parameter as {"content":"Your verification code is ##code## and must be used within 5 minutes."}.
	//
	// >
	//
	// 	- If line breaks are required in JSON-formatted data, they must meet the relevant requirements that are specified in the standard JSON protocol.
	//
	// 	- For more information about template variables, see [SMS template specifications](https://help.aliyun.com/document_detail/108253.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"code":"##code##"}
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
	// The validity period of the verification code. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	ValidTime *int64 `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
}

func (s SendSmsVerifyCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s SendSmsVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *SendSmsVerifyCodeRequest) GetAutoRetry() *int64 {
	return s.AutoRetry
}

func (s *SendSmsVerifyCodeRequest) GetCodeLength() *int64 {
	return s.CodeLength
}

func (s *SendSmsVerifyCodeRequest) GetCodeType() *int64 {
	return s.CodeType
}

func (s *SendSmsVerifyCodeRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *SendSmsVerifyCodeRequest) GetDuplicatePolicy() *int64 {
	return s.DuplicatePolicy
}

func (s *SendSmsVerifyCodeRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *SendSmsVerifyCodeRequest) GetOutId() *string {
	return s.OutId
}

func (s *SendSmsVerifyCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendSmsVerifyCodeRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *SendSmsVerifyCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendSmsVerifyCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendSmsVerifyCodeRequest) GetReturnVerifyCode() *bool {
	return s.ReturnVerifyCode
}

func (s *SendSmsVerifyCodeRequest) GetSchemeName() *string {
	return s.SchemeName
}

func (s *SendSmsVerifyCodeRequest) GetSignName() *string {
	return s.SignName
}

func (s *SendSmsVerifyCodeRequest) GetSmsUpExtendCode() *string {
	return s.SmsUpExtendCode
}

func (s *SendSmsVerifyCodeRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendSmsVerifyCodeRequest) GetTemplateParam() *string {
	return s.TemplateParam
}

func (s *SendSmsVerifyCodeRequest) GetValidTime() *int64 {
	return s.ValidTime
}

func (s *SendSmsVerifyCodeRequest) SetAutoRetry(v int64) *SendSmsVerifyCodeRequest {
	s.AutoRetry = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetCodeLength(v int64) *SendSmsVerifyCodeRequest {
	s.CodeLength = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetCodeType(v int64) *SendSmsVerifyCodeRequest {
	s.CodeType = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetCountryCode(v string) *SendSmsVerifyCodeRequest {
	s.CountryCode = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetDuplicatePolicy(v int64) *SendSmsVerifyCodeRequest {
	s.DuplicatePolicy = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetInterval(v int64) *SendSmsVerifyCodeRequest {
	s.Interval = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetOutId(v string) *SendSmsVerifyCodeRequest {
	s.OutId = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetOwnerId(v int64) *SendSmsVerifyCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetPhoneNumber(v string) *SendSmsVerifyCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetResourceOwnerAccount(v string) *SendSmsVerifyCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetResourceOwnerId(v int64) *SendSmsVerifyCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetReturnVerifyCode(v bool) *SendSmsVerifyCodeRequest {
	s.ReturnVerifyCode = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetSchemeName(v string) *SendSmsVerifyCodeRequest {
	s.SchemeName = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetSignName(v string) *SendSmsVerifyCodeRequest {
	s.SignName = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetSmsUpExtendCode(v string) *SendSmsVerifyCodeRequest {
	s.SmsUpExtendCode = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetTemplateCode(v string) *SendSmsVerifyCodeRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetTemplateParam(v string) *SendSmsVerifyCodeRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetValidTime(v int64) *SendSmsVerifyCodeRequest {
	s.ValidTime = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) Validate() error {
	return dara.Validate(s)
}
