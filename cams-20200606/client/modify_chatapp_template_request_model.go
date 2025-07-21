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
	// The category of the Viber message template. Valid values:
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
	// > This parameter applies only to Viber message templates.
	//
	// example:
	//
	// text
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CategoryChangePaused *bool   `json:"CategoryChangePaused,omitempty" xml:"CategoryChangePaused,omitempty"`
	// The components of the message template.
	//
	// >  If Category is set to AUTHENTICATION, the Type sub-parameter of the Components parameter cannot be set to HEADER. If the Type sub-parameter is set to BODY or FOOTER, you do not need to set the Text sub-parameter of the Components parameter because the value is automatically generated.
	//
	// This parameter is required.
	Components []*ModifyChatappTemplateRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
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
	// > CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 659216218162179
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// ksiekdki39ksks93939
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
	// >This attribute requires providing waba in advance to Alibaba operators to open the whitelist, otherwise it will result in template submission failure
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
	// Template name.
	//
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message template.
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE: the Line message template. This type of message template will be released later.
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
	return dara.Validate(s)
}

type ModifyChatappTemplateRequestComponents struct {
	// The note indicating that customers cannot share verification codes with others. The note is displayed in the message body. This parameter is valid if Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to BODY for a WhatsApp message template.
	//
	// example:
	//
	// false
	AddSecretRecommendation *bool `json:"AddSecretRecommendation,omitempty" xml:"AddSecretRecommendation,omitempty"`
	// The buttons. Specify this parameter only if you set the Type sub-parameter of the Components parameter to **BUTTONS**.
	//
	// >  ####
	//
	// 	- A marketing or utility WhatsApp message template can contain up to 10 buttons.
	//
	// 	- A WhatsApp message template can contain only one phone call button.
	//
	// 	- A WhatsApp message template can contain up to two URL buttons.
	//
	// 	- In a WhatsApp message template, a quick reply button cannot be used together with a phone call button or a URL button.
	Buttons []*ModifyChatappTemplateRequestComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description of the media resource.
	//
	// >  If the Type sub-parameter of the Components parameter is set to **HEADER*	- and the Format parameter is set to **IMAGE, DOCUMENT, or VIDEO**, you can specify this parameter.
	//
	// example:
	//
	// This is a video
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The carousel cards of the carousel template.
	Cards []*ModifyChatappTemplateRequestComponentsCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// The validity period of the verification code in the WhatsApp authentication template. Unit: minutes. This parameter is valid only when Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to FOOTER. The validity period of the verification code is displayed in the footer.
	//
	// example:
	//
	// 5
	CodeExpirationMinutes *int32 `json:"CodeExpirationMinutes,omitempty" xml:"CodeExpirationMinutes,omitempty"`
	// The length of the video in the Viber message template. Unit: seconds. Valid values: 0 to 600.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the document.
	//
	// >  If the Type sub-parameter of the Components parameter is set to **HEADER*	- and the Format parameter is set to **DOCUMENT**, you can specify this parameter.
	//
	// example:
	//
	// video name
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the document attached in the Viber message template.
	//
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The type of the media resource. Valid values:
	//
	// 	- **TEXT**
	//
	// 	- **IMAGE**
	//
	// 	- **DOCUMENT**
	//
	// 	- **VIDEO**
	//
	// example:
	//
	// TEXT
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Specifies whether the coupon code has an expiration time. Specify this parameter if the Type sub-parameter of the Components parameter is set to LIMITED_TIME_OFFER.
	//
	// example:
	//
	// true
	HasExpiration *bool `json:"HasExpiration,omitempty" xml:"HasExpiration,omitempty"`
	// The text of the message that you want to send.
	//
	// >  If Category is set to AUTHENTICATION, do not specify the Text sub-parameter of the Components parameter.
	//
	// example:
	//
	// hello chatapp
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The thumbnail URL of the video in the Viber message template.
	//
	// example:
	//
	// https://cdn.multiplymall.mobiapp.cloud/cloudcode/yc-165407506207478-165511576113195/20220905/ec5b9737-1507-4208-bb27-8da3958da961.jpg?x-oss-process=image/resize,w_100
	ThumbUrl *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
	// The component type. Valid values:
	//
	// 	- **BODY**
	//
	// 	- **HEADER**
	//
	// 	- **FOOTER**
	//
	// 	- **BUTTONS**
	//
	// 	- **CAROUSEL**
	//
	// 	- **LIMITED_TIME_OFFER**
	//
	// >
	//
	// 	- In a WhatsApp message template, a **Body*	- component cannot exceed 1,024 characters in length. A **HEADER*	- or **FOOTER*	- component cannot exceed 60 characters in length.
	//
	// 	- **FOOTER**, **CAROUSEL**, and **LIMITED_TIME_OFFER*	- components are not supported in Viber message templates.
	//
	// 	- In Viber message templates, media resources such as images, videos, and documents are placed in the **HEADER*	- component. If a Viber message contains text and an image, the image is placed below the text in the message received on a device.
	//
	// This parameter is required.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the media resource.
	//
	// example:
	//
	// https://img.tukuppt.com/png_preview/00/10/24/1GygxVK3F4.jpg
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
	return dara.Validate(s)
}

type ModifyChatappTemplateRequestComponentsButtons struct {
	// The text of the one-tap autofill button. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP for a WhatsApp message template.
	//
	// example:
	//
	// Autofill
	AutofillText *string `json:"AutofillText,omitempty" xml:"AutofillText,omitempty"`
	// The coupon code. It can contain only letters and digits. You can set this parameter to a variable such as $(couponCode). Specify the value of couponCode when you send a message.
	//
	// example:
	//
	// 120293
	CouponCode *string `json:"CouponCode,omitempty" xml:"CouponCode,omitempty"`
	// The Flow action.
	//
	// Valid values:
	//
	// 	- DATA_EXCHANGE
	//
	// 	- NAVIGATE
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
	// The unsubscribe button. This parameter is valid if Category is set to MARKETING and the Type sub-parameter of the Buttons parameter is set to QUICK_REPLY for a WhatsApp message template. Marketing messages will not be sent to customers if you configure message sending in the Chat App Message Service console and the customers click this button.
	//
	// example:
	//
	// false
	IsOptOut *bool `json:"IsOptOut,omitempty" xml:"IsOptOut,omitempty"`
	// The first screen in the Flow. This parameter is required if FlowAction is set to NAVIGATE.
	//
	// example:
	//
	// DETAILS
	NavigateScreen *string `json:"NavigateScreen,omitempty" xml:"NavigateScreen,omitempty"`
	// Deprecated
	//
	// The app package name that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP for a WhatsApp message template.
	//
	// example:
	//
	// com.demo
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The phone number.
	//
	// example:
	//
	// +8613888887889
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Deprecated
	//
	// The app signing key hash that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP for a WhatsApp message template.
	//
	// example:
	//
	// 29dkeke
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
	// List of supported apps.
	SupportedApps []*ModifyChatappTemplateRequestComponentsButtonsSupportedApps `json:"SupportedApps,omitempty" xml:"SupportedApps,omitempty" type:"Repeated"`
	// The text of the button.
	//
	// example:
	//
	// phone-button-text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The button type. Valid values:
	//
	// 	- **PHONE_NUMBER**: phone call button
	//
	// 	- **URL**: URL button
	//
	// 	- **QUICK_REPLY**: quick reply button
	//
	// 	- **COPY_CODE**: copy code button
	//
	// 	- **ONE_TAP**: one-tap autofill button if Category is set to AUTHENTICATION
	//
	// >
	//
	// 	- If Category is set to AUTHENTICATION for a WhatsApp message template, you can add only one button to the WhatsApp message template and you must set the Type sub-parameter of the Buttons parameter to COPY_CODE or ONE_TAP. If Type is set to COPY_CODE, the Text sub-parameter of the Buttons parameter is required. If Type is set to ONE_TAP, the Text, SignatureHash, PackageName, and AutofillText sub-parameters of the Buttons parameter are required. The value of Text is displayed if the desired app is not installed on the device. The value of Text indicates that you must manually copy the verification code.
	//
	// 	- You can add only one button to a Viber message template, and you must set the Type sub-parameter of the Buttons parameter to URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to which you are redirected when you click the URL button.
	//
	// example:
	//
	// https://www.website.com/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL type. Valid values:
	//
	// 	- **static**
	//
	// 	- **dynamic**
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
	return dara.Validate(s)
}

type ModifyChatappTemplateRequestComponentsButtonsSupportedApps struct {
	// The Whatsapp template is required when the Category is\\" Authorisation \\"and the Button Type is\\" ONE_TAP/ZERO-TAP\\", indicating the signature hash value of the Whatsapp call application.
	//
	// example:
	//
	// com.example.myapplication
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The Whatsapp template is required when the Category is\\" Authorisation \\"and the Button Type is\\" ONE_TAP/ZERO-TAP\\", indicating the signature hash value of the Whatsapp call application.
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
	// The components of the carousel card.
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
	return dara.Validate(s)
}

type ModifyChatappTemplateRequestComponentsCardsCardComponents struct {
	// The buttons. Specify this parameter only if you set the Type sub-parameter of the CardComponents parameter to BUTTONS. A carousel card can contain up to two buttons.
	Buttons []*ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The type of the media resource. This parameter is valid if the Type sub-parameter of the CardComponents parameter is set to HEADER. Valid values:
	//
	// 	- **IMAGE**
	//
	// 	- **VIDEO**
	//
	// example:
	//
	// IMAGE
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The body content of the carousel card.
	//
	// example:
	//
	// Who is the very powerful team
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The component type. Valid values:
	//
	// 	- **BODY**
	//
	// 	- **HEADER**
	//
	// 	- **BUTTONS**
	//
	// This parameter is required.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the media resource.
	//
	// example:
	//
	// https://alibaba.com/img.png
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
	return dara.Validate(s)
}

type ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons struct {
	// The phone number.
	//
	// example:
	//
	// +8613800
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The text of the button.
	//
	// example:
	//
	// Call me
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The button type. Valid values:
	//
	// 	- **PHONE_NUMBER**: phone call button
	//
	// 	- **URL**: URL button
	//
	// 	- **QUICK_REPLY**: quick reply button
	//
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to which you are redirected when you click the URL button.
	//
	// example:
	//
	// https://alibaba.com/xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL type. Valid values:
	//
	// 	- **static**
	//
	// 	- **dynamic**
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
