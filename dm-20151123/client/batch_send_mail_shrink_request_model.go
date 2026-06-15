// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *BatchSendMailShrinkRequest
	GetAccountName() *string
	SetAddressType(v int32) *BatchSendMailShrinkRequest
	GetAddressType() *int32
	SetClickTrace(v string) *BatchSendMailShrinkRequest
	GetClickTrace() *string
	SetDomainAuth(v bool) *BatchSendMailShrinkRequest
	GetDomainAuth() *bool
	SetHeaders(v string) *BatchSendMailShrinkRequest
	GetHeaders() *string
	SetIpPoolId(v string) *BatchSendMailShrinkRequest
	GetIpPoolId() *string
	SetOwnerId(v int64) *BatchSendMailShrinkRequest
	GetOwnerId() *int64
	SetReceiversShrink(v string) *BatchSendMailShrinkRequest
	GetReceiversShrink() *string
	SetReceiversName(v string) *BatchSendMailShrinkRequest
	GetReceiversName() *string
	SetReplyAddress(v string) *BatchSendMailShrinkRequest
	GetReplyAddress() *string
	SetReplyAddressAlias(v string) *BatchSendMailShrinkRequest
	GetReplyAddressAlias() *string
	SetResourceOwnerAccount(v string) *BatchSendMailShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BatchSendMailShrinkRequest
	GetResourceOwnerId() *int64
	SetTagName(v string) *BatchSendMailShrinkRequest
	GetTagName() *string
	SetTemplateContentShrink(v string) *BatchSendMailShrinkRequest
	GetTemplateContentShrink() *string
	SetTemplateName(v string) *BatchSendMailShrinkRequest
	GetTemplateName() *string
	SetUnSubscribeFilterLevel(v string) *BatchSendMailShrinkRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *BatchSendMailShrinkRequest
	GetUnSubscribeLinkType() *string
}

type BatchSendMailShrinkRequest struct {
	// The sender address configured in the management console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Valid values:
	//
	// - 0: random account
	//
	// - 1: sender address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// Valid values:
	//
	// - 1: Enables data tracking.
	//
	// - 0 (default): Disables data tracking.
	//
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	// Specifies whether to enable domain-level authentication.
	//
	// - true
	//
	// - false
	//
	// Use this parameter only for domain-level authentication. Ignore it for sender address-level authentication.
	//
	// 1. Create the address domain-auth-created-by-system@example.com in the console. Keep the prefix before @ unchanged and replace the suffix with your own domain name.
	//
	// 2.
	//
	// **API scenario**
	//
	// Set AccountName to the domain name. The recipient sees domain-auth-created-by-system@example.com as the sender.
	//
	// **SMTP scenario**
	//
	// a. Call the ModifyPWByDomain operation to set the domain password.
	//
	// b. Authenticate with the domain name and the configured password. Set the actual sender (mailfrom) to a custom address such as user@example.com. The recipient sees user@example.com as the sender.
	//
	// example:
	//
	// true
	DomainAuth *bool `json:"DomainAuth,omitempty" xml:"DomainAuth,omitempty"`
	// The email header settings.
	//
	// Both standard and non-standard fields must comply with the syntax requirements for headers defined in the standard. A maximum of 10 headers can be passed through the headers field when sending emails via API. Headers exceeding this limit are ignored. SMTP has no such limit.
	//
	// 1. Standard fields
	//
	// Message-ID, List-Unsubscribe, List-Unsubscribe-Post
	//
	// Standard fields overwrite the original values in the email header.
	//
	// 2. Non-standard fields
	//
	// Case-insensitive.
	//
	// a. Fields prefixed with X-User- (not pushed to EventBridge or Message Service (MNS). This restriction applies to API only. SMTP allows any custom fields.)
	//
	// b. Fields prefixed with X-User-Notify- (pushed to EventBridge and Message Service (MNS). Both API and SMTP are supported.)
	//
	// When pushed to EventBridge or MNS, these fields are included under the header field.
	//
	// example:
	//
	// {
	//
	//       "Message-ID": "<d52ce63e-a0d5-4f95-b6a9-e1256a44f5fb@example.net>",
	//
	//       "X-User-UID1": "UID-1-000001",
	//
	//       "X-User-UID2": "UID-2-000001",
	//
	//       "X-User-Notify-UID1": "UID-3-000001",
	//
	//       "X-User-Notify-UID2": "UID-4-000001"
	//
	//
	//
	// }
	Headers *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The ID of the dedicated IP address pool. Users who have purchased dedicated IP addresses can use this parameter to specify the outbound IP address for this email sending.
	//
	// example:
	//
	// e4xxxxxe-4xx0-4xx3-8xxa-74cxxxxx1cef
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recipient list. The number of recipients must not exceed 100. Use this parameter or ReceiversName. If both Receivers and ReceiversName are specified, ReceiversName takes precedence.
	//
	// Example: [{"To":["Jackie@example.com"],"TemplateData":{"UserName":"Jackie"}},{"To":["Tom@example.com"],"TemplateData":{"UserName":"Tom"}}].
	ReceiversShrink *string `json:"Receivers,omitempty" xml:"Receivers,omitempty"`
	// The name of a pre-created recipient list that has recipients uploaded.
	//
	// > **Note**
	//
	// > The number of recipients in the list must not exceed the remaining daily quota. Otherwise, the email sending fails.
	//
	// > Wait at least 10 minutes after triggering the task before deleting the recipient list. Otherwise, the email sending may fail.
	//
	// example:
	//
	// test2
	ReceiversName *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	// The reply-to address.
	//
	// example:
	//
	// test2***@example.net
	ReplyAddress *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	// The alias of the reply-to address.
	//
	// example:
	//
	// 小红
	ReplyAddressAlias    *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the email tag.
	//
	// example:
	//
	// test3
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The custom email content. Directly specify the content without creating a template in advance. Use this parameter or TemplateName. If both TemplateContent and TemplateName are specified, TemplateName takes precedence.
	TemplateContentShrink *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of a pre-created and approved template.
	//
	// example:
	//
	// test1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The filtering level. For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// - disabled: No filtering is applied.
	//
	// - default: Uses the default policy. Batch addresses use sender address-level filtering.
	//
	// - mailfrom: Sender address-level filtering.
	//
	// - mailfrom_domain: Sender domain-level filtering.
	//
	// - edm_id: Account-level filtering.
	//
	// example:
	//
	// mailfrom_domain
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	// The type of the generated unsubscribe link. For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// - disabled: No unsubscribe link is generated.
	//
	// - default: Uses the default policy. An unsubscribe link is generated when a batch-type sender address sends emails to specific domains, such as domains containing keywords "gmail", "yahoo", "google", "aol.com", "hotmail", "outlook", or "ymail.com".
	//
	// The display language is automatically determined based on the recipient\\"s browser settings.
	//
	// example:
	//
	// default
	UnSubscribeLinkType *string `json:"UnSubscribeLinkType,omitempty" xml:"UnSubscribeLinkType,omitempty"`
}

func (s BatchSendMailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMailShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMailShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *BatchSendMailShrinkRequest) GetAddressType() *int32 {
	return s.AddressType
}

func (s *BatchSendMailShrinkRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *BatchSendMailShrinkRequest) GetDomainAuth() *bool {
	return s.DomainAuth
}

func (s *BatchSendMailShrinkRequest) GetHeaders() *string {
	return s.Headers
}

func (s *BatchSendMailShrinkRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *BatchSendMailShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchSendMailShrinkRequest) GetReceiversShrink() *string {
	return s.ReceiversShrink
}

func (s *BatchSendMailShrinkRequest) GetReceiversName() *string {
	return s.ReceiversName
}

func (s *BatchSendMailShrinkRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *BatchSendMailShrinkRequest) GetReplyAddressAlias() *string {
	return s.ReplyAddressAlias
}

func (s *BatchSendMailShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BatchSendMailShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BatchSendMailShrinkRequest) GetTagName() *string {
	return s.TagName
}

func (s *BatchSendMailShrinkRequest) GetTemplateContentShrink() *string {
	return s.TemplateContentShrink
}

func (s *BatchSendMailShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *BatchSendMailShrinkRequest) GetUnSubscribeFilterLevel() *string {
	return s.UnSubscribeFilterLevel
}

func (s *BatchSendMailShrinkRequest) GetUnSubscribeLinkType() *string {
	return s.UnSubscribeLinkType
}

func (s *BatchSendMailShrinkRequest) SetAccountName(v string) *BatchSendMailShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetAddressType(v int32) *BatchSendMailShrinkRequest {
	s.AddressType = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetClickTrace(v string) *BatchSendMailShrinkRequest {
	s.ClickTrace = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetDomainAuth(v bool) *BatchSendMailShrinkRequest {
	s.DomainAuth = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetHeaders(v string) *BatchSendMailShrinkRequest {
	s.Headers = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetIpPoolId(v string) *BatchSendMailShrinkRequest {
	s.IpPoolId = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetOwnerId(v int64) *BatchSendMailShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetReceiversShrink(v string) *BatchSendMailShrinkRequest {
	s.ReceiversShrink = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetReceiversName(v string) *BatchSendMailShrinkRequest {
	s.ReceiversName = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetReplyAddress(v string) *BatchSendMailShrinkRequest {
	s.ReplyAddress = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetReplyAddressAlias(v string) *BatchSendMailShrinkRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetResourceOwnerAccount(v string) *BatchSendMailShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetResourceOwnerId(v int64) *BatchSendMailShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetTagName(v string) *BatchSendMailShrinkRequest {
	s.TagName = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetTemplateContentShrink(v string) *BatchSendMailShrinkRequest {
	s.TemplateContentShrink = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetTemplateName(v string) *BatchSendMailShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetUnSubscribeFilterLevel(v string) *BatchSendMailShrinkRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *BatchSendMailShrinkRequest) SetUnSubscribeLinkType(v string) *BatchSendMailShrinkRequest {
	s.UnSubscribeLinkType = &v
	return s
}

func (s *BatchSendMailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
