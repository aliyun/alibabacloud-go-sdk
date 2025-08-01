// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageWithTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *SendMessageWithTemplateRequest
	GetChannelId() *string
	SetFrom(v string) *SendMessageWithTemplateRequest
	GetFrom() *string
	SetSmsUpExtendCode(v string) *SendMessageWithTemplateRequest
	GetSmsUpExtendCode() *string
	SetTemplateCode(v string) *SendMessageWithTemplateRequest
	GetTemplateCode() *string
	SetTemplateParam(v string) *SendMessageWithTemplateRequest
	GetTemplateParam() *string
	SetTo(v string) *SendMessageWithTemplateRequest
	GetTo() *string
	SetValidityPeriod(v int64) *SendMessageWithTemplateRequest
	GetValidityPeriod() *int64
}

type SendMessageWithTemplateRequest struct {
	// The ID of the channel.
	//
	// example:
	//
	// 5739
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The signature. To query the signature, log on to the [Short Message Service (SMS) console](https://sms-intl.console.aliyun.com/overview) and navigate to the **Signatures*	- tab of the **Go China*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alicloud321
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The extension code of the MO message.
	//
	// example:
	//
	// 90999
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	// The code of the message template. To query the code, log on to the [SMS console](https://sms-intl.console.aliyun.com/overview) and navigate to the **Templates*	- tab of the **Go China*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The value of the variable in the message template. If a variable exists in the template, the parameter is required.
	//
	// example:
	//
	// {"code":"1234","product":"ytx"}
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
	// The mobile phone number to which the message is sent. You must add the country code to the beginning of the mobile phone number. Example: 861503871\\*\\*\\*\\*.
	//
	// For more information, see [Dialing codes](https://www.alibabacloud.com/help/en/sms/product-overview/dialing-codes?spm=a2c63.p38356.0.0.367279cbwQFoeM).
	//
	// This parameter is required.
	//
	// example:
	//
	// 861503871****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The validity period of the message.
	//
	// example:
	//
	// 1
	ValidityPeriod *int64 `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty"`
}

func (s SendMessageWithTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageWithTemplateRequest) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *SendMessageWithTemplateRequest) GetFrom() *string {
	return s.From
}

func (s *SendMessageWithTemplateRequest) GetSmsUpExtendCode() *string {
	return s.SmsUpExtendCode
}

func (s *SendMessageWithTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendMessageWithTemplateRequest) GetTemplateParam() *string {
	return s.TemplateParam
}

func (s *SendMessageWithTemplateRequest) GetTo() *string {
	return s.To
}

func (s *SendMessageWithTemplateRequest) GetValidityPeriod() *int64 {
	return s.ValidityPeriod
}

func (s *SendMessageWithTemplateRequest) SetChannelId(v string) *SendMessageWithTemplateRequest {
	s.ChannelId = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetFrom(v string) *SendMessageWithTemplateRequest {
	s.From = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetSmsUpExtendCode(v string) *SendMessageWithTemplateRequest {
	s.SmsUpExtendCode = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetTemplateCode(v string) *SendMessageWithTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetTemplateParam(v string) *SendMessageWithTemplateRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetTo(v string) *SendMessageWithTemplateRequest {
	s.To = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetValidityPeriod(v int64) *SendMessageWithTemplateRequest {
	s.ValidityPeriod = &v
	return s
}

func (s *SendMessageWithTemplateRequest) Validate() error {
	return dara.Validate(s)
}
