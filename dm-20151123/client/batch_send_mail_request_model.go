// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *BatchSendMailRequest
	GetAccountName() *string
	SetAddressType(v int32) *BatchSendMailRequest
	GetAddressType() *int32
	SetClickTrace(v string) *BatchSendMailRequest
	GetClickTrace() *string
	SetDomainAuth(v bool) *BatchSendMailRequest
	GetDomainAuth() *bool
	SetHeaders(v string) *BatchSendMailRequest
	GetHeaders() *string
	SetIpPoolId(v string) *BatchSendMailRequest
	GetIpPoolId() *string
	SetOwnerId(v int64) *BatchSendMailRequest
	GetOwnerId() *int64
	SetReceiversName(v string) *BatchSendMailRequest
	GetReceiversName() *string
	SetReplyAddress(v string) *BatchSendMailRequest
	GetReplyAddress() *string
	SetReplyAddressAlias(v string) *BatchSendMailRequest
	GetReplyAddressAlias() *string
	SetResourceOwnerAccount(v string) *BatchSendMailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BatchSendMailRequest
	GetResourceOwnerId() *int64
	SetTagName(v string) *BatchSendMailRequest
	GetTagName() *string
	SetTemplateName(v string) *BatchSendMailRequest
	GetTemplateName() *string
	SetUnSubscribeFilterLevel(v string) *BatchSendMailRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *BatchSendMailRequest
	GetUnSubscribeLinkType() *string
}

type BatchSendMailRequest struct {
	// The sender address configured in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// - 0: Random account
	//
	// - 1: Sender address
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// - 1: Enables the data tracking feature.
	//
	// - 0 (default): Disables the data tracking feature.
	//
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	// Enables domain-level authentication.
	//
	// - true
	//
	// - false
	//
	// Use this parameter only for domain-level authentication. Ignore it for sender address-level authentication.
	//
	// 1\\. The console creates the address \\`domain-auth-created-by-system\\@example.com\\`. Do not change the prefix before the at sign (@). Replace the domain suffix with your own domain.
	//
	// 2\\.
	//
	// **API scenario**
	//
	// Set \\`AccountName\\` to your domain. Recipients see \\`domain-auth-created-by-system\\@example.com\\` as the sender.
	//
	// **SMTP scenario**
	//
	// a. Use the \\`ModifyPWByDomain\\` API to set a password for your domain.
	//
	// b. Authenticate using your domain and the password. Set the actual sender address (\\`mailfrom\\`) to a custom address, such as \\`user\\@example.com\\`. Recipients see \\`user\\@example.com\\` as the sender.
	//
	// example:
	//
	// true
	DomainAuth *bool `json:"DomainAuth,omitempty" xml:"DomainAuth,omitempty"`
	// Message header settings.
	//
	// All fields, standard or non-standard, must follow standard header syntax. For API calls, the \\`headers\\` field supports up to 10 headers. Any headers beyond this limit are ignored. SMTP does not have a header limit.
	//
	// 1\\. Standard fields
	//
	// \\`Message-ID\\`, \\`List-Unsubscribe\\`, \\`List-Unsubscribe-Post\\`
	//
	// Standard fields overwrite existing values in the message header.
	//
	// 2\\. Non-standard fields
	//
	// Case-insensitive
	//
	// a. Start with \\`X-User-\\`. These fields are not pushed to EventBridge or Message Service. They are required only for API calls. SMTP supports any custom header.
	//
	// b. Start with \\`X-User-Notify-\\`. These fields are pushed to EventBridge and Message Service. They are supported by both API and SMTP.
	//
	// When pushed to EventBridge or Message Service, these fields appear under the \\`headers\\` object.
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
	// The ID of the dedicated IP address pool. If you purchased dedicated IP addresses, use this parameter to specify the egress IP address for sending the email.
	//
	// example:
	//
	// e4xxxxxe-4xx0-4xx3-8xxa-74cxxxxx1cef
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of a pre-created recipient list to which recipients have been uploaded.
	//
	// Note:
	//
	// The number of recipients in the list must not exceed your remaining daily quota. Otherwise, email sending fails.
	//
	// Do not delete the recipient list for at least 10 minutes after triggering the task. Otherwise, email sending may fail.
	//
	// This parameter is required.
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
	// The alias for the reply-to address.
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
	// The name of a pre-created and approved template.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The filtering level. For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// - disabled: No filtering.
	//
	// - default: Uses the default policy. Batch emails are filtered at the sender address level.
	//
	// - mailfrom: Filters at the sender address level.
	//
	// - mailfrom_domain: Filters at the email domain level.
	//
	// - edm_id: Filters at the account level.
	//
	// example:
	//
	// mailfrom_domain
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	// The type of unsubscribe link to generate. For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// - disabled: Does not generate a link.
	//
	// - default: Uses the default policy. An unsubscribe link is generated when batch emails are sent from a sender address to specific domains, such as those containing the keywords "gmail", "yahoo", "google", "aol.com", "hotmail", "outlook", or "ymail.com".
	//
	// The language of the unsubscribe link matches the recipient\\"s browser language setting.
	//
	// example:
	//
	// default
	UnSubscribeLinkType *string `json:"UnSubscribeLinkType,omitempty" xml:"UnSubscribeLinkType,omitempty"`
}

func (s BatchSendMailRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMailRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMailRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *BatchSendMailRequest) GetAddressType() *int32 {
	return s.AddressType
}

func (s *BatchSendMailRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *BatchSendMailRequest) GetDomainAuth() *bool {
	return s.DomainAuth
}

func (s *BatchSendMailRequest) GetHeaders() *string {
	return s.Headers
}

func (s *BatchSendMailRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *BatchSendMailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchSendMailRequest) GetReceiversName() *string {
	return s.ReceiversName
}

func (s *BatchSendMailRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *BatchSendMailRequest) GetReplyAddressAlias() *string {
	return s.ReplyAddressAlias
}

func (s *BatchSendMailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BatchSendMailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BatchSendMailRequest) GetTagName() *string {
	return s.TagName
}

func (s *BatchSendMailRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *BatchSendMailRequest) GetUnSubscribeFilterLevel() *string {
	return s.UnSubscribeFilterLevel
}

func (s *BatchSendMailRequest) GetUnSubscribeLinkType() *string {
	return s.UnSubscribeLinkType
}

func (s *BatchSendMailRequest) SetAccountName(v string) *BatchSendMailRequest {
	s.AccountName = &v
	return s
}

func (s *BatchSendMailRequest) SetAddressType(v int32) *BatchSendMailRequest {
	s.AddressType = &v
	return s
}

func (s *BatchSendMailRequest) SetClickTrace(v string) *BatchSendMailRequest {
	s.ClickTrace = &v
	return s
}

func (s *BatchSendMailRequest) SetDomainAuth(v bool) *BatchSendMailRequest {
	s.DomainAuth = &v
	return s
}

func (s *BatchSendMailRequest) SetHeaders(v string) *BatchSendMailRequest {
	s.Headers = &v
	return s
}

func (s *BatchSendMailRequest) SetIpPoolId(v string) *BatchSendMailRequest {
	s.IpPoolId = &v
	return s
}

func (s *BatchSendMailRequest) SetOwnerId(v int64) *BatchSendMailRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSendMailRequest) SetReceiversName(v string) *BatchSendMailRequest {
	s.ReceiversName = &v
	return s
}

func (s *BatchSendMailRequest) SetReplyAddress(v string) *BatchSendMailRequest {
	s.ReplyAddress = &v
	return s
}

func (s *BatchSendMailRequest) SetReplyAddressAlias(v string) *BatchSendMailRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *BatchSendMailRequest) SetResourceOwnerAccount(v string) *BatchSendMailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BatchSendMailRequest) SetResourceOwnerId(v int64) *BatchSendMailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BatchSendMailRequest) SetTagName(v string) *BatchSendMailRequest {
	s.TagName = &v
	return s
}

func (s *BatchSendMailRequest) SetTemplateName(v string) *BatchSendMailRequest {
	s.TemplateName = &v
	return s
}

func (s *BatchSendMailRequest) SetUnSubscribeFilterLevel(v string) *BatchSendMailRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *BatchSendMailRequest) SetUnSubscribeLinkType(v string) *BatchSendMailRequest {
	s.UnSubscribeLinkType = &v
	return s
}

func (s *BatchSendMailRequest) Validate() error {
	return dara.Validate(s)
}
