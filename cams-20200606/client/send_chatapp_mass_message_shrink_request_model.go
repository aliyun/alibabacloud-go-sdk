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
	SenderListShrink *string `json:"SenderList,omitempty" xml:"SenderList,omitempty"`
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
