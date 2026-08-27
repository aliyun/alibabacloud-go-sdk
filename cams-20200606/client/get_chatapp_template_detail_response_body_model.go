// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappTemplateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetChatappTemplateDetailResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetChatappTemplateDetailResponseBody
	GetCode() *string
	SetData(v *GetChatappTemplateDetailResponseBodyData) *GetChatappTemplateDetailResponseBody
	GetData() *GetChatappTemplateDetailResponseBodyData
	SetMessage(v string) *GetChatappTemplateDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChatappTemplateDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetChatappTemplateDetailResponseBody
	GetSuccess() *bool
}

type GetChatappTemplateDetailResponseBody struct {
	// The access denied detail information.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// { 		"category": "ACCOUNT_UPDATE", 		"name": "account_notice", 		"language": "en_US", 		"templateCode": "744c4b5c79c9432497a075bdfca3****", 		"auditStatus": "APPROVED", 		"components": "[{\\"type\\":\\"BODY\\",\\"text\\":\\"body_text$(textVariable)\\"},{\\"type\\":\\"HEADER\\",\\"formate\\":\\"IMAGE\\",\\"url\\":\\"$(linkVariable)\\"},{\\"type\\":\\"FOOTER\\",\\"text\\":\\"footer-text\\"},{\\"type\\":\\"BUTTONS\\",\\"buttons\\":[{\\"type\\":\\"PHONE_NUMBER\\",\\"text\\":\\"phone-button-text\\",\\"phone_number\\":\\"+861388888****\\"},{\\"type\\":\\"URL\\",\\"text\\":\\"url-button-text\\",\\"url\\":\\"https://www.website.com/\\"}]}]", 		"example": "{\\"textVariable\\": \\"text\\", \\"linkVariable\\": \\"link\\"}" 	}
	Data *GetChatappTemplateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetChatappTemplateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetChatappTemplateDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatappTemplateDetailResponseBody) GetData() *GetChatappTemplateDetailResponseBodyData {
	return s.Data
}

func (s *GetChatappTemplateDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatappTemplateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatappTemplateDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetChatappTemplateDetailResponseBody) SetAccessDeniedDetail(v string) *GetChatappTemplateDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetCode(v string) *GetChatappTemplateDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetData(v *GetChatappTemplateDetailResponseBodyData) *GetChatappTemplateDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetMessage(v string) *GetChatappTemplateDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetRequestId(v string) *GetChatappTemplateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetSuccess(v bool) *GetChatappTemplateDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChatappTemplateDetailResponseBodyData struct {
	// Indicates whether the current template can be used for sending. This parameter is valid only for Marketing templates.
	//
	// example:
	//
	// true
	AllowSend *bool `json:"AllowSend,omitempty" xml:"AllowSend,omitempty"`
	// The audit status. Valid values:
	//
	// - **pass**: Approved.
	//
	// - **fail**: Rejected.
	//
	// - **auditing**: Under review.
	//
	// - **unaudit**: Review suspended.
	//
	// example:
	//
	// pass
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The WhatsApp template category. Valid values:
	//
	// - **UTILITY**: transaction-related.
	//
	// - **MARKETING**: marketing template.
	//
	// - **AUTHENTICATION**: identity verification.
	//
	// Viber template category. Valid values:
	//
	// - **UTILITY**: transaction-related.
	//
	// - **MARKETING**: marketing template.
	//
	// - **AUTHENTICATION**: identity verification.
	//
	// example:
	//
	// UTILITY
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether template sending is paused when a Utility template is changed to a Marketing template.
	//
	// example:
	//
	// true
	CategoryChangePaused *bool `json:"CategoryChangePaused,omitempty" xml:"CategoryChangePaused,omitempty"`
	// The list of message template components.
	Components []*GetChatappTemplateDetailResponseBodyDataComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The variable example.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The language of the template. For detailed language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en_US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The message time-to-live when sending WhatsApp Authentication template messages.
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// hello_whatsapp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// productSetId
	//
	// example:
	//
	// 939***
	ProductSetId *string `json:"ProductSetId,omitempty" xml:"ProductSetId,omitempty"`
	// The template quality.
	//
	// - RED: Low quality.
	//
	// - YELLOW: Medium quality.
	//
	// - UNKNOWN: Quality unknown.
	//
	// - GREEN: High quality.
	//
	// example:
	//
	// GREEN
	QualityScore *string `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	// The reason for template review rejection.
	//
	// example:
	//
	// None
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the template.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The templatetype.
	//
	// - **WHATSAPP**
	//
	// - **VIBER**
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyData) GetAllowSend() *bool {
	return s.AllowSend
}

func (s *GetChatappTemplateDetailResponseBodyData) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *GetChatappTemplateDetailResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *GetChatappTemplateDetailResponseBodyData) GetCategoryChangePaused() *bool {
	return s.CategoryChangePaused
}

func (s *GetChatappTemplateDetailResponseBodyData) GetComponents() []*GetChatappTemplateDetailResponseBodyDataComponents {
	return s.Components
}

func (s *GetChatappTemplateDetailResponseBodyData) GetExample() map[string]*string {
	return s.Example
}

func (s *GetChatappTemplateDetailResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *GetChatappTemplateDetailResponseBodyData) GetMessageSendTtlSeconds() *int32 {
	return s.MessageSendTtlSeconds
}

func (s *GetChatappTemplateDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetChatappTemplateDetailResponseBodyData) GetProductSetId() *string {
	return s.ProductSetId
}

func (s *GetChatappTemplateDetailResponseBodyData) GetQualityScore() *string {
	return s.QualityScore
}

func (s *GetChatappTemplateDetailResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *GetChatappTemplateDetailResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetChatappTemplateDetailResponseBodyData) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetChatappTemplateDetailResponseBodyData) SetAllowSend(v bool) *GetChatappTemplateDetailResponseBodyData {
	s.AllowSend = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetAuditStatus(v string) *GetChatappTemplateDetailResponseBodyData {
	s.AuditStatus = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetCategory(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetCategoryChangePaused(v bool) *GetChatappTemplateDetailResponseBodyData {
	s.CategoryChangePaused = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetComponents(v []*GetChatappTemplateDetailResponseBodyDataComponents) *GetChatappTemplateDetailResponseBodyData {
	s.Components = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetExample(v map[string]*string) *GetChatappTemplateDetailResponseBodyData {
	s.Example = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetLanguage(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetMessageSendTtlSeconds(v int32) *GetChatappTemplateDetailResponseBodyData {
	s.MessageSendTtlSeconds = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetName(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetProductSetId(v string) *GetChatappTemplateDetailResponseBodyData {
	s.ProductSetId = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetQualityScore(v string) *GetChatappTemplateDetailResponseBodyData {
	s.QualityScore = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetReason(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetTemplateCode(v string) *GetChatappTemplateDetailResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetTemplateType(v string) *GetChatappTemplateDetailResponseBodyData {
	s.TemplateType = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) Validate() error {
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

type GetChatappTemplateDetailResponseBodyDataComponents struct {
	// Valid for WhatsApp templates when Category is Authentication and Component Type is Body. Indicates whether a security recommendation message that advises users not to share the verification code is displayed above the Body.
	//
	// example:
	//
	// false
	AddSecretRecommendation *bool `json:"AddSecretRecommendation,omitempty" xml:"AddSecretRecommendation,omitempty"`
	// The button list. Applicable only to the **BUTTONS*	- component.
	//
	// > WhatsApp button quantity rules:
	//
	// > - For WhatsApp templates with Category set to MARKETING or UTILITY, a maximum of 10 buttons are allowed.
	//
	// > - Only 1 PHONE_NUMBER button is allowed.
	//
	// > - A maximum of 2 URL buttons are allowed.
	//
	// > - QUICK_REPLY buttons cannot be mixed in random order with PHONE_NUMBER or URL buttons.
	Buttons []*GetChatappTemplateDetailResponseBodyDataComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The file description.
	//
	// example:
	//
	// example
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The list of carousel cards.
	Cards []*GetChatappTemplateDetailResponseBodyDataComponentsCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// The verification code validity period in minutes for WhatsApp Authentication templates. Valid only when the message type is WhatsApp, Category is Authentication, and Component Type is Footer. This information is displayed in the Footer position.
	//
	// example:
	//
	// 5
	CodeExpirationMinutes *int32 `json:"CodeExpirationMinutes,omitempty" xml:"CodeExpirationMinutes,omitempty"`
	// The video duration of a Viber video message. Valid values: 0 to 600.
	//
	// example:
	//
	// 50
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The file name.
	//
	// example:
	//
	// example
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type of a Viber file message.
	//
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The format.
	//
	// example:
	//
	// TEXT
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The latitude of the location.
	//
	// example:
	//
	// 28.001
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// The location address.
	//
	// example:
	//
	// hangzhou
	LocationAddress *string `json:"LocationAddress,omitempty" xml:"LocationAddress,omitempty"`
	// The location name.
	//
	// example:
	//
	// hangzhou
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// The longitude of the location.
	//
	// example:
	//
	// 120.002
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// The offer code expiration variable in a Limited Time Offer (LTO) template.
	//
	// example:
	//
	// $(offerExpirationTimeMs)
	OfferExpirationTimeMs *string `json:"OfferExpirationTimeMs,omitempty" xml:"OfferExpirationTimeMs,omitempty"`
	// The text of the message to be sent.
	//
	// example:
	//
	// hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The thumbnail for a Viber video message.
	//
	// example:
	//
	// https://img.png
	ThumbUrl *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
	// The component type.
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
	// > - For WhatsApp templates, the **BODY*	- component cannot exceed 1024 characters. The **HEADER*	- and **FOOTER*	- components cannot exceed 60 characters.
	//
	// > - For Viber templates, the **FOOTER**, **CAROUSEL**, and **LIMITED_TIME_OFFER*	- types are invalid.
	//
	// > - For Viber templates, images, videos, and files are placed in the **HEADER*	- (the device displays images below the text).
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The material URL.
	//
	// example:
	//
	// https://image.developer.aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// Specifies whether the offer code has an expiration time in a Limited Time Offer (LTO) template.
	//
	// example:
	//
	// true
	HasExpiration *bool `json:"hasExpiration,omitempty" xml:"hasExpiration,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponents) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponents) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetAddSecretRecommendation() *bool {
	return s.AddSecretRecommendation
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetButtons() []*GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	return s.Buttons
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetCaption() *string {
	return s.Caption
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetCards() []*GetChatappTemplateDetailResponseBodyDataComponentsCards {
	return s.Cards
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetCodeExpirationMinutes() *int32 {
	return s.CodeExpirationMinutes
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetDuration() *int32 {
	return s.Duration
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetFileName() *string {
	return s.FileName
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetFileType() *string {
	return s.FileType
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetFormat() *string {
	return s.Format
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetLatitude() *string {
	return s.Latitude
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetLocationAddress() *string {
	return s.LocationAddress
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetLocationName() *string {
	return s.LocationName
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetLongitude() *string {
	return s.Longitude
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetOfferExpirationTimeMs() *string {
	return s.OfferExpirationTimeMs
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetText() *string {
	return s.Text
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetThumbUrl() *string {
	return s.ThumbUrl
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetType() *string {
	return s.Type
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetUrl() *string {
	return s.Url
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) GetHasExpiration() *bool {
	return s.HasExpiration
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetAddSecretRecommendation(v bool) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.AddSecretRecommendation = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetButtons(v []*GetChatappTemplateDetailResponseBodyDataComponentsButtons) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Buttons = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetCaption(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Caption = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetCards(v []*GetChatappTemplateDetailResponseBodyDataComponentsCards) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Cards = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetCodeExpirationMinutes(v int32) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.CodeExpirationMinutes = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetDuration(v int32) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Duration = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFileName(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.FileName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFileType(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.FileType = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFormat(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Format = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetLatitude(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Latitude = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetLocationAddress(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.LocationAddress = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetLocationName(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.LocationName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetLongitude(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Longitude = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetOfferExpirationTimeMs(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.OfferExpirationTimeMs = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetThumbUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.ThumbUrl = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Url = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetHasExpiration(v bool) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.HasExpiration = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) Validate() error {
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

type GetChatappTemplateDetailResponseBodyDataComponentsButtons struct {
	// Required for WhatsApp templates with Category set to Authentication and Button Type set to ONE_TAP. Specifies the button text for the WhatsApp Autofill operation.
	//
	// example:
	//
	// Autofill
	AutofillText *string `json:"AutofillText,omitempty" xml:"AutofillText,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 202039ksjs
	CouponCode *string `json:"CouponCode,omitempty" xml:"CouponCode,omitempty"`
	// The extended fields.
	ExtendAttrs *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs `json:"ExtendAttrs,omitempty" xml:"ExtendAttrs,omitempty" type:"Struct"`
	// The Flow data event type. Valid values:
	//
	// - NAVIGATE: navigation.
	//
	// - DATA_EXCHANGE: data exchange.
	//
	// example:
	//
	// NAVIGATE
	FlowAction *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// Flow ID
	//
	// example:
	//
	// 3838292983
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// Valid for WhatsApp templates with Category set to Marketing and Button Type set to QUICK_REPLY. Indicates that the button is a marketing opt-out button. If a customer clicks this button and the send control operation is configured on the Chat App Message Service platform, subsequent Marketing messages will not be sent to the customer.
	//
	// example:
	//
	// false
	IsOptOut *bool `json:"IsOptOut,omitempty" xml:"IsOptOut,omitempty"`
	// The navigate screen. Required when FlowAction is set to NAVIGATE.
	//
	// example:
	//
	// DETAILS
	NavigateScreen *string `json:"NavigateScreen,omitempty" xml:"NavigateScreen,omitempty"`
	// Required for WhatsApp templates with Category set to Authentication and Button Type set to ONE_TAP. Specifies the package name for WhatsApp to launch the application.
	//
	// example:
	//
	// com.aliyun
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The phone number. Valid only when the button type is **PHONE_NUMBER**.
	//
	// example:
	//
	// +861398745****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Required for WhatsApp templates with Category set to Authentication and Button Type set to ONE_TAP. Specifies the signature hash value for WhatsApp to launch the application.
	//
	// example:
	//
	// 2993839
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
	// The applications supported by the ONE_TAP/ZERO_TAP verification code.
	SupportedApps []*GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps `json:"SupportedApps,omitempty" xml:"SupportedApps,omitempty" type:"Repeated"`
	// The display name of the button.
	//
	// example:
	//
	// example
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The button type.
	//
	// - **PHONE_NUMBER**: phone call button.
	//
	// - **URL**: web page button.
	//
	// - **QUICK_REPLY**: quick reply button.
	//
	// - **COPY_CODE**: copy verification code or coupon code.
	//
	// - **ONE_TAP**: autofill button for AUTHENTICATION templates.
	//
	// - **ZERO_TAP**: autofill button for AUTHENTICATION templates.
	//
	// - **MPM**: multi-product catalog.
	//
	// - **CATALOG**: catalog.
	//
	// - **FLOW**: open WhatsApp flow.
	//
	// > - For WhatsApp templates with Category set to AUTHENTICATION, only one button is allowed, and the type can only be COPY_CODE or ONE_TAP. If the type is COPY_CODE, Text is required. If the type is ONE_TAP, Text (displayed when the target application is not installed on the device, indicating the name of the copy verification code button), SignatureHash, PackageName, and AutofillText are required.
	//
	// > - Viber templates allow only one button, and it must be of the URL type.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL that is accessed when the link button is clicked.
	//
	// example:
	//
	// https://example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL type.
	//
	// - **static**: static.
	//
	// - **dynamic**: dynamic.
	//
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtons) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtons) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetAutofillText() *string {
	return s.AutofillText
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetCouponCode() *string {
	return s.CouponCode
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetExtendAttrs() *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	return s.ExtendAttrs
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetFlowAction() *string {
	return s.FlowAction
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetFlowId() *string {
	return s.FlowId
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetIsOptOut() *bool {
	return s.IsOptOut
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetNavigateScreen() *string {
	return s.NavigateScreen
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetPackageName() *string {
	return s.PackageName
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetSignatureHash() *string {
	return s.SignatureHash
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetSupportedApps() []*GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps {
	return s.SupportedApps
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetText() *string {
	return s.Text
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetType() *string {
	return s.Type
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetUrl() *string {
	return s.Url
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) GetUrlType() *string {
	return s.UrlType
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetAutofillText(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.AutofillText = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetCouponCode(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.CouponCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetExtendAttrs(v *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.ExtendAttrs = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetFlowAction(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.FlowAction = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetFlowId(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.FlowId = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetIsOptOut(v bool) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.IsOptOut = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetNavigateScreen(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.NavigateScreen = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetPackageName(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.PackageName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetPhoneNumber(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetSignatureHash(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.SignatureHash = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetSupportedApps(v []*GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.SupportedApps = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.Url = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetUrlType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.UrlType = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) Validate() error {
	if s.ExtendAttrs != nil {
		if err := s.ExtendAttrs.Validate(); err != nil {
			return err
		}
	}
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

type GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs struct {
	// The event type.
	//
	// example:
	//
	// nextCard
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The intent code.
	//
	// example:
	//
	// test
	IntentCode *string `json:"IntentCode,omitempty" xml:"IntentCode,omitempty"`
	// The next template language.
	//
	// example:
	//
	// en
	NextLanguageCode *string `json:"NextLanguageCode,omitempty" xml:"NextLanguageCode,omitempty"`
	// The next template code.
	//
	// example:
	//
	// 20939920093993
	NextTemplateCode *string `json:"NextTemplateCode,omitempty" xml:"NextTemplateCode,omitempty"`
	// The next template name.
	//
	// example:
	//
	// abc
	NextTemplateName *string `json:"NextTemplateName,omitempty" xml:"NextTemplateName,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) GetAction() *string {
	return s.Action
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) GetIntentCode() *string {
	return s.IntentCode
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) GetNextLanguageCode() *string {
	return s.NextLanguageCode
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) GetNextTemplateCode() *string {
	return s.NextTemplateCode
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) GetNextTemplateName() *string {
	return s.NextTemplateName
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetAction(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.Action = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetIntentCode(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.IntentCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetNextLanguageCode(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.NextLanguageCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetNextTemplateCode(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.NextTemplateCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetNextTemplateName(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.NextTemplateName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) Validate() error {
	return dara.Validate(s)
}

type GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps struct {
	// The package name.
	//
	// example:
	//
	// com.test
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The package signature hash.
	//
	// example:
	//
	// 29kdkeik939
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) GetPackageName() *string {
	return s.PackageName
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) GetSignatureHash() *string {
	return s.SignatureHash
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) SetPackageName(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps {
	s.PackageName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) SetSignatureHash(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps {
	s.SignatureHash = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) Validate() error {
	return dara.Validate(s)
}

type GetChatappTemplateDetailResponseBodyDataComponentsCards struct {
	// The list of card components.
	CardComponents []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents `json:"CardComponents,omitempty" xml:"CardComponents,omitempty" type:"Repeated"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCards) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCards) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCards) GetCardComponents() []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	return s.CardComponents
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCards) SetCardComponents(v []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) *GetChatappTemplateDetailResponseBodyDataComponentsCards {
	s.CardComponents = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCards) Validate() error {
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

type GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents struct {
	// The list of card buttons.
	Buttons []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The header type in a carousel template. Only IMAGE and VIDEO are supported. All cards must have the same header type.
	//
	// example:
	//
	// HEADER
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The card text content.
	//
	// example:
	//
	// Body
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The component type.
	//
	// example:
	//
	// HEADER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The web URL.
	//
	// example:
	//
	// https://aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) GetButtons() []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	return s.Buttons
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) GetFormat() *string {
	return s.Format
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) GetText() *string {
	return s.Text
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) GetType() *string {
	return s.Type
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) GetUrl() *string {
	return s.Url
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetButtons(v []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Buttons = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetFormat(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Format = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Url = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) Validate() error {
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

type GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons struct {
	// The phone number.
	//
	// example:
	//
	// +861380005****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The button content.
	//
	// example:
	//
	// example
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The button type for a carousel template. Valid values: URL, PHONE_NUMBER, or QUICK_REPLY.
	//
	// example:
	//
	// URL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to visit when the button is clicked.
	//
	// example:
	//
	// https://aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL type. Valid values:
	//
	//
	//
	// - static: static.
	//
	// - dynamic: dynamic.
	//
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) GetText() *string {
	return s.Text
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) GetType() *string {
	return s.Type
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) GetUrl() *string {
	return s.Url
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) GetUrlType() *string {
	return s.UrlType
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetPhoneNumber(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.Url = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetUrlType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.UrlType = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) Validate() error {
	return dara.Validate(s)
}
