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
}

type GetChatappTemplateDetailResponseBody struct {
	// Access denied details.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code.
	//
	// 	- Example: OK. This value indicates that the request is successful.
	//
	// 	- Other codes indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
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

func (s *GetChatappTemplateDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChatappTemplateDetailResponseBodyData struct {
	AllowSend *bool `json:"AllowSend,omitempty" xml:"AllowSend,omitempty"`
	// The review status of the message template. Valid values:
	//
	// 	- **pass**: The message template is approved.
	//
	// 	- **fail**: The message template is rejected.
	//
	// 	- **auditing**: The message template is being reviewed.
	//
	// 	- **unaudit**: The review is suspended.
	//
	// example:
	//
	// pass
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The category of the template when the returned value of TemplateType is WHATSAPP. Valid values:
	//
	// 	- **UTILITY**: a transactional template
	//
	// 	- **MARKETING**: a marketing template
	//
	// 	- **AUTHENTICATION**: an identity authentication template
	//
	// The category of the template when the returned value of the TemplateType parameter is VIBER. Valid values:
	//
	// 	- **text**: a template that contains only text
	//
	// 	- **image**: a template that contains only images
	//
	// 	- **text_image_button**: a template that contains text, images, and buttons
	//
	// 	- **text_button**: a template that contains text and buttons
	//
	// 	- **document**: a template that contains only files
	//
	// 	- **video**: a template that contains only videos
	//
	// 	- **text_video**: a template that contains text and videos
	//
	// 	- **text_video_button**: a template that contains text, videos, and buttons
	//
	// 	- **text_image**: a template that contains text and images
	//
	// > If Category is set to text_video_button, users cannot open a web page by clicking the button. Users can open only the video in the message. In this case, you do not need to specify the Url parameter for the URL button in the template.
	//
	// example:
	//
	// TRANSACTIONAL
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CategoryChangePaused *bool   `json:"CategoryChangePaused,omitempty" xml:"CategoryChangePaused,omitempty"`
	// The components of the message template.
	Components []*GetChatappTemplateDetailResponseBodyDataComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The examples of variables.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en_US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The validity period of the WhatsApp authentication message.
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The quality of the template.
	//
	// example:
	//
	// GREEN
	QualityScore *string `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	// The reason why the template was rejected.
	//
	// example:
	//
	// None
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
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
	// The note indicating that customers cannot share verification codes with others. The note is displayed in the message body. This parameter is valid if Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to BODY for a WhatsApp message template.
	//
	// example:
	//
	// false
	AddSecretRecommendation *bool `json:"AddSecretRecommendation,omitempty" xml:"AddSecretRecommendation,omitempty"`
	// The buttons. This parameter is returned only if the Type sub-parameter of the Components parameter is set to **BUTTONS**.
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
	Buttons []*GetChatappTemplateDetailResponseBodyDataComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description of the document.
	//
	// example:
	//
	// The new file has been uploaded.
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The carousel cards.
	Cards []*GetChatappTemplateDetailResponseBodyDataComponentsCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// The validity period of the verification code in the WhatsApp authentication template. Unit: minutes. This parameter is valid only when Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to FOOTER for a WhatsApp message template. The validity period of the verification code is displayed in the footer.
	//
	// example:
	//
	// 5
	CodeExpirationMinutes *int32 `json:"CodeExpirationMinutes,omitempty" xml:"CodeExpirationMinutes,omitempty"`
	// The length of the video in the Viber message template. Unit: seconds. Valid values: 0 to 600.
	//
	// example:
	//
	// 50
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the document.
	//
	// example:
	//
	// Express file
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the document attached in the Viber message template.
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
	// The address of the location.
	//
	// example:
	//
	// Hangzhou
	LocationAddress *string `json:"LocationAddress,omitempty" xml:"LocationAddress,omitempty"`
	// The name of the location.
	//
	// example:
	//
	// Hangzhou
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// The longitude of the location.
	//
	// example:
	//
	// 120.002
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// The variable when the coupon code expires in the limited-time offer template.
	//
	// example:
	//
	// $(offerExpirationTimeMs)
	OfferExpirationTimeMs *string `json:"OfferExpirationTimeMs,omitempty" xml:"OfferExpirationTimeMs,omitempty"`
	// The text of the message that you want to send.
	//
	// example:
	//
	// Hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The thumbnail URL of the video in the Viber message template.
	//
	// example:
	//
	// https://img.png
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
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the media resource.
	//
	// example:
	//
	// https://image.developer.aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// Indicates whether the coupon code has an expiration time in the limited-time offer template.
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
	// The text of the one-tap autofill button. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP in a WhatsApp message template.
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
	// The Flow action. Valid values: NAVIGATE and DATA_EXCHANGE.
	//
	// example:
	//
	// NAVIGATE
	FlowAction *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// The Flow ID.
	//
	// example:
	//
	// 3838292983
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The unsubscribe button. This parameter is valid if Category is set to MARKETING and the Type sub-parameter of the Buttons parameter is set to QUICK_REPLY for a WhatsApp message template. Marketing messages will not be sent to customers if you configure message sending in the Chat App Message Service console and the customers click this button.
	//
	// example:
	//
	// false
	IsOptOut *bool `json:"IsOptOut,omitempty" xml:"IsOptOut,omitempty"`
	// The first screen in the Flow. This parameter is returned if FlowAction is set to NAVIGATE.
	//
	// example:
	//
	// DETAILS
	NavigateScreen *string `json:"NavigateScreen,omitempty" xml:"NavigateScreen,omitempty"`
	// The app package name that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP in a WhatsApp message template.
	//
	// example:
	//
	// com.aliyun
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The phone number. This parameter is valid only if the Type sub-parameter of the Buttons parameter is set to **PHONE_NUMBER**.
	//
	// example:
	//
	// 861398745****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The app signing key hash that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP in a WhatsApp message template.
	//
	// example:
	//
	// 2993839
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
	// The apps that support one-tap authentication and zero-tap authentication.
	SupportedApps []*GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps `json:"SupportedApps,omitempty" xml:"SupportedApps,omitempty" type:"Repeated"`
	// The display name of the button.
	//
	// example:
	//
	// Call
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
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to which you are redirected when you click the URL button.
	//
	// example:
	//
	// https://example.com
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
	// The language of the next template.
	//
	// example:
	//
	// en
	NextLanguageCode *string `json:"NextLanguageCode,omitempty" xml:"NextLanguageCode,omitempty"`
	// The code of the next template.
	//
	// example:
	//
	// 20939920093993
	NextTemplateCode *string `json:"NextTemplateCode,omitempty" xml:"NextTemplateCode,omitempty"`
	// The name of the next template.
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
	// The app package name.
	//
	// example:
	//
	// com.test
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The app signing key hash.
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
	// The components of the carousel card.
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
	// The buttons of the carousel card.
	Buttons []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The type of the header in the carousel template. The header can only be an image or a video. The headers of all carousel cards must be the same. The type of the media resources that are included in the message. Valid values: IMGAGE and VIDEO.
	//
	// example:
	//
	// HEADER
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The text of the carousel card.
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
	// The URL.
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
	// +86138000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The button text.
	//
	// example:
	//
	// Button text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the button in the carousel template. Valid values: URL, PHONE_NUMBER, and QUICK_REQLY.
	//
	// example:
	//
	// URL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to which you are redirected when you click the URL button.
	//
	// example:
	//
	// https://aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The type of the URL. Valid values: static and dynamic.
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
