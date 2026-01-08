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
	// example:
	//
	// 示例值示例值
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
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
	// example:
	//
	// 示例值
	ContextMessageId *string `json:"ContextMessageId,omitempty" xml:"ContextMessageId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 示例值示例值
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// example:
	//
	// 示例值
	FallBackContent  *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	FallBackDuration *int32  `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// example:
	//
	// 示例值示例值
	FallBackRule     *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	FlowActionShrink *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 示例值
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// example:
	//
	// 示例值
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	MessageCampaignId *string `json:"MessageCampaignId,omitempty" xml:"MessageCampaignId,omitempty"`
	// example:
	//
	// 示例值
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The payload of the button.
	//
	// example:
	//
	// payloadtext1,payloadtext2,payloadtext3
	PayloadShrink       *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	ProductActionShrink *string `json:"ProductAction,omitempty" xml:"ProductAction,omitempty"`
	// example:
	//
	// individual
	RecipientType        *string `json:"RecipientType,omitempty" xml:"RecipientType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 示例值示例值
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateParamsShrink *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// 示例值
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
	// example:
	//
	// 示例值示例值
	TrackingData *string `json:"TrackingData,omitempty" xml:"TrackingData,omitempty"`
	Ttl          *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
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
