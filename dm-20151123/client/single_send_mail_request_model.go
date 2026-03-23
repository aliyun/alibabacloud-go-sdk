// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleSendMailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SingleSendMailRequest
	GetAccountName() *string
	SetAddressType(v int32) *SingleSendMailRequest
	GetAddressType() *int32
	SetAttachments(v []*SingleSendMailRequestAttachments) *SingleSendMailRequest
	GetAttachments() []*SingleSendMailRequestAttachments
	SetBccAddress(v string) *SingleSendMailRequest
	GetBccAddress() *string
	SetClickTrace(v string) *SingleSendMailRequest
	GetClickTrace() *string
	SetDomainAuth(v bool) *SingleSendMailRequest
	GetDomainAuth() *bool
	SetFromAlias(v string) *SingleSendMailRequest
	GetFromAlias() *string
	SetHeaders(v string) *SingleSendMailRequest
	GetHeaders() *string
	SetHtmlBody(v string) *SingleSendMailRequest
	GetHtmlBody() *string
	SetIpPoolId(v string) *SingleSendMailRequest
	GetIpPoolId() *string
	SetOwnerId(v int64) *SingleSendMailRequest
	GetOwnerId() *int64
	SetReplyAddress(v string) *SingleSendMailRequest
	GetReplyAddress() *string
	SetReplyAddressAlias(v string) *SingleSendMailRequest
	GetReplyAddressAlias() *string
	SetReplyToAddress(v bool) *SingleSendMailRequest
	GetReplyToAddress() *bool
	SetResourceOwnerAccount(v string) *SingleSendMailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SingleSendMailRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *SingleSendMailRequest
	GetSubject() *string
	SetTagName(v string) *SingleSendMailRequest
	GetTagName() *string
	SetTemplate(v *SingleSendMailRequestTemplate) *SingleSendMailRequest
	GetTemplate() *SingleSendMailRequestTemplate
	SetTextBody(v string) *SingleSendMailRequest
	GetTextBody() *string
	SetToAddress(v string) *SingleSendMailRequest
	GetToAddress() *string
	SetUnSubscribeFilterLevel(v string) *SingleSendMailRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *SingleSendMailRequest
	GetUnSubscribeLinkType() *string
}

type SingleSendMailRequest struct {
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
	Attachments []*SingleSendMailRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
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
	Template *SingleSendMailRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
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

func (s SingleSendMailRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailRequest) GoString() string {
	return s.String()
}

func (s *SingleSendMailRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SingleSendMailRequest) GetAddressType() *int32 {
	return s.AddressType
}

func (s *SingleSendMailRequest) GetAttachments() []*SingleSendMailRequestAttachments {
	return s.Attachments
}

func (s *SingleSendMailRequest) GetBccAddress() *string {
	return s.BccAddress
}

func (s *SingleSendMailRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *SingleSendMailRequest) GetDomainAuth() *bool {
	return s.DomainAuth
}

func (s *SingleSendMailRequest) GetFromAlias() *string {
	return s.FromAlias
}

func (s *SingleSendMailRequest) GetHeaders() *string {
	return s.Headers
}

func (s *SingleSendMailRequest) GetHtmlBody() *string {
	return s.HtmlBody
}

func (s *SingleSendMailRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *SingleSendMailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SingleSendMailRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *SingleSendMailRequest) GetReplyAddressAlias() *string {
	return s.ReplyAddressAlias
}

func (s *SingleSendMailRequest) GetReplyToAddress() *bool {
	return s.ReplyToAddress
}

func (s *SingleSendMailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SingleSendMailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SingleSendMailRequest) GetSubject() *string {
	return s.Subject
}

func (s *SingleSendMailRequest) GetTagName() *string {
	return s.TagName
}

func (s *SingleSendMailRequest) GetTemplate() *SingleSendMailRequestTemplate {
	return s.Template
}

func (s *SingleSendMailRequest) GetTextBody() *string {
	return s.TextBody
}

func (s *SingleSendMailRequest) GetToAddress() *string {
	return s.ToAddress
}

func (s *SingleSendMailRequest) GetUnSubscribeFilterLevel() *string {
	return s.UnSubscribeFilterLevel
}

func (s *SingleSendMailRequest) GetUnSubscribeLinkType() *string {
	return s.UnSubscribeLinkType
}

func (s *SingleSendMailRequest) SetAccountName(v string) *SingleSendMailRequest {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailRequest) SetAddressType(v int32) *SingleSendMailRequest {
	s.AddressType = &v
	return s
}

func (s *SingleSendMailRequest) SetAttachments(v []*SingleSendMailRequestAttachments) *SingleSendMailRequest {
	s.Attachments = v
	return s
}

func (s *SingleSendMailRequest) SetBccAddress(v string) *SingleSendMailRequest {
	s.BccAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetClickTrace(v string) *SingleSendMailRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailRequest) SetDomainAuth(v bool) *SingleSendMailRequest {
	s.DomainAuth = &v
	return s
}

func (s *SingleSendMailRequest) SetFromAlias(v string) *SingleSendMailRequest {
	s.FromAlias = &v
	return s
}

func (s *SingleSendMailRequest) SetHeaders(v string) *SingleSendMailRequest {
	s.Headers = &v
	return s
}

func (s *SingleSendMailRequest) SetHtmlBody(v string) *SingleSendMailRequest {
	s.HtmlBody = &v
	return s
}

func (s *SingleSendMailRequest) SetIpPoolId(v string) *SingleSendMailRequest {
	s.IpPoolId = &v
	return s
}

func (s *SingleSendMailRequest) SetOwnerId(v int64) *SingleSendMailRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleSendMailRequest) SetReplyAddress(v string) *SingleSendMailRequest {
	s.ReplyAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetReplyAddressAlias(v string) *SingleSendMailRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *SingleSendMailRequest) SetReplyToAddress(v bool) *SingleSendMailRequest {
	s.ReplyToAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetResourceOwnerAccount(v string) *SingleSendMailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleSendMailRequest) SetResourceOwnerId(v int64) *SingleSendMailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleSendMailRequest) SetSubject(v string) *SingleSendMailRequest {
	s.Subject = &v
	return s
}

func (s *SingleSendMailRequest) SetTagName(v string) *SingleSendMailRequest {
	s.TagName = &v
	return s
}

func (s *SingleSendMailRequest) SetTemplate(v *SingleSendMailRequestTemplate) *SingleSendMailRequest {
	s.Template = v
	return s
}

func (s *SingleSendMailRequest) SetTextBody(v string) *SingleSendMailRequest {
	s.TextBody = &v
	return s
}

func (s *SingleSendMailRequest) SetToAddress(v string) *SingleSendMailRequest {
	s.ToAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetUnSubscribeFilterLevel(v string) *SingleSendMailRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *SingleSendMailRequest) SetUnSubscribeLinkType(v string) *SingleSendMailRequest {
	s.UnSubscribeLinkType = &v
	return s
}

func (s *SingleSendMailRequest) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SingleSendMailRequestAttachments struct {
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

func (s SingleSendMailRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailRequestAttachments) GoString() string {
	return s.String()
}

func (s *SingleSendMailRequestAttachments) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *SingleSendMailRequestAttachments) GetAttachmentUrl() *string {
	return s.AttachmentUrl
}

func (s *SingleSendMailRequestAttachments) SetAttachmentName(v string) *SingleSendMailRequestAttachments {
	s.AttachmentName = &v
	return s
}

func (s *SingleSendMailRequestAttachments) SetAttachmentUrl(v string) *SingleSendMailRequestAttachments {
	s.AttachmentUrl = &v
	return s
}

func (s *SingleSendMailRequestAttachments) Validate() error {
	return dara.Validate(s)
}

type SingleSendMailRequestTemplate struct {
	// The variables and their values in the template.
	TemplateData map[string]*string `json:"TemplateData,omitempty" xml:"TemplateData,omitempty"`
	// The template ID.
	//
	// example:
	//
	// xxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SingleSendMailRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailRequestTemplate) GoString() string {
	return s.String()
}

func (s *SingleSendMailRequestTemplate) GetTemplateData() map[string]*string {
	return s.TemplateData
}

func (s *SingleSendMailRequestTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SingleSendMailRequestTemplate) SetTemplateData(v map[string]*string) *SingleSendMailRequestTemplate {
	s.TemplateData = v
	return s
}

func (s *SingleSendMailRequestTemplate) SetTemplateId(v string) *SingleSendMailRequestTemplate {
	s.TemplateId = &v
	return s
}

func (s *SingleSendMailRequestTemplate) Validate() error {
	return dara.Validate(s)
}
