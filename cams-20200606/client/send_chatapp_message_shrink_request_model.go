// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *SendChatappMessageShrinkRequest
	GetAdAccountId() *string
	SetCategory(v string) *SendChatappMessageShrinkRequest
	GetCategory() *string
	SetChannelType(v string) *SendChatappMessageShrinkRequest
	GetChannelType() *string
	SetContent(v string) *SendChatappMessageShrinkRequest
	GetContent() *string
	SetContextMessageId(v string) *SendChatappMessageShrinkRequest
	GetContextMessageId() *string
	SetCustSpaceId(v string) *SendChatappMessageShrinkRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *SendChatappMessageShrinkRequest
	GetCustWabaId() *string
	SetFallBackContent(v string) *SendChatappMessageShrinkRequest
	GetFallBackContent() *string
	SetFallBackDuration(v int32) *SendChatappMessageShrinkRequest
	GetFallBackDuration() *int32
	SetFallBackId(v string) *SendChatappMessageShrinkRequest
	GetFallBackId() *string
	SetFallBackRule(v string) *SendChatappMessageShrinkRequest
	GetFallBackRule() *string
	SetFlowActionShrink(v string) *SendChatappMessageShrinkRequest
	GetFlowActionShrink() *string
	SetFrom(v string) *SendChatappMessageShrinkRequest
	GetFrom() *string
	SetIsvCode(v string) *SendChatappMessageShrinkRequest
	GetIsvCode() *string
	SetLabel(v string) *SendChatappMessageShrinkRequest
	GetLabel() *string
	SetLanguage(v string) *SendChatappMessageShrinkRequest
	GetLanguage() *string
	SetMessageCampaignId(v string) *SendChatappMessageShrinkRequest
	GetMessageCampaignId() *string
	SetMessageType(v string) *SendChatappMessageShrinkRequest
	GetMessageType() *string
	SetOwnerId(v int64) *SendChatappMessageShrinkRequest
	GetOwnerId() *int64
	SetPayloadShrink(v string) *SendChatappMessageShrinkRequest
	GetPayloadShrink() *string
	SetProductActionShrink(v string) *SendChatappMessageShrinkRequest
	GetProductActionShrink() *string
	SetRecipientType(v string) *SendChatappMessageShrinkRequest
	GetRecipientType() *string
	SetResourceOwnerAccount(v string) *SendChatappMessageShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendChatappMessageShrinkRequest
	GetResourceOwnerId() *int64
	SetTag(v string) *SendChatappMessageShrinkRequest
	GetTag() *string
	SetTaskId(v string) *SendChatappMessageShrinkRequest
	GetTaskId() *string
	SetTemplateCode(v string) *SendChatappMessageShrinkRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *SendChatappMessageShrinkRequest
	GetTemplateName() *string
	SetTemplateParamsShrink(v string) *SendChatappMessageShrinkRequest
	GetTemplateParamsShrink() *string
	SetTo(v string) *SendChatappMessageShrinkRequest
	GetTo() *string
	SetTokenType(v string) *SendChatappMessageShrinkRequest
	GetTokenType() *string
	SetTrackingData(v string) *SendChatappMessageShrinkRequest
	GetTrackingData() *string
	SetTtl(v int32) *SendChatappMessageShrinkRequest
	GetTtl() *int32
	SetType(v string) *SendChatappMessageShrinkRequest
	GetType() *string
}

type SendChatappMessageShrinkRequest struct {
	// The Meta ad account ID.
	//
	// > This parameter is a test parameter that is not fully available. Ignore this parameter.
	//
	// example:
	//
	// 123123********
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// The message category (for WhatsApp direct send).
	//
	// 	Warning: Do not specify this parameter unless you are a Meta-invited customer. Otherwise, message delivery will fail.
	//
	// example:
	//
	// UTILITY
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The channel type. Valid values:
	//
	// - **whatsapp*	-
	//
	// - **messenger*	-
	//
	// - **instagram**
	//
	// - **telegram**
	//
	// <props="intl">- **viber**
	//
	// This parameter is required.
	//
	// example:
	//
	// whatsapp
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The message content.
	//
	// **WhatsApp message notes:**
	//
	// - If **messageType*	- is **text**, the **text*	- field is required and the **Caption*	- field must not be specified.
	//
	// - If **messageType*	- is **image**, the **Link*	- field is required.
	//
	// - If **messageType*	- is **video**, the **Link*	- field is required.
	//
	// - If **messageType*	- is **audio**, the **Link*	- field is required and the **Caption*	- field is invalid.
	//
	// - If **messageType*	- is **document**, the **Link*	- and **FileName*	- fields are required and the **Caption*	- field is invalid.
	//
	// - If **messageType*	- is **interactive**, the **type*	- and **action*	- fields are required.
	//
	// - If **messageType*	- is **contacts**, the **name*	- field is required.
	//
	// - If **messageType*	- is **location**, the **longitude*	- and **latitude*	- fields are required.
	//
	// - If **messageType*	- is **sticker**, the **Link*	- field is required and the **Caption*	- and **FileName*	- fields are invalid.
	//
	// - If **messageType*	- is **reaction**, the **messageId*	- and **emoji*	- fields are required.
	//
	//
	// **Messenger message notes:**
	//
	// - If **messageType*	- is **text**, the **text*	- field is required.
	//
	// - If **messageType*	- is **image**, **video**, **audio**, or **document**, the **link*	- field is required.
	//
	// **Instagram message notes:**
	//
	// - If **messageType*	- is **text**, the **text*	- field is required.
	//
	// - If **messageType*	- is **image**, **video**, or **audio**, the **link*	- field is required.
	//
	//
	// <props="intl">**Viber message notes:**
	//
	// <props="intl">- If **messageType*	- is **text**, the **text*	- field is required.
	//
	// <props="intl">- If **messageType*	- is **image**, the **link*	- field is required.
	//
	// <props="intl">- If **messageType*	- is **video**, the **link**, **thumbnail**, **fileSize**, and **duration*	- fields are required.
	//
	// <props="intl">- If **messageType*	- is **document**, the **link**, **fileName**, and **fileType*	- fields are required.
	//
	// <props="intl">- If **messageType*	- is **text_button**, the **text**, **caption**, and **action*	- fields are required.
	//
	// <props="intl">- If **messageType*	- is **text_image_button**, the **text**, **link**, **caption**, and **action*	- fields are required.
	//
	// <props="intl">- If **messageType*	- is **text_video**, the **text**, **link**, **thumbnail**, **fileSize**, and **duration*	- fields are required.
	//
	// <props="intl">- If **messageType*	- is **text_video_button**, the **text**, **link**, **thumbnail**, **fileSize**, **duration**, and **caption*	- fields are required, and the **action*	- field must not be empty.
	//
	// example:
	//
	// {
	//
	//   "text": "hello,whatsapp",
	//
	//   "link": "https://*******",
	//
	//   "caption": "****",
	//
	//   "fileName": "****"
	//
	// }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the message to reply to. This is the ID of a previously sent or received message.
	//
	// example:
	//
	// 61851ccb2f1365b16aee****
	ContextMessageId *string `json:"ContextMessageId,omitempty" xml:"ContextMessageId,omitempty"`
	// The SpaceId of the ISV sub-customer, or the direct customer instance ID. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// example:
	//
	// cams-8c8*********
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The ISV customer WABA ID. This parameter is deprecated. Use CustSpaceId instead, which is the direct customer instance ID. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// example:
	//
	// cams-8c8*********
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The custom fallback content. This parameter is for the China site (Chinese mainland). China site users can ignore this parameter.
	//
	// example:
	//
	// Fallback SMS
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// The fallback trigger time. This parameter is for the international site. China site users can ignore this parameter. <props="intl">If the message does not return a delivered receipt within the specified time, fallback is triggered. If this parameter is not specified, fallback is not triggered based on time and occurs only when the message fails to send or a failure status report is received. Unit: seconds. Minimum value: 60. Maximum value: 43200.
	//
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The fallback policy ID. This parameter is for the China site (Chinese mainland). China site users can ignore this parameter. <props="intl">You can view the policy ID on the [**Fallback Policy**](https://chatapp.console.alibabacloud.com/FallbackStrategy) page.
	//
	// example:
	//
	// S0****
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The fallback rule. This parameter is for the international site. China site users can ignore this parameter.
	//
	// <props="intl">Valid values:
	//
	// <props="intl">- **undelivered**: fallback is triggered when the message cannot be delivered to the device (template and parameter validation must pass at the sending stage; template bans or number bans are not validated). This rule is used by default if the parameter value is empty.
	//
	// <props="intl">- **sentFailed**: fallback is also triggered when template or template variable validation fails. Only the channelType, type, messageType, to, and from (existence check) parameters are strictly validated.
	//
	// example:
	//
	// undelivered
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// The Flow message object.
	FlowActionShrink *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// The sender number.
	//
	// - If ChannelType is set to **whatsapp**, this is the phone number registered and bindded with WhatsApp. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **WABA Management*	- > **Phone Number Management*	- page.
	//
	// - If ChannelType is set to **messenger**, this is the Page ID. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Public Page*	- page.
	//
	// - If ChannelType is set to **instagram**, this is the Instagram professional account ID (Account ID). You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Professional Account*	- page.
	//
	// <props="intl">- If ChannelType is set to **viber**, this is the Viber Service ID. You can view it on the [**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Service ID Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861387777****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Deprecated
	//
	// The ISV verification code used to verify whether a RAM user is authorized by the ISV. This parameter is deprecated and can be ignored.
	//
	// example:
	//
	// 123123******
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The Viber message type. This parameter is for the international site. China site users can ignore this parameter.
	//
	// <props="intl">Valid values:
	//
	// <props="intl">- **pormotion**: marketing or promotional messages.
	//
	// <props="intl">- **transaction**: notification messages.
	//
	// example:
	//
	// promotion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language. For a list of language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The campaign message ID.
	//
	// > This parameter is a test parameter that is not fully available. Ignore this parameter.
	//
	// example:
	//
	// 123123********
	MessageCampaignId *string `json:"MessageCampaignId,omitempty" xml:"MessageCampaignId,omitempty"`
	// The detailed message type when Type is set to message. Valid values:
	//
	// <details>
	//
	// <summary>WHATSAPP</summary>
	//
	// - text: text message.
	//
	// - image: image message.
	//
	// - video: video message.
	//
	// - audio: audio message.
	//
	// - document: document message.
	//
	// - interactive: interactive message.
	//
	// - location: location message.
	//
	// - contacts: contacts message.
	//
	// - reaction: reaction message.
	//
	// - sticker: sticker message.
	//
	// - typing_indicator: typing indicator message.
	//
	// - pin: pin or unpin message (available only for group messages).
	//
	// - carousel: carousel message.
	//
	// </details>
	//
	// <details>
	//
	// <summary>VIBER</summary>
	//
	// - text: text message.
	//
	// - image: image message.
	//
	// - text_image_button: text + image + button message.
	//
	// - text_button: text + button message.
	//
	// - document: document message.
	//
	// - video: video message.
	//
	// - text_video: text + video message.
	//
	// - text_video_button: text + video + button message.
	//
	// - text_image: text + image message.
	//
	// </details>
	//
	//
	// <details>
	//
	// <summary>MESSENGER / INSTAGRAM</summary>
	//
	// - text: text message.
	//
	// - image: image message.
	//
	// - video: video message.
	//
	// - document: document message.
	//
	// - audio: audio message.
	//
	// - interactive: interactive message.
	//
	// - couponTemplate: coupon template message.
	//
	// - regularTemplate: regular template message.
	//
	// - quickReply: quick reply message.
	//
	// - buttonTemplate: button template message.
	//
	// </details>
	//
	// <details>
	//
	// <summary>TELEGRAM</summary>
	//
	// - text: text message.
	//
	// - image: image message.
	//
	// - video: video message.
	//
	// - audio: audio message.
	//
	// - document: document message.
	//
	// - location: location message.
	//
	// - gif: animated GIF message.
	//
	// - sticker: sticker message.
	//
	// </details>
	//
	// example:
	//
	// text
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The collection of button trigger messages.
	//
	// example:
	//
	// payloadtext1,payloadtext2,payloadtext3
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The product information. This parameter applies only to the WhatsApp channel type and refers to the product information you uploaded on Meta.
	ProductActionShrink *string `json:"ProductAction,omitempty" xml:"ProductAction,omitempty"`
	// The recipient type. Valid values:
	//
	// - individual: an individual.
	//
	// - group: a group.
	//
	// example:
	//
	// individual
	RecipientType        *string `json:"RecipientType,omitempty" xml:"RecipientType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information. Custom tag information for Viber message delivery.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The custom task ID.
	//
	// example:
	//
	// 10000****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The template code. You can view the template code on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Template Design*	- page.
	//
	// example:
	//
	// 1119***************
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name. You can view the template name on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Template Design*	- page.
	//
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The collection of template parameters.
	TemplateParamsShrink *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The recipient number.
	//
	// - If ChannelType is set to **whatsapp**, this is the phone number of the message recipient.
	//
	// - If ChannelType is set to **messenger**, this is the Page-Scoped User ID generated when the user interacts with the Facebook page.
	//
	// - If ChannelType is set to **instagram**, this is the Instagram User ID generated when the user interacts with the Instagram business or creator account.
	//
	// <props="intl">- If ChannelType is set to **viber**, this is the phone number of the message recipient.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861388988****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The token type.
	//
	// > This parameter is a test parameter that is not fully available. Ignore this parameter.
	//
	// example:
	//
	// bearer
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
	// The custom tracking data passed in for Viber message types. This parameter is for the international site. China site users can ignore this parameter.
	//
	// example:
	//
	// Tracking Data
	TrackingData *string `json:"TrackingData,omitempty" xml:"TrackingData,omitempty"`
	// The timeout period for Viber message delivery. This parameter is for the international site. China site users can ignore this parameter. <props="intl">Unit: seconds. Valid values: 30 to 1209600.
	//
	// example:
	//
	// 50
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The message type. Valid values:
	//
	// - template: a message template that has been approved in the console. This type of message can be sent at any time.
	//
	// - message: a message in any format. This type of message can be sent only within 24 hours after the last message is received from the user.
	//
	// 	Notice: If Type is set to template, you must specify TemplateCode. If Type is set to message, you must specify MessageType.
	//
	// This parameter is required.
	//
	// example:
	//
	// message
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendChatappMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMessageShrinkRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *SendChatappMessageShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *SendChatappMessageShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *SendChatappMessageShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *SendChatappMessageShrinkRequest) GetContextMessageId() *string {
	return s.ContextMessageId
}

func (s *SendChatappMessageShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *SendChatappMessageShrinkRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *SendChatappMessageShrinkRequest) GetFallBackContent() *string {
	return s.FallBackContent
}

func (s *SendChatappMessageShrinkRequest) GetFallBackDuration() *int32 {
	return s.FallBackDuration
}

func (s *SendChatappMessageShrinkRequest) GetFallBackId() *string {
	return s.FallBackId
}

func (s *SendChatappMessageShrinkRequest) GetFallBackRule() *string {
	return s.FallBackRule
}

func (s *SendChatappMessageShrinkRequest) GetFlowActionShrink() *string {
	return s.FlowActionShrink
}

func (s *SendChatappMessageShrinkRequest) GetFrom() *string {
	return s.From
}

func (s *SendChatappMessageShrinkRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *SendChatappMessageShrinkRequest) GetLabel() *string {
	return s.Label
}

func (s *SendChatappMessageShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *SendChatappMessageShrinkRequest) GetMessageCampaignId() *string {
	return s.MessageCampaignId
}

func (s *SendChatappMessageShrinkRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *SendChatappMessageShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendChatappMessageShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *SendChatappMessageShrinkRequest) GetProductActionShrink() *string {
	return s.ProductActionShrink
}

func (s *SendChatappMessageShrinkRequest) GetRecipientType() *string {
	return s.RecipientType
}

func (s *SendChatappMessageShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendChatappMessageShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendChatappMessageShrinkRequest) GetTag() *string {
	return s.Tag
}

func (s *SendChatappMessageShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SendChatappMessageShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendChatappMessageShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SendChatappMessageShrinkRequest) GetTemplateParamsShrink() *string {
	return s.TemplateParamsShrink
}

func (s *SendChatappMessageShrinkRequest) GetTo() *string {
	return s.To
}

func (s *SendChatappMessageShrinkRequest) GetTokenType() *string {
	return s.TokenType
}

func (s *SendChatappMessageShrinkRequest) GetTrackingData() *string {
	return s.TrackingData
}

func (s *SendChatappMessageShrinkRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *SendChatappMessageShrinkRequest) GetType() *string {
	return s.Type
}

func (s *SendChatappMessageShrinkRequest) SetAdAccountId(v string) *SendChatappMessageShrinkRequest {
	s.AdAccountId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetCategory(v string) *SendChatappMessageShrinkRequest {
	s.Category = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetChannelType(v string) *SendChatappMessageShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetContent(v string) *SendChatappMessageShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetContextMessageId(v string) *SendChatappMessageShrinkRequest {
	s.ContextMessageId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetCustSpaceId(v string) *SendChatappMessageShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetCustWabaId(v string) *SendChatappMessageShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackContent(v string) *SendChatappMessageShrinkRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackDuration(v int32) *SendChatappMessageShrinkRequest {
	s.FallBackDuration = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackId(v string) *SendChatappMessageShrinkRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackRule(v string) *SendChatappMessageShrinkRequest {
	s.FallBackRule = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFlowActionShrink(v string) *SendChatappMessageShrinkRequest {
	s.FlowActionShrink = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFrom(v string) *SendChatappMessageShrinkRequest {
	s.From = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetIsvCode(v string) *SendChatappMessageShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetLabel(v string) *SendChatappMessageShrinkRequest {
	s.Label = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetLanguage(v string) *SendChatappMessageShrinkRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetMessageCampaignId(v string) *SendChatappMessageShrinkRequest {
	s.MessageCampaignId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetMessageType(v string) *SendChatappMessageShrinkRequest {
	s.MessageType = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetOwnerId(v int64) *SendChatappMessageShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetPayloadShrink(v string) *SendChatappMessageShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetProductActionShrink(v string) *SendChatappMessageShrinkRequest {
	s.ProductActionShrink = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetRecipientType(v string) *SendChatappMessageShrinkRequest {
	s.RecipientType = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetResourceOwnerAccount(v string) *SendChatappMessageShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetResourceOwnerId(v int64) *SendChatappMessageShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTag(v string) *SendChatappMessageShrinkRequest {
	s.Tag = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTaskId(v string) *SendChatappMessageShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTemplateCode(v string) *SendChatappMessageShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTemplateName(v string) *SendChatappMessageShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTemplateParamsShrink(v string) *SendChatappMessageShrinkRequest {
	s.TemplateParamsShrink = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTo(v string) *SendChatappMessageShrinkRequest {
	s.To = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTokenType(v string) *SendChatappMessageShrinkRequest {
	s.TokenType = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTrackingData(v string) *SendChatappMessageShrinkRequest {
	s.TrackingData = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTtl(v int32) *SendChatappMessageShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetType(v string) *SendChatappMessageShrinkRequest {
	s.Type = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
