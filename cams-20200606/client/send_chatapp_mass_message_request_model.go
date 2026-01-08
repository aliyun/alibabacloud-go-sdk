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
	SetOwnerId(v int64) *SendChatappMassMessageRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SendChatappMassMessageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendChatappMassMessageRequest
	GetResourceOwnerId() *int64
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
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 示例值示例值示例值
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FallBackContent  *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	FallBackDuration *int32  `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
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
	// 示例值示例值
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// example:
	//
	// 示例值示例值
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Language             *string                                    `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId              *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                    `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                     `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SenderList           []*SendChatappMassMessageRequestSenderList `json:"SenderList,omitempty" xml:"SenderList,omitempty" type:"Repeated"`
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
	// 示例值示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 46
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

func (s *SendChatappMassMessageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendChatappMassMessageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendChatappMassMessageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *SendChatappMassMessageRequest) SetOwnerId(v int64) *SendChatappMassMessageRequest {
	s.OwnerId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetResourceOwnerAccount(v string) *SendChatappMassMessageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetResourceOwnerId(v int64) *SendChatappMassMessageRequest {
	s.ResourceOwnerId = &v
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
	if s.SenderList != nil {
		for _, item := range s.SenderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendChatappMassMessageRequestSenderList struct {
	FlowAction     *SendChatappMassMessageRequestSenderListFlowAction    `json:"FlowAction,omitempty" xml:"FlowAction,omitempty" type:"Struct"`
	Payload        []*string                                             `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	ProductAction  *SendChatappMassMessageRequestSenderListProductAction `json:"ProductAction,omitempty" xml:"ProductAction,omitempty" type:"Struct"`
	TemplateParams map[string]*string                                    `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// example:
	//
	// 示例值示例值示例值
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

type SendChatappMassMessageRequestSenderListFlowAction struct {
	FlowActionData map[string]interface{} `json:"FlowActionData,omitempty" xml:"FlowActionData,omitempty"`
	// example:
	//
	// 示例值
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
	Sections []*SendChatappMassMessageRequestSenderListProductActionSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值
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

type SendChatappMassMessageRequestSenderListProductActionSections struct {
	ProductItems []*SendChatappMassMessageRequestSenderListProductActionSectionsProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值
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

type SendChatappMassMessageRequestSenderListProductActionSectionsProductItems struct {
	// example:
	//
	// 示例值示例值
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
