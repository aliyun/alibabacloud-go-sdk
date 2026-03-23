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
	// The sender address configured in the Direct Mail console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***@example.net
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The address type. Valid values:
	//
	// 0: Random account
	//
	// 1: Sender address
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// This feature is available only in the latest software development kit (SDK). It is not supported by OpenAPI or signature mechanisms. For more information, see [How do I send an email with attachments using an SDK?]().
	Attachments []*SingleSendMailShrinkRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// - The list of blind carbon copy (BCC) recipients.
	//
	// - A copy of the email is sent to each BCC address. The BCC information is not visible to any recipient, including those in the ToAddress and BccAddress fields.
	//
	// - To protect the privacy of BCC recipients, email tracking is disabled by default for emails sent to BCC addresses. This means that behavioral data, such as open rates and click-through rates, is not recorded for BCC emails. However, billing, sending details, and sending status statistics are the same as for regular emails.
	//
	// - You can specify up to two BCC recipients for each email.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1@example.com,2@example.com
	BccAddress *string `json:"BccAddress,omitempty" xml:"BccAddress,omitempty"`
	// 1: Enables data tracking.
	//
	// 0 (default): Disables data tracking.
	//
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	// Enable domain-level authentication.
	//
	// - true
	//
	// - false
	//
	// Use this only for domain-level authentication. Ignore it for sender address-level authentication.
	//
	// 1\\. Create the \\`domain-auth-created-by-system\\@example.com\\` address in the console. Keep the prefix before the at sign (@) fixed and use your own domain as the suffix.
	//
	// 2\\.
	//
	// **API scenario**
	//
	// Set \\`AccountName\\` to your domain. The recipient sees the sender as \\`domain-auth-created-by-system\\@example.com\\`.
	//
	// **SMTP scenario**
	//
	// a. Set the domain password using the \\`ModifyPWByDomain\\` API.
	//
	// b. Authenticate using the domain and the set password. For the actual sender, pass a custom address, such as \\`user\\@example.com\\`, in the \\`mailfrom\\` field. The recipient sees the sender as \\`user\\@example.com\\`.
	//
	// example:
	//
	// true
	DomainAuth *bool `json:"DomainAuth,omitempty" xml:"DomainAuth,omitempty"`
	// The nickname of the sender. The nickname must be fewer than 15 characters.
	//
	// For example, if you set the nickname to "Xiao Hong" and the sender address is test\\*\\*\\*@example.net, the recipient sees the sender as "Xiao Hong" \\<test\\*\\*\\*@example.net>.
	//
	// example:
	//
	// 小红
	FromAlias *string `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	// Message header settings
	//
	// Both standard and non-standard fields must follow the syntax rules for message headers. The API supports a maximum of 10 headers in the headers field. Any headers exceeding this limit are ignored. SMTP, however, does not have this limit.
	//
	// 1\\. Standard fields
	//
	// Message-ID, List-Unsubscribe, List-Unsubscribe-Post
	//
	// Standard fields overwrite existing values in the message header.
	//
	// 2\\. Non-standard fields
	//
	// Case-insensitive
	//
	// a. Fields that start with X-User- (These are not pushed to the EventBridge event bus or Message Service MNS. They are required only for the API, whereas SMTP supports any custom header.)
	//
	// b. Fields that start with X-User-Notify- (These are pushed to the EventBridge event bus and Message Service MNS, and are supported by both the API and SMTP.)
	//
	// When pushed to EventBridge or MNS, these fields appear in the header field.
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
	// Note: Specify HtmlBody or TextBody.
	//
	// - The size of the parameter passed in a URL is limited to approximately 80 KB.
	//
	// - The new SDK limits the body parameter to approximately 8 MB (Java 1.4.0 and later, Python 3 1.4.0 and later, PHP 1.4.0 and later).
	//
	// example:
	//
	// body
	HtmlBody *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	// The ID of the dedicated IP address pool. If you purchased dedicated IP addresses, use this parameter to specify the outbound IP address for the current email. For more information, see [Dedicated IPs]().
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
	// The nickname of the reply-to address.
	//
	// example:
	//
	// 小红
	ReplyAddressAlias *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	// Specifies whether to use the reply-to address configured in the console. The reply-to address must be verified. Valid values: true and false.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ReplyToAddress       *bool   `json:"ReplyToAddress,omitempty" xml:"ReplyToAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The subject of the email. The subject cannot exceed 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Subject
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The email tag that you create in the Direct Mail console. Use tags to classify email batches and query the sending status of each batch. If email tracking is enabled, you must specify an email tag.
	//
	// The tag can be 1 to 128 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The template information for sending template-based emails.
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The text body of the email.
	//
	// Note: Specify HtmlBody or TextBody.
	//
	// - The size of the parameter passed in a URL is limited to approximately 80 KB.
	//
	// - The new SDK limits the body parameter to approximately 8 MB (Java 1.4.0 and later, Python 3 1.4.0 and later, PHP 1.4.0 and later).
	//
	// example:
	//
	// body
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// The destination address. To specify multiple addresses, separate them with commas (,). You can specify a maximum of 100 addresses. Recipient groups are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	// The filtering level. For more information, see [Unsubscribe link generation and filtering mechanism]().
	//
	// disabled: No filtering.
	//
	// default: The default policy is used. Batch addresses are filtered at the sender address level.
	//
	// mailfrom: Filters at the sender address level.
	//
	// mailfrom_domain: Filters at the email domain level.
	//
	// edm_id: Filters at the account level.
	//
	// example:
	//
	// mailfrom_domain
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	// disabled: No link is generated.
	//
	// default: The default policy is used. An unsubscribe link is generated for batch emails sent to specific domains, such as domains that contain keywords like "gmail", "yahoo",
	//
	// "google", "aol.com", "hotmail",
	//
	// "outlook", or "ymail.com". For more information, see [Unsubscribe link generation and filtering mechanism]().
	//
	// The display language is automatically detected based on the recipient\\"s browser settings.
	//
	// "outlook", or "ymail.com". For more information, see [Unsubscribe link generation and filtering mechanism]().
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
	// This feature is available only in the latest SDK. It is not supported by OpenAPI or signature mechanisms.
	//
	// example:
	//
	// test.txt
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// This feature is available only in the latest SDK. It is not supported by OpenAPI or signature mechanisms.
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
