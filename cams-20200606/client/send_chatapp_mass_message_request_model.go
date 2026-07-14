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
	// The channel type. Valid values:
	//
	// - **whatsapp**
	//
	// - **messenger**
	//
	// - **instagram**
	//
	// <props="intl">- **viber**
	//
	// This parameter is required.
	//
	// example:
	//
	// whatsapp
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The ISV sub-customer SpaceId or direct customer instance ID. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) page.
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
	// The custom fallback content. This parameter is for the China site (Chinese). China site users can ignore this parameter.
	//
	// example:
	//
	// Fallback SMS
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// The fallback trigger time. This parameter is for the international site. China site users can ignore this parameter. <props="intl">If no delivery receipt is returned within the specified time, the fallback is triggered. If this parameter is not specified, the fallback is triggered only when the message fails to send or a failure status report is received. Unit: seconds. Minimum value: 60. Maximum value: 43200.
	//
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The fallback policy ID. This parameter is for the China site (Chinese). China site users can ignore this parameter. <props="intl">You can view the policy ID on the [**Fallback Policy**](https://chatapp.console.alibabacloud.com/FallbackStrategy) page.
	//
	// example:
	//
	// S0****
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The fallback rule. This parameter is for the international site. China site users can ignore this parameter.
	//
	// <props="intl">Valid values:
	//
	// <props="intl">- **undelivered**: the fallback is triggered when the message cannot be delivered to the device. During sending, the template and parameters must pass validation. Blocked templates or blocked numbers are not validated. This rule is used by default if the parameter value is empty.
	//
	// <props="intl">- **sentFailed**: the fallback is triggered when validation of the template or template variables fails. Only channelType, type, messageType, to, and from (whether it exists) are strictly validated.
	//
	// example:
	//
	// undelivered
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// The sender phone number.
	//
	// - If ChannelType is **whatsapp**, this is the phone number registered and bindng with WhatsApp. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **WABA Management*	- > **Phone Number Management*	- page.
	//
	// - If ChannelType is **messenger**, this is the Page ID. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Public Page*	- page.
	//
	// - If ChannelType is **instagram**, this is the Instagram professional Account ID. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Professional Account*	- page.
	//
	// <props="intl">- If ChannelType is **viber**, this is the Viber Service ID. You can view it on the [**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Service ID Management*	- page.
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
	// skdi3kksloslikd****
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
	// This parameter is required.
	//
	// example:
	//
	// en
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of recipient phone numbers.
	SenderList []*SendChatappMassMessageRequestSenderList `json:"SenderList,omitempty" xml:"SenderList,omitempty" type:"Repeated"`
	// The tag information. Custom tag information for Viber message sending.
	//
	// example:
	//
	// Tag
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
	// The timeout period for Viber message sending. This parameter is for the international site. China site users can ignore this parameter. <props="intl">Unit: seconds. Valid values: 30 to 1209600.
	//
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
	// The Flow message object.
	FlowAction *SendChatappMassMessageRequestSenderListFlowAction `json:"FlowAction,omitempty" xml:"FlowAction,omitempty" type:"Struct"`
	// The list of button trigger message identifiers.
	Payload []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	// The product information. This parameter applies only to WhatsApp channels and refers to the product information you uploaded on Meta.
	ProductAction *SendChatappMassMessageRequestSenderListProductAction `json:"ProductAction,omitempty" xml:"ProductAction,omitempty" type:"Struct"`
	// example:
	//
	// individual
	RecipientType *string `json:"RecipientType,omitempty" xml:"RecipientType,omitempty"`
	// The collection of template parameters.
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The recipient phone number.
	//
	// - If ChannelType is **whatsapp**, this is the phone number of the message recipient.
	//
	// - If ChannelType is **messenger**, this is the Page-Scoped User ID generated when the user interacts with the Facebook page.
	//
	// - If ChannelType is **instagram**, this is the Instagram User ID generated when the user interacts with the Instagram business or creator account.
	//
	// <props="intl">- If ChannelType is **viber**, this is the phone number of the message recipient.
	//
	// example:
	//
	// 861386666****
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

func (s *SendChatappMassMessageRequestSenderList) GetRecipientType() *string {
	return s.RecipientType
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

func (s *SendChatappMassMessageRequestSenderList) SetRecipientType(v string) *SendChatappMassMessageRequestSenderList {
	s.RecipientType = &v
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
	// The collection of flow default parameters.
	FlowActionData map[string]interface{} `json:"FlowActionData,omitempty" xml:"FlowActionData,omitempty"`
	// The custom flow token information.
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
	// The list of product categories. A maximum of 10 categories and 30 products are supported.
	Sections []*SendChatappMassMessageRequestSenderListProductActionSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// The product catalog ID. You can obtain this ID by calling the [ListProductCatalog](https://help.aliyun.com/document_detail/2539783.html) operation.
	//
	// example:
	//
	// skkks99****
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
	// The list of product information.
	ProductItems []*SendChatappMassMessageRequestSenderListProductActionSectionsProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Repeated"`
	// The category name. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Catalog Management*	- > **Product Management*	- page, or obtain it by calling the [ListProduct](https://help.aliyun.com/document_detail/2557786.html) operation.
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
	// The product ID. You can view it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Catalog Management*	- > **Product Management*	- page, or obtain it by calling the [ListProduct](https://help.aliyun.com/document_detail/2557786.html) operation.
	//
	// example:
	//
	// ksi3****
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
