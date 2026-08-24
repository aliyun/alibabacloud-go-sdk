// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleSendMailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SingleSendMailShrinkRequest
	GetAccountName() *string
	SetAddressType(v int32) *SingleSendMailShrinkRequest
	GetAddressType() *int32
	SetAttachments(v []*SingleSendMailShrinkRequestAttachments) *SingleSendMailShrinkRequest
	GetAttachments() []*SingleSendMailShrinkRequestAttachments
	SetBccAddress(v string) *SingleSendMailShrinkRequest
	GetBccAddress() *string
	SetClickTrace(v string) *SingleSendMailShrinkRequest
	GetClickTrace() *string
	SetDomainAuth(v bool) *SingleSendMailShrinkRequest
	GetDomainAuth() *bool
	SetFromAlias(v string) *SingleSendMailShrinkRequest
	GetFromAlias() *string
	SetHeaders(v string) *SingleSendMailShrinkRequest
	GetHeaders() *string
	SetHtmlBody(v string) *SingleSendMailShrinkRequest
	GetHtmlBody() *string
	SetIpPoolId(v string) *SingleSendMailShrinkRequest
	GetIpPoolId() *string
	SetOwnerId(v int64) *SingleSendMailShrinkRequest
	GetOwnerId() *int64
	SetReplyAddress(v string) *SingleSendMailShrinkRequest
	GetReplyAddress() *string
	SetReplyAddressAlias(v string) *SingleSendMailShrinkRequest
	GetReplyAddressAlias() *string
	SetReplyToAddress(v bool) *SingleSendMailShrinkRequest
	GetReplyToAddress() *bool
	SetResourceOwnerAccount(v string) *SingleSendMailShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SingleSendMailShrinkRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *SingleSendMailShrinkRequest
	GetSubject() *string
	SetTagName(v string) *SingleSendMailShrinkRequest
	GetTagName() *string
	SetTemplateShrink(v string) *SingleSendMailShrinkRequest
	GetTemplateShrink() *string
	SetTextBody(v string) *SingleSendMailShrinkRequest
	GetTextBody() *string
	SetToAddress(v string) *SingleSendMailShrinkRequest
	GetToAddress() *string
	SetUnSubscribeFilterLevel(v string) *SingleSendMailShrinkRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *SingleSendMailShrinkRequest
	GetUnSubscribeLinkType() *string
}

type SingleSendMailShrinkRequest struct {
	// The sender address configured in the management console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***@example.net
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The address type. Valid values:
	//
	// - 0: random account
	//
	// - 1: sender address
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// Supported only when using the new SDK. Not supported through OpenAPI or signature mechanism methods. For more information, refer to [How do I send emails with attachments through the SDK?](https://help.aliyun.com/document_detail/2937843.html).
	Attachments []*SingleSendMailShrinkRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// - Specifies the BCC (blind carbon copy) recipient list for the email.
	//
	// - The system sends a copy identical to the main email content to each BCC address. The BCC information is not visible to any recipients (including ToAddress and BccAddress).
	//
	// - To protect the privacy of BCC recipients, email tracking features are disabled by default for BCC emails. This means the system does not record behavioral data such as open rates or click-through rates for BCC emails. However, billing for sending volume, sending details, and sending status statistics remain consistent with regular emails.
	//
	// - A maximum of 2 BCC recipients can be specified per send.
	//
	// Note: The SingleSendMail operation does not support the Cc (carbon copy) field. Use SMTP if you need this feature.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1@example.com,2@example.com
	BccAddress *string `json:"BccAddress,omitempty" xml:"BccAddress,omitempty"`
	// Specifies whether to enable data tracking. Valid values:
	//
	// - 1: Enable data tracking.
	//
	// - 0 (default): Disable data tracking.
	//
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	// Specifies whether to enable domain-level authentication. Valid values:
	//
	// - true
	//
	// - false
	//
	// Use this parameter only for domain-level authentication. Ignore it for sender address-level authentication.
	//
	// 1. Create the address domain-auth-created-by-system@example.com in the console. Keep the prefix before @ unchanged and use your own domain name as the suffix.
	//
	// 2.
	//
	// **API scenario**
	//
	// Set AccountName to a custom sender address for the domain. The recipient sees the custom sender address as the sender.
	//
	// **SMTP scenario**
	//
	// a. Set the domain password through the ModifyPWByDomain operation.
	//
	// b. Authenticate using the domain name and the configured password. Pass a custom address such as user@example.com as the actual sender (mailfrom). The recipient sees user@example.com as the sender.
	//
	// example:
	//
	// true
	DomainAuth *bool `json:"DomainAuth,omitempty" xml:"DomainAuth,omitempty"`
	// The sender nickname. The value cannot exceed 15 characters in length.
	//
	// For example, if the sender nickname is set to "Jane" and the sender address is test***@example.net, the recipient sees the sender address as "Jane" test***@example.net.
	//
	// example:
	//
	// Jane
	FromAlias *string `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
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
	// a. Fields prefixed with X-User- (not pushed to EventBridge or Message Service MNS. This is an API-only requirement. SMTP allows any custom fields.)
	//
	// b. Fields prefixed with X-User-Notify- (pushed to EventBridge and Message Service MNS. Both API and SMTP are supported.)
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
	// The HTML body of the email.
	//
	// Note: HtmlBody and TextBody are used for different types of email content. You must specify one of them.
	//
	// - The size limit for URL-based parameter passing is approximately 80 KB.
	//
	// - The size limit for Body-based parameter passing with the new SDK is approximately 8 MB (Java 1.4.0 or later, Python3 1.4.0 or later, PHP 1.4.0 or later).
	//
	// example:
	//
	// body
	HtmlBody *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	// The ID of the dedicated IP address pool. Users who have purchased dedicated IP addresses can use this parameter to specify the outbound IP address for this email. For more information, refer to [Dedicated IP](https://help.aliyun.com/document_detail/2932088.html).
	//
	// example:
	//
	// e4xxxxxe-4xx0-4xx3-8xxa-74cxxxxx1cef
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The reply-to address.
	//
	// example:
	//
	// test2***@example.net
	ReplyAddress *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	// The reply-to address nickname.
	//
	// example:
	//
	// Jane
	ReplyAddressAlias *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	// Specifies whether to use the reply-to address configured in the management console (the address must be verified). Valid values: true or false.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ReplyToAddress       *bool   `json:"ReplyToAddress,omitempty" xml:"ReplyToAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The email subject. The value cannot exceed 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Subject
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The tag created in the DirectMail console. Tags are used to categorize email batches. You can query the sending status of each batch by tag. If the email tracking feature is enabled, you must use an email tag when sending emails.
	//
	// The value must be 1 to 128 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The template information for template-based sending.
	//
	// When sending with a template, the HtmlBody and TextBody values are ignored.
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The text body of the email.
	//
	// Note: HtmlBody and TextBody are used for different types of email content. You must specify one of them.
	//
	// - The size limit for URL-based parameter passing is approximately 80 KB.
	//
	// - The size limit for Body-based parameter passing with the new SDK is approximately 8 MB (Java 1.4.0 or later, Python3 1.4.0 or later, PHP 1.4.0 or later).
	//
	// example:
	//
	// body
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// The destination address. You can specify multiple email addresses separated by commas. A maximum of 100 addresses are supported (mailing lists are supported).
	//
	// This parameter is required.
	//
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	// The filtering level. For more information, refer to [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// Valid values:
	//
	// - disabled: No filtering is applied.
	//
	// - default: The default policy is used. Batch addresses use sender address-level filtering.
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
	// The type of unsubscribe link. Valid values:
	//
	// - disabled: No unsubscribe link is generated.
	//
	// - default: The default policy is used. An unsubscribe link is generated when emails are sent from batch-type sender addresses to specific domains, such as those containing keywords "gmail", "yahoo", "google", "aol.com", "hotmail", "outlook", or "ymail.com". For more information, refer to [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// The display language is automatically detected based on the recipient\\"s browser settings.
	//
	// example:
	//
	// default
	UnSubscribeLinkType *string `json:"UnSubscribeLinkType,omitempty" xml:"UnSubscribeLinkType,omitempty"`
}

func (s SingleSendMailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailShrinkRequest) GoString() string {
	return s.String()
}

func (s *SingleSendMailShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SingleSendMailShrinkRequest) GetAddressType() *int32 {
	return s.AddressType
}

func (s *SingleSendMailShrinkRequest) GetAttachments() []*SingleSendMailShrinkRequestAttachments {
	return s.Attachments
}

func (s *SingleSendMailShrinkRequest) GetBccAddress() *string {
	return s.BccAddress
}

func (s *SingleSendMailShrinkRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *SingleSendMailShrinkRequest) GetDomainAuth() *bool {
	return s.DomainAuth
}

func (s *SingleSendMailShrinkRequest) GetFromAlias() *string {
	return s.FromAlias
}

func (s *SingleSendMailShrinkRequest) GetHeaders() *string {
	return s.Headers
}

func (s *SingleSendMailShrinkRequest) GetHtmlBody() *string {
	return s.HtmlBody
}

func (s *SingleSendMailShrinkRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *SingleSendMailShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SingleSendMailShrinkRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *SingleSendMailShrinkRequest) GetReplyAddressAlias() *string {
	return s.ReplyAddressAlias
}

func (s *SingleSendMailShrinkRequest) GetReplyToAddress() *bool {
	return s.ReplyToAddress
}

func (s *SingleSendMailShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SingleSendMailShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SingleSendMailShrinkRequest) GetSubject() *string {
	return s.Subject
}

func (s *SingleSendMailShrinkRequest) GetTagName() *string {
	return s.TagName
}

func (s *SingleSendMailShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *SingleSendMailShrinkRequest) GetTextBody() *string {
	return s.TextBody
}

func (s *SingleSendMailShrinkRequest) GetToAddress() *string {
	return s.ToAddress
}

func (s *SingleSendMailShrinkRequest) GetUnSubscribeFilterLevel() *string {
	return s.UnSubscribeFilterLevel
}

func (s *SingleSendMailShrinkRequest) GetUnSubscribeLinkType() *string {
	return s.UnSubscribeLinkType
}

func (s *SingleSendMailShrinkRequest) SetAccountName(v string) *SingleSendMailShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetAddressType(v int32) *SingleSendMailShrinkRequest {
	s.AddressType = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetAttachments(v []*SingleSendMailShrinkRequestAttachments) *SingleSendMailShrinkRequest {
	s.Attachments = v
	return s
}

func (s *SingleSendMailShrinkRequest) SetBccAddress(v string) *SingleSendMailShrinkRequest {
	s.BccAddress = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetClickTrace(v string) *SingleSendMailShrinkRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetDomainAuth(v bool) *SingleSendMailShrinkRequest {
	s.DomainAuth = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetFromAlias(v string) *SingleSendMailShrinkRequest {
	s.FromAlias = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetHeaders(v string) *SingleSendMailShrinkRequest {
	s.Headers = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetHtmlBody(v string) *SingleSendMailShrinkRequest {
	s.HtmlBody = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetIpPoolId(v string) *SingleSendMailShrinkRequest {
	s.IpPoolId = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetOwnerId(v int64) *SingleSendMailShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetReplyAddress(v string) *SingleSendMailShrinkRequest {
	s.ReplyAddress = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetReplyAddressAlias(v string) *SingleSendMailShrinkRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetReplyToAddress(v bool) *SingleSendMailShrinkRequest {
	s.ReplyToAddress = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetResourceOwnerAccount(v string) *SingleSendMailShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetResourceOwnerId(v int64) *SingleSendMailShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetSubject(v string) *SingleSendMailShrinkRequest {
	s.Subject = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetTagName(v string) *SingleSendMailShrinkRequest {
	s.TagName = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetTemplateShrink(v string) *SingleSendMailShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetTextBody(v string) *SingleSendMailShrinkRequest {
	s.TextBody = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetToAddress(v string) *SingleSendMailShrinkRequest {
	s.ToAddress = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetUnSubscribeFilterLevel(v string) *SingleSendMailShrinkRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetUnSubscribeLinkType(v string) *SingleSendMailShrinkRequest {
	s.UnSubscribeLinkType = &v
	return s
}

func (s *SingleSendMailShrinkRequest) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SingleSendMailShrinkRequestAttachments struct {
	// Supported only when using the new SDK. Not supported through OpenAPI or signature mechanism methods.
	//
	// example:
	//
	// test.txt
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// Supported only when using the new SDK. Not supported through OpenAPI or signature mechanism methods.
	//
	// example:
	//
	// C:\\Users\\Downloads\\test.txt
	AttachmentUrl *string `json:"AttachmentUrl,omitempty" xml:"AttachmentUrl,omitempty"`
}

func (s SingleSendMailShrinkRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailShrinkRequestAttachments) GoString() string {
	return s.String()
}

func (s *SingleSendMailShrinkRequestAttachments) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *SingleSendMailShrinkRequestAttachments) GetAttachmentUrl() *string {
	return s.AttachmentUrl
}

func (s *SingleSendMailShrinkRequestAttachments) SetAttachmentName(v string) *SingleSendMailShrinkRequestAttachments {
	s.AttachmentName = &v
	return s
}

func (s *SingleSendMailShrinkRequestAttachments) SetAttachmentUrl(v string) *SingleSendMailShrinkRequestAttachments {
	s.AttachmentUrl = &v
	return s
}

func (s *SingleSendMailShrinkRequestAttachments) Validate() error {
	return dara.Validate(s)
}
