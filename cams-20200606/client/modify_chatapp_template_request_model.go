// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChatappTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ModifyChatappTemplateRequest
	GetCategory() *string
	SetCategoryChangePaused(v bool) *ModifyChatappTemplateRequest
	GetCategoryChangePaused() *bool
	SetComponents(v []*ModifyChatappTemplateRequestComponents) *ModifyChatappTemplateRequest
	GetComponents() []*ModifyChatappTemplateRequestComponents
	SetCustSpaceId(v string) *ModifyChatappTemplateRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *ModifyChatappTemplateRequest
	GetCustWabaId() *string
	SetExample(v map[string]*string) *ModifyChatappTemplateRequest
	GetExample() map[string]*string
	SetIsvCode(v string) *ModifyChatappTemplateRequest
	GetIsvCode() *string
	SetLanguage(v string) *ModifyChatappTemplateRequest
	GetLanguage() *string
	SetMessageSendTtlSeconds(v int32) *ModifyChatappTemplateRequest
	GetMessageSendTtlSeconds() *int32
	SetTemplateCode(v string) *ModifyChatappTemplateRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *ModifyChatappTemplateRequest
	GetTemplateName() *string
	SetTemplateType(v string) *ModifyChatappTemplateRequest
	GetTemplateType() *string
}

type ModifyChatappTemplateRequest struct {
	// The templatetype is immutable.
	//
	// example:
	//
	// text
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CategoryChangePaused *bool   `json:"CategoryChangePaused,omitempty" xml:"CategoryChangePaused,omitempty"`
	// A list of message template components.
	//
	// > When Category is AUTHENTICATION, Components cannot contain a node with Type set to HEADER. If Type is BODY or FOOTER, the Text content is empty and is automatically generated.
	//
	// This parameter is required.
	Components []*ModifyChatappTemplateRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The Space ID of the ISV sub-customer, or the instance ID of a direct customer.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business Account (WABA) ID of the independent software vendor (ISV) customer.
	//
	// > This parameter is deprecated. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The template example.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The ISV verification code. This code is used to verify that the RAM user is authorized by the ISV.
	//
	// example:
	//
	// ksiekdki39ksks93939
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The template language. For a list of language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The time-to-live (TTL) of the template message in seconds.
	//
	// - For AUTHENTICATION templates, the value ranges from 30 to 900.
	//
	// - For UTILITY templates, the value ranges from 30 to 43,200.
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
	// The template type.
	//
	// - **WHATSAPP**
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ModifyChatappTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyChatappTemplateRequest) GetCategoryChangePaused() *bool {
	return s.CategoryChangePaused
}

func (s *ModifyChatappTemplateRequest) GetComponents() []*ModifyChatappTemplateRequestComponents {
	return s.Components
}

func (s *ModifyChatappTemplateRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ModifyChatappTemplateRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *ModifyChatappTemplateRequest) GetExample() map[string]*string {
	return s.Example
}

func (s *ModifyChatappTemplateRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *ModifyChatappTemplateRequest) GetLanguage() *string {
	return s.Language
}

func (s *ModifyChatappTemplateRequest) GetMessageSendTtlSeconds() *int32 {
	return s.MessageSendTtlSeconds
}

func (s *ModifyChatappTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ModifyChatappTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifyChatappTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ModifyChatappTemplateRequest) SetCategory(v string) *ModifyChatappTemplateRequest {
	s.Category = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetCategoryChangePaused(v bool) *ModifyChatappTemplateRequest {
	s.CategoryChangePaused = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetComponents(v []*ModifyChatappTemplateRequestComponents) *ModifyChatappTemplateRequest {
	s.Components = v
	return s
}

func (s *ModifyChatappTemplateRequest) SetCustSpaceId(v string) *ModifyChatappTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetCustWabaId(v string) *ModifyChatappTemplateRequest {
	s.CustWabaId = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetExample(v map[string]*string) *ModifyChatappTemplateRequest {
	s.Example = v
	return s
}

func (s *ModifyChatappTemplateRequest) SetIsvCode(v string) *ModifyChatappTemplateRequest {
	s.IsvCode = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetLanguage(v string) *ModifyChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetMessageSendTtlSeconds(v int32) *ModifyChatappTemplateRequest {
	s.MessageSendTtlSeconds = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetTemplateCode(v string) *ModifyChatappTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetTemplateName(v string) *ModifyChatappTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetTemplateType(v string) *ModifyChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *ModifyChatappTemplateRequest) Validate() error {
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

type ModifyChatappTemplateRequestComponents struct {
	// For WhatsApp templates, this is valid when Category is AUTHENTICATION and Component Type is Body. It indicates that a message is displayed above the body, advising users not to share the verification code.
	//
	// example:
	//
	// false
	AddSecretRecommendation *bool `json:"AddSecretRecommendation,omitempty" xml:"AddSecretRecommendation,omitempty"`
	// A list of buttons. This applies only to the `BUTTONS` component.
	//
	// > Number of buttons for WhatsApp:
	//
	// >
	//
	// > - For MARKETING or UTILITY templates, WhatsApp allows a maximum of 10 buttons.
	//
	// >
	//
	// > - Only one button of the PHONE_NUMBER type is allowed.
	//
	// >
	//
	// > - A maximum of two buttons of the URL type are allowed.
	//
	// >
	//
	// > - QUICK_REPLY buttons cannot be mixed with PHONE_NUMBER or URL buttons.
	Buttons []*ModifyChatappTemplateRequestComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description.
	//
	// > You can add a description when Type is HEADER and Format is IMAGE, `DOCUMENT`, or `VIDEO`.
	//
	// example:
	//
	// This is a video
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// A list of cards for the Carousel template.
	Cards []*ModifyChatappTemplateRequestComponentsCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// The validity period of the verification code for a WhatsApp AUTHENTICATION template, in minutes. This is valid only for WhatsApp messages when Category is AUTHENTICATION and Component Type is Footer. This information is displayed in the footer.
	//
	// example:
	//
	// 5
	CodeExpirationMinutes *int32 `json:"CodeExpirationMinutes,omitempty" xml:"CodeExpirationMinutes,omitempty"`
	// This field is invalid.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The file name.
	//
	// > Specify the file name when Type is HEADER and `Format` is `DOCUMENT`.
	//
	// example:
	//
	// video
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This field is invalid.
	//
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The media resource type.
	//
	// - **TEXT**: Text
	//
	// - **IMAGE**: Image
	//
	// - **DOCUMENT**: Document
	//
	// - **VIDEO**: Video
	//
	// example:
	//
	// TEXT
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Indicates whether the coupon code has an expiration time. Used when type is LIMITED_TIME_OFFER.
	//
	// example:
	//
	// true
	HasExpiration *bool `json:"HasExpiration,omitempty" xml:"HasExpiration,omitempty"`
	// The text of the message.
	//
	// > If Category is AUTHENTICATION, this property is empty.
	//
	// example:
	//
	// hello chatapp
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This field is invalid.
	//
	// example:
	//
	// https://cdn.multiplymall.mobiapp.cloud/cloudcode/yc-165407506207478-165511576113195/20220905/ec5b9737-1507-4208-bb27-8da3958da961.jpg?x-oss-process=image/resize,w_100
	ThumbUrl *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
	// The component type:
	//
	// - **BODY**
	//
	// - **HEADER**
	//
	// - **FOOTER**
	//
	// - **BUTTONS**
	//
	// - **CAROUSEL**
	//
	// - **LIMITED_TIME_OFFER**
	//
	// > 	- For WhatsApp templates, the `BODY` component cannot exceed 1,024 characters. The `HEADER` and `FOOTER` components cannot exceed 60 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the media asset.
	//
	// example:
	//
	// https://img.****.com/png_preview/00/10/24/1GygxVK3F4.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ModifyChatappTemplateRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponents) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponents) GetAddSecretRecommendation() *bool {
	return s.AddSecretRecommendation
}

func (s *ModifyChatappTemplateRequestComponents) GetButtons() []*ModifyChatappTemplateRequestComponentsButtons {
	return s.Buttons
}

func (s *ModifyChatappTemplateRequestComponents) GetCaption() *string {
	return s.Caption
}

func (s *ModifyChatappTemplateRequestComponents) GetCards() []*ModifyChatappTemplateRequestComponentsCards {
	return s.Cards
}

func (s *ModifyChatappTemplateRequestComponents) GetCodeExpirationMinutes() *int32 {
	return s.CodeExpirationMinutes
}

func (s *ModifyChatappTemplateRequestComponents) GetDuration() *int32 {
	return s.Duration
}

func (s *ModifyChatappTemplateRequestComponents) GetFileName() *string {
	return s.FileName
}

func (s *ModifyChatappTemplateRequestComponents) GetFileType() *string {
	return s.FileType
}

func (s *ModifyChatappTemplateRequestComponents) GetFormat() *string {
	return s.Format
}

func (s *ModifyChatappTemplateRequestComponents) GetHasExpiration() *bool {
	return s.HasExpiration
}

func (s *ModifyChatappTemplateRequestComponents) GetText() *string {
	return s.Text
}

func (s *ModifyChatappTemplateRequestComponents) GetThumbUrl() *string {
	return s.ThumbUrl
}

func (s *ModifyChatappTemplateRequestComponents) GetType() *string {
	return s.Type
}

func (s *ModifyChatappTemplateRequestComponents) GetUrl() *string {
	return s.Url
}

func (s *ModifyChatappTemplateRequestComponents) SetAddSecretRecommendation(v bool) *ModifyChatappTemplateRequestComponents {
	s.AddSecretRecommendation = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetButtons(v []*ModifyChatappTemplateRequestComponentsButtons) *ModifyChatappTemplateRequestComponents {
	s.Buttons = v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetCaption(v string) *ModifyChatappTemplateRequestComponents {
	s.Caption = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetCards(v []*ModifyChatappTemplateRequestComponentsCards) *ModifyChatappTemplateRequestComponents {
	s.Cards = v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetCodeExpirationMinutes(v int32) *ModifyChatappTemplateRequestComponents {
	s.CodeExpirationMinutes = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetDuration(v int32) *ModifyChatappTemplateRequestComponents {
	s.Duration = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetFileName(v string) *ModifyChatappTemplateRequestComponents {
	s.FileName = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetFileType(v string) *ModifyChatappTemplateRequestComponents {
	s.FileType = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetFormat(v string) *ModifyChatappTemplateRequestComponents {
	s.Format = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetHasExpiration(v bool) *ModifyChatappTemplateRequestComponents {
	s.HasExpiration = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetText(v string) *ModifyChatappTemplateRequestComponents {
	s.Text = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetThumbUrl(v string) *ModifyChatappTemplateRequestComponents {
	s.ThumbUrl = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetType(v string) *ModifyChatappTemplateRequestComponents {
	s.Type = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetUrl(v string) *ModifyChatappTemplateRequestComponents {
	s.Url = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) Validate() error {
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

type ModifyChatappTemplateRequestComponentsButtons struct {
	// Required for WhatsApp templates when Category is AUTHENTICATION and Button Type is ONE_TAP or ZERO_TAP. This is the button text for the WhatsApp Autofill action.
	//
	// example:
	//
	// Autofill
	AutofillText *string `json:"AutofillText,omitempty" xml:"AutofillText,omitempty"`
	// The coupon code. It supports only letters and numbers. You can pass a variable, such as \\`$(couponCode)\\`, and provide the actual coupon code at the time of sending.
	//
	// example:
	//
	// 120293
	CouponCode *string `json:"CouponCode,omitempty" xml:"CouponCode,omitempty"`
	// The Flow data event type. Valid values:
	//
	// - DATA_EXCHANGE: Data exchange.
	//
	// - NAVIGATE: Navigation.
	//
	// example:
	//
	// NAVIGATE
	FlowAction *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// The Flow ID.
	//
	// example:
	//
	// 664597077870605
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// For WhatsApp templates, this is valid when Category is Marketing and Button type is QUICK_REPLY. It indicates that the button is a marketing opt-out button. If a customer clicks this button and a sending control action is configured in ChatApp, subsequent marketing messages are not sent to the customer.
	//
	// example:
	//
	// false
	IsOptOut *bool `json:"IsOptOut,omitempty" xml:"IsOptOut,omitempty"`
	// The screen to navigate to. This is required when FlowAction is NAVIGATE.
	//
	// example:
	//
	// DETAILS
	NavigateScreen *string `json:"NavigateScreen,omitempty" xml:"NavigateScreen,omitempty"`
	// Deprecated
	//
	// Use the properties under SupportedApps.
	//
	// example:
	//
	// com.demo
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The phone number.
	//
	// example:
	//
	// +861388888****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Deprecated
	//
	// Use the properties under SupportedApps.
	//
	// example:
	//
	// 29dkeke
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
	// A list of supported apps.
	SupportedApps []*ModifyChatappTemplateRequestComponentsButtonsSupportedApps `json:"SupportedApps,omitempty" xml:"SupportedApps,omitempty" type:"Repeated"`
	// The button text.
	//
	// example:
	//
	// phone-button-text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The button type.
	//
	// - **PHONE_NUMBER**: Call button
	//
	// - **URL**: URL button
	//
	// - **QUICK_REPLY**: Quick reply button
	//
	// - **COPY_CODE**: Copy verification code or coupon code
	//
	// - **ONE_TAP**: Backfill button for AUTHENTICATION templates
	//
	// - **ZERO_TAP**: Backfill button for AUTHENTICATION templates
	//
	// - **MPM**: Multi-product message
	//
	// - **CATALOG**: Catalog
	//
	// - **FLOW**: Open WhatsApp flow
	//
	// > 	- For WhatsApp templates where Category is AUTHENTICATION, only one button is allowed, and its type must be COPY_CODE or ONE_TAP. If the type is COPY_CODE, Text is required. If the type is ONE_TAP, the Text (the name of the copy code button, displayed when the target application is not installed), SignatureHash, PackageName, and AutofillText properties are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to visit when the button is clicked.
	//
	// example:
	//
	// https://www.website.com/***
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL type.
	//
	// - **static**
	//
	// - **dynamic**
	//
	// example:
	//
	// dynamic
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s ModifyChatappTemplateRequestComponentsButtons) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsButtons) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetAutofillText() *string {
	return s.AutofillText
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetCouponCode() *string {
	return s.CouponCode
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetFlowAction() *string {
	return s.FlowAction
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetFlowId() *string {
	return s.FlowId
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetIsOptOut() *bool {
	return s.IsOptOut
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetNavigateScreen() *string {
	return s.NavigateScreen
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetPackageName() *string {
	return s.PackageName
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetSignatureHash() *string {
	return s.SignatureHash
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetSupportedApps() []*ModifyChatappTemplateRequestComponentsButtonsSupportedApps {
	return s.SupportedApps
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetText() *string {
	return s.Text
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetType() *string {
	return s.Type
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetUrl() *string {
	return s.Url
}

func (s *ModifyChatappTemplateRequestComponentsButtons) GetUrlType() *string {
	return s.UrlType
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetAutofillText(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.AutofillText = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetCouponCode(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.CouponCode = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetFlowAction(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.FlowAction = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetFlowId(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.FlowId = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetIsOptOut(v bool) *ModifyChatappTemplateRequestComponentsButtons {
	s.IsOptOut = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetNavigateScreen(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.NavigateScreen = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetPackageName(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.PackageName = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetPhoneNumber(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetSignatureHash(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.SignatureHash = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetSupportedApps(v []*ModifyChatappTemplateRequestComponentsButtonsSupportedApps) *ModifyChatappTemplateRequestComponentsButtons {
	s.SupportedApps = v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetText(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.Text = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetType(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.Type = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetUrl(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.Url = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetUrlType(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.UrlType = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) Validate() error {
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

type ModifyChatappTemplateRequestComponentsButtonsSupportedApps struct {
	// Required for WhatsApp templates when Category is AUTHENTICATION and Button Type is ONE_TAP or ZERO_TAP. This is the package name of the application launched by WhatsApp.
	//
	// example:
	//
	// com.example.myapplication
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// Required for WhatsApp templates when Category is AUTHENTICATION and Button Type is ONE_TAP or ZERO_TAP. This is the signature hash of the application launched by WhatsApp.
	//
	// example:
	//
	// fk39kd93ks9
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
}

func (s ModifyChatappTemplateRequestComponentsButtonsSupportedApps) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsButtonsSupportedApps) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsButtonsSupportedApps) GetPackageName() *string {
	return s.PackageName
}

func (s *ModifyChatappTemplateRequestComponentsButtonsSupportedApps) GetSignatureHash() *string {
	return s.SignatureHash
}

func (s *ModifyChatappTemplateRequestComponentsButtonsSupportedApps) SetPackageName(v string) *ModifyChatappTemplateRequestComponentsButtonsSupportedApps {
	s.PackageName = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtonsSupportedApps) SetSignatureHash(v string) *ModifyChatappTemplateRequestComponentsButtonsSupportedApps {
	s.SignatureHash = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtonsSupportedApps) Validate() error {
	return dara.Validate(s)
}

type ModifyChatappTemplateRequestComponentsCards struct {
	// A list of controls in the Carousel card.
	//
	// This parameter is required.
	CardComponents []*ModifyChatappTemplateRequestComponentsCardsCardComponents `json:"CardComponents,omitempty" xml:"CardComponents,omitempty" type:"Repeated"`
}

func (s ModifyChatappTemplateRequestComponentsCards) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsCards) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsCards) GetCardComponents() []*ModifyChatappTemplateRequestComponentsCardsCardComponents {
	return s.CardComponents
}

func (s *ModifyChatappTemplateRequestComponentsCards) SetCardComponents(v []*ModifyChatappTemplateRequestComponentsCardsCardComponents) *ModifyChatappTemplateRequestComponentsCards {
	s.CardComponents = v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCards) Validate() error {
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

type ModifyChatappTemplateRequestComponentsCardsCardComponents struct {
	// A list of buttons. This applies only to the BUTTONS component. Each Carousel card can have a maximum of two buttons.
	Buttons []*ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The media resource type. This is valid when Type is HEADER.
	//
	// - **IMAGE**
	//
	// - **VIDEO**
	//
	// example:
	//
	// IMAGE
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The content of the BODY in the Carousel card.
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
	// The URL of the media asset.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ModifyChatappTemplateRequestComponentsCardsCardComponents) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsCardsCardComponents) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) GetButtons() []*ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	return s.Buttons
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) GetFormat() *string {
	return s.Format
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) GetText() *string {
	return s.Text
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) GetType() *string {
	return s.Type
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) GetUrl() *string {
	return s.Url
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetButtons(v []*ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Buttons = v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetFormat(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Format = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetText(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Text = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetType(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Type = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetUrl(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Url = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) Validate() error {
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

type ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons struct {
	// The phone number.
	//
	// example:
	//
	// +861368893****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The button text.
	//
	// example:
	//
	// Call me
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The button type.
	//
	// - **PHONE_NUMBER**: Call button
	//
	// - **URL**: URL button
	//
	// - **QUICK_REPLY**: Quick reply button
	//
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to visit when the button is clicked.
	//
	// example:
	//
	// https://alibaba.com/xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL type.
	//
	// - **static**
	//
	// - **dynamic**
	//
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) GetText() *string {
	return s.Text
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) GetType() *string {
	return s.Type
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) GetUrl() *string {
	return s.Url
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) GetUrlType() *string {
	return s.UrlType
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetPhoneNumber(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetText(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Text = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetType(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Type = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetUrl(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Url = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetUrlType(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.UrlType = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) Validate() error {
	return dara.Validate(s)
}
