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
	// The objects of the message template.
	//
	// This parameter is required.
	CardObjects []*SendCardSmsRequestCardObjects `json:"CardObjects,omitempty" xml:"CardObjects,omitempty" type:"Repeated"`
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_70
	CardTemplateCode *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	// The code of the digital message template that applies when the card message is rolled back. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// example:
	//
	// SMS_003
	DigitalTemplateCode *string `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	// The variables of the digital message template.
	//
	// > If you need to add line breaks to the JSON template, make sure that the format is valid.
	//
	// example:
	//
	// {\\"msg\\",\\"xxxd\\"}
	DigitalTemplateParam *string `json:"DigitalTemplateParam,omitempty" xml:"DigitalTemplateParam,omitempty"`
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
	// 38d76c9b-4a9a-4c89-afae-61fd8e0e****
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The signature. You can view the template code in the **Signature*	- column on the **Signaturess*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > The signature must be approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The code of the text message template that applies when the card message is rolled back. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved. If you set the **FallbackType*	- parameter to **SMS**, this parameter is required.
	//
	// example:
	//
	// SIER_TEST_01
	SmsTemplateCode *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	// The variables of the text message template.
	//
	// > If you need to add line breaks to the JSON template, make sure that the format is valid.
	//
	// example:
	//
	// {\\"uri\\":\\"Zg11tZ\\"}
	SmsTemplateParam *string `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	// The extension code of the upstream message. Upstream messages are messages sent to the communication service provider. Upstream messages are used to customize a service, complete an inquiry, or send a request. You are charged for sending upstream messages based on the billing standards of the service provider.
	//
	// > If you do not need upstream messages, ignore this parameter.
	//
	// example:
	//
	// 1
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	// The code of the text message template.
	//
	// Log on to the Alibaba Cloud SMS console. In the left-side navigation pane, click **Go Globe*	- or **Go China**. You can view the message template in the **Template Code*	- column on the **Message Templates*	- tab.
	//
	// > The message templates must be created on the Go Globe page and approved.
	//
	// example:
	//
	// SMS_2322****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The variables of the message template. Format: JSON.
	//
	// > If you need to add line breaks to the JSON template, make sure that the format is valid.
	//
	// example:
	//
	// {
	//
	//       \\"code\\": \\"1111\\"
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
	return dara.Validate(s)
}

type SendCardSmsRequestCardObjects struct {
	// The URL to which the message is redirected if the message fails to be rendered.
	//
	// example:
	//
	// https://alibaba.com
	CustomUrl *string `json:"customUrl,omitempty" xml:"customUrl,omitempty"`
	// The variables. Special characters, such as $ and {}, do not need to be entered.
	//
	// example:
	//
	// {\\"param3\\":\\"three\\",\\"param1\\":\\"one\\",\\"param2\\":\\"two\\"}
	DyncParams *string `json:"dyncParams,omitempty" xml:"dyncParams,omitempty"`
	// The mobile phone number.
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
