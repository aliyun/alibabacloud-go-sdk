// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetMessageType(v string) *SendChatappMessageShrinkRequest
	GetMessageType() *string
	SetPayloadShrink(v string) *SendChatappMessageShrinkRequest
	GetPayloadShrink() *string
	SetProductActionShrink(v string) *SendChatappMessageShrinkRequest
	GetProductActionShrink() *string
	SetRecipientType(v string) *SendChatappMessageShrinkRequest
	GetRecipientType() *string
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
	SetTrackingData(v string) *SendChatappMessageShrinkRequest
	GetTrackingData() *string
	SetTtl(v int32) *SendChatappMessageShrinkRequest
	GetTtl() *int32
	SetType(v string) *SendChatappMessageShrinkRequest
	GetType() *string
}

type SendChatappMessageShrinkRequest struct {
	// The channel type. Valid values:
	//
	// 	- **whatsapp**
	//
	// 	- **viber**
	//
	// 	- **line*	- (under development)
	//
	// This parameter is required.
	//
	// example:
	//
	// whatsapp
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The message content.
	//
	// **Notes on WhatsApp messages:**
	//
	// 	- If you set **messageType*	- to **text**, you must specify **text*	- and must not specify **Caption**.
	//
	// 	- If you set **messageType*	- to **image**, you must specify **Link**.
	//
	// 	- If you set **messageType*	- to **video**, you must specify **Link**.
	//
	// 	- If you set **messageType*	- to **audio**, **Link*	- is required and **Caption*	- is invalid.
	//
	// 	- If you set **messageType*	- to **document**, **Link*	- and **FileName*	- are required and **Caption*	- is invalid.
	//
	// 	- If you set **messageType*	- to **interactive**, you must specify **type*	- and **action**.
	//
	// 	- If you set **messageType*	- to **contacts**, you must specify **name**.
	//
	// 	- If you set **messageType*	- to **location**, you must specify **longitude*	- and **latitude**.
	//
	// 	- If you set **messageType*	- to **sticker**, you must specify **Link**, and **Caption*	- and **FileName*	- are invalid.
	//
	// 	- If you set **messageType*	- to **reaction**, you must specify **messageId*	- and **emoji**.
	//
	// **Notes on Viber messages:**
	//
	// 	- If you set **messageType*	- to **text**, you must specify **text**.
	//
	// 	- If you set **messageType*	- to **image**, you must specify **link**.
	//
	// 	- If you set **messageType*	- to **video**, you must specify **link**, **thumbnail**, **fileSize**, and **duration**.
	//
	// 	- If you set **messageType*	- to **document**, you must specify **link**, **fileName**, and **fileType**.
	//
	// 	- If you set **messageType*	- to **text_button**, you must specify **text**, **caption**, and **action**.
	//
	// 	- If you set **messageType*	- to **text_image_button**, you must specify **text**, **link**, **caption**, and **action**.
	//
	// 	- If you set **messageType*	- to **text_video**, you must specify **text**, **link**, **thumbnail**, **fileSize**, and **duration**.
	//
	// 	- If you set **messageType*	- to **text_video_button**, you must specify **text**, **link**, **thumbnail**, **fileSize**, **duration**, and **caption**. In addition, you must not specify **action**.
	//
	// example:
	//
	// {\\"text\\": \\"hello whatsapp\\", \\"link\\": \\"\\", \\"caption\\": \\"\\", \\"fileName\\": \\"\\" }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the reply message.
	//
	// example:
	//
	// 61851ccb2f1365b16aee****
	ContextMessageId *string `json:"ContextMessageId,omitempty" xml:"ContextMessageId,omitempty"`
	// The space ID of the user.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business Account (WABA) ID of the RAM user within the independent software vendor (ISV) account.
	//
	// >  CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The content of the fallback message.
	//
	// example:
	//
	// This is a fallback message.
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// Specifies the period of time after which the fallback message is sent if the message receipt that indicates the message is delivered to clients is not received. If this parameter is left empty, the fallback message is sent only when the **message fails to be sent*	- or **the message receipt that indicates the message is not delivered to clients*	- is received. Unit: seconds. Valid values: 60 to 43200.
	//
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The ID of the fallback policy. You can create a fallback policy and view the information in the Chat App Message Service console.
	//
	// example:
	//
	// S_000001
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The fallback rule. Valid values:
	//
	// 	- **undelivered**: A fallback is triggered if the message is not delivered to clients. When the message is being sent, the template parameters are verified. If the parameters fail to pass the verification, the message fails to be sent. Whether the template and phone number are prohibited is not verified. By default, this value is used when FallBackRule is left empty.
	//
	// 	- **sentFailed**: A fallback is triggered even if the template parameters including variables fail to pass the verification. If the channelType, type, messageType, to, and from parameters fail to pass the verification, a fallback is not triggered.
	//
	// example:
	//
	// undelivered
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// The Flow action.
	FlowActionShrink *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// The mobile phone number of the message sender.
	//
	// >  You can specify a mobile phone number that is registered for a WhatsApp account and is approved in the Chat App Message Service console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1360000****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ISV verification code. This parameter is used to verify whether the RAM user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The type of the Viber message. This parameter is required if ChannelType is set to viber. Valid values:
	//
	// 	- **promotion**
	//
	// 	- **transaction**
	//
	// example:
	//
	// promotion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language that is used in the message template. This parameter is required only if you set the Type parameter to **template**. For more information about language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The specific type of the message. This parameter is required only if you set the Type parameter to **message**.
	//
	// **Valid values of MessageType when you set the ChannelType parameter to whatsapp:**
	//
	// 	- **text**: a text message.
	//
	// 	- **image**: an image message.
	//
	// 	- **video**: a video message.
	//
	// 	- **audio**: an audio message.
	//
	// 	- **document**: a document message.
	//
	// 	- **interactive**: an interactive message.
	//
	// 	- **contacts**: a contact message.
	//
	// 	- **location**: a location message.
	//
	// 	- **sticker**: a sticker message.
	//
	// 	- **reaction**: a reaction message.
	//
	// **Valid values of MessageType when you set the ChannelType parameter to viber:**
	//
	// 	- **text**: a text message.
	//
	// 	- **image**: an image message.
	//
	// 	- **video**: a video message.
	//
	// 	- **document**: a document message.
	//
	// 	- **text_button**: a message that contains the text and button media objects.
	//
	// 	- **text_image_button**: a message that contains multiple media objects, including the text, image, and button.
	//
	// 	- **text_video**: a message that contains the text and video media objects.
	//
	// 	- **text_video_button**: a message that contains multiple media objects, including text, video, and button.
	//
	// 	- **text_image**: a message that contains the text and image media objects.
	//
	// > For more information, see [Parameters of a message template](https://help.aliyun.com/document_detail/454530.html).
	//
	// example:
	//
	// text
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The payload of the button.
	//
	// example:
	//
	// payloadtext1,payloadtext2,payloadtext3
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The information about the products included in the WhatsApp catalog message or multi-product message (MPM).
	ProductActionShrink *string `json:"ProductAction,omitempty" xml:"ProductAction,omitempty"`
	RecipientType       *string `json:"RecipientType,omitempty" xml:"RecipientType,omitempty"`
	// The tag information of the Viber message.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 100000001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The code of the message template. This parameter is required only if you set the Type parameter to **template**.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The variables of the message template.
	TemplateParamsShrink *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The mobile phone number of the message receiver.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The tracking data of the Viber message.
	//
	// example:
	//
	// tracking_id:123456
	TrackingData *string `json:"TrackingData,omitempty" xml:"TrackingData,omitempty"`
	// The timeout period for sending the Viber message. Valid values: 30 to 1209600. Unit: seconds.
	//
	// example:
	//
	// 50
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The message type. Valid values:
	//
	// 	- **template**: the template message. A template message is sent based on a template that is created and approved in the Chat App Message Service console. You can send template messages based on your business requirements.
	//
	// 	- **message**: the custom message. You can send a custom WhatsApp message to a user only within 24 hours after you receive the last message from the user. This limit does not apply to custom Viber messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// template
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendChatappMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageShrinkRequest) GoString() string {
	return s.String()
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

func (s *SendChatappMessageShrinkRequest) GetMessageType() *string {
	return s.MessageType
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

func (s *SendChatappMessageShrinkRequest) GetTrackingData() *string {
	return s.TrackingData
}

func (s *SendChatappMessageShrinkRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *SendChatappMessageShrinkRequest) GetType() *string {
	return s.Type
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

func (s *SendChatappMessageShrinkRequest) SetMessageType(v string) *SendChatappMessageShrinkRequest {
	s.MessageType = &v
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
