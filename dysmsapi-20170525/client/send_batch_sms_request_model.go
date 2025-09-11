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
	// The extension field of the external record. The value is a string that contains no more than 256 characters.
	//
	// > The parameter is optional.
	//
	// example:
	//
	// abcdefg
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The mobile number of the recipient. Format:
	//
	// 	- Message delivery to the Chinese mainland: +/+86/0086/86 or an 11-digit mobile number without a prefix. Example: 1590000\\*\\*\\*\\*.
	//
	// 	- Message delivery to countries or regions outside the Chinese mainland: Dialing code + Mobile number. Example: 852000012\\*\\*\\*\\*.
	//
	// > We recommend that you call the SendSms operation to send verification codes.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1590000****","1350000****"]
	PhoneNumberJson      *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature.
	//
	// Log on to the Alibaba Cloud SMS console. In the left-side navigation pane, click **Go Globe*	- or **Go China**. You can view the signature in the **Signature*	- column on the **Signatures*	- tab.
	//
	// > The signatures must be approved and correspond to the mobile numbers in sequence.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["Aliyun","Alibaba"]
	SignNameJson *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	// The extension code of the MO message. Format: JSON array.
	//
	// > The parameter is optional.
	//
	// example:
	//
	// ["90999","90998"]
	SmsUpExtendCodeJson *string `json:"SmsUpExtendCodeJson,omitempty" xml:"SmsUpExtendCodeJson,omitempty"`
	// The code of the message template.
	//
	// Log on to the Alibaba Cloud SMS console. In the left-side navigation pane, click **Go Globe*	- or **Go China**. You can view the message template in the **Template Code*	- column on the **Message Templates*	- tab.
	//
	// > The message templates must be created on the Go Globe page and approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_15255****
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
