// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSingleSendMailAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SingleSendMailAdvanceRequest
	GetAccountName() *string
	SetAddressType(v int32) *SingleSendMailAdvanceRequest
	GetAddressType() *int32
	SetAttachments(v []*SingleSendMailAdvanceRequestAttachments) *SingleSendMailAdvanceRequest
	GetAttachments() []*SingleSendMailAdvanceRequestAttachments
	SetBccAddress(v string) *SingleSendMailAdvanceRequest
	GetBccAddress() *string
	SetClickTrace(v string) *SingleSendMailAdvanceRequest
	GetClickTrace() *string
	SetDomainAuth(v bool) *SingleSendMailAdvanceRequest
	GetDomainAuth() *bool
	SetFromAlias(v string) *SingleSendMailAdvanceRequest
	GetFromAlias() *string
	SetHeaders(v string) *SingleSendMailAdvanceRequest
	GetHeaders() *string
	SetHtmlBody(v string) *SingleSendMailAdvanceRequest
	GetHtmlBody() *string
	SetIpPoolId(v string) *SingleSendMailAdvanceRequest
	GetIpPoolId() *string
	SetOwnerId(v int64) *SingleSendMailAdvanceRequest
	GetOwnerId() *int64
	SetReplyAddress(v string) *SingleSendMailAdvanceRequest
	GetReplyAddress() *string
	SetReplyAddressAlias(v string) *SingleSendMailAdvanceRequest
	GetReplyAddressAlias() *string
	SetReplyToAddress(v bool) *SingleSendMailAdvanceRequest
	GetReplyToAddress() *bool
	SetResourceOwnerAccount(v string) *SingleSendMailAdvanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SingleSendMailAdvanceRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *SingleSendMailAdvanceRequest
	GetSubject() *string
	SetTagName(v string) *SingleSendMailAdvanceRequest
	GetTagName() *string
	SetTemplate(v *SingleSendMailAdvanceRequestTemplate) *SingleSendMailAdvanceRequest
	GetTemplate() *SingleSendMailAdvanceRequestTemplate
	SetTextBody(v string) *SingleSendMailAdvanceRequest
	GetTextBody() *string
	SetToAddress(v string) *SingleSendMailAdvanceRequest
	GetToAddress() *string
	SetUnSubscribeFilterLevel(v string) *SingleSendMailAdvanceRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *SingleSendMailAdvanceRequest
	GetUnSubscribeLinkType() *string
}

type SingleSendMailAdvanceRequest struct {
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
	Attachments []*SingleSendMailAdvanceRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
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
	Template *SingleSendMailAdvanceRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
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

func (s SingleSendMailAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SingleSendMailAdvanceRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SingleSendMailAdvanceRequest) GetAddressType() *int32 {
	return s.AddressType
}

func (s *SingleSendMailAdvanceRequest) GetAttachments() []*SingleSendMailAdvanceRequestAttachments {
	return s.Attachments
}

func (s *SingleSendMailAdvanceRequest) GetBccAddress() *string {
	return s.BccAddress
}

func (s *SingleSendMailAdvanceRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *SingleSendMailAdvanceRequest) GetDomainAuth() *bool {
	return s.DomainAuth
}

func (s *SingleSendMailAdvanceRequest) GetFromAlias() *string {
	return s.FromAlias
}

func (s *SingleSendMailAdvanceRequest) GetHeaders() *string {
	return s.Headers
}

func (s *SingleSendMailAdvanceRequest) GetHtmlBody() *string {
	return s.HtmlBody
}

func (s *SingleSendMailAdvanceRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *SingleSendMailAdvanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SingleSendMailAdvanceRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *SingleSendMailAdvanceRequest) GetReplyAddressAlias() *string {
	return s.ReplyAddressAlias
}

func (s *SingleSendMailAdvanceRequest) GetReplyToAddress() *bool {
	return s.ReplyToAddress
}

func (s *SingleSendMailAdvanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SingleSendMailAdvanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SingleSendMailAdvanceRequest) GetSubject() *string {
	return s.Subject
}

func (s *SingleSendMailAdvanceRequest) GetTagName() *string {
	return s.TagName
}

func (s *SingleSendMailAdvanceRequest) GetTemplate() *SingleSendMailAdvanceRequestTemplate {
	return s.Template
}

func (s *SingleSendMailAdvanceRequest) GetTextBody() *string {
	return s.TextBody
}

func (s *SingleSendMailAdvanceRequest) GetToAddress() *string {
	return s.ToAddress
}

func (s *SingleSendMailAdvanceRequest) GetUnSubscribeFilterLevel() *string {
	return s.UnSubscribeFilterLevel
}

func (s *SingleSendMailAdvanceRequest) GetUnSubscribeLinkType() *string {
	return s.UnSubscribeLinkType
}

func (s *SingleSendMailAdvanceRequest) SetAccountName(v string) *SingleSendMailAdvanceRequest {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetAddressType(v int32) *SingleSendMailAdvanceRequest {
	s.AddressType = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetAttachments(v []*SingleSendMailAdvanceRequestAttachments) *SingleSendMailAdvanceRequest {
	s.Attachments = v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetBccAddress(v string) *SingleSendMailAdvanceRequest {
	s.BccAddress = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetClickTrace(v string) *SingleSendMailAdvanceRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetDomainAuth(v bool) *SingleSendMailAdvanceRequest {
	s.DomainAuth = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetFromAlias(v string) *SingleSendMailAdvanceRequest {
	s.FromAlias = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetHeaders(v string) *SingleSendMailAdvanceRequest {
	s.Headers = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetHtmlBody(v string) *SingleSendMailAdvanceRequest {
	s.HtmlBody = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetIpPoolId(v string) *SingleSendMailAdvanceRequest {
	s.IpPoolId = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetOwnerId(v int64) *SingleSendMailAdvanceRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetReplyAddress(v string) *SingleSendMailAdvanceRequest {
	s.ReplyAddress = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetReplyAddressAlias(v string) *SingleSendMailAdvanceRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetReplyToAddress(v bool) *SingleSendMailAdvanceRequest {
	s.ReplyToAddress = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetResourceOwnerAccount(v string) *SingleSendMailAdvanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetResourceOwnerId(v int64) *SingleSendMailAdvanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetSubject(v string) *SingleSendMailAdvanceRequest {
	s.Subject = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetTagName(v string) *SingleSendMailAdvanceRequest {
	s.TagName = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetTemplate(v *SingleSendMailAdvanceRequestTemplate) *SingleSendMailAdvanceRequest {
	s.Template = v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetTextBody(v string) *SingleSendMailAdvanceRequest {
	s.TextBody = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetToAddress(v string) *SingleSendMailAdvanceRequest {
	s.ToAddress = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetUnSubscribeFilterLevel(v string) *SingleSendMailAdvanceRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetUnSubscribeLinkType(v string) *SingleSendMailAdvanceRequest {
	s.UnSubscribeLinkType = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) Validate() error {
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

type SingleSendMailAdvanceRequestAttachments struct {
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
	AttachmentUrlObject io.Reader `json:"AttachmentUrl,omitempty" xml:"AttachmentUrl,omitempty"`
}

func (s SingleSendMailAdvanceRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailAdvanceRequestAttachments) GoString() string {
	return s.String()
}

func (s *SingleSendMailAdvanceRequestAttachments) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *SingleSendMailAdvanceRequestAttachments) GetAttachmentUrlObject() io.Reader {
	return s.AttachmentUrlObject
}

func (s *SingleSendMailAdvanceRequestAttachments) SetAttachmentName(v string) *SingleSendMailAdvanceRequestAttachments {
	s.AttachmentName = &v
	return s
}

func (s *SingleSendMailAdvanceRequestAttachments) SetAttachmentUrlObject(v io.Reader) *SingleSendMailAdvanceRequestAttachments {
	s.AttachmentUrlObject = v
	return s
}

func (s *SingleSendMailAdvanceRequestAttachments) Validate() error {
	return dara.Validate(s)
}

type SingleSendMailAdvanceRequestTemplate struct {
	// The variables and their values in the template.
	TemplateData map[string]*string `json:"TemplateData,omitempty" xml:"TemplateData,omitempty"`
	// The template ID.
	//
	// example:
	//
	// xxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SingleSendMailAdvanceRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailAdvanceRequestTemplate) GoString() string {
	return s.String()
}

func (s *SingleSendMailAdvanceRequestTemplate) GetTemplateData() map[string]*string {
	return s.TemplateData
}

func (s *SingleSendMailAdvanceRequestTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SingleSendMailAdvanceRequestTemplate) SetTemplateData(v map[string]*string) *SingleSendMailAdvanceRequestTemplate {
	s.TemplateData = v
	return s
}

func (s *SingleSendMailAdvanceRequestTemplate) SetTemplateId(v string) *SingleSendMailAdvanceRequestTemplate {
	s.TemplateId = &v
	return s
}

func (s *SingleSendMailAdvanceRequestTemplate) Validate() error {
	return dara.Validate(s)
}
