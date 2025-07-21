// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappTemplateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetChatappTemplateDetailRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *GetChatappTemplateDetailRequest
	GetCustWabaId() *string
	SetIsvCode(v string) *GetChatappTemplateDetailRequest
	GetIsvCode() *string
	SetLanguage(v string) *GetChatappTemplateDetailRequest
	GetLanguage() *string
	SetTemplateCode(v string) *GetChatappTemplateDetailRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *GetChatappTemplateDetailRequest
	GetTemplateName() *string
	SetTemplateType(v string) *GetChatappTemplateDetailRequest
	GetTemplateType() *string
}

type GetChatappTemplateDetailRequest struct {
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// >  CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The independent software vendor (ISV) verification code. This parameter is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en_US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// ****4b5c79c9432497a075bdfca36bf5
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Name of a template.
	//
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message template. Valid values:
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE (developing)
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetChatappTemplateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetChatappTemplateDetailRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *GetChatappTemplateDetailRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *GetChatappTemplateDetailRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetChatappTemplateDetailRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetChatappTemplateDetailRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetChatappTemplateDetailRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetChatappTemplateDetailRequest) SetCustSpaceId(v string) *GetChatappTemplateDetailRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetCustWabaId(v string) *GetChatappTemplateDetailRequest {
	s.CustWabaId = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetIsvCode(v string) *GetChatappTemplateDetailRequest {
	s.IsvCode = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetLanguage(v string) *GetChatappTemplateDetailRequest {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetTemplateCode(v string) *GetChatappTemplateDetailRequest {
	s.TemplateCode = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetTemplateName(v string) *GetChatappTemplateDetailRequest {
	s.TemplateName = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetTemplateType(v string) *GetChatappTemplateDetailRequest {
	s.TemplateType = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) Validate() error {
	return dara.Validate(s)
}
