// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatappTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowCategoryChange(v bool) *CreateChatappTemplateShrinkRequest
	GetAllowCategoryChange() *bool
	SetCategory(v string) *CreateChatappTemplateShrinkRequest
	GetCategory() *string
	SetCategoryChangePaused(v bool) *CreateChatappTemplateShrinkRequest
	GetCategoryChangePaused() *bool
	SetComponentsShrink(v string) *CreateChatappTemplateShrinkRequest
	GetComponentsShrink() *string
	SetCustSpaceId(v string) *CreateChatappTemplateShrinkRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *CreateChatappTemplateShrinkRequest
	GetCustWabaId() *string
	SetExampleShrink(v string) *CreateChatappTemplateShrinkRequest
	GetExampleShrink() *string
	SetIsvCode(v string) *CreateChatappTemplateShrinkRequest
	GetIsvCode() *string
	SetLanguage(v string) *CreateChatappTemplateShrinkRequest
	GetLanguage() *string
	SetMessageSendTtlSeconds(v int32) *CreateChatappTemplateShrinkRequest
	GetMessageSendTtlSeconds() *int32
	SetName(v string) *CreateChatappTemplateShrinkRequest
	GetName() *string
	SetProductSetId(v string) *CreateChatappTemplateShrinkRequest
	GetProductSetId() *string
	SetTemplateType(v string) *CreateChatappTemplateShrinkRequest
	GetTemplateType() *string
}

type CreateChatappTemplateShrinkRequest struct {
	// Specifies whether to allow Facebook to automatically change the template category (to improve the template approval rate). This property is valid only when TemplateType is set to WHATSAPP.
	//
	// 	Notice: This property has been deprecated. WhatsApp no longer supports this property.</notice>
	//
	// example:
	//
	// true
	AllowCategoryChange *bool `json:"AllowCategoryChange,omitempty" xml:"AllowCategoryChange,omitempty"`
	// WhatsApp template category. Valid values:
	//
	// - **UTILITY**: transaction-related.
	//
	// - **MARKETING**: marketing.
	//
	// - **AUTHENTICATION**: identity verification.
	//
	// Viber template category. Valid values:
	//
	// - **UTILITY**: transaction-related.
	//
	// - **MARKETING**: marketing.
	//
	// - **AUTHENTICATION**: identity verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// UTILITY
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to pause template sending when a Utility template is changed to Marketing type. This property is valid only for WhatsApp templates.
	//
	// example:
	//
	// false
	CategoryChangePaused *bool `json:"CategoryChangePaused,omitempty" xml:"CategoryChangePaused,omitempty"`
	// The list of message template components.
	//
	// > When Category=AUTHENTICATION, Components cannot contain nodes with Type=HEADER. When Type=BODY or FOOTER, the Text content must be empty.
	//
	// This parameter is required.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The SpaceId of the ISV sub-customer or the direct customer instance ID.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The ISV customer WabaId.
	//
	// > Deprecated parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The example for creating the template.
	//
	// example:
	//
	// hello_whatsapp
	ExampleShrink *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// Deprecated
	//
	// The ISV verification code, used to verify whether the sub-account is authorized by the ISV.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The template language. For detailed language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The time-to-live (TTL) for template messages in WhatsApp.
	//
	// - AUTHENTICATION: valid values range from 30 to 900.
	//
	// - UTILITY: valid values range from 30 to 43200.
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// hello_whatsapp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// productSetId
	//
	// example:
	//
	// 9928**
	ProductSetId *string `json:"ProductSetId,omitempty" xml:"ProductSetId,omitempty"`
	// The templatetype. Valid values:
	//
	// - **WHATSAPP**
	//
	// - **VIBER**
	//
	// This parameter is required.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateChatappTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateShrinkRequest) GetAllowCategoryChange() *bool {
	return s.AllowCategoryChange
}

func (s *CreateChatappTemplateShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateChatappTemplateShrinkRequest) GetCategoryChangePaused() *bool {
	return s.CategoryChangePaused
}

func (s *CreateChatappTemplateShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *CreateChatappTemplateShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreateChatappTemplateShrinkRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *CreateChatappTemplateShrinkRequest) GetExampleShrink() *string {
	return s.ExampleShrink
}

func (s *CreateChatappTemplateShrinkRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *CreateChatappTemplateShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateChatappTemplateShrinkRequest) GetMessageSendTtlSeconds() *int32 {
	return s.MessageSendTtlSeconds
}

func (s *CreateChatappTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateChatappTemplateShrinkRequest) GetProductSetId() *string {
	return s.ProductSetId
}

func (s *CreateChatappTemplateShrinkRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateChatappTemplateShrinkRequest) SetAllowCategoryChange(v bool) *CreateChatappTemplateShrinkRequest {
	s.AllowCategoryChange = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetCategory(v string) *CreateChatappTemplateShrinkRequest {
	s.Category = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetCategoryChangePaused(v bool) *CreateChatappTemplateShrinkRequest {
	s.CategoryChangePaused = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetComponentsShrink(v string) *CreateChatappTemplateShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetCustSpaceId(v string) *CreateChatappTemplateShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetCustWabaId(v string) *CreateChatappTemplateShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetExampleShrink(v string) *CreateChatappTemplateShrinkRequest {
	s.ExampleShrink = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetIsvCode(v string) *CreateChatappTemplateShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetLanguage(v string) *CreateChatappTemplateShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetMessageSendTtlSeconds(v int32) *CreateChatappTemplateShrinkRequest {
	s.MessageSendTtlSeconds = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetName(v string) *CreateChatappTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetProductSetId(v string) *CreateChatappTemplateShrinkRequest {
	s.ProductSetId = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetTemplateType(v string) *CreateChatappTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
