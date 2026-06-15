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
	// `0`: A random account.
	//
	// `1`: A sender address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// This feature is available only through the latest SDKs. It is not supported for OpenAPI calls or signature-based authentication. For more information, see [How do I send an email with an attachment by using an SDK?](https://help.aliyun.com/document_detail/2937843.html).
	Attachments []*SingleSendMailRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// - A comma-separated list of BCC recipients.
	//
	// - The system sends a copy of the email to each BCC recipient. The BCC information is hidden from all recipients, including those specified in `ToAddress` and `BccAddress`.
	//
	// - To protect privacy, email tracking features (such as open and click tracking) are disabled for emails sent to BCC recipients. However, billing and sending status are still tracked.
	//
	// - A maximum of two BCC recipients are allowed per request.
	//
	// Note: The `SingleSendMail` API operation does not support a CC field. To send carbon copies, use SMTP.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1@example.com,2@example.com
	BccAddress *string `json:"BccAddress,omitempty" xml:"BccAddress,omitempty"`
	// Specifies whether to enable click tracking. Valid values: `"1"` enables click tracking, and `"0"` disables it (default).
	//
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	// Specifies whether to enable domain-level authentication.
	//
	// - `true`
	//
	// - `false`
	//
	// This parameter is used only for domain-level authentication. Ignore it for sender address-level authentication.
	//
	// 1\\. Create the address `domain-auth-created-by-system@example.com` in the console. The prefix must be fixed, and the suffix must be your domain.
	//
	// 2\\.
	//
	// **API scenario**
	//
	// Set `AccountName` to your domain. Recipients will see the sender as `domain-auth-created-by-system@example.com`.
	//
	// **SMTP scenario**
	//
	// a. Call the `ModifyPWByDomain` API operation to set a password for the domain.
	//
	// b. Authenticate with the domain and the configured password. Pass a custom address, such as `user@example.com`, as the actual sender in the `MAIL FROM` command. Recipients will see `user@example.com` as the sender.
	//
	// example:
	//
	// true
	DomainAuth *bool `json:"DomainAuth,omitempty" xml:"DomainAuth,omitempty"`
	// The sender name. It must be 15 characters or shorter.
	//
	// For example, if you set the sender name to "Xiaohong" and the sender address is `test***@example.net`, the recipient sees the sender as "Xiaohong" \\<test\\*\\*\\*@example.net>.
	//
	// example:
	//
	// Jane
	FromAlias *string `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	// Custom email header settings.
	//
	// Both standard and non-standard fields must comply with standard header syntax. You can specify up to 10 headers for an API call. Excess headers are ignored. This limit does not apply to SMTP.
	//
	// 1\\. Standard fields
	//
	// `Message-ID`, `List-Unsubscribe`, `List-Unsubscribe-Post`
	//
	// Standard fields overwrite existing values in the email header.
	//
	// 2\\. Non-standard fields
	//
	// Case-insensitive.
	//
	// a. Fields starting with `X-User-`: These are not pushed to EventBridge or Message Service (MNS). This prefix is required only for API calls, not for SMTP.
	//
	// b. Fields starting with `X-User-Notify-`: These are pushed to EventBridge and MNS. This is supported for both API and SMTP calls.
	//
	// When pushed to EventBridge or MNS, the header object will contain these fields.
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
	// Note: You must specify either `HtmlBody` or `TextBody`.
	//
	// - The size of the body is limited to approximately 80 KB when passed as a URL parameter.
	//
	// - For recent SDKs (Java 1.4.0+, Python 3 1.4.0+, and PHP 1.4.0+), the request body is limited to approximately 8 MB.
	//
	// example:
	//
	// body
	HtmlBody *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	// The ID of the dedicated IP pool. If you have purchased dedicated IPs, you can use this parameter to select which dedicated IP pool to use for sending the email. For more information, see [Dedicated IP](https://help.aliyun.com/document_detail/2932088.html).
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
	// The name displayed for the reply-to address.
	//
	// example:
	//
	// Jane
	ReplyAddressAlias *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	// Specifies whether to use the default reply-to address configured in the console. This address must be verified. Valid values: true, false.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ReplyToAddress       *bool   `json:"ReplyToAddress,omitempty" xml:"ReplyToAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The subject of the email, with a maximum length of 256 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Subject
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// A tag for categorizing email batches, which you can create in the Direct Mail console. Tags allow you to query the sending status of each batch and are required if you enable email tracking. The tag must be 1 to 128 characters long and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The template information for sending a templated email.
	Template *SingleSendMailRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	// The text body of the email.
	//
	// Note: You must specify either `HtmlBody` or `TextBody`.
	//
	// - The size of the body is limited to approximately 80 KB when passed as a URL parameter.
	//
	// - For recent SDKs (Java 1.4.0+, Python 3 1.4.0+, and PHP 1.4.0+), the request body is limited to approximately 8 MB.
	//
	// example:
	//
	// body
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// The destination email address(es). To specify multiple addresses, separate them with commas (up to 100).
	//
	// This parameter is required.
	//
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	// The filtering level. For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// `disabled`: No filtering.
	//
	// `default`: Uses the default policy. For batch addresses, filtering is applied at the sender address level.
	//
	// `mailfrom`: Filters at the sender address level.
	//
	// `mailfrom_domain`: Filters at the sender domain level.
	//
	// `edm_id`: Filters at the account level.
	//
	// example:
	//
	// mailfrom_domain
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	// `disabled`: Does not generate an unsubscribe link.
	//
	// `default`: Uses the default policy. For batch sender addresses, an unsubscribe link is generated when sending to specific domains containing keywords such as "gmail", "yahoo",
	//
	// "google", "aol.com", "hotmail",
	//
	// "outlook", and "ymail.com". For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// The display language is automatically determined based on the recipient\\"s browser settings.
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
	// The filename of the attachment.
	//
	// example:
	//
	// test.txt
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// The local file path of the attachment that the SDK will use.
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
	// The variables and their values for the template.
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
