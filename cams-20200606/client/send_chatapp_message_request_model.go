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
	// > This parameter is for internal testing, is not generally available, and can be ignored.
	//
	// example:
	//
	// 123123********
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// The message category for direct WhatsApp sending.
	//
	// 	Warning:
	//
	// Specify this parameter only if you are a Meta-invited customer. Otherwise, the message may fail to send.
	//
	// example:
	//
	// UTILITY
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The channel type. Valid values:
	//
	// - **whatsapp**
	//
	// - **messenger**
	//
	// - **instagram**
	//
	// - **telegram**
	//
	// <props="intl">
	//
	// - **viber**
	//
	// This parameter is required.
	//
	// example:
	//
	// whatsapp
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The message content, in a JSON-formatted string.
	//
	// **Notes for WhatsApp messages:**
	//
	// - If `MessageType` is `text`, the `text` field is required, and the `Caption` field is not supported.
	//
	// - If `MessageType` is `image`, the `Link` field is required.
	//
	// - If `MessageType` is `video`, the `Link` field is required.
	//
	// - If `MessageType` is `audio`, the `Link` field is required. The `Caption` field is not supported.
	//
	// - If `MessageType` is `document`, the `Link` and `FileName` fields are required. The `Caption` field is not supported.
	//
	// - If `MessageType` is `interactive`, the `type` and `action` fields are required.
	//
	// - If `MessageType` is `contacts`, the `name` field is required.
	//
	// - If `MessageType` is `location`, the `longitude` and `latitude` fields are required.
	//
	// - If `MessageType` is `sticker`, the `Link` field is required. The `Caption` and `FileName` fields are not supported.
	//
	// - If `MessageType` is `reaction`, the `messageId` and `emoji` fields are required.
	//
	// **Notes for Messenger messages:**
	//
	// - If `MessageType` is `text`, the `text` field is required.
	//
	// - If `MessageType` is `image`, `video`, `audio`, or `document`, the `link` field is required.
	//
	// **Notes for Instagram messages:**
	//
	// - If `MessageType` is `text`, the `text` field is required.
	//
	// - If `MessageType` is `image`, `video`, or `audio`, the `link` field is required.
	//
	// <props="intl">
	//
	// **Notes for Viber messages:**
	//
	//
	//
	// <props="intl">
	//
	// - If `MessageType` is `text`, the `text` field is required.
	//
	//
	//
	// <props="intl">
	//
	// - If `MessageType` is `image`, the `link` field is required.
	//
	//
	//
	// <props="intl">
	//
	// - If `MessageType` is `video`, the `link`, `thumbnail`, `fileSize`, and `duration` fields are required.
	//
	//
	//
	// <props="intl">
	//
	// - If `MessageType` is `document`, the `link`, `fileName`, and `fileType` fields are required.
	//
	//
	//
	// <props="intl">
	//
	// - If `MessageType` is `text_button`, the `text`, `caption`, and `action` fields are required.
	//
	//
	//
	// <props="intl">
	//
	// - If `MessageType` is `text_image_button`, the `text`, `link`, `caption`, and `action` fields are required.
	//
	//
	//
	// <props="intl">
	//
	// - If `MessageType` is `text_video`, the `text`, `link`, `thumbnail`, `fileSize`, and `duration` fields are required.
	//
	//
	//
	// <props="intl">
	//
	// - If `MessageType` is `text_video_button`, the `text`, `link`, `thumbnail`, `fileSize`, `duration`, and `caption` fields are required. The `action` field is not supported.
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
	// The ID of the message to which you are replying.
	//
	// example:
	//
	// 61851ccb2f1365b16aee****
	ContextMessageId *string `json:"ContextMessageId,omitempty" xml:"ContextMessageId,omitempty"`
	// The Space ID of the ISV\\"s sub-account. For a direct customer, this is the Instance ID. You can find the ID on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// example:
	//
	// cams-8c8*********
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// **Deprecated.*	- Use `CustSpaceId` instead. The WABA ID of an ISV\\"s customer. For a direct customer, this is the Instance ID. You can find the ID on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// example:
	//
	// cams-8c8*********
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The custom content of the fallback message. This parameter is available only on the International Site and can be ignored if you are using the China site.
	//
	// example:
	//
	// Fallback SMS
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// The duration after which a fallback is triggered. This parameter is available only on the International Site and can be ignored if you are using the China site.<props="intl"> If a delivery receipt is not returned within the specified period, a fallback is triggered. If this parameter is omitted, a fallback is triggered only if the message fails to send or a failed delivery receipt is returned. Unit: seconds. The value must be between 60 and 43200.
	//
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The ID of the fallback strategy. This parameter is available only on the International Site and can be ignored if you are using the China site.<props="intl"> You can find the strategy ID on the [**Fallback Policy**](https://chatapp.console.alibabacloud.com/FallbackStrategy) page.
	//
	// example:
	//
	// S0****
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The fallback rule. This parameter is available only on the International Site and can be ignored if you are using the China site.
	//
	// <props="intl">Valid values:
	//
	// <props="intl">
	//
	// - **undelivered**: A fallback is triggered if message delivery fails. The template and parameters must be valid at the time of sending. Blocked templates or phone numbers are not validated. This is the default rule if the parameter is empty.
	//
	//
	//
	// <props="intl">
	//
	// - **sentFailed**: A fallback is triggered if the message fails parameter validation, such as for the template or template parameters. Only the existence of `channelType`, `type`, `messageType`, `to`, and `from` is strictly validated.
	//
	// example:
	//
	// undelivered
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// The Flow message object.
	FlowAction *SendChatappMessageRequestFlowAction `json:"FlowAction,omitempty" xml:"FlowAction,omitempty" type:"Struct"`
	// The sender\\"s number or ID.
	//
	// - If `ChannelType` is **whatsapp**, this is the phone number registered with WhatsApp. You can find the number on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **WABA Management*	- > **Phone Number Management*	- page.
	//
	// - If `ChannelType` is **messenger**, this is the Facebook Page ID. You can find this ID on your <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Facebook Page*	- page.
	//
	// - If `ChannelType` is **instagram**, this is the Instagram professional account ID (Account ID). You can find the ID on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Professional Account*	- page.
	//
	// <props="intl">
	//
	// - If `ChannelType` is **viber**, this is the Viber service ID (Service ID). You can find the ID on the [**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Service Number Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861387777****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Deprecated
	//
	// **Deprecated.*	- A verification code used to authorize an ISV\\"s sub-account. You can ignore this parameter.
	//
	// example:
	//
	// 123123******
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The Viber message type. This parameter is available only on the International Site and can be ignored if you are using the China site.
	//
	// <props="intl">Valid values:
	//
	// <props="intl">
	//
	// - **promotion**: A promotional or marketing message.
	//
	//
	//
	// <props="intl">
	//
	// - **transaction**: A notification message.
	//
	// example:
	//
	// promotion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language of the message template. For a list of supported languages and their corresponding codes, see [language code](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the message campaign.
	//
	// > This parameter is for internal testing, is not generally available, and can be ignored.
	//
	// example:
	//
	// 123123********
	MessageCampaignId *string `json:"MessageCampaignId,omitempty" xml:"MessageCampaignId,omitempty"`
	// The message type to use when `Type` is set to `message`. The valid values vary based on the channel type:
	//
	// <details>
	//
	// <summary>
	//
	// WHATSAPP
	//
	// </summary>
	//
	// - `text`: A text message.
	//
	// - `image`: An image message.
	//
	// - `video`: A video message.
	//
	// - `audio`: An audio message.
	//
	// - `document`: A document message.
	//
	// - `interactive`: An interactive message.
	//
	// - `location`: A location message.
	//
	// - `contacts`: A contacts message.
	//
	// - `reaction`: A reaction message.
	//
	// - `sticker`: A sticker message.
	//
	// - `typing_indicator`: A typing indicator message.
	//
	// - `pin`: A message to pin or unpin. This type is available only for group messages.
	//
	// - `carousel`: A carousel message.
	//
	// </details>
	//
	// <details>
	//
	// <summary>
	//
	// VIBER
	//
	// </summary>
	//
	// - `text`: A text message.
	//
	// - `image`: An image message.
	//
	// - `text_image_button`: A message with text, an image, and a button.
	//
	// - `text_button`: A message with text and a button.
	//
	// - `document`: A document message.
	//
	// - `video`: A video message.
	//
	// - `text_video`: A message with text and a video.
	//
	// - `text_video_button`: A message with text, a video, and a button.
	//
	// - `text_image`: A message with text and an image.
	//
	// </details>
	//
	// <details>
	//
	// <summary>
	//
	// MESSENGER / INSTAGRAM
	//
	// </summary>
	//
	// - `text`: A text message.
	//
	// - `image`: An image message.
	//
	// - `video`: A video message.
	//
	// - `document`: A document message.
	//
	// - `audio`: An audio message.
	//
	// - `interactive`: An interactive message.
	//
	// - `couponTemplate`: A coupon template message.
	//
	// - `regularTemplate`: A regular template message.
	//
	// - `quickReply`: A quick reply message.
	//
	// - `buttonTemplate`: A button template message.
	//
	// </details>
	//
	// <details>
	//
	// <summary>
	//
	// TELEGRAM
	//
	// </summary>
	//
	// - `text`: A text message.
	//
	// - `image`: An image message.
	//
	// - `video`: A video message.
	//
	// - `audio`: An audio message.
	//
	// - `document`: A document message.
	//
	// - `location`: A location message.
	//
	// - `gif`: An animated GIF message.
	//
	// - `sticker`: A sticker message.
	//
	// </details>
	//
	// example:
	//
	// text
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// An array of custom data strings that are sent to your webhook when a user clicks a corresponding button.
	//
	// example:
	//
	// payloadtext1,payloadtext2,payloadtext3
	Payload []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	// Product information that you have uploaded to Meta. This parameter applies to WhatsApp channels only.
	ProductAction *SendChatappMessageRequestProductAction `json:"ProductAction,omitempty" xml:"ProductAction,omitempty" type:"Struct"`
	// The recipient type. Valid values:
	//
	// - `individual`: A single recipient.
	//
	// - `group`: A group.
	//
	// example:
	//
	// individual
	RecipientType        *string `json:"RecipientType,omitempty" xml:"RecipientType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A custom tag for the Viber message.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// A custom task ID.
	//
	// example:
	//
	// 10000****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The message template code. You can find the code on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Template Design*	- page.
	//
	// example:
	//
	// 1119***************
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name. You can find the template name on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Template Design*	- page.
	//
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The parameters for the message template.
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The recipient\\"s number or ID.
	//
	// - If `ChannelType` is **whatsapp**, this is the recipient\\"s phone number.
	//
	// - If `ChannelType` is **messenger**, this is a Page-Scoped User ID (PSID) generated when a user interacts with your Facebook Page.
	//
	// - If `ChannelType` is **instagram**, this is an Instagram-Scoped User ID (IGSID) generated when a user interacts with your Instagram business or creator account.
	//
	// <props="intl">
	//
	// - If `ChannelType` is **viber**, this is the recipient\\"s phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861388988****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The token type.
	//
	// > This parameter is for internal testing, is not generally available, and can be ignored.
	//
	// example:
	//
	// bearer
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
	// Custom tracking data for a Viber message. This parameter is available only on the International Site and can be ignored if you are using the China site.
	//
	// example:
	//
	// Tracking Data
	TrackingData *string `json:"TrackingData,omitempty" xml:"TrackingData,omitempty"`
	// The time-to-live (TTL) for a Viber message. This parameter is available only on the International Site and can be ignored if you are using the China site.<props="intl"> Unit: seconds. The value must be between 30 and 1209600.
	//
	// example:
	//
	// 50
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The message type. Valid values:
	//
	// - `template`: A message template approved in the console. You can send this type of message at any time.
	//
	// - `message`: A message of any format. You can send this type of message only within 24 hours of receiving the last message from a user.
	//
	// 	Notice:
	//
	// If you set `Type` to `template`, you must set the `TemplateCode` parameter. If you set `Type` to `message`, you must set the `MessageType` parameter.
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
	// A collection of default flow parameters.
	FlowActionData map[string]interface{} `json:"FlowActionData,omitempty" xml:"FlowActionData,omitempty"`
	// The custom flow token.
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
	// A list of product categories. You can specify up to 10 categories and a total of 30 products.
	Sections []*SendChatappMessageRequestProductActionSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// The product catalog ID. You can get this ID by calling the [ListProductCatalog](https://help.aliyun.com/document_detail/2539783.html) operation.
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
	// A list of product information.
	ProductItems []*SendChatappMessageRequestProductActionSectionsProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Repeated"`
	// The category name. You can find the name on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Catalog Management*	- > **Product Management*	- page, or obtain it by calling the [ListProduct](https://help.aliyun.com/document_detail/2557786.html) operation.
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
	// The product ID. You can find the ID on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Catalog Management*	- > **Product Management*	- page, or obtain it by calling the [ListProduct](https://help.aliyun.com/document_detail/2557786.html) operation.
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
