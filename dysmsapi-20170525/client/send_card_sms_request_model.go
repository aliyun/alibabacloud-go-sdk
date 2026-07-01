// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCardSmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCardObjects(v []*SendCardSmsRequestCardObjects) *SendCardSmsRequest
	GetCardObjects() []*SendCardSmsRequestCardObjects
	SetCardTemplateCode(v string) *SendCardSmsRequest
	GetCardTemplateCode() *string
	SetDigitalTemplateCode(v string) *SendCardSmsRequest
	GetDigitalTemplateCode() *string
	SetDigitalTemplateParam(v string) *SendCardSmsRequest
	GetDigitalTemplateParam() *string
	SetFallbackType(v string) *SendCardSmsRequest
	GetFallbackType() *string
	SetOutId(v string) *SendCardSmsRequest
	GetOutId() *string
	SetSignName(v string) *SendCardSmsRequest
	GetSignName() *string
	SetSmsTemplateCode(v string) *SendCardSmsRequest
	GetSmsTemplateCode() *string
	SetSmsTemplateParam(v string) *SendCardSmsRequest
	GetSmsTemplateParam() *string
	SetSmsUpExtendCode(v string) *SendCardSmsRequest
	GetSmsUpExtendCode() *string
	SetTemplateCode(v string) *SendCardSmsRequest
	GetTemplateCode() *string
	SetTemplateParam(v string) *SendCardSmsRequest
	GetTemplateParam() *string
}

type SendCardSmsRequest struct {
	// The card message objects.
	//
	// This parameter is required.
	//
	// example:
	//
	// SendCardSms
	CardObjects []*SendCardSmsRequestCardObjects `json:"CardObjects,omitempty" xml:"CardObjects,omitempty" type:"Repeated"`
	// The code of the card message template. On the [Template Management](https://dysms.console.aliyun.com/domestic/card) page of the **Card Messages*	- module in the console, select the code of an approved card message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_70
	CardTemplateCode *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	// The code of the fallback digital message template. This parameter is required if you set **FallbackType*	- to **DIGITALSMS**.
	//
	// You can view the digital message template list on the [Template Management](https://dysms.console.aliyun.com/domestic/digit) page of the **Digital Messages*	- module in the console.
	//
	// > The template must be added and approved.
	//
	// example:
	//
	// DIGITAL_SMS_31359****
	DigitalTemplateCode *string `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	// The actual values of the variables in the fallback digital message template. This parameter is required if the digital message template specified by **DigitalTemplateCode*	- contains variables.
	//
	// > If the JSON value contains line breaks, follow the standard JSON protocol.
	//
	// example:
	//
	// {"msg","xxxd"}
	DigitalTemplateParam *string `json:"DigitalTemplateParam,omitempty" xml:"DigitalTemplateParam,omitempty"`
	// The fallback type. Valid values:
	//
	// - **SMS**: Falls back to a text message for phone numbers that do not support card messages.
	//
	// - **DIGITALSMS**: Falls back to a digital message for phone numbers that do not support card messages.
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
	// 38d76c9b-4a9a-4c89-afae-61fd8e0e****
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The signature name. You can call the [QuerySmsSignList](https://help.aliyun.com/document_detail/419282.html) operation to query the signatures applied for under the current account or view the signature list in the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/domestic/text/sign).
	//
	// > The signature must be approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The code of the fallback text message template. This parameter is required if you set **FallbackType*	- to **SMS**.
	//
	// You can call the [QuerySmsTemplateList](https://help.aliyun.com/document_detail/419288.html) operation to query the templates applied for under the current account or view the template list in the [SMS console](https://dysms.console.aliyun.com/domestic/text/template).
	//
	// > The template must be added and approved.
	//
	// example:
	//
	// SMS_48068****
	SmsTemplateCode *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	// The actual values of the variables in the fallback text message template. This parameter is required if the text message template specified by **SmsTemplateCode*	- contains variables.
	//
	// > If the JSON value contains line breaks, follow the standard JSON protocol.
	//
	// example:
	//
	// {"jifen":"积分"}
	SmsTemplateParam *string `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	// The extension code of the MO message. An MO message is a message sent to the communications service provider to customize a service, perform a query, or handle other business. The message is charged at the standard rate of the carrier.
	//
	// > If you do not have such requirements, ignore this parameter.
	//
	// example:
	//
	// 1
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	// The code of the custom content template.
	//
	// The custom content is sent to the recipient as a text message template combined with a card parsing link. Log on to the [SMS console](https://dysms.console.aliyun.com/overview), choose **Domestic Messages*	- or **International/HK/MO/TW Messages**, and view the **Template Code*	- on the **Template Management*	- tab.
	//
	// > - The template code must be added and approved. To send international or Hong Kong, Macao, or Taiwan messages, use an international or Hong Kong, Macao, or Taiwan message template.
	//
	// > - For example, if the selected text message template is "You have a new message" and the card parsing link is `1*.cn/2**d`, the final content is `You have a new message 1*.cn/2**d`. Test the message and control the word count before sending.
	//
	// example:
	//
	// SMS_2322****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The actual values of the variables in the custom content template. This parameter is required if the message template specified by **TemplateCode*	- contains variables.
	//
	// > If the JSON value contains line breaks, follow the standard JSON protocol.
	//
	// example:
	//
	// {
	//
	//       "code": "1111"
	//
	// }
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
}

func (s SendCardSmsRequest) String() string {
	return dara.Prettify(s)
}

func (s SendCardSmsRequest) GoString() string {
	return s.String()
}

func (s *SendCardSmsRequest) GetCardObjects() []*SendCardSmsRequestCardObjects {
	return s.CardObjects
}

func (s *SendCardSmsRequest) GetCardTemplateCode() *string {
	return s.CardTemplateCode
}

func (s *SendCardSmsRequest) GetDigitalTemplateCode() *string {
	return s.DigitalTemplateCode
}

func (s *SendCardSmsRequest) GetDigitalTemplateParam() *string {
	return s.DigitalTemplateParam
}

func (s *SendCardSmsRequest) GetFallbackType() *string {
	return s.FallbackType
}

func (s *SendCardSmsRequest) GetOutId() *string {
	return s.OutId
}

func (s *SendCardSmsRequest) GetSignName() *string {
	return s.SignName
}

func (s *SendCardSmsRequest) GetSmsTemplateCode() *string {
	return s.SmsTemplateCode
}

func (s *SendCardSmsRequest) GetSmsTemplateParam() *string {
	return s.SmsTemplateParam
}

func (s *SendCardSmsRequest) GetSmsUpExtendCode() *string {
	return s.SmsUpExtendCode
}

func (s *SendCardSmsRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendCardSmsRequest) GetTemplateParam() *string {
	return s.TemplateParam
}

func (s *SendCardSmsRequest) SetCardObjects(v []*SendCardSmsRequestCardObjects) *SendCardSmsRequest {
	s.CardObjects = v
	return s
}

func (s *SendCardSmsRequest) SetCardTemplateCode(v string) *SendCardSmsRequest {
	s.CardTemplateCode = &v
	return s
}

func (s *SendCardSmsRequest) SetDigitalTemplateCode(v string) *SendCardSmsRequest {
	s.DigitalTemplateCode = &v
	return s
}

func (s *SendCardSmsRequest) SetDigitalTemplateParam(v string) *SendCardSmsRequest {
	s.DigitalTemplateParam = &v
	return s
}

func (s *SendCardSmsRequest) SetFallbackType(v string) *SendCardSmsRequest {
	s.FallbackType = &v
	return s
}

func (s *SendCardSmsRequest) SetOutId(v string) *SendCardSmsRequest {
	s.OutId = &v
	return s
}

func (s *SendCardSmsRequest) SetSignName(v string) *SendCardSmsRequest {
	s.SignName = &v
	return s
}

func (s *SendCardSmsRequest) SetSmsTemplateCode(v string) *SendCardSmsRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *SendCardSmsRequest) SetSmsTemplateParam(v string) *SendCardSmsRequest {
	s.SmsTemplateParam = &v
	return s
}

func (s *SendCardSmsRequest) SetSmsUpExtendCode(v string) *SendCardSmsRequest {
	s.SmsUpExtendCode = &v
	return s
}

func (s *SendCardSmsRequest) SetTemplateCode(v string) *SendCardSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendCardSmsRequest) SetTemplateParam(v string) *SendCardSmsRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendCardSmsRequest) Validate() error {
	if s.CardObjects != nil {
		for _, item := range s.CardObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendCardSmsRequestCardObjects struct {
	// 渲染失败后跳转链接。
	//
	// example:
	//
	// https://alibaba.com
	CustomUrl *string `json:"customUrl,omitempty" xml:"customUrl,omitempty"`
	// 动态参数。动参变量不需要${}
	//
	// example:
	//
	// {"param3":"李四3","param1":"李四","param2":"李四2"}
	DyncParams *string `json:"dyncParams,omitempty" xml:"dyncParams,omitempty"`
	// 接收卡片短信的手机号码。
	//
	// example:
	//
	// 1390000****
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s SendCardSmsRequestCardObjects) String() string {
	return dara.Prettify(s)
}

func (s SendCardSmsRequestCardObjects) GoString() string {
	return s.String()
}

func (s *SendCardSmsRequestCardObjects) GetCustomUrl() *string {
	return s.CustomUrl
}

func (s *SendCardSmsRequestCardObjects) GetDyncParams() *string {
	return s.DyncParams
}

func (s *SendCardSmsRequestCardObjects) GetMobile() *string {
	return s.Mobile
}

func (s *SendCardSmsRequestCardObjects) SetCustomUrl(v string) *SendCardSmsRequestCardObjects {
	s.CustomUrl = &v
	return s
}

func (s *SendCardSmsRequestCardObjects) SetDyncParams(v string) *SendCardSmsRequestCardObjects {
	s.DyncParams = &v
	return s
}

func (s *SendCardSmsRequestCardObjects) SetMobile(v string) *SendCardSmsRequestCardObjects {
	s.Mobile = &v
	return s
}

func (s *SendCardSmsRequestCardObjects) Validate() error {
	return dara.Validate(s)
}
