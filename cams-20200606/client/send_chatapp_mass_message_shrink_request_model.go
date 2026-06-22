// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMassMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *SendChatappMassMessageShrinkRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *SendChatappMassMessageShrinkRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *SendChatappMassMessageShrinkRequest
	GetCustWabaId() *string
	SetFallBackContent(v string) *SendChatappMassMessageShrinkRequest
	GetFallBackContent() *string
	SetFallBackDuration(v int32) *SendChatappMassMessageShrinkRequest
	GetFallBackDuration() *int32
	SetFallBackId(v string) *SendChatappMassMessageShrinkRequest
	GetFallBackId() *string
	SetFallBackRule(v string) *SendChatappMassMessageShrinkRequest
	GetFallBackRule() *string
	SetFrom(v string) *SendChatappMassMessageShrinkRequest
	GetFrom() *string
	SetIsvCode(v string) *SendChatappMassMessageShrinkRequest
	GetIsvCode() *string
	SetLabel(v string) *SendChatappMassMessageShrinkRequest
	GetLabel() *string
	SetLanguage(v string) *SendChatappMassMessageShrinkRequest
	GetLanguage() *string
	SetOwnerId(v int64) *SendChatappMassMessageShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SendChatappMassMessageShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendChatappMassMessageShrinkRequest
	GetResourceOwnerId() *int64
	SetSenderListShrink(v string) *SendChatappMassMessageShrinkRequest
	GetSenderListShrink() *string
	SetTag(v string) *SendChatappMassMessageShrinkRequest
	GetTag() *string
	SetTaskId(v string) *SendChatappMassMessageShrinkRequest
	GetTaskId() *string
	SetTemplateCode(v string) *SendChatappMassMessageShrinkRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *SendChatappMassMessageShrinkRequest
	GetTemplateName() *string
	SetTtl(v int64) *SendChatappMassMessageShrinkRequest
	GetTtl() *int64
}

type SendChatappMassMessageShrinkRequest struct {
	// The channel type. Valid values:
	//
	// - **whatsapp**
	//
	// - **messenger**
	//
	// - **instagram**
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
	// The Space ID of the ISV sub-customer, or the instance ID for a direct customer. View it on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// example:
	//
	// cams-8c8*********
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business Account (WABA) ID of the Independent Software Vendor (ISV) customer. This is a deprecated parameter. Use CustSpaceId instead, which is the direct customer\\"s instance ID. View the ID on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// example:
	//
	// cams-8c8*********
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The custom fallback content. This parameter is for the international site (alibabacloud.com). You can ignore it for the China site (aliyun.com).
	//
	// example:
	//
	// Fallback SMS
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// The time to trigger a fallback. This parameter is for the international site (alibabacloud.com). You can ignore it for the China site (aliyun.com).<props="intl">If a delivery receipt is not returned within the specified time, a fallback is triggered. If you leave this parameter empty, the fallback is not determined by time. A fallback is triggered only when the message fails to be sent or a failed status report is received. Unit: seconds. Minimum value: 60. Maximum value: 43200.
	//
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The fallback policy ID. This parameter is for the international site (alibabacloud.com). You can ignore it for the China site (aliyun.com).<props="intl">View the policy ID on the [**Fallback Policy**](https://chatapp.console.alibabacloud.com/FallbackStrategy) page.
	//
	// example:
	//
	// S0****
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The fallback rule. This parameter is for the international site (alibabacloud.com). You can ignore it for the China site (aliyun.com). <props="intl">Valid values:
	//
	// <props="intl">
	//
	// - **undelivered**: A fallback is triggered if the message cannot be delivered to the recipient. This rule requires that the template and parameters pass verification before sending. The rule does not apply if the message is blocked from sending, for example, due to a blacklisted template or phone number. This is the default rule if this parameter is empty.
	//
	//
	//
	// <props="intl">
	//
	// - **sentFailed**: A fallback is triggered if the template, template variables, or other parameters fail verification. Only the channelType, type, messageType, to, and the existence of the from parameter are strictly verified.
	//
	// example:
	//
	// undelivered
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// The sender\\"s number.
	//
	// - If ChannelType is **whatsapp**, this is the phone number registered and bound with WhatsApp. View the number on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **WABA Management*	- > **Phone Number Management*	- page.
	//
	// - If ChannelType is **messenger**, this is the Page ID. View the ID on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Facebook Homepage*	- page.
	//
	// - If ChannelType is **instagram**, this is the Instagram professional account ID. View the ID on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Professional Account*	- page.
	//
	// <props="intl">
	//
	// - If ChannelType is **viber**, this is the Viber service ID. View the ID on the [**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Service Number Management&#x20;**&#x70;age.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861387777****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Deprecated
	//
	// ISV verification code, used to verify if the user is authorized by an ISV. This parameter is deprecated.
	//
	// example:
	//
	// skdi3kksloslikd****
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The Viber message type. This parameter is for the international site (alibabacloud.com). You can ignore it for the China site (aliyun.com). <props="intl">Valid values:
	//
	// <props="intl">
	//
	// - **promotion**: marketing messages.
	//
	//
	//
	// <props="intl">
	//
	// - **transaction**: notification messages.
	//
	// example:
	//
	// promotion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language. For language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
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
	// A list of recipients.
	SenderListShrink *string `json:"SenderList,omitempty" xml:"SenderList,omitempty"`
	// A custom tag for a Viber message.
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
	// The template code. View the template code on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Template Design*	- page.
	//
	// example:
	//
	// 1119***************
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name. View the template name on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Manage*	- > **Template Design*	- page.
	//
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The timeout period for sending a Viber message. This parameter is for the international site (alibabacloud.com). You can ignore it for the China site (aliyun.com).<props="intl">Unit: seconds. Valid values: 30 to 1209600.
	//
	// example:
	//
	// 46
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s SendChatappMassMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *SendChatappMassMessageShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *SendChatappMassMessageShrinkRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *SendChatappMassMessageShrinkRequest) GetFallBackContent() *string {
	return s.FallBackContent
}

func (s *SendChatappMassMessageShrinkRequest) GetFallBackDuration() *int32 {
	return s.FallBackDuration
}

func (s *SendChatappMassMessageShrinkRequest) GetFallBackId() *string {
	return s.FallBackId
}

func (s *SendChatappMassMessageShrinkRequest) GetFallBackRule() *string {
	return s.FallBackRule
}

func (s *SendChatappMassMessageShrinkRequest) GetFrom() *string {
	return s.From
}

func (s *SendChatappMassMessageShrinkRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *SendChatappMassMessageShrinkRequest) GetLabel() *string {
	return s.Label
}

func (s *SendChatappMassMessageShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *SendChatappMassMessageShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendChatappMassMessageShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendChatappMassMessageShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendChatappMassMessageShrinkRequest) GetSenderListShrink() *string {
	return s.SenderListShrink
}

func (s *SendChatappMassMessageShrinkRequest) GetTag() *string {
	return s.Tag
}

func (s *SendChatappMassMessageShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SendChatappMassMessageShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendChatappMassMessageShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SendChatappMassMessageShrinkRequest) GetTtl() *int64 {
	return s.Ttl
}

func (s *SendChatappMassMessageShrinkRequest) SetChannelType(v string) *SendChatappMassMessageShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetCustSpaceId(v string) *SendChatappMassMessageShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetCustWabaId(v string) *SendChatappMassMessageShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackContent(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackDuration(v int32) *SendChatappMassMessageShrinkRequest {
	s.FallBackDuration = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackId(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackRule(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackRule = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFrom(v string) *SendChatappMassMessageShrinkRequest {
	s.From = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetIsvCode(v string) *SendChatappMassMessageShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetLabel(v string) *SendChatappMassMessageShrinkRequest {
	s.Label = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetLanguage(v string) *SendChatappMassMessageShrinkRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetOwnerId(v int64) *SendChatappMassMessageShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetResourceOwnerAccount(v string) *SendChatappMassMessageShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetResourceOwnerId(v int64) *SendChatappMassMessageShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetSenderListShrink(v string) *SendChatappMassMessageShrinkRequest {
	s.SenderListShrink = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTag(v string) *SendChatappMassMessageShrinkRequest {
	s.Tag = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTaskId(v string) *SendChatappMassMessageShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTemplateCode(v string) *SendChatappMassMessageShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTemplateName(v string) *SendChatappMassMessageShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTtl(v int64) *SendChatappMassMessageShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
