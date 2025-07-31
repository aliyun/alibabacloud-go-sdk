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
	SetTrackingData(v string) *SendChatappMessageRequest
	GetTrackingData() *string
	SetTtl(v int32) *SendChatappMessageRequest
	GetTtl() *int32
	SetType(v string) *SendChatappMessageRequest
	GetType() *string
}

type SendChatappMessageRequest struct {
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
	FallBackRule *string                              `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	FlowAction   *SendChatappMessageRequestFlowAction `json:"FlowAction,omitempty" xml:"FlowAction,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	From *string `json:"From,omitempty" xml:"From,omitempty"`
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
	// 示例值
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The payload of the button.
	//
	// example:
	//
	// payloadtext1,payloadtext2,payloadtext3
	Payload       []*string                               `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	ProductAction *SendChatappMessageRequestProductAction `json:"ProductAction,omitempty" xml:"ProductAction,omitempty" type:"Struct"`
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
	TemplateName   *string            `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	To *string `json:"To,omitempty" xml:"To,omitempty"`
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
	FlowActionData map[string]interface{} `json:"FlowActionData,omitempty" xml:"FlowActionData,omitempty"`
	// example:
	//
	// 示例值示例值
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
	Sections []*SendChatappMessageRequestProductActionSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值示例值
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
	ProductItems []*SendChatappMessageRequestProductActionSectionsProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值
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
	// example:
	//
	// 示例值示例值
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
