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
	// The code type of the URLs.
	//
	// 	- **1**: group texting
	//
	// 	- **2**: personalization
	//
	// example:
	//
	// 2
	CardCodeType *int32 `json:"CardCodeType,omitempty" xml:"CardCodeType,omitempty"`
	// The type of the short URLs.
	//
	// 	- 1: standard short code.
	//
	// 	- 2: custom short code.
	//
	// > If the **CardLinkType*	- is not specified, standard short codes are generated. If you need to generate custom short codes, contact Alibaba Cloud SMS technical support.
	//
	// example:
	//
	// 1
	CardLinkType *int32 `json:"CardLinkType,omitempty" xml:"CardLinkType,omitempty"`
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_****
	CardTemplateCode *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	// The variables of the message template.
	//
	// example:
	//
	// [{},{}]
	CardTemplateParamJson *string `json:"CardTemplateParamJson,omitempty" xml:"CardTemplateParamJson,omitempty"`
	// The custom short code. It can contain 4 to 8 digits or letters.
	//
	// > If the CardLinkType parameter is set to 2, the CustomShortCodeJson parameter is required.
	//
	// example:
	//
	// abCde
	CustomShortCodeJson *string `json:"CustomShortCodeJson,omitempty" xml:"CustomShortCodeJson,omitempty"`
	// The original domain name. You must submit domain names for approval in advance.
	//
	// >
	//
	// 	- If the **CardLinkType*	- parameter is set to **2**, the **Domain*	- parameter is required.
	//
	// 	- The **Domain*	- parameter cannot exceed 100 characters in length. If the parameter is not specified, a default domain name is used.
	//
	// example:
	//
	// xxx.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The extension field.
	//
	// example:
	//
	// BC20220608102511660860762****
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The mobile phone numbers of recipients, custom identifiers, or system identifiers.
	//
	// >
	//
	// 	- A maximum of 10,000 mobile phone numbers are supported.
	//
	// 	- You can enter custom identifier. Each identifier can be a maximum of 60 characters in length.
	//
	// 	- You can apply for a maximum of 10 OPPO templates at a time.
	//
	// example:
	//
	// [\\"1390000****
	//
	// \\",\\"1370000****
	//
	// \\"]
	PhoneNumberJson *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	// The signature. You can view the template code in the **Signature*	- column on the **Signaturess*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > The signatures must be approved and correspond to the mobile numbers in sequence.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"aliyun\\", \\"aliyun2\\"]
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
