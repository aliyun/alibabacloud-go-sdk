// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *SendChatappMessageRequest
	GetAdAccountId() *string
	SetCategory(v string) *SendChatappMessageRequest
	GetCategory() *string
	SetChannelType(v string) *SendChatappMessageRequest
	GetChannelType() *string
	SetContent(v string) *SendChatappMessageRequest
	GetContent() *string
	SetContextMessageId(v string) *SendChatappMessageRequest
	GetContextMessageId() *string
	SetCustSpaceId(v string) *SendChatappMessageRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *SendChatappMessageRequest
	GetCustWabaId() *string
	SetFallBackContent(v string) *SendChatappMessageRequest
	GetFallBackContent() *string
	SetFallBackDuration(v int32) *SendChatappMessageRequest
	GetFallBackDuration() *int32
	SetFallBackId(v string) *SendChatappMessageRequest
	GetFallBackId() *string
	SetFallBackRule(v string) *SendChatappMessageRequest
	GetFallBackRule() *string
	SetFlowAction(v *SendChatappMessageRequestFlowAction) *SendChatappMessageRequest
	GetFlowAction() *SendChatappMessageRequestFlowAction
	SetFrom(v string) *SendChatappMessageRequest
	GetFrom() *string
	SetIsvCode(v string) *SendChatappMessageRequest
	GetIsvCode() *string
	SetLabel(v string) *SendChatappMessageRequest
	GetLabel() *string
	SetLanguage(v string) *SendChatappMessageRequest
	GetLanguage() *string
	SetMessageCampaignId(v string) *SendChatappMessageRequest
	GetMessageCampaignId() *string
	SetMessageType(v string) *SendChatappMessageRequest
	GetMessageType() *string
	SetOwnerId(v int64) *SendChatappMessageRequest
	GetOwnerId() *int64
	SetPayload(v []*string) *SendChatappMessageRequest
	GetPayload() []*string
	SetProductAction(v *SendChatappMessageRequestProductAction) *SendChatappMessageRequest
	GetProductAction() *SendChatappMessageRequestProductAction
	SetRecipientType(v string) *SendChatappMessageRequest
	GetRecipientType() *string
	SetResourceOwnerAccount(v string) *SendChatappMessageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendChatappMessageRequest
	GetResourceOwnerId() *int64
	SetTag(v string) *SendChatappMessageRequest
	GetTag() *string
	SetTaskId(v string) *SendChatappMessageRequest
	GetTaskId() *string
	SetTemplateCode(v string) *SendChatappMessageRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *SendChatappMessageRequest
	GetTemplateName() *string
	SetTemplateParams(v map[string]*string) *SendChatappMessageRequest
	GetTemplateParams() map[string]*string
	SetTo(v string) *SendChatappMessageRequest
	GetTo() *string
	SetTokenType(v string) *SendChatappMessageRequest
	GetTokenType() *string
	SetTrackingData(v string) *SendChatappMessageRequest
	GetTrackingData() *string
	SetTtl(v int32) *SendChatappMessageRequest
	GetTtl() *int32
	SetType(v string) *SendChatappMessageRequest
	GetType() *string
}

type SendChatappMessageRequest struct {
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
	FlowAction *SendChatappMessageRequestFlowAction `json:"FlowAction,omitempty" xml:"FlowAction,omitempty" type:"Struct"`
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
	Payload []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	// The product information. This parameter applies only to the WhatsApp channel type and refers to the product information you uploaded on Meta.
	ProductAction *SendChatappMessageRequestProductAction `json:"ProductAction,omitempty" xml:"ProductAction,omitempty" type:"Struct"`
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
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
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

func (s SendChatappMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *SendChatappMessageRequest) GetCategory() *string {
	return s.Category
}

func (s *SendChatappMessageRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *SendChatappMessageRequest) GetContent() *string {
	return s.Content
}

func (s *SendChatappMessageRequest) GetContextMessageId() *string {
	return s.ContextMessageId
}

func (s *SendChatappMessageRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *SendChatappMessageRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *SendChatappMessageRequest) GetFallBackContent() *string {
	return s.FallBackContent
}

func (s *SendChatappMessageRequest) GetFallBackDuration() *int32 {
	return s.FallBackDuration
}

func (s *SendChatappMessageRequest) GetFallBackId() *string {
	return s.FallBackId
}

func (s *SendChatappMessageRequest) GetFallBackRule() *string {
	return s.FallBackRule
}

func (s *SendChatappMessageRequest) GetFlowAction() *SendChatappMessageRequestFlowAction {
	return s.FlowAction
}

func (s *SendChatappMessageRequest) GetFrom() *string {
	return s.From
}

func (s *SendChatappMessageRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *SendChatappMessageRequest) GetLabel() *string {
	return s.Label
}

func (s *SendChatappMessageRequest) GetLanguage() *string {
	return s.Language
}

func (s *SendChatappMessageRequest) GetMessageCampaignId() *string {
	return s.MessageCampaignId
}

func (s *SendChatappMessageRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *SendChatappMessageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendChatappMessageRequest) GetPayload() []*string {
	return s.Payload
}

func (s *SendChatappMessageRequest) GetProductAction() *SendChatappMessageRequestProductAction {
	return s.ProductAction
}

func (s *SendChatappMessageRequest) GetRecipientType() *string {
	return s.RecipientType
}

func (s *SendChatappMessageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendChatappMessageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendChatappMessageRequest) GetTag() *string {
	return s.Tag
}

func (s *SendChatappMessageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SendChatappMessageRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendChatappMessageRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SendChatappMessageRequest) GetTemplateParams() map[string]*string {
	return s.TemplateParams
}

func (s *SendChatappMessageRequest) GetTo() *string {
	return s.To
}

func (s *SendChatappMessageRequest) GetTokenType() *string {
	return s.TokenType
}

func (s *SendChatappMessageRequest) GetTrackingData() *string {
	return s.TrackingData
}

func (s *SendChatappMessageRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *SendChatappMessageRequest) GetType() *string {
	return s.Type
}

func (s *SendChatappMessageRequest) SetAdAccountId(v string) *SendChatappMessageRequest {
	s.AdAccountId = &v
	return s
}

func (s *SendChatappMessageRequest) SetCategory(v string) *SendChatappMessageRequest {
	s.Category = &v
	return s
}

func (s *SendChatappMessageRequest) SetChannelType(v string) *SendChatappMessageRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMessageRequest) SetContent(v string) *SendChatappMessageRequest {
	s.Content = &v
	return s
}

func (s *SendChatappMessageRequest) SetContextMessageId(v string) *SendChatappMessageRequest {
	s.ContextMessageId = &v
	return s
}

func (s *SendChatappMessageRequest) SetCustSpaceId(v string) *SendChatappMessageRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SendChatappMessageRequest) SetCustWabaId(v string) *SendChatappMessageRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackContent(v string) *SendChatappMessageRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackDuration(v int32) *SendChatappMessageRequest {
	s.FallBackDuration = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackId(v string) *SendChatappMessageRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackRule(v string) *SendChatappMessageRequest {
	s.FallBackRule = &v
	return s
}

func (s *SendChatappMessageRequest) SetFlowAction(v *SendChatappMessageRequestFlowAction) *SendChatappMessageRequest {
	s.FlowAction = v
	return s
}

func (s *SendChatappMessageRequest) SetFrom(v string) *SendChatappMessageRequest {
	s.From = &v
	return s
}

func (s *SendChatappMessageRequest) SetIsvCode(v string) *SendChatappMessageRequest {
	s.IsvCode = &v
	return s
}

func (s *SendChatappMessageRequest) SetLabel(v string) *SendChatappMessageRequest {
	s.Label = &v
	return s
}

func (s *SendChatappMessageRequest) SetLanguage(v string) *SendChatappMessageRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMessageRequest) SetMessageCampaignId(v string) *SendChatappMessageRequest {
	s.MessageCampaignId = &v
	return s
}

func (s *SendChatappMessageRequest) SetMessageType(v string) *SendChatappMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendChatappMessageRequest) SetOwnerId(v int64) *SendChatappMessageRequest {
	s.OwnerId = &v
	return s
}

func (s *SendChatappMessageRequest) SetPayload(v []*string) *SendChatappMessageRequest {
	s.Payload = v
	return s
}

func (s *SendChatappMessageRequest) SetProductAction(v *SendChatappMessageRequestProductAction) *SendChatappMessageRequest {
	s.ProductAction = v
	return s
}

func (s *SendChatappMessageRequest) SetRecipientType(v string) *SendChatappMessageRequest {
	s.RecipientType = &v
	return s
}

func (s *SendChatappMessageRequest) SetResourceOwnerAccount(v string) *SendChatappMessageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendChatappMessageRequest) SetResourceOwnerId(v int64) *SendChatappMessageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendChatappMessageRequest) SetTag(v string) *SendChatappMessageRequest {
	s.Tag = &v
	return s
}

func (s *SendChatappMessageRequest) SetTaskId(v string) *SendChatappMessageRequest {
	s.TaskId = &v
	return s
}

func (s *SendChatappMessageRequest) SetTemplateCode(v string) *SendChatappMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMessageRequest) SetTemplateName(v string) *SendChatappMessageRequest {
	s.TemplateName = &v
	return s
}

func (s *SendChatappMessageRequest) SetTemplateParams(v map[string]*string) *SendChatappMessageRequest {
	s.TemplateParams = v
	return s
}

func (s *SendChatappMessageRequest) SetTo(v string) *SendChatappMessageRequest {
	s.To = &v
	return s
}

func (s *SendChatappMessageRequest) SetTokenType(v string) *SendChatappMessageRequest {
	s.TokenType = &v
	return s
}

func (s *SendChatappMessageRequest) SetTrackingData(v string) *SendChatappMessageRequest {
	s.TrackingData = &v
	return s
}

func (s *SendChatappMessageRequest) SetTtl(v int32) *SendChatappMessageRequest {
	s.Ttl = &v
	return s
}

func (s *SendChatappMessageRequest) SetType(v string) *SendChatappMessageRequest {
	s.Type = &v
	return s
}

func (s *SendChatappMessageRequest) Validate() error {
	if s.FlowAction != nil {
		if err := s.FlowAction.Validate(); err != nil {
			return err
		}
	}
	if s.ProductAction != nil {
		if err := s.ProductAction.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendChatappMessageRequestFlowAction struct {
	// The collection of Flow default parameters.
	FlowActionData map[string]interface{} `json:"FlowActionData,omitempty" xml:"FlowActionData,omitempty"`
	// The custom Flow token information.
	//
	// example:
	//
	// kde****
	FlowToken *string `json:"FlowToken,omitempty" xml:"FlowToken,omitempty"`
}

func (s SendChatappMessageRequestFlowAction) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageRequestFlowAction) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequestFlowAction) GetFlowActionData() map[string]interface{} {
	return s.FlowActionData
}

func (s *SendChatappMessageRequestFlowAction) GetFlowToken() *string {
	return s.FlowToken
}

func (s *SendChatappMessageRequestFlowAction) SetFlowActionData(v map[string]interface{}) *SendChatappMessageRequestFlowAction {
	s.FlowActionData = v
	return s
}

func (s *SendChatappMessageRequestFlowAction) SetFlowToken(v string) *SendChatappMessageRequestFlowAction {
	s.FlowToken = &v
	return s
}

func (s *SendChatappMessageRequestFlowAction) Validate() error {
	return dara.Validate(s)
}

type SendChatappMessageRequestProductAction struct {
	// The list of product categories (up to 10 categories and 30 products).
	Sections []*SendChatappMessageRequestProductActionSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// The product catalog ID. You can obtain this ID by calling the [ListProductCatalog](https://help.aliyun.com/document_detail/2539783.html) operation.
	//
	// example:
	//
	// skkks99****
	ThumbnailProductRetailerId *string `json:"ThumbnailProductRetailerId,omitempty" xml:"ThumbnailProductRetailerId,omitempty"`
}

func (s SendChatappMessageRequestProductAction) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageRequestProductAction) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequestProductAction) GetSections() []*SendChatappMessageRequestProductActionSections {
	return s.Sections
}

func (s *SendChatappMessageRequestProductAction) GetThumbnailProductRetailerId() *string {
	return s.ThumbnailProductRetailerId
}

func (s *SendChatappMessageRequestProductAction) SetSections(v []*SendChatappMessageRequestProductActionSections) *SendChatappMessageRequestProductAction {
	s.Sections = v
	return s
}

func (s *SendChatappMessageRequestProductAction) SetThumbnailProductRetailerId(v string) *SendChatappMessageRequestProductAction {
	s.ThumbnailProductRetailerId = &v
	return s
}

func (s *SendChatappMessageRequestProductAction) Validate() error {
	if s.Sections != nil {
		for _, item := range s.Sections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendChatappMessageRequestProductActionSections struct {
	// The list of product information.
	ProductItems []*SendChatappMessageRequestProductActionSectionsProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Repeated"`
	// The category name. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Catalog Management*	- > **Product Management*	- page, or obtain it by calling the [ListProduct](https://help.aliyun.com/document_detail/2557786.html) operation.
	//
	// example:
	//
	// abcd
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SendChatappMessageRequestProductActionSections) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageRequestProductActionSections) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequestProductActionSections) GetProductItems() []*SendChatappMessageRequestProductActionSectionsProductItems {
	return s.ProductItems
}

func (s *SendChatappMessageRequestProductActionSections) GetTitle() *string {
	return s.Title
}

func (s *SendChatappMessageRequestProductActionSections) SetProductItems(v []*SendChatappMessageRequestProductActionSectionsProductItems) *SendChatappMessageRequestProductActionSections {
	s.ProductItems = v
	return s
}

func (s *SendChatappMessageRequestProductActionSections) SetTitle(v string) *SendChatappMessageRequestProductActionSections {
	s.Title = &v
	return s
}

func (s *SendChatappMessageRequestProductActionSections) Validate() error {
	if s.ProductItems != nil {
		for _, item := range s.ProductItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendChatappMessageRequestProductActionSectionsProductItems struct {
	// The product ID. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Catalog Management*	- > **Product Management*	- page, or obtain it by calling the [ListProduct](https://help.aliyun.com/document_detail/2557786.html) operation.
	//
	// example:
	//
	// ksi3****
	ProductRetailerId *string `json:"ProductRetailerId,omitempty" xml:"ProductRetailerId,omitempty"`
}

func (s SendChatappMessageRequestProductActionSectionsProductItems) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageRequestProductActionSectionsProductItems) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequestProductActionSectionsProductItems) GetProductRetailerId() *string {
	return s.ProductRetailerId
}

func (s *SendChatappMessageRequestProductActionSectionsProductItems) SetProductRetailerId(v string) *SendChatappMessageRequestProductActionSectionsProductItems {
	s.ProductRetailerId = &v
	return s
}

func (s *SendChatappMessageRequestProductActionSectionsProductItems) Validate() error {
	return dara.Validate(s)
}
