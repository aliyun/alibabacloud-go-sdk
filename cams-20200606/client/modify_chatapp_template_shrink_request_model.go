// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChatappTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ModifyChatappTemplateShrinkRequest
	GetCategory() *string
	SetCategoryChangePaused(v bool) *ModifyChatappTemplateShrinkRequest
	GetCategoryChangePaused() *bool
	SetComponentsShrink(v string) *ModifyChatappTemplateShrinkRequest
	GetComponentsShrink() *string
	SetCustSpaceId(v string) *ModifyChatappTemplateShrinkRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *ModifyChatappTemplateShrinkRequest
	GetCustWabaId() *string
	SetExampleShrink(v string) *ModifyChatappTemplateShrinkRequest
	GetExampleShrink() *string
	SetIsvCode(v string) *ModifyChatappTemplateShrinkRequest
	GetIsvCode() *string
	SetLanguage(v string) *ModifyChatappTemplateShrinkRequest
	GetLanguage() *string
	SetMessageSendTtlSeconds(v int32) *ModifyChatappTemplateShrinkRequest
	GetMessageSendTtlSeconds() *int32
	SetTemplateCode(v string) *ModifyChatappTemplateShrinkRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *ModifyChatappTemplateShrinkRequest
	GetTemplateName() *string
	SetTemplateType(v string) *ModifyChatappTemplateShrinkRequest
	GetTemplateType() *string
}

type ModifyChatappTemplateShrinkRequest struct {
	// The templatetype cannot be modified.
	//
	// example:
	//
	// text
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to pause sending when a Utility template is changed to Marketing type.
	//
	// example:
	//
	// 120
	CategoryChangePaused *bool `json:"CategoryChangePaused,omitempty" xml:"CategoryChangePaused,omitempty"`
	// The list of message template components.
	//
	// > When Category is set to AUTHENTICATION, Components cannot contain a node with Type set to HEADER. When Type is set to BODY or FOOTER and the Text content is empty, the content is automatically generated.
	//
	// This parameter is required.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The SpaceId of the ISV sub-customer or the instance ID of a direct customer.
	//
	// example:
	//
	// 28251486512358****
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
	// The examples for creating the template.
	ExampleShrink *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The ISV verification code used to verify whether the RAM user is authorized by the ISV.
	//
	// example:
	//
	// ksiekdki39ksks93939
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The template language. For detailed language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The validity period for sending template messages in WhatsApp.
	//
	// - AUTHENTICATION: valid values are 30 to 900.
	//
	// - UTILITY: valid values are 30 to 43200.
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The message template code.
	//
	// example:
	//
	// 8472929283883
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name.
	//
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The templatetype.
	//
	// - **WHATSAPP**
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ModifyChatappTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyChatappTemplateShrinkRequest) GetCategoryChangePaused() *bool {
	return s.CategoryChangePaused
}

func (s *ModifyChatappTemplateShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *ModifyChatappTemplateShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ModifyChatappTemplateShrinkRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *ModifyChatappTemplateShrinkRequest) GetExampleShrink() *string {
	return s.ExampleShrink
}

func (s *ModifyChatappTemplateShrinkRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *ModifyChatappTemplateShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *ModifyChatappTemplateShrinkRequest) GetMessageSendTtlSeconds() *int32 {
	return s.MessageSendTtlSeconds
}

func (s *ModifyChatappTemplateShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ModifyChatappTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifyChatappTemplateShrinkRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ModifyChatappTemplateShrinkRequest) SetCategory(v string) *ModifyChatappTemplateShrinkRequest {
	s.Category = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetCategoryChangePaused(v bool) *ModifyChatappTemplateShrinkRequest {
	s.CategoryChangePaused = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetComponentsShrink(v string) *ModifyChatappTemplateShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetCustSpaceId(v string) *ModifyChatappTemplateShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetCustWabaId(v string) *ModifyChatappTemplateShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetExampleShrink(v string) *ModifyChatappTemplateShrinkRequest {
	s.ExampleShrink = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetIsvCode(v string) *ModifyChatappTemplateShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetLanguage(v string) *ModifyChatappTemplateShrinkRequest {
	s.Language = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetMessageSendTtlSeconds(v int32) *ModifyChatappTemplateShrinkRequest {
	s.MessageSendTtlSeconds = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetTemplateCode(v string) *ModifyChatappTemplateShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetTemplateName(v string) *ModifyChatappTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetTemplateType(v string) *ModifyChatappTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
