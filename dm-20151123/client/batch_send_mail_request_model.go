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
	SetReceivers(v []*BatchSendMailRequestReceivers) *BatchSendMailRequest
	GetReceivers() []*BatchSendMailRequestReceivers
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
	SetTemplateContent(v *BatchSendMailRequestTemplateContent) *BatchSendMailRequest
	GetTemplateContent() *BatchSendMailRequestTemplateContent
	SetTemplateName(v string) *BatchSendMailRequest
	GetTemplateName() *string
	SetUnSubscribeFilterLevel(v string) *BatchSendMailRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *BatchSendMailRequest
	GetUnSubscribeLinkType() *string
}

type BatchSendMailRequest struct {
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
	// Valid values:
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
	// b. Authenticate with the domain name and the configured password. Pass a custom address such as user@example.com as the actual sender (mailfrom). The recipient sees user@example.com as the sender.
	//
	// example:
	//
	// true
	DomainAuth *bool `json:"DomainAuth,omitempty" xml:"DomainAuth,omitempty"`
	// The email header settings.
	//
	// Both standard and non-standard fields must comply with the syntax requirements for headers defined in the standard. A maximum of 10 headers can be passed through the headers field when sending emails via API. Headers that exceed this limit are ignored. SMTP has no such limit.
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
	// a. Fields prefixed with X-User- (not pushed to EventBridge or Message Service (MNS). This restriction applies only to API. SMTP allows any custom fields.)
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
	// The ID of the dedicated IP address pool. Users who have purchased dedicated IP addresses can use this parameter to specify the outbound IP address for this email sending task.
	//
	// example:
	//
	// e4xxxxxe-4xx0-4xx3-8xxa-74cxxxxx1cef
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recipient list. The number of recipients cannot exceed 100. Specify this parameter or ReceiversName. If both Receivers and ReceiversName are specified, ReceiversName takes precedence.
	//
	// Example: [{"To":["Jackie@example.com"],"TemplateData":{"UserName":"Jackie"}},{"To":["Tom@example.com"],"TemplateData":{"UserName":"Tom"}}].
	Receivers []*BatchSendMailRequestReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// The name of a pre-created recipient list that has recipients uploaded.
	//
	// Note:
	//
	// The number of recipients in the list must not exceed the remaining daily quota. Otherwise, the email sending fails.
	//
	// Do not delete the recipient list until at least 10 minutes after the task is triggered. Otherwise, the email sending may fail.
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
	// The tag name of the email.
	//
	// example:
	//
	// test3
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The custom email content. You can directly specify the content without creating a template in advance. Specify this parameter or TemplateName. If both TemplateContent and TemplateName are specified, TemplateName takes precedence.
	TemplateContent *BatchSendMailRequestTemplateContent `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty" type:"Struct"`
	// The name of a pre-created and approved template.
	//
	// example:
	//
	// test1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The filtering level. For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// Valid values:
	//
	// - disabled: No filtering is applied.
	//
	// - default: The default policy is used. Batch addresses are filtered at the sender address level.
	//
	// - mailfrom: Filtering at the sender address level.
	//
	// - mailfrom_domain: Filtering at the sender domain level.
	//
	// - edm_id: Filtering at the account level.
	//
	// example:
	//
	// mailfrom_domain
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	// The type of the generated unsubscribe link. For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// Valid values:
	//
	// - disabled: No unsubscribe link is generated.
	//
	// - default: The default policy is used. An unsubscribe link is generated when emails are sent from a batch-type sender address to specific domains that contain keywords such as "gmail", "yahoo", "google", "aol.com", "hotmail", "outlook", or "ymail.com".
	//
	// The display language is automatically determined based on the recipient\\"s browser settings.
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

func (s *BatchSendMailRequest) GetReceivers() []*BatchSendMailRequestReceivers {
	return s.Receivers
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

func (s *BatchSendMailRequest) GetTemplateContent() *BatchSendMailRequestTemplateContent {
	return s.TemplateContent
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

func (s *BatchSendMailRequest) SetReceivers(v []*BatchSendMailRequestReceivers) *BatchSendMailRequest {
	s.Receivers = v
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

func (s *BatchSendMailRequest) SetTemplateContent(v *BatchSendMailRequestTemplateContent) *BatchSendMailRequest {
	s.TemplateContent = v
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
	if s.Receivers != nil {
		for _, item := range s.Receivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TemplateContent != nil {
		if err := s.TemplateContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchSendMailRequestReceivers struct {
	// The email template parameters. This parameter is of the JSON map type.
	TemplateData map[string]*string `json:"TemplateData,omitempty" xml:"TemplateData,omitempty"`
	// The recipient list. This parameter is of the array type.
	To []*string `json:"To,omitempty" xml:"To,omitempty" type:"Repeated"`
}

func (s BatchSendMailRequestReceivers) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMailRequestReceivers) GoString() string {
	return s.String()
}

func (s *BatchSendMailRequestReceivers) GetTemplateData() map[string]*string {
	return s.TemplateData
}

func (s *BatchSendMailRequestReceivers) GetTo() []*string {
	return s.To
}

func (s *BatchSendMailRequestReceivers) SetTemplateData(v map[string]*string) *BatchSendMailRequestReceivers {
	s.TemplateData = v
	return s
}

func (s *BatchSendMailRequestReceivers) SetTo(v []*string) *BatchSendMailRequestReceivers {
	s.To = v
	return s
}

func (s *BatchSendMailRequestReceivers) Validate() error {
	return dara.Validate(s)
}

type BatchSendMailRequestTemplateContent struct {
	// The display name of the sender.
	//
	// example:
	//
	// Jackie
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The HTML body of the email.
	//
	// Note: HtmlBody and TextBody are used for different types of email content. You must specify one of them.
	//
	// The new SDK uses Body for parameter passing with a size limit of approximately 8 MB (Java 1.4.0 and later, Python3 1.4.0 and later, PHP 1.4.0 and later).
	//
	// example:
	//
	// <h1>全场九折，仅限今日</h1>
	HtmlBody *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	// The subject of the email.
	//
	// example:
	//
	// 黑色星期五，专属折扣来袭
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The plain text body of the email.
	//
	// Note: HtmlBody and TextBody are used for different types of email content. You must specify one of them.
	//
	// The new SDK uses Body for parameter passing with a size limit of approximately 8 MB (Java 1.4.0 and later, Python3 1.4.0 and later, PHP 1.4.0 and later).
	//
	// example:
	//
	// 全场九折，仅限今日
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
}

func (s BatchSendMailRequestTemplateContent) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMailRequestTemplateContent) GoString() string {
	return s.String()
}

func (s *BatchSendMailRequestTemplateContent) GetAlias() *string {
	return s.Alias
}

func (s *BatchSendMailRequestTemplateContent) GetHtmlBody() *string {
	return s.HtmlBody
}

func (s *BatchSendMailRequestTemplateContent) GetSubject() *string {
	return s.Subject
}

func (s *BatchSendMailRequestTemplateContent) GetTextBody() *string {
	return s.TextBody
}

func (s *BatchSendMailRequestTemplateContent) SetAlias(v string) *BatchSendMailRequestTemplateContent {
	s.Alias = &v
	return s
}

func (s *BatchSendMailRequestTemplateContent) SetHtmlBody(v string) *BatchSendMailRequestTemplateContent {
	s.HtmlBody = &v
	return s
}

func (s *BatchSendMailRequestTemplateContent) SetSubject(v string) *BatchSendMailRequestTemplateContent {
	s.Subject = &v
	return s
}

func (s *BatchSendMailRequestTemplateContent) SetTextBody(v string) *BatchSendMailRequestTemplateContent {
	s.TextBody = &v
	return s
}

func (s *BatchSendMailRequestTemplateContent) Validate() error {
	return dara.Validate(s)
}
