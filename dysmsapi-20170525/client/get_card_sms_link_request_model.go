// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardSmsLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCardCodeType(v int32) *GetCardSmsLinkRequest
	GetCardCodeType() *int32
	SetCardLinkType(v int32) *GetCardSmsLinkRequest
	GetCardLinkType() *int32
	SetCardTemplateCode(v string) *GetCardSmsLinkRequest
	GetCardTemplateCode() *string
	SetCardTemplateParamJson(v string) *GetCardSmsLinkRequest
	GetCardTemplateParamJson() *string
	SetCustomShortCodeJson(v string) *GetCardSmsLinkRequest
	GetCustomShortCodeJson() *string
	SetDomain(v string) *GetCardSmsLinkRequest
	GetDomain() *string
	SetOutId(v string) *GetCardSmsLinkRequest
	GetOutId() *string
	SetPhoneNumberJson(v string) *GetCardSmsLinkRequest
	GetPhoneNumberJson() *string
	SetSignNameJson(v string) *GetCardSmsLinkRequest
	GetSignNameJson() *string
}

type GetCardSmsLinkRequest struct {
	// The encoding type of the short URL for the card message. Valid values:
	//
	// - 1: bulk sending.
	//
	// - 2: personalized.
	//
	// example:
	//
	// 2
	CardCodeType *int32 `json:"CardCodeType,omitempty" xml:"CardCodeType,omitempty"`
	// The type of the short URL for the card message. Valid values:
	//
	// - 1: standard short URL.
	//
	// - 2: custom short URL.
	//
	// > If **CardLinkType*	- is left empty, the default value is standard short URL. To generate a custom short URL, contact Alibaba Cloud operations to register in advance.
	//
	// example:
	//
	// 1
	CardLinkType *int32 `json:"CardLinkType,omitempty" xml:"CardLinkType,omitempty"`
	// The code of the card message template. In the console, go to the [Card Messages > Template Management](https://dysms.console.aliyun.com/domestic/card) page and select the code of an approved card message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_****
	CardTemplateCode *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	// The variables of the card message template.
	//
	// example:
	//
	// [{\\"customUrl\\":\\"https://alibaba.com\\",\\"dyncParams\\": \\"{\\\\\\"动参key\\\\\\":\\\\\\"动参value\\\\\\"}\\"},{\\"customUrl\\":\\"https://alibaba.com\\",\\"dyncParams\\": \\"{\\\\\\"动参key\\\\\\":\\\\\\"动参value\\\\\\"}\\"}]
	CardTemplateParamJson *string `json:"CardTemplateParamJson,omitempty" xml:"CardTemplateParamJson,omitempty"`
	// The custom short code. The value must be 4 to 8 digits or letters.
	//
	// > This parameter is required when the generation type is custom short URL.
	//
	// example:
	//
	// abCde
	CustomShortCodeJson *string `json:"CustomShortCodeJson,omitempty" xml:"CustomShortCodeJson,omitempty"`
	// The short URL domain assigned to the sending account. The domain must be registered in advance.
	//
	// > - When **CardLinkType*	- is set to **2**, the **Domain*	- parameter is required.
	//
	// > - If the **Domain*	- parameter is left empty, the system default domain is used. The value can be up to 100 characters in length.
	//
	// example:
	//
	// xxx.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The external extension field.
	//
	// example:
	//
	// BC20220608102511660860762****
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The phone number, user ID, or internal system identifier.
	//
	// > - Supports up to 10,000 phone numbers.
	//
	// > - You can also specify a custom identifier of up to 60 characters.
	//
	// > - For OPPO templates, you can submit up to 10 requests at a time.
	//
	// example:
	//
	// [\\"1390000****
	//
	// \\",\\"1370000****
	//
	// \\"]
	PhoneNumberJson *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	// The signature name of the SMS message.
	//
	// In the console, go to the [Domestic Messages > Signature Management](https://dysms.console.aliyun.com/domestic/text/sign) tab and view the name in the **Signature Name*	- column. You can also call the [QuerySmsSignList](https://www.alibabacloud.com/help/en/sms/developer-reference/api-dysmsapi-2017-05-25-querysmssignlist) operation to view SMS signature names.
	//
	// > The signature must be added and approved. The number of SMS signatures must match the number of phone numbers, and each signature must correspond to a phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"阿里云\\", \\"阿里云2\\"]
	SignNameJson *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
}

func (s GetCardSmsLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsLinkRequest) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkRequest) GetCardCodeType() *int32 {
	return s.CardCodeType
}

func (s *GetCardSmsLinkRequest) GetCardLinkType() *int32 {
	return s.CardLinkType
}

func (s *GetCardSmsLinkRequest) GetCardTemplateCode() *string {
	return s.CardTemplateCode
}

func (s *GetCardSmsLinkRequest) GetCardTemplateParamJson() *string {
	return s.CardTemplateParamJson
}

func (s *GetCardSmsLinkRequest) GetCustomShortCodeJson() *string {
	return s.CustomShortCodeJson
}

func (s *GetCardSmsLinkRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetCardSmsLinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *GetCardSmsLinkRequest) GetPhoneNumberJson() *string {
	return s.PhoneNumberJson
}

func (s *GetCardSmsLinkRequest) GetSignNameJson() *string {
	return s.SignNameJson
}

func (s *GetCardSmsLinkRequest) SetCardCodeType(v int32) *GetCardSmsLinkRequest {
	s.CardCodeType = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetCardLinkType(v int32) *GetCardSmsLinkRequest {
	s.CardLinkType = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetCardTemplateCode(v string) *GetCardSmsLinkRequest {
	s.CardTemplateCode = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetCardTemplateParamJson(v string) *GetCardSmsLinkRequest {
	s.CardTemplateParamJson = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetCustomShortCodeJson(v string) *GetCardSmsLinkRequest {
	s.CustomShortCodeJson = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetDomain(v string) *GetCardSmsLinkRequest {
	s.Domain = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetOutId(v string) *GetCardSmsLinkRequest {
	s.OutId = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetPhoneNumberJson(v string) *GetCardSmsLinkRequest {
	s.PhoneNumberJson = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetSignNameJson(v string) *GetCardSmsLinkRequest {
	s.SignNameJson = &v
	return s
}

func (s *GetCardSmsLinkRequest) Validate() error {
	return dara.Validate(s)
}
