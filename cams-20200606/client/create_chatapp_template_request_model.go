// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatappTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowCategoryChange(v bool) *CreateChatappTemplateRequest
	GetAllowCategoryChange() *bool
	SetCategory(v string) *CreateChatappTemplateRequest
	GetCategory() *string
	SetCategoryChangePaused(v bool) *CreateChatappTemplateRequest
	GetCategoryChangePaused() *bool
	SetComponents(v []*CreateChatappTemplateRequestComponents) *CreateChatappTemplateRequest
	GetComponents() []*CreateChatappTemplateRequestComponents
	SetCustSpaceId(v string) *CreateChatappTemplateRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *CreateChatappTemplateRequest
	GetCustWabaId() *string
	SetExample(v map[string]*string) *CreateChatappTemplateRequest
	GetExample() map[string]*string
	SetIsvCode(v string) *CreateChatappTemplateRequest
	GetIsvCode() *string
	SetLanguage(v string) *CreateChatappTemplateRequest
	GetLanguage() *string
	SetMessageSendTtlSeconds(v int32) *CreateChatappTemplateRequest
	GetMessageSendTtlSeconds() *int32
	SetName(v string) *CreateChatappTemplateRequest
	GetName() *string
	SetProductSetId(v string) *CreateChatappTemplateRequest
	GetProductSetId() *string
	SetTemplateType(v string) *CreateChatappTemplateRequest
	GetTemplateType() *string
}

type CreateChatappTemplateRequest struct {
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
	Components []*CreateChatappTemplateRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
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
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
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

func (s CreateChatappTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequest) GetAllowCategoryChange() *bool {
	return s.AllowCategoryChange
}

func (s *CreateChatappTemplateRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateChatappTemplateRequest) GetCategoryChangePaused() *bool {
	return s.CategoryChangePaused
}

func (s *CreateChatappTemplateRequest) GetComponents() []*CreateChatappTemplateRequestComponents {
	return s.Components
}

func (s *CreateChatappTemplateRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreateChatappTemplateRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *CreateChatappTemplateRequest) GetExample() map[string]*string {
	return s.Example
}

func (s *CreateChatappTemplateRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *CreateChatappTemplateRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateChatappTemplateRequest) GetMessageSendTtlSeconds() *int32 {
	return s.MessageSendTtlSeconds
}

func (s *CreateChatappTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateChatappTemplateRequest) GetProductSetId() *string {
	return s.ProductSetId
}

func (s *CreateChatappTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateChatappTemplateRequest) SetAllowCategoryChange(v bool) *CreateChatappTemplateRequest {
	s.AllowCategoryChange = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetCategory(v string) *CreateChatappTemplateRequest {
	s.Category = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetCategoryChangePaused(v bool) *CreateChatappTemplateRequest {
	s.CategoryChangePaused = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetComponents(v []*CreateChatappTemplateRequestComponents) *CreateChatappTemplateRequest {
	s.Components = v
	return s
}

func (s *CreateChatappTemplateRequest) SetCustSpaceId(v string) *CreateChatappTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetCustWabaId(v string) *CreateChatappTemplateRequest {
	s.CustWabaId = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetExample(v map[string]*string) *CreateChatappTemplateRequest {
	s.Example = v
	return s
}

func (s *CreateChatappTemplateRequest) SetIsvCode(v string) *CreateChatappTemplateRequest {
	s.IsvCode = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetLanguage(v string) *CreateChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetMessageSendTtlSeconds(v int32) *CreateChatappTemplateRequest {
	s.MessageSendTtlSeconds = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetName(v string) *CreateChatappTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetProductSetId(v string) *CreateChatappTemplateRequest {
	s.ProductSetId = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetTemplateType(v string) *CreateChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateChatappTemplateRequest) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateChatappTemplateRequestComponents struct {
	// Valid for WhatsApp templates when Category is AUTHENTICATION and Component Type is Body. Displays a prompt above the Body advising not to share the verification code with others.
	//
	// example:
	//
	// true
	AddSecretRecommendation *bool `json:"AddSecretRecommendation,omitempty" xml:"AddSecretRecommendation,omitempty"`
	// The button list. Applicable only to **BUTTONS*	- components.
	//
	// > WhatsApp button quantity rules:
	//
	// > - For WhatsApp templates with Category MARKETING/UTILITY, a maximum of 10 buttons are allowed.
	//
	// > - Only 1 PHONE_NUMBER button is allowed.
	//
	// > - A maximum of 2 URL buttons are allowed.
	//
	// > - QUICK_REPLY buttons cannot appear in mixed order with PHONE_NUMBER/URL buttons.
	//
	// > Viber button quantity rules:
	//
	// > - Only URL type is supported, and only one button is allowed.
	//
	// > - When the HEADER contains a VIDEO, the button type is URL, but you cannot set a URL address.
	Buttons []*CreateChatappTemplateRequestComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description of the file.
	//
	// example:
	//
	// This is a video
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The Carousel template card list.
	Cards []*CreateChatappTemplateRequestComponentsCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// The validity period (in minutes) of the verification code for WhatsApp AUTHENTICATION templates. Valid only for WhatsApp messages when Category is AUTHENTICATION and Component Type is Footer (displayed in the Footer position).
	//
	// example:
	//
	// 5
	CodeExpirationMinutes *int32 `json:"CodeExpirationMinutes,omitempty" xml:"CodeExpirationMinutes,omitempty"`
	// The duration (in seconds) of Viber video messages. Valid values: 0 to 600.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// Express delivery video
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type for Viber file messages.
	//
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The media resource type.
	//
	// - **TEXT**: text
	//
	//
	//
	// - **IMAGE**: image
	//
	// - **DOCUMENT**: document
	//
	// - **VIDEO**: video
	//
	// example:
	//
	// TEXT
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Specifies whether the coupon code has an expiration time. This parameter is used when type = LIMITED_TIME_OFFER.
	//
	// example:
	//
	// true
	HasExpiration *bool `json:"HasExpiration,omitempty" xml:"HasExpiration,omitempty"`
	// The text of the message to be sent.
	//
	// > For WHATSAPP type, this property value is empty when Category=AUTHENTICATION.
	//
	// example:
	//
	// hello whatsapp
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The thumbnail for Viber video messages.
	//
	// example:
	//
	// https://cdn.multiplymall.mobiapp.cloud/yunmall/B-LM-LMALL202207130001/20220730/d712a057-a6af-4513-bbe6-7ee57ea60983.png?x-oss-process=image/resize,w_100
	ThumbUrl *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
	// The component type. Valid values:
	//
	// - **BODY**
	//
	// - **HEADER**
	//
	// - **FOOTER**
	//
	//  - **BUTTONS**
	//
	// - **CAROUSEL**
	//
	// - **LIMITED_TIME_OFFER**
	//
	// > - For WhatsApp templates, the **BODY*	- component cannot exceed 1024 characters. The **HEADER*	- and **FOOTER*	- components cannot exceed 60 characters.
	//
	// > - For Viber templates, the **FOOTER**, **CAROUSEL**, and **LIMITED_TIME_OFFER*	- types are invalid.
	//
	// > - For Viber templates, images, videos, and files are placed in the **HEADER*	- (the device displays images below the text). Text is placed in the **BODY**.
	//
	// This parameter is required.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The media resource path.
	//
	// > For Viber type, the recommended image size is 800 px × 800 px.
	//
	// example:
	//
	// https://image.developer.aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateChatappTemplateRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponents) GetAddSecretRecommendation() *bool {
	return s.AddSecretRecommendation
}

func (s *CreateChatappTemplateRequestComponents) GetButtons() []*CreateChatappTemplateRequestComponentsButtons {
	return s.Buttons
}

func (s *CreateChatappTemplateRequestComponents) GetCaption() *string {
	return s.Caption
}

func (s *CreateChatappTemplateRequestComponents) GetCards() []*CreateChatappTemplateRequestComponentsCards {
	return s.Cards
}

func (s *CreateChatappTemplateRequestComponents) GetCodeExpirationMinutes() *int32 {
	return s.CodeExpirationMinutes
}

func (s *CreateChatappTemplateRequestComponents) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateChatappTemplateRequestComponents) GetFileName() *string {
	return s.FileName
}

func (s *CreateChatappTemplateRequestComponents) GetFileType() *string {
	return s.FileType
}

func (s *CreateChatappTemplateRequestComponents) GetFormat() *string {
	return s.Format
}

func (s *CreateChatappTemplateRequestComponents) GetHasExpiration() *bool {
	return s.HasExpiration
}

func (s *CreateChatappTemplateRequestComponents) GetText() *string {
	return s.Text
}

func (s *CreateChatappTemplateRequestComponents) GetThumbUrl() *string {
	return s.ThumbUrl
}

func (s *CreateChatappTemplateRequestComponents) GetType() *string {
	return s.Type
}

func (s *CreateChatappTemplateRequestComponents) GetUrl() *string {
	return s.Url
}

func (s *CreateChatappTemplateRequestComponents) SetAddSecretRecommendation(v bool) *CreateChatappTemplateRequestComponents {
	s.AddSecretRecommendation = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetButtons(v []*CreateChatappTemplateRequestComponentsButtons) *CreateChatappTemplateRequestComponents {
	s.Buttons = v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetCaption(v string) *CreateChatappTemplateRequestComponents {
	s.Caption = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetCards(v []*CreateChatappTemplateRequestComponentsCards) *CreateChatappTemplateRequestComponents {
	s.Cards = v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetCodeExpirationMinutes(v int32) *CreateChatappTemplateRequestComponents {
	s.CodeExpirationMinutes = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetDuration(v int32) *CreateChatappTemplateRequestComponents {
	s.Duration = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFileName(v string) *CreateChatappTemplateRequestComponents {
	s.FileName = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFileType(v string) *CreateChatappTemplateRequestComponents {
	s.FileType = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFormat(v string) *CreateChatappTemplateRequestComponents {
	s.Format = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetHasExpiration(v bool) *CreateChatappTemplateRequestComponents {
	s.HasExpiration = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetText(v string) *CreateChatappTemplateRequestComponents {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetThumbUrl(v string) *CreateChatappTemplateRequestComponents {
	s.ThumbUrl = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetType(v string) *CreateChatappTemplateRequestComponents {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetUrl(v string) *CreateChatappTemplateRequestComponents {
	s.Url = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) Validate() error {
	if s.Buttons != nil {
		for _, item := range s.Buttons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Cards != nil {
		for _, item := range s.Cards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateChatappTemplateRequestComponentsButtons struct {
	// Required for WhatsApp templates when Category is AUTHENTICATION and Button Type is ONE_TAP/ZERO_TAP. The button text for the WhatsApp Autofill operation.
	//
	// example:
	//
	// Autofill
	AutofillText *string `json:"AutofillText,omitempty" xml:"AutofillText,omitempty"`
	// The coupon code value. Only letters and numbers are supported. You can pass in a variable such as $(couponCode) and provide the actual coupon code when sending.
	//
	// example:
	//
	// 120293
	CouponCode *string `json:"CouponCode,omitempty" xml:"CouponCode,omitempty"`
	// The Flow data event type. Valid values:
	//
	// - DATA_EXCHANGE: data exchange.
	//
	// - NAVIGATE: navigation.
	//
	// example:
	//
	// NAVIGATE
	FlowAction *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// Flow ID。
	//
	// example:
	//
	// 479884093605****
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// Valid for WhatsApp templates when Category is Marketing and Button type is QUICK_REPLY. Indicates the button is a marketing opt-out button. If the customer clicks this button and send control is configured on ChatApp, subsequent Marketing messages will not be sent to the customer.
	//
	// example:
	//
	// false
	IsOptOut *bool `json:"IsOptOut,omitempty" xml:"IsOptOut,omitempty"`
	// The navigate screen. Required when FlowAction=NAVIGATE.
	//
	// example:
	//
	// DETAILS
	NavigateScreen *string `json:"NavigateScreen,omitempty" xml:"NavigateScreen,omitempty"`
	// Use the properties under SupportedApps instead.
	//
	// example:
	//
	// com.demo
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The phone number. Valid only when the button type is **PHONE_NUMBER**.
	//
	// example:
	//
	// +861368897****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Use the properties under SupportedApps instead.
	//
	// example:
	//
	// wi299382
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
	// The list of supported applications.
	SupportedApps []*CreateChatappTemplateRequestComponentsButtonsSupportedApps `json:"SupportedApps,omitempty" xml:"SupportedApps,omitempty" type:"Repeated"`
	// The display name of the button.
	//
	// example:
	//
	// Call Me
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The button type.
	//
	// - **PHONE_NUMBER**: dial phone button
	//
	// - **URL**: web button
	//
	// - **QUICK_REPLY**: quick reply button
	//
	// - **COPY_CODE**: copy verification code or coupon code
	//
	// - **ONE_TAP**: autofill button for AUTHENTICATION templates
	//
	// - **ZERO_TAP**: autofill button for AUTHENTICATION templates
	//
	// - **MPM**: multi-product catalog
	//
	// - **CATALOG**: catalog
	//
	// - **FLOW**: open WhatsApp flow
	//
	// > - For WhatsApp templates with Category AUTHENTICATION, only one button is allowed, and the type can only be COPY_CODE/ONE_TAP. When COPY_CODE is selected, Text is required. When ONE_TAP is selected, Text (displayed when the target application is not installed on the device, representing the copy verification code button name), SignatureHash, PackageName, and AutofillText are required.
	//
	// > - Viber templates allow only one Button, and it must be URL type.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to visit when the link button is clicked.
	//
	// example:
	//
	// https://example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL type.
	//
	// - **static**: static
	//
	// - **dynamic**: dynamic
	//
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsButtons) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsButtons) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetAutofillText() *string {
	return s.AutofillText
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetCouponCode() *string {
	return s.CouponCode
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetFlowAction() *string {
	return s.FlowAction
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetFlowId() *string {
	return s.FlowId
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetIsOptOut() *bool {
	return s.IsOptOut
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetNavigateScreen() *string {
	return s.NavigateScreen
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetPackageName() *string {
	return s.PackageName
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetSignatureHash() *string {
	return s.SignatureHash
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetSupportedApps() []*CreateChatappTemplateRequestComponentsButtonsSupportedApps {
	return s.SupportedApps
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetText() *string {
	return s.Text
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetType() *string {
	return s.Type
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetUrl() *string {
	return s.Url
}

func (s *CreateChatappTemplateRequestComponentsButtons) GetUrlType() *string {
	return s.UrlType
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetAutofillText(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.AutofillText = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetCouponCode(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.CouponCode = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetFlowAction(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.FlowAction = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetFlowId(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.FlowId = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetIsOptOut(v bool) *CreateChatappTemplateRequestComponentsButtons {
	s.IsOptOut = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetNavigateScreen(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.NavigateScreen = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetPackageName(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.PackageName = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetPhoneNumber(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetSignatureHash(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.SignatureHash = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetSupportedApps(v []*CreateChatappTemplateRequestComponentsButtonsSupportedApps) *CreateChatappTemplateRequestComponentsButtons {
	s.SupportedApps = v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetText(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetType(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetUrl(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.Url = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetUrlType(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.UrlType = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) Validate() error {
	if s.SupportedApps != nil {
		for _, item := range s.SupportedApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateChatappTemplateRequestComponentsButtonsSupportedApps struct {
	// Required for WhatsApp templates when Category is AUTHENTICATION and Button Type is ONE_TAP/ZERO_TAP. The package name of the application invoked by WhatsApp.
	//
	// example:
	//
	// com.kuaidian.waimaistaff
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// Required for WhatsApp templates when Category is AUTHENTICATION and Button Type is ONE_TAP/ZERO_TAP. The signature hash value for the application invoked by WhatsApp.
	//
	// example:
	//
	// ieid83kdiek
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsButtonsSupportedApps) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsButtonsSupportedApps) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsButtonsSupportedApps) GetPackageName() *string {
	return s.PackageName
}

func (s *CreateChatappTemplateRequestComponentsButtonsSupportedApps) GetSignatureHash() *string {
	return s.SignatureHash
}

func (s *CreateChatappTemplateRequestComponentsButtonsSupportedApps) SetPackageName(v string) *CreateChatappTemplateRequestComponentsButtonsSupportedApps {
	s.PackageName = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtonsSupportedApps) SetSignatureHash(v string) *CreateChatappTemplateRequestComponentsButtonsSupportedApps {
	s.SignatureHash = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtonsSupportedApps) Validate() error {
	return dara.Validate(s)
}

type CreateChatappTemplateRequestComponentsCards struct {
	// The list of components in the Carousel card.
	CardComponents []*CreateChatappTemplateRequestComponentsCardsCardComponents `json:"CardComponents,omitempty" xml:"CardComponents,omitempty" type:"Repeated"`
}

func (s CreateChatappTemplateRequestComponentsCards) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsCards) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsCards) GetCardComponents() []*CreateChatappTemplateRequestComponentsCardsCardComponents {
	return s.CardComponents
}

func (s *CreateChatappTemplateRequestComponentsCards) SetCardComponents(v []*CreateChatappTemplateRequestComponentsCardsCardComponents) *CreateChatappTemplateRequestComponentsCards {
	s.CardComponents = v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCards) Validate() error {
	if s.CardComponents != nil {
		for _, item := range s.CardComponents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateChatappTemplateRequestComponentsCardsCardComponents struct {
	// The button list. Applicable only to BUTTONS components. Each Carousel card can have a maximum of two buttons.
	Buttons []*CreateChatappTemplateRequestComponentsCardsCardComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The media resource type. Valid when Type = HEADER.
	//
	// - **IMAGE**: image
	//
	// - **VIDEO**: video
	//
	// example:
	//
	// IMAGE
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The BODY content in the Carousel card.
	//
	// example:
	//
	// Who is the very powerful team
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The component type. Valid values:
	//
	// - **BODY**
	//
	// - **HEADER**
	//
	// - **BUTTONS**
	//
	// This parameter is required.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The material path.
	//
	// example:
	//
	// https://alibaba.com/img.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsCardsCardComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsCardsCardComponents) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) GetButtons() []*CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	return s.Buttons
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) GetFormat() *string {
	return s.Format
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) GetText() *string {
	return s.Text
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) GetType() *string {
	return s.Type
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) GetUrl() *string {
	return s.Url
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetButtons(v []*CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Buttons = v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetFormat(v string) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Format = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetText(v string) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetType(v string) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetUrl(v string) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Url = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) Validate() error {
	if s.Buttons != nil {
		for _, item := range s.Buttons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateChatappTemplateRequestComponentsCardsCardComponentsButtons struct {
	// The phone number.
	//
	// example:
	//
	// +86138007****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The button text.
	//
	// example:
	//
	// Call me
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The button type.
	//
	// - **PHONE_NUMBER**: dial phone button
	//
	// - **URL**: web button
	//
	// - **QUICK_REPLY**: quick reply button
	//
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL that is accessed when the button is clicked.
	//
	// example:
	//
	// https://alibaba.com/xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL type.
	//
	// - **static**: Static.
	//
	// - **dynamic**: Dynamic.
	//
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) GetText() *string {
	return s.Text
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) GetType() *string {
	return s.Type
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) GetUrl() *string {
	return s.Url
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) GetUrlType() *string {
	return s.UrlType
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetPhoneNumber(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetText(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetType(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetUrl(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Url = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetUrlType(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.UrlType = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) Validate() error {
	return dara.Validate(s)
}
