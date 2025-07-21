// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMassMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *SendChatappMassMessageRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *SendChatappMassMessageRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *SendChatappMassMessageRequest
	GetCustWabaId() *string
	SetFallBackContent(v string) *SendChatappMassMessageRequest
	GetFallBackContent() *string
	SetFallBackDuration(v int32) *SendChatappMassMessageRequest
	GetFallBackDuration() *int32
	SetFallBackId(v string) *SendChatappMassMessageRequest
	GetFallBackId() *string
	SetFallBackRule(v string) *SendChatappMassMessageRequest
	GetFallBackRule() *string
	SetFrom(v string) *SendChatappMassMessageRequest
	GetFrom() *string
	SetIsvCode(v string) *SendChatappMassMessageRequest
	GetIsvCode() *string
	SetLabel(v string) *SendChatappMassMessageRequest
	GetLabel() *string
	SetLanguage(v string) *SendChatappMassMessageRequest
	GetLanguage() *string
	SetSenderList(v []*SendChatappMassMessageRequestSenderList) *SendChatappMassMessageRequest
	GetSenderList() []*SendChatappMassMessageRequestSenderList
	SetTag(v string) *SendChatappMassMessageRequest
	GetTag() *string
	SetTaskId(v string) *SendChatappMassMessageRequest
	GetTaskId() *string
	SetTemplateCode(v string) *SendChatappMassMessageRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *SendChatappMassMessageRequest
	GetTemplateName() *string
	SetTtl(v int64) *SendChatappMassMessageRequest
	GetTtl() *int64
}

type SendChatappMassMessageRequest struct {
	// The type of the channel. Valid values:
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
	// Fallback message
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// Specifies the period of time after which the fallback message is sent if the message receipt that indicates the message is delivered to clients is not received. If this parameter is left empty, the fallback message is sent only when the message fails to be sent or the message receipt that indicates the message is not delivered to clients is received. Unit: seconds. Valid values: 60 to 43200.
	//
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The ID of the fallback policy.
	//
	// example:
	//
	// S00001
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
	// The mobile phone number of the message sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861387777****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ISV verification code. This parameter is used to verify whether the RAM user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The type of the Viber message. Valid values:
	//
	// 	- **promotion**
	//
	// 	- **transaction**
	//
	// example:
	//
	// promotion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language. For more information about language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The mobile phone numbers of the message receivers.
	//
	// This parameter is required.
	SenderList []*SendChatappMassMessageRequestSenderList `json:"SenderList,omitempty" xml:"SenderList,omitempty" type:"Repeated"`
	// The tag information when the ChannelType parameter is set to viber.
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
	// The template code.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca36bf5
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The timeout period for sending messages when the ChannelType parameter is set to viber. Valid values: 30 to 1209600. Unit: seconds.
	//
	// example:
	//
	// 50
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s SendChatappMassMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *SendChatappMassMessageRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *SendChatappMassMessageRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *SendChatappMassMessageRequest) GetFallBackContent() *string {
	return s.FallBackContent
}

func (s *SendChatappMassMessageRequest) GetFallBackDuration() *int32 {
	return s.FallBackDuration
}

func (s *SendChatappMassMessageRequest) GetFallBackId() *string {
	return s.FallBackId
}

func (s *SendChatappMassMessageRequest) GetFallBackRule() *string {
	return s.FallBackRule
}

func (s *SendChatappMassMessageRequest) GetFrom() *string {
	return s.From
}

func (s *SendChatappMassMessageRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *SendChatappMassMessageRequest) GetLabel() *string {
	return s.Label
}

func (s *SendChatappMassMessageRequest) GetLanguage() *string {
	return s.Language
}

func (s *SendChatappMassMessageRequest) GetSenderList() []*SendChatappMassMessageRequestSenderList {
	return s.SenderList
}

func (s *SendChatappMassMessageRequest) GetTag() *string {
	return s.Tag
}

func (s *SendChatappMassMessageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SendChatappMassMessageRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendChatappMassMessageRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SendChatappMassMessageRequest) GetTtl() *int64 {
	return s.Ttl
}

func (s *SendChatappMassMessageRequest) SetChannelType(v string) *SendChatappMassMessageRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetCustSpaceId(v string) *SendChatappMassMessageRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetCustWabaId(v string) *SendChatappMassMessageRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFallBackContent(v string) *SendChatappMassMessageRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFallBackDuration(v int32) *SendChatappMassMessageRequest {
	s.FallBackDuration = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFallBackId(v string) *SendChatappMassMessageRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFallBackRule(v string) *SendChatappMassMessageRequest {
	s.FallBackRule = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFrom(v string) *SendChatappMassMessageRequest {
	s.From = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetIsvCode(v string) *SendChatappMassMessageRequest {
	s.IsvCode = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetLabel(v string) *SendChatappMassMessageRequest {
	s.Label = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetLanguage(v string) *SendChatappMassMessageRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetSenderList(v []*SendChatappMassMessageRequestSenderList) *SendChatappMassMessageRequest {
	s.SenderList = v
	return s
}

func (s *SendChatappMassMessageRequest) SetTag(v string) *SendChatappMassMessageRequest {
	s.Tag = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetTaskId(v string) *SendChatappMassMessageRequest {
	s.TaskId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetTemplateCode(v string) *SendChatappMassMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetTemplateName(v string) *SendChatappMassMessageRequest {
	s.TemplateName = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetTtl(v int64) *SendChatappMassMessageRequest {
	s.Ttl = &v
	return s
}

func (s *SendChatappMassMessageRequest) Validate() error {
	return dara.Validate(s)
}

type SendChatappMassMessageRequestSenderList struct {
	// The Flow action.
	FlowAction *SendChatappMassMessageRequestSenderListFlowAction `json:"FlowAction,omitempty" xml:"FlowAction,omitempty" type:"Struct"`
	// The payload of the button.
	Payload []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	// The information about the product.
	ProductAction *SendChatappMassMessageRequestSenderListProductAction `json:"ProductAction,omitempty" xml:"ProductAction,omitempty" type:"Struct"`
	// The parameters of the template.
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The mobile phone number of the message receiver.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861388988****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendChatappMassMessageRequestSenderList) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderList) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderList) GetFlowAction() *SendChatappMassMessageRequestSenderListFlowAction {
	return s.FlowAction
}

func (s *SendChatappMassMessageRequestSenderList) GetPayload() []*string {
	return s.Payload
}

func (s *SendChatappMassMessageRequestSenderList) GetProductAction() *SendChatappMassMessageRequestSenderListProductAction {
	return s.ProductAction
}

func (s *SendChatappMassMessageRequestSenderList) GetTemplateParams() map[string]*string {
	return s.TemplateParams
}

func (s *SendChatappMassMessageRequestSenderList) GetTo() *string {
	return s.To
}

func (s *SendChatappMassMessageRequestSenderList) SetFlowAction(v *SendChatappMassMessageRequestSenderListFlowAction) *SendChatappMassMessageRequestSenderList {
	s.FlowAction = v
	return s
}

func (s *SendChatappMassMessageRequestSenderList) SetPayload(v []*string) *SendChatappMassMessageRequestSenderList {
	s.Payload = v
	return s
}

func (s *SendChatappMassMessageRequestSenderList) SetProductAction(v *SendChatappMassMessageRequestSenderListProductAction) *SendChatappMassMessageRequestSenderList {
	s.ProductAction = v
	return s
}

func (s *SendChatappMassMessageRequestSenderList) SetTemplateParams(v map[string]*string) *SendChatappMassMessageRequestSenderList {
	s.TemplateParams = v
	return s
}

func (s *SendChatappMassMessageRequestSenderList) SetTo(v string) *SendChatappMassMessageRequestSenderList {
	s.To = &v
	return s
}

func (s *SendChatappMassMessageRequestSenderList) Validate() error {
	return dara.Validate(s)
}

type SendChatappMassMessageRequestSenderListFlowAction struct {
	// The default parameter of the Flow.
	FlowActionData map[string]interface{} `json:"FlowActionData,omitempty" xml:"FlowActionData,omitempty"`
	// The information about the Flow token.
	//
	// example:
	//
	// kde****
	FlowToken *string `json:"FlowToken,omitempty" xml:"FlowToken,omitempty"`
}

func (s SendChatappMassMessageRequestSenderListFlowAction) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderListFlowAction) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderListFlowAction) GetFlowActionData() map[string]interface{} {
	return s.FlowActionData
}

func (s *SendChatappMassMessageRequestSenderListFlowAction) GetFlowToken() *string {
	return s.FlowToken
}

func (s *SendChatappMassMessageRequestSenderListFlowAction) SetFlowActionData(v map[string]interface{}) *SendChatappMassMessageRequestSenderListFlowAction {
	s.FlowActionData = v
	return s
}

func (s *SendChatappMassMessageRequestSenderListFlowAction) SetFlowToken(v string) *SendChatappMassMessageRequestSenderListFlowAction {
	s.FlowToken = &v
	return s
}

func (s *SendChatappMassMessageRequestSenderListFlowAction) Validate() error {
	return dara.Validate(s)
}

type SendChatappMassMessageRequestSenderListProductAction struct {
	// The products. Up to 30 products and 10 categories can be added.
	Sections []*SendChatappMassMessageRequestSenderListProductActionSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// The retailer ID of the product.
	//
	// example:
	//
	// skkks999393
	ThumbnailProductRetailerId *string `json:"ThumbnailProductRetailerId,omitempty" xml:"ThumbnailProductRetailerId,omitempty"`
}

func (s SendChatappMassMessageRequestSenderListProductAction) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderListProductAction) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderListProductAction) GetSections() []*SendChatappMassMessageRequestSenderListProductActionSections {
	return s.Sections
}

func (s *SendChatappMassMessageRequestSenderListProductAction) GetThumbnailProductRetailerId() *string {
	return s.ThumbnailProductRetailerId
}

func (s *SendChatappMassMessageRequestSenderListProductAction) SetSections(v []*SendChatappMassMessageRequestSenderListProductActionSections) *SendChatappMassMessageRequestSenderListProductAction {
	s.Sections = v
	return s
}

func (s *SendChatappMassMessageRequestSenderListProductAction) SetThumbnailProductRetailerId(v string) *SendChatappMassMessageRequestSenderListProductAction {
	s.ThumbnailProductRetailerId = &v
	return s
}

func (s *SendChatappMassMessageRequestSenderListProductAction) Validate() error {
	return dara.Validate(s)
}

type SendChatappMassMessageRequestSenderListProductActionSections struct {
	// The products.
	ProductItems []*SendChatappMassMessageRequestSenderListProductActionSectionsProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Repeated"`
	// The name of the category.
	//
	// example:
	//
	// abcd
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SendChatappMassMessageRequestSenderListProductActionSections) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderListProductActionSections) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderListProductActionSections) GetProductItems() []*SendChatappMassMessageRequestSenderListProductActionSectionsProductItems {
	return s.ProductItems
}

func (s *SendChatappMassMessageRequestSenderListProductActionSections) GetTitle() *string {
	return s.Title
}

func (s *SendChatappMassMessageRequestSenderListProductActionSections) SetProductItems(v []*SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) *SendChatappMassMessageRequestSenderListProductActionSections {
	s.ProductItems = v
	return s
}

func (s *SendChatappMassMessageRequestSenderListProductActionSections) SetTitle(v string) *SendChatappMassMessageRequestSenderListProductActionSections {
	s.Title = &v
	return s
}

func (s *SendChatappMassMessageRequestSenderListProductActionSections) Validate() error {
	return dara.Validate(s)
}

type SendChatappMassMessageRequestSenderListProductActionSectionsProductItems struct {
	// The retailer ID of the product.
	//
	// example:
	//
	// ksi399d8
	ProductRetailerId *string `json:"ProductRetailerId,omitempty" xml:"ProductRetailerId,omitempty"`
}

func (s SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) GetProductRetailerId() *string {
	return s.ProductRetailerId
}

func (s *SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) SetProductRetailerId(v string) *SendChatappMassMessageRequestSenderListProductActionSectionsProductItems {
	s.ProductRetailerId = &v
	return s
}

func (s *SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) Validate() error {
	return dara.Validate(s)
}
