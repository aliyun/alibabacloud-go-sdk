// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMessageRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetMessageType(v string) *SendChatappMessageRequest
	GetMessageType() *string
	SetPayload(v []*string) *SendChatappMessageRequest
	GetPayload() []*string
	SetProductAction(v *SendChatappMessageRequestProductAction) *SendChatappMessageRequest
	GetProductAction() *SendChatappMessageRequestProductAction
	SetRecipientType(v string) *SendChatappMessageRequest
	GetRecipientType() *string
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
	SetTrackingData(v string) *SendChatappMessageRequest
	GetTrackingData() *string
	SetTtl(v int32) *SendChatappMessageRequest
	GetTtl() *int32
	SetType(v string) *SendChatappMessageRequest
	GetType() *string
}

type SendChatappMessageRequest struct {
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
	FlowAction *SendChatappMessageRequestFlowAction `json:"FlowAction,omitempty" xml:"FlowAction,omitempty" type:"Struct"`
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
	Payload []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	// The information about the products included in the WhatsApp catalog message or multi-product message (MPM).
	ProductAction *SendChatappMessageRequestProductAction `json:"ProductAction,omitempty" xml:"ProductAction,omitempty" type:"Struct"`
	RecipientType *string                                 `json:"RecipientType,omitempty" xml:"RecipientType,omitempty"`
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
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
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

func (s SendChatappMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageRequest) GoString() string {
	return s.String()
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

func (s *SendChatappMessageRequest) GetMessageType() *string {
	return s.MessageType
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

func (s *SendChatappMessageRequest) GetTrackingData() *string {
	return s.TrackingData
}

func (s *SendChatappMessageRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *SendChatappMessageRequest) GetType() *string {
	return s.Type
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

func (s *SendChatappMessageRequest) SetMessageType(v string) *SendChatappMessageRequest {
	s.MessageType = &v
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
	return dara.Validate(s)
}

type SendChatappMessageRequestFlowAction struct {
	// The default parameter of the Flow.
	FlowActionData map[string]interface{} `json:"FlowActionData,omitempty" xml:"FlowActionData,omitempty"`
	// The Flow token.
	//
	// example:
	//
	// 1122***
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
	// The products. Up to 30 products and 10 categories can be added.
	Sections []*SendChatappMessageRequestProductActionSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// The retailer ID of the product.
	//
	// example:
	//
	// S238SK
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
	return dara.Validate(s)
}

type SendChatappMessageRequestProductActionSections struct {
	// The products.
	ProductItems []*SendChatappMessageRequestProductActionSectionsProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Repeated"`
	// The name of the category.
	//
	// example:
	//
	// Test
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
	return dara.Validate(s)
}

type SendChatappMessageRequestProductActionSectionsProductItems struct {
	// The retailer ID of the product.
	//
	// example:
	//
	// 9I39E9E
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
