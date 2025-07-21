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
	SetTemplateType(v string) *CreateChatappTemplateShrinkRequest
	GetTemplateType() *string
}

type CreateChatappTemplateShrinkRequest struct {
	// Deprecated
	//
	// Specifies whether to allow Facebook to automatically change the directory of the template. If you set this parameter to true, the review success rate of the template is improved. This parameter is valid only when TemplateType is set to WHATSAPP.
	//
	// example:
	//
	// true
	AllowCategoryChange *bool `json:"AllowCategoryChange,omitempty" xml:"AllowCategoryChange,omitempty"`
	// The category of the template if TemplateType is set to WHATSAPP. Valid values:
	//
	// 	- **UTILITY**: the transaction template
	//
	// 	- **MARKETING**: the marketing template
	//
	// 	- **AUTHENTICATION**: the authentication template
	//
	// The category of the template if TemplateType is set to VIBER. Valid values:
	//
	// 	- **text**: the template that contains only text
	//
	// 	- **image**: the template that contains only images
	//
	// 	- **text_image_button**: the template that contains text, images, and buttons
	//
	// 	- **text_button**: the template that contains text and buttons
	//
	// 	- **document**: the template that contains only documents
	//
	// 	- **video**: the template that contains only videos
	//
	// 	- **text_video**: the template that contains text and videos
	//
	// 	- **text_video_button**: the template that contains text, videos, and buttons
	//
	// 	- **text_image**: the template that contains text and images
	//
	// This parameter is required.
	//
	// example:
	//
	// The code of the message template.
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CategoryChangePaused *bool   `json:"CategoryChangePaused,omitempty" xml:"CategoryChangePaused,omitempty"`
	// The components of the message template.
	//
	// >  If Category is set to AUTHENTICATION, the Type sub-parameter of the Components parameter cannot be set to HEADER. If the Type sub-parameter is set to BODY or FOOTER, the Text sub-parameter of the Components parameter must be empty.
	//
	// This parameter is required.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// > CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	ExampleShrink *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the user is authorized by the ISV account.
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
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Validity period of authentication template message sending in WhatsApp
	//
	// > This attribute requires providing waba in advance to Alibaba operators to open the whitelist, otherwise it will result in template submission failure
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The name of the message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// hello_whatsapp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the message template.
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE: the Line message template. This type of message template will be released later.
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

func (s *CreateChatappTemplateShrinkRequest) SetTemplateType(v string) *CreateChatappTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
