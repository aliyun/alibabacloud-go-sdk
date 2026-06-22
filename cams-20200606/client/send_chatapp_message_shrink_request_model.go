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
	FlowActionShrink *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
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
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// Product information that you have uploaded to Meta. This parameter applies to WhatsApp channels only.
	ProductActionShrink *string `json:"ProductAction,omitempty" xml:"ProductAction,omitempty"`
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
	TemplateParamsShrink *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
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
